package controllers

import (
	"idealist/app/crud"
	"idealist/app/models"
	"idealist/pkg/utils"
	"idealist/platform/database"

	"github.com/gofiber/fiber/v2"
)

// Idea creation for users.
// @Description Create a new idea
// @Summary Create a new idea
// @Tags Idea
// @Accept json
// @Produce json
// @Success 201 {string} status "ok"
// @Router /v1/idea/ [post]
func CreateIdea(c *fiber.Ctx) error {
	// Initialize idea instance
	schema := &crud.IdeaCreate{}

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
		println("Error getting member")
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   error.Error(),
		})
	}

	if member == nil {
		println("Error getting member #2")
		// Return status 404 and error message.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Member not found!",
		})
	}

	memberID := member.ID

	// Create a new idea.
	idea := &models.Idea{
		Title:        schema.Title,
		Body:         schema.Body,
		AuthorID:     memberID,
		ProjectID:    schema.ProjectID,
		BrainStormID: schema.BrainStormID,
	}

	// Create a new idea.
	error = db.Create(&idea).Error

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
		"msg":   "Idea created successfully!",
	})
}

// Brainstorm creation for users.
// @Description Create a new brainstorm with some ideas
// @Summary Create a new brainstorm
// @Tags Brainstorm
// @Accept json
// @Produce json
// @Success 201 {string} status "ok"
// @Router /v1/brainstorm/ [post]
func CreateBrainstorm(c *fiber.Ctx) error {
	// Initialize idea instance
	schema := &crud.BrainstormCreate{}

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
		println("Error getting member")
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   error.Error(),
		})
	}

	if member == nil {
		println("Error getting member #2")
		// Return status 404 and error message.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Member not found!",
		})
	}

	memberID := member.ID

	// Create a new brainstorm.
	brainstorm := &models.BrainStorm{
		Name:      schema.Name,
		ProjectID: schema.ProjectID,
		AuthorID:  memberID,
	}

	ideas := schema.Ideas

	for _, idea := range ideas {
		brainstorm.Ideas = append(brainstorm.Ideas, &models.Idea{
			Title: idea.Title,
			Body:  idea.Body,

			// A user starts a brainstorm with some of his ideas, then
			// its shared and other users can add their ideas to it later.
			AuthorID: memberID,

			// Ideas can remain associated with the project if the brainstorm
			// gets deleted
			ProjectID: schema.ProjectID,
		})
	}

	// Create a new idea.
	error = db.Create(&brainstorm).Error

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
		"msg":   "Brainstorm created successfully!",
	})
}
