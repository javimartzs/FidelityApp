package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var (
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string

	AdminUser string
	AdminPass string
	AdminRole string
)

func InitConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")

	AdminUser = os.Getenv("ADMIN_USER")
	AdminPass = os.Getenv("ADMIN_PASS")
	AdminRole = os.Getenv("ADMIN_ROLE")
}

// Funcion para Hashear las contraseñas de los usuarios
func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password")
	}

	return string(hashedPassword), nil
}
