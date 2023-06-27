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
	route.Get("/member", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.GetMembers)          // get all members
	route.Post("/member", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.CreateMember)       // create a new member
	route.Post("/role", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.CreateRole)           // create a new role
	route.Patch("/role/assign", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.AssignRole)   // assign role to member
	route.Patch("/role/unassign", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.RemoveRole) // unassign role from member
	route.Post("/rank", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.CreateRank)           // create a new rank
	route.Get("/rank", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.GetRanks)              // get all ranks
	route.Post("/user/sign/out", middleware.JWTProtected(), middleware.JWTExpirationChecker(), controllers.UserSignOut) // de-authorization user
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)                                      // renew Access & Refresh tokens
}
