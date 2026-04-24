package main

import (
	"log"
	"pool/internal/booking"
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

	bookingService := booking.NewBookingService(dbConnection)
	bookingHandler := booking.NewUserHandlerFrom(bookingService)
	bookingHandler.RegisterEndPoints(router)

	router.Run()
}
