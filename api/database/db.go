package database

import (
	"api/config"
	"api/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	DB.AutoMigrate(
		models.Admin{},
		models.Client{},
	)

	createAdminUser()
}

func createAdminUser() {
	// Verificar si el usuario administrador ya existe
	var existingAdmin models.Admin
	if err := DB.Where("username = ?", config.AdminUser).First(&existingAdmin).Error; err == nil {
		log.Printf("Admin user already exists: %s", config.AdminUser)
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Fatalf("Failed to check for existing admin user: %v", err)
	}

	// Hashear la contraseña del administrador
	hashedPassword, err := config.HashPassword(config.AdminPass)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Crear el objeto admin con la contraseña hasheada
	admin := models.Admin{
		Username: config.AdminUser,
		Password: hashedPassword,
		Role:     config.AdminRole,
	}

	// Guardar el usuario administrador en la base de datos
	if err := DB.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	log.Printf("Admin user created: %s with role %s", config.AdminUser, config.AdminRole)
}
