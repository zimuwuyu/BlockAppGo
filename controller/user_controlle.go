package controller

import (
	config "BlockApp/conf"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type UserController struct {
}

type Creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = map[string]string{
	"admin": "password123",
}

// 生成JWT Token
func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // 2小时过期
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.Config.System.GetJwtSecret())
}

func (user *UserController) UserLogin(ctx *gin.Context) {
	var creds Creds
	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 校验用户名密码
	if pwd, exists := users[creds.Username]; !exists || pwd != creds.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 生成Token
	token, err := generateToken(creds.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func protectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected area!"})
}
