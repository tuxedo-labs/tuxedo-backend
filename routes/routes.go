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
	// /users/update

	// blog
	app.Get("/blog", handler.GetBlog)
	app.Post("/blog", handler.PostBlog)
}

func AutoMigrate() {
	config.RunMigrate(&entity.Users{})
	config.RunMigrate(&entity.Contacts{})
	config.RunMigrate(&entity.VerifyToken{})
	config.RunMigrate(&entity.Blog{})
}
