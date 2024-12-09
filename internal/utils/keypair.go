package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenExpireDuration = time.Hour * 24
)

// KeyPair 封装了 RSA 密钥对相关操作
type KeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// LoadKeyPair 从文件加载 RSA 密钥对
func LoadKeyPair(privateKeyPath, publicKeyPath string) (*KeyPair, error) {
	privateKey, err := LoadPrivateKey(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	publicKey, err := LoadPublicKey(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load public key: %w", err)
	}

	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

// GenerateToken 生成 RS256 JWT Token
func (kp *KeyPair) GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(TokenExpireDuration).Unix(), // 设置过期时间 24 小时
		"iat": time.Now().Unix(),                          // 签发时间
	})
	return token.SignedString(kp.PrivateKey)
}

// ParseToken 验证并解析 RS256 JWT Token
func (kp *KeyPair) ParseToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return kp.PublicKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(float64); ok {
			return uint(id), nil
		}
	}

	return 0, errors.New("invalid token claims")
}

// 加载私钥
func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// 加载公钥
func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("invalid public key")
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}
