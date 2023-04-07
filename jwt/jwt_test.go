package jwt_test

import (
	"fmt"
	"testing"

	"github.com/r1cebucket/gopkg/jwt"
	"github.com/r1cebucket/gopkg/log"
)

func init() {
	log.Setup("debug")
}

type User struct {
	Username string
	Password string
}

func TestToken(t *testing.T) {
	// u := User{
	// 	Username: "testuser",
	// 	Password: "123",
	// }
	token, err := jwt.GenToken("hello")
	if err != nil {
		log.Err(err).Msg("failed to gen token")
		t.Error()
	}
	log.Info().Msg(token)

	payload, err := jwt.ParseToken(token)
	if err != nil {
		log.Err(err).Msg("failed to parse token")
		t.Error()
	}
	log.Info().Msg(fmt.Sprintf("%T", payload.Data))

}
