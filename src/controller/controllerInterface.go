package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Add(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	SetFile(c *fiber.Ctx) error
}
