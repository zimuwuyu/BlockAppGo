package controller

import (
	"BlockApp/db"
	"BlockApp/middleware"
	"BlockApp/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserController struct {
}

type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"123456"`
}

// LoginResponse 登录响应结构体
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// UserLogin 用户登录
// @Summary 用户登录
// @Description 通过用户名和密码进行登录，并返回 JWT Token
// @Tags User
// @Accept json
// @Produce json
// @Param request body LoginRequest true "用户登录信息"
// @Success 200 {object} map[string]string "返回 Token"
// @Failure 400 {object} map[string]string "请求错误"
// @Failure 404 {object} map[string]string "用户不存在"
// @Failure 500 {object} map[string]string "服务器错误"
// @Router /v1/login [post]
func (uc *UserController) UserLogin(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}
	var user = model.User{Name: req.Username}
	// 校验用户名密码
	result := db.PgsqlDB.First(&user, "name = ?", req.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// 使用bcrypt库对密码进行校验
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		return
	}
	// 生成Token
	accessToken, refreshToken, err := middleware.GenerateTokens(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

// UserRegister 用户注册
// @Summary 用户注册(暂时不开放)
// @Description 通过用户名和密码注册新用户(暂时不开放)
// @Tags User
// @Accept json
// @Produce json
// @Param request body LoginRequest true "用户注册信息"
// @Success 200 {object} map[string]string "注册成功"
// @Failure 400 {object} map[string]string "请求错误"
// @Failure 500 {object} map[string]string "服务器错误"
// @Router /v1/register [post]
func (uc *UserController) UserRegister(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	var user = model.User{Name: req.Username}
	// 校验用户名密码
	result := db.PgsqlDB.First(&user, "name = ?", req.Username)
	if result.Error == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User is exited"})
		return
	}
	_password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate password"})
		return
	}
	user.PassWord = string(_password)
	user.Role = "USER"
	if err := db.PgsqlDB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

}

func (uc *UserController) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

// User 接口
// @Summary 刷新token
// @Description 刷新token
// @Tags User
// @Produce json
// @Param request body RefreshRequest true "刷新token"
// @Success 200 {object} map[string]string "返回 Token"
// @Failure 400 {object} map[string]string "请求错误"
// @Failure 401 {object} map[string]string "无效的 Refresh Token"
// @Failure 500 {object} map[string]string "服务器错误"
// @Router /v1/refreshToken [post]
func (uc *UserController) RefreshToken(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求格式错误"})
		return
	}

	// 解析 Refresh Token
	claims, err := middleware.ParseToken(req.RefreshToken)
	if err != nil || !claims.IsRefresh {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 Refresh Token"})
		return
	}

	// 生成新的 Access Token 和 Refresh Token
	newAccessToken, newRefreshToken, err := middleware.GenerateTokens(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成新令牌"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
