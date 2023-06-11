package controllers

import (
	"net/http"
	"os"
	"starterkit-go-auth/initializer"
	"starterkit-go-auth/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Register
func Register(c *gin.Context)  {
	// Ambil data yg dikirim dari body
	var body struct {
		Nama string
		Phone string
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Gagal mengambil data!",
		})
		return
	}

	// Simpan password dengan hash
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 20)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Gagal menyimpan password dengan hash!",
		})
		return
	}

	// Simpan data yang dikirim ke database
	user := models.User{
		Nama: body.Nama,
		Phone: body.Phone,
		Email: body.Email,
		Password: string(hash),
	}
	result := initializer.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Gagal menyimpan data user!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// LOGIN
func Login(c *gin.Context)  {
	// Ambil data yg dikirim dari body
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Gagal mengambil data!",
		})
		return
	}

	// Kirim data dari body
	var user models.User
	initializer.DB.First(&user, "email = ?", body.Email)

	if user.ID ==0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Email atau Password salah!",
		})
		return
	}

	// Cek apakah password yg di input dengan yg di BD sama atau tidak
	err :=  bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Email atau Password salah!",
		})
		return
	}

	// Generate token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"expired": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Token gagal dibuat!",
		})
		return
	}

	// Tampilkan token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30,"","",false,true)
	c.JSON(http.StatusOK, gin.H{})
}

func User(c *gin.Context)  {
	user, _:= c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}