package handler

import (
	"university/auth"
	"university/database"
	"university/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AddStudents(ctx *fiber.Ctx) error {
	student := new(model.Student)
	err := ctx.BodyParser(student)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Failed to Body Parse",
		})
	}
	validate := validator.New()
	if errValidate := validate.Struct(student); errValidate != nil {
		return ctx.JSON(fiber.Map{
			"message": "Failed to Validate",
		})
	}
	newStudent := model.Student{
		StudentName:    student.StudentName,
		StudentId:      student.StudentId,
		PhoneNo:        student.PhoneNo,
		Grade:          student.Grade,
		DepartmentName: student.DepartmentName,
		Degree:         student.Degree,
		Address:        student.Address,
	}
	errCreateStudent := database.DB.Create(&newStudent).Error
	if errCreateStudent != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Create Student",
		})
	}
	return ctx.JSON(newStudent)
}
func GetStudentById(ctx *fiber.Ctx) error {
	var student model.Student
	account := new(auth.Account)
	err := ctx.BodyParser(account)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Parse JSON!",
		})
	}
	validate := validator.New()
	if errValidate := validate.Struct(student); errValidate != nil {
		return ctx.JSON(fiber.Map{
			"message": "Failed to Validate",
		})
	}
	err = database.DB.First(&student, "student_id = ?", account.Id).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to get student from Database",
		})
	}
	studentResponse := model.Student{
		StudentId:      student.StudentId,
		StudentName:    student.StudentName,
		PhoneNo:        student.PhoneNo,
		Grade:          student.Grade,
		DepartmentName: student.DepartmentName,
		Degree:         student.Degree,
		Address:        student.Address,
	}
	return ctx.JSON(studentResponse)
}
