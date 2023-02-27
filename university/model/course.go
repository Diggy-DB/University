package model

type Course struct {
	CourseName     string `json:"course_name" validate:"required"`
	CourseNo       string `json:"course_no" validate:"required"`
	TeacherId      string `json:"teacher_id" validate:"required"`
	Credit         int64  `json:"credit" validate:"required"`
	DepartmentName string `json:"department_name" validate:"required"`
}
