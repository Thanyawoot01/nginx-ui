package backup

import (
	"os/exec"
	"fmt"
	"os"
	"path/filepath"

	"github.com/0xJacky/Nginx-UI/internal/nginx"
	"github.com/0xJacky/Nginx-UI/settings"
	"github.com/uozi-tech/cosy"
	"github.com/uozi-tech/cosy/logger"
	cosysettings "github.com/uozi-tech/cosy/settings"
)

func dumpMariaDB(destDir string) error {

	dbHost := settings.DatabaseSettings.Host
	dbUser := settings.DatabaseSettings.User
	dbPass := settings.DatabaseSettings.Password
	dbName := settings.DatabaseSettings.Name

	sqlFile := filepath.Join(destDir, "database.sql")

	cmd := exec.Command(
		"mysqldump",
		"-h", dbHost,
		"-u", dbUser,
		fmt.Sprintf("-p%s", dbPass),
		dbName,
	)

	outFile, err := os.Create(sqlFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd.Stdout = outFile
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
// backupNginxUIFiles backs up the nginx-ui configuration and database files
func backupNginxUIFiles(destDir string) error {

	configPath := cosysettings.ConfPath
	if configPath == "" {
		return ErrConfigPathEmpty
	}

	destConfigPath := filepath.Join(destDir, "app.ini")

	if err := copyFile(configPath, destConfigPath); err != nil {
		return cosy.WrapErrorWithParams(ErrCopyConfigFile, err.Error())
	}

	dbName := settings.DatabaseSettings.GetName()
	dbFile := dbName + ".db"

	dbDir := filepath.Dir(configPath)
	dbPath := filepath.Join(dbDir, dbFile)

	// Try MariaDB dump first
	if err := dumpMariaDB(destDir); err != nil {

		logger.Warn("MariaDB dump failed: %v", err)

		// fallback sqlite
		if _, err := os.Stat(dbPath); err == nil {

			destDBPath := filepath.Join(destDir, dbFile)

			if err := copyFile(dbPath, destDBPath); err != nil {
				return cosy.WrapErrorWithParams(ErrCopyDBFile, err.Error())
			}
		}
	}

	return nil
}

// backupNginxFiles backs up the nginx configuration directory
func backupNginxFiles(destDir string) error {
	// Get nginx config directory
	nginxConfigDir := nginx.GetConfPath()
	if nginxConfigDir == "" {
		return ErrNginxConfigDirEmpty
	}

	// Copy nginx config directory
	if err := copyDirectory(nginxConfigDir, destDir); err != nil {
		return cosy.WrapErrorWithParams(ErrCopyNginxConfigDir, err.Error())
	}

	return nil
}

// writeHashInfoFile creates a hash information file for verification
func writeHashInfoFile(hashFilePath string, info HashInfo) error {
	content := fmt.Sprintf("nginx-ui_hash: %s\nnginx_hash: %s\ntimestamp: %s\nversion: %s\n",
		info.NginxUIHash, info.NginxHash, info.Timestamp, info.Version)

	if err := os.WriteFile(hashFilePath, []byte(content), 0644); err != nil {
		return cosy.WrapErrorWithParams(ErrCreateHashFile, err.Error())
	}

	return nil
}
