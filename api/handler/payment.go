package handler

import (
	pb "api-gateway/genproto/payments"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePayment godoc
// @Summary Creates payment
// @Description Adds a new payment
// @Tags payment
// @Security ApiKeyAuth
// @Param data body payments.NewPayment true "New payment"
// @Success 201 {object} string "Payment created"
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /payments [post]
func (h *Handler) CreatePayment(c *gin.Context) {
	h.Logger.Info("CreatePayment handler is invoked")

	var req pb.NewPayment
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	message, err := json.Marshal(&req)
	if err != nil {
		handleError(c, h, err, "error serializing payment", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	err = h.KafkaProducer.Produce(ctx, h.TopicPaymentCreated, []byte(message))
	if err != nil {
		handleError(c, h, err, "error creating payment", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("CreatePayment handler is completed")
	c.JSON(http.StatusCreated, "Payment created")
}

// GetPayment godoc
// @Summary Gets payment
// @Description Gets payment
// @Tags payment
// @Security ApiKeyAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} payments.Payment
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /payments/{id} [get]
func (h *Handler) GetPayment(c *gin.Context) {
	h.Logger.Info("GetPayment handler is invoked")

	id := c.Param("id")
	if id == "" {
		handleError(c, h, nil, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Payment.GetPayment(ctx, &pb.ID{Id: id})
	if err != nil {
		handleError(c, h, err, "error finding payment", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("GetPayment handler is completed")
	c.JSON(http.StatusOK, resp)
}

// FetchPayments godoc
// @Summary Fetches payments
// @Description Fetches payments
// @Tags payment
// @Security ApiKeyAuth
// @Param page query int true "Page number"
// @Param limit query int true "Number of items per page"
// @Success 200 {object} payments.PaymentsList
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /payments/all [get]
func (h *Handler) FetchPayments(c *gin.Context) {
	h.Logger.Info("FetchPayments handler is invoked")

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

	resp, err := h.Payment.ListPayments(ctx, &pb.Pagination{Page: page, Limit: limit})
	if err != nil {
		handleError(c, h, err, "error fetching payments", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("FetchPayments handler is completed")
	c.JSON(http.StatusOK, resp)
}
