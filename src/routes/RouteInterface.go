package routes

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/src/controller"
	"go-fiber-api/src/models"
	"go-fiber-api/src/utils/response"
	"go-fiber-api/src/utils/security"
)

type Route interface {
	RegisterRoutes(app *fiber.App)
	GetController() *controller.Controller
}

func PreRoute(c *fiber.Ctx) error {
	c.Locals("signedUser", models.SignedUser{UserName: "guest", Role: "guest", Id: "0"})
	return c.Next()
}

func AuthRequired(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte("secret"),
		},
		TokenLookup: "header:Authorization",
		ContextKey:  "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return response.Unauthorized(c, response.GetError(err))
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := security.SignedUser(ctx)
			ctx.Locals("signedUser", user)
			return ctx.Next()
		},
	})(ctx)
}

func RoleRequired(roles []string) fiber.Handler {

	return func(c *fiber.Ctx) error {
		signedUser := c.Locals("signedUser").(models.SignedUser)
		if signedUser.Role == "guest" {
			return c.Next()
		}
		for _, role := range roles {
			if role == signedUser.Role {
				return c.Next()
			} else if role == "all" {
				return c.Next()
			}
		}
		return response.Forbidden(c, response.UserError("Forbidden role"))
	}
}
