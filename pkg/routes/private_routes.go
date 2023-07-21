package routes

import (
	"idealist/app/controllers"
	"idealist/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/user/sign/out", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.UserSignOut) // de-authorization user
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)                                      // renew Access & Refresh tokens

	route.Post("/brainstorm", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.CreateBrainstorm) // create new brainstorm
	route.Post("/idea", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.CreateIdea)             // create new idea
	route.Post("/project", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.CreateProject)       // create new project
	route.Get("/project/owned", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.GetOwnedProjects)
}
