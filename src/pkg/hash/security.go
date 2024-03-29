package hash

import (
	"github.com/ysfgrl/go-fiber-api/src/models"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, *models.Error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", models.GetError(err)
	}
	return string(hashed), nil
}

func VerifyPassword(hashed string, password string) *models.Error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return models.GetError(err)
	}
	return nil
}

//
//func NewToken(user user_repository.User, JwtSecretKey string) (string, *models2.Error) {
//	claims := jwt.MapClaims{
//		"id":       user.ID.Hex(),
//		"userName": user.UserName,
//		"role":     user.Role,
//		"iss":      user.ID.Hex(),
//		"iat":      time.Now().Unix(),
//		"exp":      time.Now().Add(time.Minute * 5).Unix(),
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	jwt, err := token.SignedString([]byte(JwtSecretKey))
//	if err != nil {
//		return "", response.GetError(err)
//	}
//	return jwt, nil
//}
//
//func validateSignedMethod(token *jwt.Token) (interface{}, error) {
//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//	}
//	return "JwtSecretKey", nil
//}
//
//func SignedUser(c *fiber.Ctx) models2.SignedUser {
//	user := c.Locals("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//	return models2.SignedUser{
//		Id:       claims["id"].(string),
//		UserName: claims["userName"].(string),
//		Role:     claims["role"].(string),
//	}
//}
