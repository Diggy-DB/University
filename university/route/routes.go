package route

import (
	"university/handler"
	"university/middleware"

	"github.com/gofiber/fiber/v2"
)

func AccessRoute(app *fiber.App) {
	// Welcome
	app.Get("/", handler.Welcome)

	// Token
	app.Get("/token", middleware.LoginHandler)

	// Account
	app.Get("/users", middleware.CheckToken, middleware.CheckManager, handler.GetAllAccounts)
	app.Post("/user", middleware.CheckToken, middleware.CheckManager, handler.CreateAccount)

	// Student
	app.Get("/student", middleware.CheckToken, middleware.CheckStudent, handler.GetStudentById)
	app.Post("/student", middleware.CheckToken, middleware.CheckManager, handler.AddStudents)

	// Teacher
	app.Get("/teacher", middleware.CheckToken, middleware.CheckTeacher, handler.GetStudentById)
	app.Post("/teacher", middleware.CheckToken, middleware.CheckManager, handler.AddStudents)

	// Department
	app.Get("/department", middleware.CheckToken, handler.ShowAllDepartment)
	app.Post("/department", middleware.CheckToken, handler.AddDepartment)

	// Course
	app.Get("/course", middleware.CheckToken, handler.ShowAllCourse)
	app.Post("/course", middleware.CheckToken, handler.AddCourse)

	// Chosen
	app.Get("/chosencourse", middleware.CheckToken, middleware.CheckStudent, handler.ShowAllChosenCourse)
	app.Post("/chosencourse", middleware.CheckToken, middleware.CheckManager, handler.AddChosenCourse)
}
