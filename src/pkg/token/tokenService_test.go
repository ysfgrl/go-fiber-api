package token

import (
	"testing"
)

func TestJwtSecurity_CreateToken(t *testing.T) {
	t.Log("Testting ..")
	payload := UserPayload{
		UserID:   "myID",
		UserName: "myUsername",
		Role:     "MyRole",
	}

	token, err := CreateToken(payload)
	if err != nil {
		err.PrintConsole()
		t.Fatalf("%v", err)
	}
	t.Logf(token)
	parsed, err := VerifyToken(token)
	if err != nil {
		err.PrintConsole()
		t.Fatalf("%v", err)
	}
	t.Logf(parsed.UserName)

}
