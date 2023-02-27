package handler

import (
	"university/auth"
	"university/database"
	"university/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllAccounts(ctx *fiber.Ctx) error {
	var accounts []auth.Account
	result := database.DB.Debug().Find(&accounts)
	if result.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to get all accounts",
		})
	}
	return ctx.JSON(accounts)
}
func CreateAccount(ctx *fiber.Ctx) error {
	user := new(auth.Account)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Parse JSON!",
		})
	}
	account := auth.Account{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}
	account.Password = utils.HashPassword(user.Password)
	errCreateUser := database.DB.Create(&account).Error
	if errCreateUser != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Create Account!",
		})
	}
	return ctx.JSON(account)
}
