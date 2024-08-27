package token

import (
	"fmt"
	"log"
	"strings"
	"time"

	pb "api/genproto/auth"

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
	AccessToken  string
	RefreshToken string
}

var tokenKey = "my_secret_key"

func GenerateJWTToken(user *pb.Users) *Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["email"] = user.Email
	accessClaims["role"] = "user"
	accessClaims["id"] = user.UserId
	accessClaims["phone_number"] = user.PhoneNumber
	accessClaims["iat"] = time.Now().Unix()
	accessClaims["exp"] = time.Now().Add(48 * time.Hour).Unix() // Expires in 48 hours
	access, err := accessToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("Error while generating access token: ", err)
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["email"] = user.Email
	refreshClaims["role"] = "user"
	refreshClaims["id"] = user.UserId
	refreshClaims["phone_number"] = user.PhoneNumber
	refreshClaims["iat"] = time.Now().Unix()
	refreshClaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("Error while generating refresh token: ", err)
	}

	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, fmt.Errorf("invalid token or claims")
	}

	return claims, nil
}

func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtHandler.SigningKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, fmt.Errorf("invalid token or claims")
	}

	return claims, nil
}

func GetIdFromToken(tokenStr string) (string, error) {
	// Check if token is missing or incorrectly formatted
	if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
		return "", fmt.Errorf("missing or malformed JWT")
	}

	// Remove the "Bearer " prefix
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	// Parse the token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("error parsing token: %w", err)
	}

	// Validate the token and extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token or claims")
	}

	// Log the extracted claims for debugging
	log.Printf("Extracted Claims: %v", claims)

	// Extract and return the user ID
	userId := cast.ToString(claims["id"])
	if userId == "" {
		return "", fmt.Errorf("user ID not found in token")
	}

	return userId, nil
}
