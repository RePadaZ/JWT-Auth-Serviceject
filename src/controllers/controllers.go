package controllers

import (
	config "JWT-Auth-Serviceject/src/confing"
	"JWT-Auth-Serviceject/src/database"
	"JWT-Auth-Serviceject/src/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Хэшируем пароль перед сохранением
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Проверяем пароль
func checkPassword(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

// Генерируем JWT-токен
func generateToken(email string) (string, error) {

	cfg := config.AppConfig.JWT // Достаём настройки JWT

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * time.Duration(cfg.Expire)).Unix(),
	})

	return token.SignedString(cfg.Secret)
}

// Register Регистрация пользователя
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	hashedPassword, _ := hashPassword(user.Password)
	user.Password = hashedPassword

	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь создан"})
}

// Login Вход пользователя
func Login(c *gin.Context) {
	var user models.User
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	database.DB.Where("email = ?", input.Email).First(&user)
	if user.ID == 0 || !checkPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	token, _ := generateToken(user.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
