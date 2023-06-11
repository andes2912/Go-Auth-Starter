package middleware

import (
	"fmt"
	"net/http"
	"os"
	"starterkit-go-auth/initializer"
	"starterkit-go-auth/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func MiddlewareAuth(c *gin.Context)  {
	// Ambil cookie jika tersedia
	tokenString, err :=c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Cek expired token
		if float64(time.Now().Unix()) > claims["expired"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// cari user berdasarkan token
		var user models.User
		initializer.DB.First(&user, claims["sub "])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}