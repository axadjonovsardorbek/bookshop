package token

import (
	"api/config"
	"log/slog"

	"net/http"

	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/spf13/cast"
)

type JWTHandler struct {
	Sub        string
	Exp        string
	Iat        string
	Role       string
	SigningKey string
	Token      string
}

type Tokens struct {
	AccessToken string

	RefreshToken string
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	cfg := config.Load()
	var (
		token *jwt.Token
		err error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.TOKENKEY), nil
	}

	token, err = jwt.Parse(tokenStr, keyFunc)

	if err != nil {

		return nil, err

	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !(ok && token.Valid) {

		return nil, err

	}

	return claims, nil

}

func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {

	token, err := jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {

		return []byte(jwtHandler.SigningKey), nil

	})

	if err != nil {

		return nil, err

	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !(ok && token.Valid) {

		slog.Error("invalid jwt token")

		return nil, err

	}

	return claims, nil

}

func GetIdFromToken(r *http.Request) (string, int) {

	var softToken string

	token := r.Header.Get("Authorization")

	if token == "" {

		return "unauthorized", http.StatusUnauthorized

	} else if strings.Contains(token, "Bearer") {

		softToken = strings.TrimPrefix(token, "Bearer ")

	} else {

		softToken = token

	}

	claims, err := ExtractClaim(softToken)

	if err != nil {

		return "unauthorized", http.StatusUnauthorized

	}

	return cast.ToString(claims["username"]), 0

}