package handler

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/grading-system-golang/internal/models"
	"log"
	"net/http"
	"time"
)

const (
	signingKey = "asdjvaasdf1123iVDFoasdv"
	salt       = "gjdbsjkgdfg134kjdsfgbkj"
	tokenTTL   = 12 * time.Hour
)

type Claims struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
	jwt.RegisteredClaims
}

type SignInForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *fiber.Ctx) error {

	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	user.PasswordHash = generateHashPasswordHash(user.PasswordHash)

	userID, err := h.Service.AddUser(user)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error registering user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"user_id": userID})
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var form SignInForm

	err := c.BodyParser(&form)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := h.Service.GetUserByUsername(form.Username)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	if user.PasswordHash != generateHashPasswordHash(form.Password) {
		log.Println(errors.New("invalid credentials, wrong password"))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	token, err := GenerateToken(user)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error generating token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"access_token": token})
}

func GenerateToken(user models.User) (string, error) {
	expTime := time.Now().Add(tokenTTL)
	claims := &Claims{
		UserID: user.ID,
		RoleID: user.RoleID,
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

func ParseToken(accessToken string) (*Claims, error) {
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
