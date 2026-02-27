package settings

import "fmt"

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

var DatabaseSettings = &Database{
	Host:     "127.0.0.1",
	Port:     3306,
	User:     "nginx_ui",
	Password: "",
	Name:     "nginx_ui",
}

func (d *Database) GetName() string {
	if d.Name == "" {
		d.Name = "nginx_ui"
	}
	return d.Name
}

// DSN returns MySQL/MariaDB Data Source Name
func (d *Database) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.Name)
}
