package routes

import (
	"management-helper/app/controllers"
	"management-helper/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Get("/member", middleware.JWTProtected(), controllers.GetMembers)          // get all members
	route.Post("/member", middleware.JWTProtected(), controllers.CreateMember)       // create a new member
	route.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut) // de-authorization user
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens
}
