package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zehan12/hotel-reservation/backend/types"
)

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{Id: "1", FirstName: "zehan", LastName: "khan"}
	return c.JSON(u)
}
