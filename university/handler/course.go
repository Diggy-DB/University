package handler

import (
	"university/database"
	"university/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AddCourse(ctx *fiber.Ctx) error {
	course := new(model.Course)
	err := ctx.BodyParser(course)
	if err != nil {
		return ctx.Status(400).JSON("Failed to Body Parse!")
	}
	validate := validator.New()
	errValidate := validate.Struct(course)
	if errValidate != nil {
		return ctx.Status(400).JSON("Failed to Validate!")
	}
	newCourse := model.Course{
		CourseName:     course.CourseName,
		CourseNo:       course.CourseNo,
		TeacherId:      course.TeacherId,
		Credit:         course.Credit,
		DepartmentName: course.DepartmentName,
	}
	err = database.DB.Create(&newCourse).Error
	if err != nil {
		return ctx.Status(400).JSON("Failed to Create Course!")
	}
	return ctx.JSON(newCourse)
}
func ShowAllCourse(ctx *fiber.Ctx) error {
	var courses []model.Course
	err := database.DB.Find(&courses)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to get all courses",
		})
	}
	return ctx.JSON(courses)
}
