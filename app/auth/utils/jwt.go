package utils

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/auth/conf"
	"github.com/golang-jwt/jwt/v5"
)

func String2Hex(s string) ([]byte, error) {
	if s == "" {
		return nil, errors.New("empty secret")
	}
	privateKeyBytes, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return privateKeyBytes, nil
}

func SignToken(userUuid string) (string, string, error) {
	tokenExpireAt := time.Now().Add(time.Duration(conf.GetConf().JWT.TokenExpire) * time.Minute)
	refreshTokenExpireAt := time.Now().Add(time.Duration(conf.GetConf().JWT.RefreshTokenExpire) * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA,
		jwt.MapClaims{
			"uuid": userUuid,
			"exp":  tokenExpireAt.Unix(),
		},
	)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA,
		jwt.MapClaims{
			"uuid": userUuid,
			"exp":  refreshTokenExpireAt.Unix(),
		},
	)

	privateKeyBytes, err := String2Hex(conf.GetConf().JWT.PrivateSecret)
	if err != nil {
		return "", "", err
	}

	privateKey := ed25519.PrivateKey(privateKeyBytes)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString(privateKey)
	if err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil

}

func VerifyToken(tokenString string) (string, error) {

	publicKeyBytes, err := String2Hex(conf.GetConf().JWT.PublicSecret)
	if err != nil {
		return "", err
	}

	publicKey := ed25519.PublicKey(publicKeyBytes)

	// 解析和验证 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的是 EdDSA 签名方法
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})

	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}
	expireAt := claims["exp"].(float64)
	if time.Now().Unix() > int64(expireAt) {
		klog.Errorf("token expired")
		return "", errors.New("token expired")
	}

	return claims["uuid"].(string), nil
}
