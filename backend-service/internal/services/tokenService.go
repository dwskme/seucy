package services

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	SecretKey    []byte
	TokenExpires time.Duration
}

type UserClaims struct {
	Identifier string `json:"identifier"`
	jwt.RegisteredClaims
}

func NewTokenService(secretKey []byte, tokenExpires time.Duration) *TokenService {
	return &TokenService{
		SecretKey:    secretKey,
		TokenExpires: tokenExpires,
	}
}

func (s *TokenService) GetTokenFromBearerHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("invalid authorization header format")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	return token, nil
}

func (s *TokenService) GenerateToken(identifier string) (string, error) {
	expires := s.TokenExpires
	claims := UserClaims{
		identifier,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "system",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *TokenService) ValidateToken(tokenStr string) (*UserClaims, bool, error) {
	key := s.SecretKey
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, false, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, true, nil
	}

	return nil, false, errors.New("token is invalid")
}
