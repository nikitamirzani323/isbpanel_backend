package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/isbpanel_backend/controllers"
	"github.com/nikitamirzani323/isbpanel_backend/middleware"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("/ipaddress", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      "data",
			"BASEURL":     c.BaseURL(),
			"HOSTNAME":    c.Hostname(),
			"IP":          c.IP(),
			"IPS":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomain":   c.Subdomains(),
		})
	})
	app.Get("/dashboard", monitor.New())

	app.Post("/api/login", controllers.CheckLogin)
	app.Post("/api/valid", middleware.JWTProtected(), controllers.Home)
	app.Post("/api/alladmin", middleware.JWTProtected(), controllers.Adminhome)
	app.Post("/api/detailadmin", middleware.JWTProtected(), controllers.AdminDetail)
	app.Post("/api/saveadmin", middleware.JWTProtected(), controllers.AdminSave)

	app.Post("/api/alladminrule", middleware.JWTProtected(), controllers.Adminrulehome)
	app.Post("/api/saveadminrule", middleware.JWTProtected(), controllers.AdminruleSave)

	app.Post("/api/pasaran", middleware.JWTProtected(), controllers.Pasaranhome)
	app.Post("/api/pasaransave", middleware.JWTProtected(), controllers.Pasaransave)
	app.Post("/api/keluaran", middleware.JWTProtected(), controllers.Keluaranhome)
	app.Post("/api/keluaransave", middleware.JWTProtected(), controllers.Keluaransave)
	app.Post("/api/keluarandelete", middleware.JWTProtected(), controllers.Keluarandelete)

	app.Post("/api/prediksi", middleware.JWTProtected(), controllers.Prediksihome)
	app.Post("/api/prediksisave", middleware.JWTProtected(), controllers.Prediksisave)
	app.Post("/api/prediksidelete", middleware.JWTProtected(), controllers.Prediksidelete)

	app.Post("/api/tafsirmimpi", middleware.JWTProtected(), controllers.Tafsirmimpihome)
	app.Post("/api/tafsirmimpisave", middleware.JWTProtected(), controllers.Tafsirmimpisave)

	app.Post("/api/news", middleware.JWTProtected(), controllers.Newshome)
	app.Post("/api/newssave", middleware.JWTProtected(), controllers.Newssave)
	app.Post("/api/newsdelete", middleware.JWTProtected(), controllers.Newsdelete)
	app.Post("/api/categorynews", middleware.JWTProtected(), controllers.Categoryhome)
	app.Post("/api/categorynewssave", middleware.JWTProtected(), controllers.Categorysave)
	app.Post("/api/categorynewsdelete", middleware.JWTProtected(), controllers.Categorydelete)

	app.Post("/api/movie", middleware.JWTProtected(), controllers.Moviehome)
	app.Post("/api/moviealbum", middleware.JWTProtected(), controllers.Moviecloud)
	app.Post("/api/movieupload", middleware.JWTProtected(), controllers.Movieuploadcloud)
	app.Post("/api/genremovie", middleware.JWTProtected(), controllers.Genrehome)
	app.Post("/api/genremoviesave", middleware.JWTProtected(), controllers.Genresave)
	app.Post("/api/genremoviedelete", middleware.JWTProtected(), controllers.Genredelete)
	return app
}
