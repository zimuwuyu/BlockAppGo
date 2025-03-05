package middleware

import (
	config "BlockApp/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var secretKey = []byte(config.Config.Jwt.Key)

type CustomClaims struct {
	UserID    string `json:"user_id"`
	IsRefresh bool   `json:"is_refresh"`
	jwt.RegisteredClaims
}

// 解析 JWT 并验证
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("无效的签名算法: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 解析成功后，校验是否过期
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, fmt.Errorf("token 已过期")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("无效的 token")
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供 JWT 令牌"})
			c.Abort()
			return
		}

		claims, err := ParseToken(strings.TrimPrefix(tokenString, "Bearer "))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 JWT 令牌"})
			c.Abort()
			return
		}

		if claims.IsRefresh {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "不能使用 Refresh Token 访问接口"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

// 生成 JWT 令牌（包含 Access Token 和 Refresh Token）
func GenerateTokens(userID string) (accessToken string, refreshToken string, err error) {
	// Access Token 30分钟有效
	accessClaims := CustomClaims{
		UserID:    userID,
		IsRefresh: false, // 这是访问令牌
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken, err = generateJWT(accessClaims)
	if err != nil {
		return "", "", err
	}

	// Refresh Token 7天有效
	refreshClaims := CustomClaims{
		UserID:    userID,
		IsRefresh: true, // 这是刷新令牌
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7天过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken, err = generateJWT(refreshClaims)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// 生成 JWT Token
func generateJWT(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
