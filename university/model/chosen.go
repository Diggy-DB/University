package model

type Chosen struct {
	StudentId string `json:"student_id" validate:"required"`
	CourseNo  string `json:"course_no" validate:"required"`
	TeacherId string `json:"teacher_id" validate:"required"`
	Grade     int64  `json:"grade" validate:"required"`
}
