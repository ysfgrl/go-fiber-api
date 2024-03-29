package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/pkg/response"
	"github.com/ysfgrl/go-fiber-api/src/pkg/token"
)

func PreRoute(c *fiber.Ctx) error {
	c.Locals("user", token.UserPayload{UserName: "guest", Role: "guest", UserID: "0"})
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
			return response.Unauthorized(c, models.GetError(err))
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			//user := token.SignedUser(ctx)
			//ctx.Locals("user", user)
			return ctx.Next()
		},
	})(ctx)
}

func RoleRequired(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		signedUser := c.Locals("signedUser").(token.UserPayload)
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
		return response.Forbidden(c, models.UserError("Forbidden role"))
	}
}
