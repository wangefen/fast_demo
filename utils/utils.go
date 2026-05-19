package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 对明文密码进行 bcrypt 加密，返回哈希值和错误
func HashPassword(pwd string) (string, error) {
	// GenerateFromPassword 接收字节数组，第二个参数是加密成本（越大越安全也越慢）
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

// GenerateJWT 根据用户名生成 JWT 令牌，用于登录后的身份验证
func GenerateJWT(username string) (string, error) {
	// 创建 JWT：指定签名算法（HS256）和载荷数据（用户名 + 72 小时过期时间）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,                              // 自定义字段：用户名
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 过期时间：72 小时后
	})

	// 用密钥对 JWT 进行签名，生成最终的 token 字符串
	signedToken, err := token.SignedString([]byte("secret"))
	return "Bearer" + signedToken, err
}
