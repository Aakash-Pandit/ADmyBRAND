package services

import (
	"github.com/Aakash-Pandit/ADmyBRAND/models"
	"github.com/Aakash-Pandit/ADmyBRAND/storage"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(context *fiber.Ctx) error {
	users := &[]models.User{}

	db := storage.GetDatabase()
	err := db.Find(users).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(users)
}

func GetUserByID(context *fiber.Ctx) error {
	id := context.Params("id")
	user := &models.User{}

	db := storage.GetDatabase()
	err := db.Where("id = ?", id).First(user).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(context *fiber.Ctx) error {

	db := storage.GetDatabase()
	user := &models.User{}

	err := context.BodyParser(user)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	errors := models.ValidateUserStruct(*user)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err = db.Create(user).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	user := &models.User{}

	err := context.BodyParser(&user)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	errors := models.ValidateUserStruct(*user)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := storage.GetDatabase()

	err = db.Where("id = ?", id).Updates(user).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"detail": err.Error(),
		})
	}

	_ = db.Where("id = ?", id).First(user).Error
	return context.Status(fiber.StatusOK).JSON(user)
}

func DeleteUser(context *fiber.Ctx) error {
	user := &models.User{}
	id := context.Params("id")
	if id == "" {
		context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	db := storage.GetDatabase()
	err := db.Where("id = ?", id).First(user).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	db.Where("id = ?", id).Delete(user)

	return context.Status(fiber.StatusNoContent).JSON(&fiber.Map{})
}
