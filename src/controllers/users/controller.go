package users

import (
	"github.com/gofiber/fiber/v2"
	"go.architecture/core/configs"
	"go.architecture/core/helpers"
	"go.architecture/core/models"
)

func List(ctx *fiber.Ctx) error {
	var users []models.Users
	configs.DBM.Find(&users)
	// Serialize the users
	var serializedUsers []models.UsersDTO = make([]models.UsersDTO, 0)
	for _, result := range users {
		serializedUsers = append(serializedUsers, models.UsersDTO{
			ID:        result.ID,
			Name:      result.Name,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		})
	}
	// Return the user as a JSON response
	return ctx.JSON(serializedUsers)
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
