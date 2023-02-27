package model

type Student struct {
	StudentName    string `json:"student_name" validate:"required"`
	StudentId      string `json:"student_id" validate:"required"`
	PhoneNo        string `json:"phone_no" validate:"required"`
	Grade          int64  `json:"grade"`
	DepartmentName string `json:"department_name" validate:"required"`
	Degree         string `json:"degree" validate:"required"`
	Address        string `json:"address" validate:"required"`
}
