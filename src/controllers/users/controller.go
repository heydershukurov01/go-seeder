package users

import (
	"github.com/gofiber/fiber/v2"
	"go.architecture/core/configs"
	"go.architecture/core/helpers"
	"go.architecture/core/models"
)

// List all users
func List(ctx *fiber.Ctx) error {
	var users []models.Users
	configs.DBM.Find(&users)
	// Serialize the users
	var serializedUsers []models.UsersDTO
	for _, result := range users {
		var serializedUser models.UsersDTO = models.UsersDTO{
			ID:        result.ID,
			Name:      result.Name,
			Surname:   result.Surname,
			Age:       result.Age,
			Email:     result.Email,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
		// Log the user
		serializedUsers = append(serializedUsers, serializedUser)
	}
	// Return the user as a JSON response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users retrieved successfully",
		"users":   serializedUsers,
	})
}
func Create(ctx *fiber.Ctx) error {
	var userRequest models.UsersDSO
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}
	// Create the user in the database Version 1
	//userModel := models.Users{
	//	Name:     userRequest.Name,
	//	Surname:  userRequest.Surname,
	//	Email:    userRequest.Email,
	//	Age:      userRequest.Age,
	//	Password: userRequest.Password,
	//}
	// Create the user in the database Version 2 (using helpers.Spread)
	userModel := models.Users{}
	helpers.Spread(userRequest, &userModel)
	result := configs.DBM.Create(&userModel)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success response
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    userModel,
	})
}
func Update(ctx *fiber.Ctx) error {
	var userRequest models.UsersDSO
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}
	// Find the user in the database
	var user models.Users
	result := configs.DBM.First(&user, ctx.Params("id"))
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// Update the user in the database
	helpers.Spread(userRequest, &user)
	result = configs.DBM.Updates(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success response
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    user.ID,
	})
}
func Delete(ctx *fiber.Ctx) error {
	// Find the user in the database
	var user models.Users
	result := configs.DBM.First(&user, ctx.Params("id"))
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// Delete the user in the database
	result = configs.DBM.Delete(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
		"user":    user.ID,
	})
}
