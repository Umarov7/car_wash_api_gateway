package handler

import (
	pb "api-gateway/genproto/bookings"
	"api-gateway/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBooking godoc
// @Summary Creates booking
// @Description Adds a new booking
// @Tags booking
// @Security ApiKeyAuth
// @Param data body models.BookingCreate true "New booking"
// @Success 201 {object} bookings.CreateResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /bookings [post]
func (h *Handler) CreateBooking(c *gin.Context) {
	h.Logger.Info("CreateBooking handler is invoked")

	id, err := getUserID(c)
	if err != nil {
		handleError(c, h, err, "invalid user", http.StatusUnauthorized)
		return
	}

	var req models.BookingCreate
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Booking.CreateBooking(ctx, &pb.NewBooking{
		UserId:        id,
		ProviderId:    req.ProviderID,
		ServiceId:     req.ServiceID,
		Status:        req.Status,
		ScheduledTime: req.ScheduledTime,
		Location: &pb.Location{
			Address:   req.Location.Address,
			City:      req.Location.City,
			Country:   req.Location.Country,
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		},
		TotalPrice: req.TotalPrice,
	})
	if err != nil {
		handleError(c, h, err, "error creating booking", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("CreateBooking handler is completed")
	c.JSON(http.StatusOK, resp)
}

// GetBooking godoc
// @Summary Gets booking
// @Description Gets booking
// @Tags booking
// @Security ApiKeyAuth
// @Param id path string true "Booking ID"
// @Success 200 {object} bookings.Booking
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /bookings/{id} [get]
func (h *Handler) GetBooking(c *gin.Context) {
	h.Logger.Info("GetBooking handler is invoked")

	id := c.Param("id")
	if id == "" {
		handleError(c, h, nil, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Booking.GetBooking(ctx, &pb.ID{Id: id})
	if err != nil {
		handleError(c, h, err, "error finding booking", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("GetBooking handler is completed")
	c.JSON(http.StatusOK, resp)
}

// UpdateBooking godoc
// @Summary Updates booking
// @Description Updates booking
// @Tags booking
// @Security ApiKeyAuth
// @Param id path string true "Booking ID"
// @Param data body models.BookingUpdate true "New booking data"
// @Success 200 {object} bookings.UpdateResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /bookings/{id} [put]
func (h *Handler) UpdateBooking(c *gin.Context) {
	h.Logger.Info("UpdateBooking handler is invoked")

	id := c.Param("id")
	if id == "" {
		handleError(c, h, nil, "invalid data format", http.StatusBadRequest)
		return
	}

	var req models.BookingUpdate
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Booking.UpdateBooking(ctx, &pb.NewData{
		Id:            id,
		Status:        req.Status,
		ScheduledTime: req.ScheduledTime,
		Location: &pb.Location{
			Address:   req.Location.Address,
			City:      req.Location.City,
			Country:   req.Location.Country,
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		},
		TotalPrice: req.TotalPrice,
	})
	if err != nil {
		handleError(c, h, err, "error updating booking", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("UpdateBooking handler is completed")
	c.JSON(http.StatusOK, resp)
}

// CancelBooking godoc
// @Summary Cancels booking
// @Description Cancels booking
// @Tags booking
// @Security ApiKeyAuth
// @Param id path string true "Booking ID"
// @Success 200 {object} string "Booking canceled successfully"
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /bookings/{id} [delete]
func (h *Handler) CancelBooking(c *gin.Context) {
	h.Logger.Info("CancelBooking handler is invoked")

	id := c.Param("id")
	if id == "" {
		handleError(c, h, nil, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	_, err := h.Booking.CancelBooking(ctx, &pb.ID{Id: id})
	if err != nil {
		handleError(c, h, err, "error canceling booking", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("CancelBooking handler is completed")
	c.JSON(http.StatusOK, "Booking canceled successfully")
}

// FetchBookings godoc
// @Summary Fetches bookings
// @Description Fetches bookings
// @Tags booking
// @Security ApiKeyAuth
// @Param page query int true "Page number"
// @Param limit query int true "Number of items per page"
// @Success 200 {object} bookings.BookingsList
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /bookings/all [get]
func (h *Handler) FetchBookings(c *gin.Context) {
	h.Logger.Info("FetchBookings handler is invoked")

	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := parseIntQueryParam(pageStr)
	if err != nil {
		handleError(c, h, err, "invalid pagination parameter", http.StatusBadRequest)
		return
	}

	limit, err := parseIntQueryParam(limitStr)
	if err != nil {
		handleError(c, h, err, "invalid pagination parameter", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Booking.ListBookings(ctx, &pb.Pagination{Page: page, Limit: limit})
	if err != nil {
		handleError(c, h, err, "error fetching bookings", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("FetchBookings handler is completed")
	c.JSON(http.StatusOK, resp)
}
