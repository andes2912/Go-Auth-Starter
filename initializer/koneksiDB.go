package initializer

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func KoneksiDB()  {
	var err error

	dsn := os.Getenv("DATABASE")
  DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Koneksi ke database gagal!")
	}
}