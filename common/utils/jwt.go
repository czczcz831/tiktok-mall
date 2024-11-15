package utils

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
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

func SignToken(userUuid string, privateKeyString string, tokenExpire int, refreshTokenExpire int) (string, string, error) {
	tokenExpireAt := time.Now().Add(time.Duration(tokenExpire) * time.Hour)
	refreshTokenExpireAt := time.Now().Add(time.Duration(refreshTokenExpire) * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA,
		jwt.MapClaims{
			"uuid": userUuid,
			"exp":  tokenExpireAt.Unix(),
			"rt":   false,
		},
	)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA,
		jwt.MapClaims{
			"uuid": userUuid,
			"exp":  refreshTokenExpireAt.Unix(),
			"rt":   true,
		},
	)

	privateKeyBytes, err := String2Hex(privateKeyString)
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

func VerifyToken(tokenString string, publicKeyHexString string) (string, bool, error) {

	publicKeyBytes, err := String2Hex(publicKeyHexString)
	if err != nil {
		return "", false, err
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
		return "", false, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false, err
	}
	expireAt := claims["exp"].(float64)
	if time.Now().Unix() > int64(expireAt) {
		klog.Errorf("token expired")
		return "", false, errors.New("token expired")
	}

	isRefreshToken := claims["rt"].(bool)

	return claims["uuid"].(string), isRefreshToken, nil
}
