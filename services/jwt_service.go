package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtSercice struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtSercice {
	return &jwtSercice{
		secretKey: "secret-key",
		issure:    "book-api",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtSercice) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer: s.issure,
			IssuedAt: time.Now().Unix(),
		},
	} 

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtSercice) ValidateToken (token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("inavalid token %v", token)
		}
		return []byte(s.secretKey), nil
	})

	return err == nil
}
