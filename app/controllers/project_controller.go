package controllers

import (
	"idealist/app/crud"
	"idealist/app/models"
	"idealist/pkg/utils"
	"idealist/platform/database"

	"github.com/gofiber/fiber/v2"
)

// Project creation for users.
// @Description Create a new project
// @Summary Create a new project
// @Tags Project
// @Accept json
// @Produce json
// @Success 201 {string} status "ok"
// @Router /v1/idea/ [post]
func CreateProject(c *fiber.Ctx) error {
	// Initialize idea instance
	schema := &crud.ProjectCreate{}

	// Checking received data from JSON body.
	if err := c.BodyParser(schema); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db := database.DB

	// Get the user id from the token.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		println("Error extracting token metadata: ")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id := claims.UserID

	// Get the user from the database.
	member := &models.Member{}
	error := db.Where(&models.Member{
		SubClassID: id,
	}).First(&member).Error

	if error != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   error.Error(),
		})
	}

	if member == nil {
		// Return status 404 and error message.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Member not found!",
		})
	}

	memberID := member.ID

	// Create a new idea.
	project := &models.Project{
		Name:        schema.Name,
		Description: schema.Description,
		Icon:        schema.Icon,
		OwnerID:     memberID,
	}

	// Create a new idea.
	error = db.Create(&project).Error

	if error != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   error.Error(),
		})
	}

	// Return status 201 and success message.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "Project created successfully!",
	})
}

// User's projects.
// @Description Get all projects owned by the user
// @Summary Get all projects owned by the user
// @Tags Project
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/idea/ [get]
func GetOwnedProjects(c *fiber.Ctx) error {
	// Create database connection.
	db := database.DB

	// Get the user id from the token.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		println("Error extracting token metadata: ")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	id := claims.UserID

	// Get the user from the database.
	member := &models.Member{}
	error := db.Where(&models.Member{
		SubClassID: id,
	}).First(&member).Error

	if error != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   error.Error(),
		})
	}

	if member == nil {
		// Return status 404 and error message.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Member not found!",
		})
	}

	memberID := member.ID

	// Get all projects owned by the user.
	projects := []models.Project{}
	error = db.Where(&models.Project{
		OwnerID: memberID,
	}).Find(&projects).Error

	if error != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   error.Error(),
		})
	}

	// Return status 200 and success message.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":    false,
		"projects": projects,
	})
}
