package auth

import (
	"demo4/internal/config"
	"demo4/internal/model"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(user model.User) (string, error) {
	tokenExpiry := config.GetConfig().Auth.TokenExpiry
	expirationTime := time.Now().Add(time.Duration(tokenExpiry) * time.Second)
	claims := &Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetConfig().Auth.JwtSecret)) //用密钥生成token
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	log.Printf("[JWT] Using secret2: %s ", tokenString)
	log.Printf("[JWT] Using secret: %s ", (config.GetConfig().Auth.JwtSecret))
	//ParseWithClaims 验证的是 token 的签名合法性，而不是比较 token 字符串是否与某个 “配置文件中的 token” 完全一致。
	//如果你的代码中，Keyfunc 返回的是配置文件中的 正确密钥，那么任何用此密钥签名的 token 都会验证通过，无论 token 字符串本身是否与某个 “预设值” 一致。
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().Auth.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
