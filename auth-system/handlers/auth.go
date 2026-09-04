package handlers

import (
	"context"
	"os"

	"auth-system/config"
	"auth-system/models"
	"auth-system/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Register(c *fiber.Ctx) error {
	user := models.User{}

	err := c.BodyParser(&user)

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	if user.Email == "" || user.Password == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "email and password required",
		})
	}

	hashed, err := utils.HashPassword(user.Password)

	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"msg": "could not hash password",
		})
	}

	user.Password = hashed

	_, err = config.UserCollection.InsertOne(
		context.Background(),
		user,
	)

	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"msg": "could not save user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"msg": "saved",
	})
}

func Login(c *fiber.Ctx) error {
	user := models.LoginRequest{}
	storedUser := models.User{}

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "invalid request body",
		})
	}

	find := bson.M{
		"email": user.Email,
	}

	result := config.UserCollection.FindOne(
		context.Background(),
		find,
	)

	err = result.Decode(&storedUser)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"msg": "user not found",
		})
	}

	err = utils.CheckPassword(
		user.Password,
		storedUser.Password,
	)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"msg": "wrong password",
		})
	}

	token, err := utils.GenerateToken(
		storedUser.ID,
		os.Getenv("JWT_SECRET"),
		24,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msg": "could not generate token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"msg":   "successful",
		"token": token,
	})
}