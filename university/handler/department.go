package handler

import (
	"university/database"
	"university/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AddDepartment(ctx *fiber.Ctx) error {
	department := new(model.Department)
	err := ctx.BodyParser(department)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Body Parse",
		})
	}
	validate := validator.New()
	errValidate := validate.Struct(department)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Validate!",
		})
	}
	err = database.DB.Create(&department).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Create Department",
		})
	}
	return ctx.JSON(department)
}
func ShowAllDepartment(ctx *fiber.Ctx) error {
	var departments []model.Department
	err := database.DB.Find(&departments).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to get all departments",
		})
	}
	return ctx.JSON(departments)
}
