package router

import (
	"net/http"

	"github.com/0xJacky/Nginx-UI/api/analytic"
	"github.com/0xJacky/Nginx-UI/api/audit"
	"github.com/0xJacky/Nginx-UI/api/backup"
	"github.com/0xJacky/Nginx-UI/api/certificate"
	"github.com/0xJacky/Nginx-UI/api/cluster"
	"github.com/0xJacky/Nginx-UI/api/config"
	"github.com/0xJacky/Nginx-UI/api/crypto"
	dnsapi "github.com/0xJacky/Nginx-UI/api/dns"
	"github.com/0xJacky/Nginx-UI/api/event"
	"github.com/0xJacky/Nginx-UI/api/external_notify"
	"github.com/0xJacky/Nginx-UI/api/geolite"
	"github.com/0xJacky/Nginx-UI/api/license"
	"github.com/0xJacky/Nginx-UI/api/llm"
	"github.com/0xJacky/Nginx-UI/api/nginx"
	nginxLog "github.com/0xJacky/Nginx-UI/api/nginx_log"
	"github.com/0xJacky/Nginx-UI/api/notification"
	"github.com/0xJacky/Nginx-UI/api/pages"
	"github.com/0xJacky/Nginx-UI/api/public"
	"github.com/0xJacky/Nginx-UI/api/settings"
	"github.com/0xJacky/Nginx-UI/api/sites"
	"github.com/0xJacky/Nginx-UI/api/streams"
	"github.com/0xJacky/Nginx-UI/api/system"
	"github.com/0xJacky/Nginx-UI/api/template"
	"github.com/0xJacky/Nginx-UI/api/terminal"
	"github.com/0xJacky/Nginx-UI/api/upstream"
	"github.com/0xJacky/Nginx-UI/api/user"
	"github.com/0xJacky/Nginx-UI/internal/middleware"
	"github.com/0xJacky/Nginx-UI/mcp"
	"github.com/0xJacky/Nginx-UI/model"
	"github.com/gin-gonic/gin"
	"github.com/uozi-tech/cosy"
	"github.com/uozi-tech/cosy/debug"
)

func InitRouter() {
	r := cosy.GetEngine()

	r.Use(audit.LoggingMiddleware())

	r.SetTrustedProxies(nil)

	// Add CORS middleware to allow all origins
	r.Use(middleware.CORS())

	initEmbedRoute(r)

	pages.InitRouter(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	})

	mcp.InitRouter(r)

	root := r.Group("/api", middleware.IPWhiteList())
	{
		public.InitRouter(root)
		crypto.InitPublicRouter(root)
		user.InitAuthRouter(root)
		license.InitRouter(root)

		system.InitPublicRouter(root)
		system.InitSelfCheckRouter(root)
		backup.InitRouter(root)

		// Local-only routes (no proxy) - authorization required
		local := root.Group("/", middleware.AuthRequired())
		{
			llm.InitLocalRouter(local)
		}

		// Authorization required and not websocket request
		g := root.Group("/", middleware.AuthRequired(), middleware.Proxy())
		{
			// Shared (All authenticated users)
			user.InitUserRouter(g) // Profile
			analytic.InitRouter(g)
			llm.InitRouter(g)
			g.GET("/geolite/status", geolite.GetStatus)

			// Admin only
			adminOnly := g.Group("/", middleware.RequireRole(model.RoleAdmin))
			{
				debug.InitRouter(adminOnly)
				user.InitManageUserRouter(adminOnly)
				system.InitPrivateRouter(adminOnly)
				settings.InitRouter(adminOnly)
				backup.InitAutoBackupRouter(adminOnly)
				cluster.InitRouter(adminOnly)
				notification.InitRouter(adminOnly)
				external_notify.InitRouter(adminOnly)
			}

			// Admin and WebDev
			webDevShared := g.Group("/", middleware.RequireRole(model.RoleAdmin, model.RoleWebDev))
			{
				nginx.InitRouter(webDevShared)
				sites.InitRouter(webDevShared)
				streams.InitRouter(webDevShared)
				config.InitRouter(webDevShared)
				template.InitRouter(webDevShared)
				certificate.InitCertificateRouter(webDevShared)
				certificate.InitDNSCredentialRouter(webDevShared)
				certificate.InitAcmeUserRouter(webDevShared)
				dnsapi.InitRouter(webDevShared)
				nginxLog.InitRouter(webDevShared)
			}
		}

		// Authorization required and websocket request
		w := root.Group("/", middleware.AuthRequired(), middleware.ProxyWs())
		{
			// Shared
			analytic.InitWebSocketRouter(w)
			event.InitRouter(w)

			// Admin only
			adminOnlyWs := w.Group("/", middleware.RequireRole(model.RoleAdmin))
			{
				system.InitWebSocketRouter(adminOnlyWs)
				cluster.InitWebSocketRouter(adminOnlyWs)
				adminOnlyWs.GET("/geolite/download", geolite.DownloadGeoLiteDB)

				// Terminal (also require SecureSession)
				o := adminOnlyWs.Group("", middleware.RequireSecureSession())
				{
					terminal.InitRouter(o)
				}
			}

			// Admin and WebDev
			webDevSharedWs := w.Group("/", middleware.RequireRole(model.RoleAdmin, model.RoleWebDev))
			{
				certificate.InitCertificateWebSocketRouter(webDevSharedWs)
				nginxLog.InitWebSocketRouter(webDevSharedWs)
				upstream.InitRouter(webDevSharedWs)
				nginx.InitWebSocketRouter(webDevSharedWs)
			}
		}
	}
}
