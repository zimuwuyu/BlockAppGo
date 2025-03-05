package middleware

import (
	config "BlockApp/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var secretKey = []byte(config.Config.Jwt.Key)

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// 解析 JWT 并验证
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保使用 HMAC 作为签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("无效的签名算法: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 断言 claims 类型
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 显式校验 token 是否过期
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, fmt.Errorf("token 已过期")
		}
		return claims, nil
	}
	return nil, fmt.Errorf("无效的 token")
}

// Gin 中间件：验证 JWT
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头部
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供 JWT 令牌"})
			c.Abort()
			return
		}

		// 解析 token
		claims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 JWT 令牌"})
			c.Abort()
			return
		}

		// 存储用户信息到上下文
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}

// 生成 JWT 令牌
func GenerateToken(userID string) (string, error) {
	fmt.Println(config.Config.Jwt)
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Config.Jwt.TimeOut))), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                               // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                               // 立即生效
		},
	}

	// 生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
