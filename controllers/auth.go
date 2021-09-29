package controllers

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jhonwick88/pintarshop/config"
	"github.com/jhonwick88/pintarshop/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var createUser CreateUserInput
	if err := c.ShouldBindJSON(&createUser); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	user := models.User{
		Username: html.EscapeString(strings.TrimSpace(createUser.Username)),
		Password: createUser.Password,
		NickName: html.EscapeString(strings.TrimSpace(createUser.NickName)),
		Email:    html.EscapeString(strings.TrimSpace(createUser.Email)),
		Role:     1,
	}
	err := user.Validate("")
	if err != nil {
		config.FailWithMessage(err.Error(), c)
	}
	userCreated, err := user.SaveUser(models.DB)
	if err != nil {
		config.FailWithMessage("Created user failed", c)
		return
	}
	config.OkWithDetailed(userCreated, "Register success", c)
}
func Login(c *gin.Context) {
	var userLoginInput LoginUserInput
	if err := c.ShouldBindJSON(&userLoginInput); err != nil {
		config.FailWithMessage(err.Error(), c)
	}
	user := models.User{
		Email:    userLoginInput.Email,
		Password: userLoginInput.Password,
	}
	err := user.Validate("login")
	if err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	var userx models.User
	if err := models.DB.Where("email = ?", user.Email).First(&userx).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	err = models.VerifyPassword(userx.Password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		config.FailWithMessage(err.Error(), c)
		return
	}
	token, _ := CreateToken(userx.ID)
	esteh := gin.H{"user": userx, "token": token}
	config.OkWithDetailed(esteh, "Login successfully", c)
}
func CreateToken(user_id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(r *gin.Context) string {
	// := r.URL.Query()
	token, _ := r.GetQuery("token")
	if token != "" {
		return token
	}
	bearerToken := r.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(r *gin.Context) (uint, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
