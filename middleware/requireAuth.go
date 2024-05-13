package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/backEnGO/initializers"
	"github.com/backEnGO/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("in middleware")
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("Utilisateur sans token")
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token not found",
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED signing method: %v", token.Header["mauvais token"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			// on peu ajouter la supression en base du token si il est expiré ici
		}
		var User models.User
		initializers.DB.First(&User, claims["userId"])

		if User.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", User)
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
