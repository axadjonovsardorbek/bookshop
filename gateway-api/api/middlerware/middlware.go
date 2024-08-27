package middlerware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	token "api/api/token"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	Key = "notsecred"
)

func NewAuth(enforce *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allow, err := CheckPermission(ctx.FullPath(),ctx.Request, enforce)
		if err != nil {
			valid, _ := err.(jwt.ValidationError)
			if valid.Errors == jwt.ValidationErrorExpired {
				RequireRefresh(ctx)
			} else {
				RequirePermission(ctx)
			}
		} else if !allow {
			RequirePermission(ctx)
		}
		ctx.Next()
	}

}

func GetRole(r *http.Request) (string, error) {
	var (
		claims jwt.MapClaims
		err    error
	)

	jwtToken := r.Header.Get("Authorization")
	if jwtToken == "" {
		return "unauthorized", nil
	} else if strings.Contains(jwtToken, "Basic") {
		return "unauthorized", nil
	}
	claims, err = token.ExtractClaim(jwtToken)
	log.Printf("Claims--------------------------------: %+v\n", claims)

	if err != nil {
		log.Println("Error while extracting claims: ", err)
		return "unauthorized", err
	}

	return claims["role"].(string), nil
}

func GetUserIdAndUsername(r *http.Request) (string, string, error) {
	var (
		claims jwt.MapClaims
		err    error
	)

	jwtToken := r.Header.Get("Authorization")
	if jwtToken == "" {
		return "unauthorized", "", nil
	} else if strings.Contains(jwtToken, "Basic") {
		return "unauthorized", "", nil
	}
	claims, err = token.ExtractClaim(jwtToken)
	log.Printf("Claims--------------------------------: %+v\n", claims)

	if err != nil {
		log.Println("Error while extracting claims: ", err)
		return "unauthorized", "", err
	}

	return claims["userid"].(string), claims["username"].(string), nil
}

func CheckPermission(path string, r *http.Request, enforcer *casbin.Enforcer) (bool, error) {
	role, err := GetRole(r)
	fmt.Println(role)
	if err != nil {
		log.Println("Error while getting role from token: ", err)
		return false, err
	}
	method := r.Method

	allowed, err := enforcer.Enforce(role, path, method)
	if err != nil {
		log.Println("Error while comparing role from csv list: ", err)
		return false, err
	}

	return allowed, nil
}

func InvalidToken(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "Invalid token !!!",
	})
}

func RequirePermission(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "Permission denied",
	})
}

	func RequireRefresh(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Access token expired",
		})
	}