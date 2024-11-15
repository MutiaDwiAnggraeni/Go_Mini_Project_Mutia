package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Name   string `json:"name"`
	UserID int    `json:"userID"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginHandler(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	oldPassword := user.Password

	result := config.DB.First(&user, "email = ?", user.Email)
	if result.Error != nil {
		if user.ID == 0 {
			return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Email salah", nil})
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Database error", nil})
	}
	match := CheckPasswordHash(oldPassword, user.Password)
	if !match {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "password salah", nil})
	}
	token, err := GenerateJWT(int(user.ID), user.Nama)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{false, "Gagal generate token", nil})
	}
	return c.JSON(200, models.BaseResponse{true, "Sukses", map[string]string{
		"token": token,
	}})
}

func RegisterHandler(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	hash, _ := HashPassword(user.Password)
	user.Password = hash
	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(500, models.BaseResponse{false, "Gagal register", nil})
	}
	return c.JSON(200, models.BaseResponse{true, "Sukses", user})
}

func GenerateJWT(userID int, name string) (string, error) {
	claims := &jwtCustomClaims{
		name,
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("alta"))
	if err != nil {
		return "", err
	}
	return t, nil
}
