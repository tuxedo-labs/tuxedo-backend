package routes

import (
	"tuxedo/config"
	"tuxedo/handler"
	"tuxedo/middleware"
	"tuxedo/models/entity"

	"github.com/gofiber/fiber/v2"
)

var auth = middleware.Auth
var admin = middleware.AdminRole

func SetupRouter(r *fiber.App) {
	app := r.Group("/api")
	// authentication
	app.Post("/auth/login", handler.Login)
	app.Post("/auth/register", handler.Register)
	app.Post("/auth/verify-token", handler.VerifyCode)
	app.Post("/auth/resend-verify-token", handler.ResendVerifyRequest)

	//users
	app.Get("/users/profile", auth, handler.GetProfile)
	app.Put("/users/update", auth, handler.UpdateProfile)

	// blog
	app.Get("/blog", handler.GetBlog)
	app.Get("/blog/:id", handler.GetBlogByID)
	app.Post("/blog", auth, handler.PostBlog)
	app.Put("/blog/:id", auth, handler.UpdateBlog)
	app.Delete("/blog/:id", handler.DeleteBlog)
}

func AutoMigrate() {
	config.RunMigrate(&entity.Users{})
	config.RunMigrate(&entity.Contacts{})
	config.RunMigrate(&entity.VerifyToken{})
	config.RunMigrate(&entity.Blog{})
}
