package handler

import (
	"university/database"
	"university/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AddChosenCourse(ctx *fiber.Ctx) error {
	course := new(model.Chosen)
	err := ctx.BodyParser(course)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Body Parse",
		})
	}
	validate := validator.New()
	errValidate := validate.Struct(course)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to validate!",
		})
	}
	err = database.DB.Create(&course).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Create Department",
		})
	}
	return ctx.JSON(course)
}
func ShowAllChosenCourse(ctx *fiber.Ctx) error {
	var courses []model.Chosen
	err := database.DB.Find(&courses).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to get all departments",
		})
	}
	return ctx.JSON(courses)
}
