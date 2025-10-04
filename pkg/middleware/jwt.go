package middleware

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Ansalps/BrotoStack/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func OtpAuthMiddleware(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		utils.ErrorResponse(c, 401, "missing authorization", nil)
		return
	}
	tokenString = tokenString[len("Bearer "):]
	secret_key := os.Getenv("secret_key")
	if secret_key == "" {
		utils.ErrorResponse(c, 500, "missing secret key", nil)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing algorithm: %v", token.Header["alg"])
		}
		return []byte(secret_key), nil
	})
	if err != nil {
		utils.ErrorResponse(c, 401, "error in parsing token", err)
		c.Abort()
		return
	}
	if !token.Valid {
		utils.ErrorResponse(c, 401, "invalid token", nil)
		c.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.ErrorResponse(c, 401, "error in claims", nil)
		c.Abort()
		return
	}
	if claims["iss"] != "my-auth-server" && claims["otp_verfied"] != true {
		utils.ErrorResponse(c, 401, "jwt claims are different", nil)
		c.Abort()
		return
	}
	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		utils.ErrorResponse(c, 401, "jwt expired", nil)
		c.Abort()
		return
	}
	v, ok1 := claims["email"].(string)
	if !ok1 {
		utils.ErrorResponse(c, 400, "failed to assert email", nil)
		c.Abort()
		return
	}
	c.Set("email", v)
	c.Next()
}
func CreateTokenForResetPassword(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":        email,
			"iss":          "my-auth-server",
			"otp_verified": true,
			"exp":          time.Now().Add(time.Minute * 15).Unix(),
		})
	tokenString, err := token.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func GenerateAccessToken(id string) (string, error) {
	AccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":             id,
			"iss":             "my-auth-server",
			"is_logged_in":    true,
			"is_access_token": true,
			"exp":             time.Now().Add(time.Minute * 5).Unix(),
		})
	AccessTokenString, err := AccessToken.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return "", err
	}
	return AccessTokenString, nil
}

func GenerateRefreshToken(id string) (string, error) {
	RefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":              id,
			"iss":              "my-auth-server",
			"is_logged_in":     true,
			"is_refresh_token": true,
			"exp":              time.Now().Add(time.Hour * 24).Unix(),
		})
	RefreshTokenString, err := RefreshToken.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return "", err
	}
	return RefreshTokenString, nil
}
func UserLoginAuthMiddleware(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		utils.ErrorResponse(c, 401, "missing authorization", nil)
		c.Abort()
	}
	tokenString = tokenString[len("Bearer "):]
	secret_key := os.Getenv("secret_key")
	if secret_key == "" {
		utils.ErrorResponse(c, 500, "missing secret key", nil)
		c.Abort()
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			utils.ErrorResponse(c, 401, "invalid signing algorithm", nil)
			c.Abort()
		}
		return []byte(secret_key), nil
	})
	if err != nil {
		// Check if error is due to token expiration
		if errors.Is(err, jwt.ErrTokenExpired) {
			utils.ErrorResponse(c, 1090, "token has expired", err)
			c.Abort()
		}
		utils.ErrorResponse(c, 401, "error in parsing token", err)
		c.Abort()
	}
	if !token.Valid {
		utils.ErrorResponse(c, 401, "invalid token", nil)
		c.Abort()
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.ErrorResponse(c, 401, "error in claims", nil)
		c.Abort()
	}
	if claims["iss"] != "my-auth-server" || claims["is_logged_in"] != true || claims["is_access_token"] != true {
		utils.ErrorResponse(c, 401, "jwt claims are different", nil)
		c.Abort()
	}
	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		utils.ErrorResponse(c, 401, "jwt expired", nil)
		c.Abort()
	}
	v, ok1 := claims["sub"].(string)
	if !ok1 {
		utils.ErrorResponse(c, 400, "failed to assert email", nil)
	}
	c.Set("id", v)
	c.Next()
}
func AccessRegenerator(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		utils.ErrorResponse(c, 401, "missing authorization", nil)
		return
	}
	tokenString = tokenString[len("Bearer "):]
	secret_key := os.Getenv("secret_key")
	if secret_key == "" {
		utils.ErrorResponse(c, 500, "missing secret key", nil)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			utils.ErrorResponse(c, 401, "invalid signing algorithm", nil)
			c.Abort()
		}
		return []byte(secret_key), nil
	})
	if err != nil {
		// Check if error is due to token expiration
		if errors.Is(err, jwt.ErrTokenExpired) {
			utils.ErrorResponse(c, 1090, "token has expired", err)
			c.Abort()
		}
		utils.ErrorResponse(c, 401, "error in parsing token", err)
		c.Abort()
	}
	if !token.Valid {
		utils.ErrorResponse(c, 401, "invalid token", nil)
		c.Abort()
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.ErrorResponse(c, 401, "error in claims", nil)
		c.Abort()
	}
	if claims["iss"] != "my-auth-server" || claims["is_logged_in"] != true || claims["is_refresh_token"] != true {
		utils.ErrorResponse(c, 401, "jwt claims are different", nil)
		c.Abort()
	}
	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		utils.ErrorResponse(c, 401, "jwt expired", nil)
		c.Abort()
	}
	v, ok1 := claims["sub"].(string)
	if !ok1 {
		utils.ErrorResponse(c, 400, "failed to assert email", nil)
		return
	}
	AccssTokenString, err := GenerateAccessToken(v)
	if err != nil {
		utils.ErrorResponse(c, 500, "failed to generate access token", nil)
		return
	}
	data := map[string]string{
		"access_token": AccssTokenString,
	}
	utils.SuccessResponse(c, 200, "generated access token from refresh token", data)
}
