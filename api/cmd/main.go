package main

import (
	"api/config"
	"api/database"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitConfig()
	database.ConnectToDatabase()

	r := gin.Default()

	r.Run(":8080")
}
