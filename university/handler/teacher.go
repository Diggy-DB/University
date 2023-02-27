package handler

import (
	"university/auth"
	"university/database"
	"university/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AddTeacher(ctx *fiber.Ctx) error {
	teacher := new(model.Teacher)
	err := ctx.BodyParser(teacher)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Body Parse!",
		})
	}
	validate := validator.New()
	if errValidate := validate.Struct(teacher); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Validate!",
		})
	}
	newTeacher := model.Teacher{
		TeacherName:    teacher.TeacherId,
		TeacherId:      teacher.TeacherName,
		DepartmentName: teacher.DepartmentName,
		Position:       teacher.Position,
		Age:            teacher.Age,
		Gender:         teacher.Gender,
		PhoneNumber:    teacher.PhoneNumber,
	}
	err = database.DB.Create(&newTeacher).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Add Teacher!",
		})
	}
	return ctx.JSON(newTeacher)
}

func GetTeacherById(ctx *fiber.Ctx) error {
	var account auth.Account
	err := ctx.BodyParser(account)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Body Parse!",
		})
	}
	validate := validator.New()
	if errValidate := validate.Struct(account); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Validate!",
		})
	}
	var teacher model.Teacher
	err = database.DB.Find(&teacher, "teacher_id = ?", account.Id).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Find teacher",
		})
	}
	teacherResponse := model.Teacher{
		TeacherName:    teacher.TeacherName,
		TeacherId:      teacher.TeacherId,
		DepartmentName: teacher.DepartmentName,
		Position:       teacher.Position,
		Age:            teacher.Age,
		Gender:         teacher.Gender,
		PhoneNumber:    teacher.PhoneNumber,
	}
	return ctx.JSON(teacherResponse)
}
