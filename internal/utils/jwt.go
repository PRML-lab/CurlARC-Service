package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/api/idtoken"
)

type Claims struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

func VerifyGoogleIDToken(ctx context.Context, idToken string) (*idtoken.Payload, error) {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	if googleClientID == "" {
		return nil, errors.New("GOOGLE_CLIENT_ID is not set")
	}
	payload, err := idtoken.Validate(ctx, idToken, googleClientID)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func GenerateBackendAccessToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"uid": userId,
		"exp": jwt.NewNumericDate(jwt.TimeFunc().Add(72 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("BACKEND_ACCESS_TOKEN_SECRET")
	if secret == "" {
		return "", errors.New("BACKEND_ACCESS_TOKEN_SECRET is not set")
	}

	return token.SignedString([]byte(secret))
}

// JWTを解析してその署名を検証します。
func ParseBackendAccessToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	secret := os.Getenv("BACKEND_ACCESS_TOKEN_SECRET")

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// 署名アルゴリズムが一致しているか確認
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method: " + token.Header["alg"].(string))
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return claims, nil
}
