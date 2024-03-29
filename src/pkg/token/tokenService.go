package token

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ysfgrl/go-fiber-api/src/config"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"time"
)

func CreateToken(payload UserPayload) (string, *models.Error) {
	claims := tokenPayload{
		UserPayload: payload,
		ExpiredAt:   time.Now().Add(config.AppConf.Token.Expire),
		IssuedAt:    time.Now(),
		NotBefore:   time.Now(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(config.AppConf.Token.PrivateKey))
	if err != nil {
		return "", models.GetError(err)
	}
	return jwtToken, nil
}
func VerifyToken(token string) (*UserPayload, *models.Error) {
	parsed, err := jwt.ParseWithClaims(token, &tokenPayload{}, parse)
	if err != nil {
		return nil, models.GetError(err)
	}

	if claims, ok := parsed.Claims.(*tokenPayload); ok {
		return &claims.UserPayload, nil
	}
	return nil, models.GetError(err)
}

func parse(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(config.AppConf.Token.PrivateKey), nil
}

func SignedUser(c *fiber.Ctx) UserPayload {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(tokenPayload)
	return claims.UserPayload
}
