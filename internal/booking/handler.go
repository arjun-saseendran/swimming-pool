package booking

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	groupName      string
	bookingService BookingService
}

func NewUserHandler(bookingService BookingService) *BookingHandler {
	return &BookingHandler{"api/booking", bookingService}
}

func (handler *BookingHandler)RegisterEndPoints(r *gin.Engine){
	bookingGroup := r.Group(handler.groupName)

	bookingGroup.POST("",handler.CreateBooking)
	}

func (handler *BookingHandler) CreateBooking(ctx *gin.Context){
	bookingData := NewInputBooking()
	err := ctx.BindJSON(&bookingData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err":"failed bind booking data"})
	return
	}

	newBooking, err := handler.bookingService.Create(bookingData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err":"failed to create new booking"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg":"booking created successfully", "data":newBooking})
}
