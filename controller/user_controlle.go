package controller

import (
	config "BlockApp/conf"
	"BlockApp/db"
	"BlockApp/model"
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

// 生成JWT Token
func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // 2小时过期
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.Config.System.GetJwtSecret())
}

// UserLogin 用户登录
// @Summary 用户登录
// @Description 通过用户名和密码进行登录，并返回 JWT Token
// @Tags User
// @Accept json
// @Produce json
// @Param request body Creds true "用户登录信息"
// @Success 200 {object} map[string]string "返回 Token"
// @Failure 400 {object} map[string]string "请求错误"
// @Failure 404 {object} map[string]string "用户不存在"
// @Failure 500 {object} map[string]string "服务器错误"
// @Router /v1/login [post]
func (uc *UserController) UserLogin(ctx *gin.Context) {
	var creds Creds
	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	var user = model.User{Name: creds.Username}
	// 校验用户名密码
	result := db.PgsqlDB.First(&user, "name = ?", creds.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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
