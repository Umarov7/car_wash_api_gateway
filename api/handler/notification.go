package handler

import (
	pbn "api-gateway/genproto/notifications"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateNotification godoc
// @Summary Creates notification
// @Description Adds a new notification
// @Tags notification
// @Security ApiKeyAuth
// @Param data body notifications.NewNotification true "Receiver ID, Title and Message"
// @Success 201 {object} string "Notification created"
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /notifications [post]
func (h *Handler) CreateNotification(c *gin.Context) {
	h.Logger.Info("CreateNotification handler is invoked")

	var req pbn.NewNotification
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	message, err := json.Marshal(&req)
	if err != nil {
		handleError(c, h, err, "error serializing notification", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	err = h.KafkaProducer.Produce(ctx, h.TopicNotificationCreated, []byte(message))
	if err != nil {
		handleError(c, h, err, "error creating notification", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("CreateNotification handler is completed")
	c.JSON(http.StatusCreated, "Notification created")
}

// GetNotification godoc
// @Summary Gets notification
// @Description Gets notification
// @Tags notification
// @Security ApiKeyAuth
// @Param id path string true "Notification ID"
// @Success 200 {object} notifications.Notification
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /notifications/{id} [get]
func (h *Handler) GetNotification(c *gin.Context) {
	h.Logger.Info("GetNotification handler is invoked")

	id := c.Param("id")
	if id == "" {
		handleError(c, h, nil, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Notification.GetNotification(ctx, &pbn.ID{Id: id})
	if err != nil {
		handleError(c, h, err, "error finding notification", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("GetNotification handler is completed")
	c.JSON(http.StatusOK, resp)
}
