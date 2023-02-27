package model

type Department struct {
	DepartmentName string `json:"department_name" validate:"required"`
	DepartmentNo   string `json:"department_no" validate:"required"`
	BuildingNo     string `json:"building_no" validate:"required"`
	PhoneNo        string `json:"phone_no" validate:"required"`
	School         string `json:"school" validate:"required"`
}
