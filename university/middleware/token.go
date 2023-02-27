package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
	"university/auth"
	"university/database"
	"university/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(ctx *fiber.Ctx) error {
	userRequest := new(request.Login)
	err := ctx.BodyParser(userRequest)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to Body Parse!",
		})
	}
	validate := validator.New()
	errValidate := validate.Struct(userRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to validate",
		})
	}
	var user auth.Account
	err = database.DB.First(&user, "username = ?", user.Username).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to get account!",
		})
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Not the Correct Password",
		})
	}
	bytes := make([]byte, 1000)
	_, err = rand.Read(bytes)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to make token string",
		})
	}
	token := hex.EncodeToString(bytes)
	user.Token = token
	user.Expiration = time.Now().Add(time.Minute * 5)
	return ctx.JSON(fmt.Sprintf("Bearer %s", token))
}
func CheckToken(ctx *fiber.Ctx) error {
	authorization := ctx.GetReqHeaders()
	idToken := strings.TrimSpace(strings.Replace(authorization["token"], "Bearer ", "", 1))
	var user auth.Account
	if u := database.DB.Find(&user, "token = ?", idToken).Error; u != nil {
		return ctx.Status(400).JSON("Token Wrong!")
	}
	if user.Expiration.Before(time.Now()) {
		return ctx.Next()
	}
	return ctx.Status(400).JSON("Expired Token!")
}
func CheckTeacher(ctx *fiber.Ctx) error {
	authorization := ctx.GetReqHeaders()
	idToken := strings.TrimSpace(strings.Replace(authorization["token"], "Bearer ", "", 1))
	var user auth.Account
	if u := database.DB.Find(&user, "token = ?", idToken).Error; u != nil {
		return ctx.Status(400).JSON("Token Wrong!")
	}
	if user.Role == "Teacher" {
		return ctx.Next()
	}
	return ctx.Status(400).JSON("It isn't teacher!")
}
func CheckStudent(ctx *fiber.Ctx) error {
	authorization := ctx.GetReqHeaders()
	idToken := strings.TrimSpace(strings.Replace(authorization["token"], "Bearer ", "", 1))
	var user auth.Account
	if u := database.DB.Find(&user, "token = ?", idToken).Error; u != nil {
		return ctx.Status(400).JSON("Token Wrong!")
	}
	if user.Role == "Student" {
		return ctx.Next()
	}
	return ctx.Status(400).JSON("It isn't teacher!")
}
func CheckManager(ctx *fiber.Ctx) error {
	authorization := ctx.GetReqHeaders()
	idToken := strings.TrimSpace(strings.Replace(authorization["token"], "Bearer ", "", 1))
	var user auth.Account
	if u := database.DB.Find(&user, "token = ?", idToken).Error; u != nil {
		return ctx.Status(400).JSON("Token Wrong!")
	}
	if user.Role == "Manager" {
		return ctx.Next()
	}
	return ctx.Status(400).JSON("It isn't Manager!")
}
