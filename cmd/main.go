package main

import (
	"log"
	db "pool/internal/db"
	"pool/internal/user"

	"github.com/gin-gonic/gin"
)



func main() {
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect data base: %v", err)
	}
	router := gin.Default()

	userService := user.NewUserService(dbConnection)

	userHandler := user.NewUserHandlerFrom(userService)
	userHandler.RegisterEndpoints(router)

	router.Run()
}
