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
	fmt.Println("Loading key pair...", privateKeyPath, publicKeyPath)
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

type TokenClaims struct {
	UID uint64 `json:"uid"` // 用户 ID
}

// GenerateToken 生成 RS256 JWT Token
func (kp *KeyPair) GenerateToken(tc TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"uid": tc.UID,
		"exp": time.Now().Add(TokenExpireDuration).Unix(), // 设置过期时间 24 小时
		"iat": time.Now().Unix(),                          // 签发时间
	})
	return token.SignedString(kp.PrivateKey)
}

// ParseToken 验证并解析 RS256 JWT Token
func (kp *KeyPair) ParseToken(tokenString string) (TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return kp.PublicKey, nil
	})
	tc := TokenClaims{}
	if err != nil {
		return tc, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid, ok := claims["uid"].(float64)
		if !ok {
			return tc, errors.New("invalid token claims")
		}
		tc.UID = (uint64)(uid)
	}
	return tc, nil
}

// 加载私钥
func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("invalid private key")
	}
	pk, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pk.(*rsa.PrivateKey), nil
}

// 加载公钥
func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key")
	}
	pk, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pk.(*rsa.PublicKey), nil
}
