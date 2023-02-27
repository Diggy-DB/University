package model

type Teacher struct {
	TeacherName    string `json:"teacher_name" validate:"required"`
	TeacherId      string `json:"teacher_id" validate:"required"`
	DepartmentName string `json:"department_name" validate:"required"`
	Position       string `json:"position" validate:"required"`
	Age            int64  `json:"age" validate:"required"`
	Gender         string `json:"gender" validate:"required"`
	PhoneNumber    string `json:"phone_number" validate:"required"`
}
