package database

import (
	"log"
	"university/auth"
	"university/model"
)

func Migrate() {
	err := DB.AutoMigrate(&model.Student{})
	if err != nil {
		log.Fatal("Failed to migrate student!")
	}
	err = DB.AutoMigrate(&model.Course{})
	if err != nil {
		log.Fatal("Failed to migrate Course!")
	}
	err = DB.AutoMigrate(&model.Teacher{})
	if err != nil {
		log.Fatal("Failed to migrate Teacher!")
	}
	err = DB.AutoMigrate(&model.Department{})
	if err != nil {
		log.Fatal("Failed to migrate Department!")
	}
	err = DB.AutoMigrate(&model.Chosen{})
	if err != nil {
		log.Fatal("Failed to migrate Chosen Course!")
	}
	err = DB.AutoMigrate(&auth.Account{})
	if err != nil {
		log.Fatal("Failed to migrate Account!")
	}
}
