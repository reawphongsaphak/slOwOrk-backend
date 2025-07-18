package services

import (
	"context"
	"fmt"
	"time"

	"github.com/reawphongsaphak/slOwOrk-Backend/config"
	"github.com/reawphongsaphak/slOwOrk-Backend/database"
	"github.com/reawphongsaphak/slOwOrk-Backend/models"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(user models.LoginUser) (string, error) {
	cfg := config.LoadConfig()
	client := database.ConnectDB()
	coll := client.Database(cfg.DB_NAME).Collection("users")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var foundUser models.User
	err := coll.FindOne(ctx, bson.M{"email":user.Email}).Decode(&foundUser)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return "", fmt.Errorf("incorrect password")
	}
	
	claims := jwt.MapClaims {
		"email" : foundUser.Email,
		"role"	: foundUser.Detail.Role,
		"exp"	: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return signedString, nil

}