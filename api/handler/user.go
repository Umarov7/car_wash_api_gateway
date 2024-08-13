package handler

import (
	pb "api-gateway/genproto/user"
	"api-gateway/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary Gets profile
// @Description Retrieves user profile
// @Tags user
// @Security ApiKeyAuth
// @Success 200 {object} user.Profile
// @Failure 401 {object} string "Invalid user"
// @Failure 500 {object} string "Server error while processing request"
// @Router /users/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	h.Logger.Info("GetProfile handler is invoked")

	id, err := getUserID(c)
	if err != nil {
		handleError(c, h, err, "invalid user", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.User.GetProfile(ctx, &pb.ID{Id: id})
	if err != nil {
		handleError(c, h, err, "error getting user profile", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("GetProfile handler is completed")
	c.JSON(http.StatusOK, resp)
}

// UpdateProfile godoc
// @Summary Updates profile
// @Description Updates user profile
// @Tags user
// @Security ApiKeyAuth
// @Param data body models.UserUpdate true "New user data"
// @Success 200 {object} user.UpdateResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 401 {object} string "Invalid user"
// @Failure 500 {object} string "Server error while processing request"
// @Router /users/profile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	h.Logger.Info("UpdateProfile handler is invoked")

	id, err := getUserID(c)
	if err != nil {
		handleError(c, h, err, "invalid user", http.StatusUnauthorized)
		return
	}

	var req models.UserUpdate
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.User.UpdateProfile(ctx, &pb.NewData{
		Id:          id,
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		handleError(c, h, err, "error updating user profile", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("UpdateProfile handler is completed")
	c.JSON(http.StatusOK, resp)
}
