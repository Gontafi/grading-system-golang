package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/grading-system-golang/internal/models"
	"log"
	"time"
)

const (
	signingKey = "asdjvaasdf1123iVDFoasdv"
	salt       = "gjdbsjkgdfg134kjdsfgbkj"
	tokenTTL   = 12 * time.Hour
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *ServiceV1) RegisterUser(user models.User) (int, error) {
	user.PasswordHash = generateHashPasswordHash(user.PasswordHash)

	id, err := s.AddUser(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceV1) GetTokenFromUser(username string, password string) (string, error) {

	user, err := s.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if user.PasswordHash != generateHashPasswordHash(password) {
		log.Println(errors.New("invalid credentials, wrong password"))
		return "", err
	}

	token, err := GenerateToken(user)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}
func GenerateToken(user models.User) (string, error) {
	expTime := time.Now().Add(tokenTTL)
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateHashPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *ServiceV1) ParseToken(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
}
