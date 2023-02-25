package users

import (
	"github.com/gofiber/fiber/v2"
	"go.architecture/core/configs"
	_ "go.architecture/core/exceptions"
	"go.architecture/core/helpers"
	_ "go.architecture/core/interfaces"
	"go.architecture/core/models"
)

// List Get List of users
// @Summary Get list of users
// @Description Get all users
// @Accept  json
// @Produce  json
// @Tags User
// @Success 200 {object} interfaces.SuccessResponse{data=[]models.UsersDTO}
// @Failure 400 {object} exceptions.HTTPError
// @Failure 404 {object} exceptions.HTTPError
// @Failure 500 {object} exceptions.HTTPError
// @Router /api/v1/users [get]
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
		"data":    serializedUsers,
	})
}

// Create new User
// @Summary Post Add new user
// @Description Creates a new user
// @Accept  json
// @Produce  json
// @Tags User
// @Success 200 {object} interfaces.SuccessResponse{data=models.UsersDSO}
// @Failure 400 {object} exceptions.HTTPError
// @Failure 404 {object} exceptions.HTTPError
// @Failure 500 {object} exceptions.HTTPError
// @Router /api/v1/users [post]
// @Param request body models.UsersDSO true "query params"
func Create(ctx *fiber.Ctx) error {
	var userRequest models.UsersDSO
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  400,
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
			"status":  400,
		})
	}
	// Return success response
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    userModel,
	})
}

// Update User
// @Summary Update User by id
// @Description Gets ID and updates user
// @Accept  json
// @Produce  json
// @Tags User
// @Success 200 {object} interfaces.SuccessResponse
// @Failure 400 {object} exceptions.HTTPError
// @Failure 404 {object} exceptions.HTTPError
// @Failure 500 {object} exceptions.HTTPError
// @Router /api/v1/users [put]
// @Param id  path int true "User ID"
// @Param request body models.UsersDSO true "query params"
func Update(ctx *fiber.Ctx) error {
	var userRequest models.UsersDSO
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  500,
		})
	}
	// Find the user in the database
	var user models.Users
	result := configs.DBM.First(&user, ctx.Params("id"))
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"status":  404,
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

// Delete User
// @Summary Delete User by id
// @Description Gets ID and deletes user
// @Accept  json
// @Produce  json
// @Tags User
// @Success 200 {object} interfaces.SuccessResponse
// @Failure 400 {object} exceptions.HTTPError
// @Failure 404 {object} exceptions.HTTPError
// @Failure 500 {object} exceptions.HTTPError
// @Router /api/v1/users [delete]
// @Param id  path int true "User ID"
func Delete(ctx *fiber.Ctx) error {
	// Find the user in the database
	var user models.Users
	result := configs.DBM.First(&user, ctx.Params("id"))
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"status":  404,
		})
	}
	// Delete the user in the database
	result = configs.DBM.Delete(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
			"status":  400,
		})
	}
	// Return success response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
		"data":    user.ID,
	})
}
