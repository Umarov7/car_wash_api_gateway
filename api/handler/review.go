package handler

import (
	pb "api-gateway/genproto/reviews"
	"api-gateway/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReview godoc
// @Summary Create review
// @Description Creates a new review
// @Tags review
// @Security ApiKeyAuth
// @Param data body models.ReviewCreate true "Review"
// @Success 201 {object} reviews.CreateResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /reviews [post]
func (h *Handler) CreateReview(c *gin.Context) {
	h.Logger.Info("CreateReview handler is invoked")

	id, err := getUserID(c)
	if err != nil {
		handleError(c, h, err, "invalid user", http.StatusUnauthorized)
		return
	}

	var req models.ReviewCreate
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Review.CreateReview(ctx, &pb.NewReview{
		UserId:     id,
		BookingId:  req.BookingID,
		ProviderId: req.ProviderID,
		Rating:     req.Rating,
		Comment:    req.Comment,
	})
	if err != nil {
		handleError(c, h, err, "error creating review", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("CreateReview handler is completed")
	c.JSON(http.StatusCreated, resp)
}

// UpdateReview godoc
// @Summary Update review
// @Description Updates review
// @Tags review
// @Security ApiKeyAuth
// @Param id path string true "Review ID"
// @Param data body models.ReviewUpdate true "Review"
// @Success 200 {object} reviews.UpdateResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /reviews/{id} [put]
func (h *Handler) UpdateReview(c *gin.Context) {
	h.Logger.Info("UpdateReview handler is invoked")

	id := c.Param("id")
	if id == "" {
		handleError(c, h, nil, "invalid data format", http.StatusBadRequest)
		return
	}

	var req models.ReviewUpdate
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Review.UpdateReview(ctx, &pb.NewData{
		Id:      id,
		Rating:  req.Rating,
		Comment: req.Comment,
	})
	if err != nil {
		handleError(c, h, err, "error updating review", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("UpdateReview handler is completed")
	c.JSON(http.StatusOK, resp)
}

// DeleteReview godoc
// @Summary Delete review
// @Description Deletes review
// @Tags review
// @Security ApiKeyAuth
// @Param id path string true "Review ID"
// @Success 200 {object} string "Review deleted successfully"
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /reviews/{id} [delete]
func (h *Handler) DeleteReview(c *gin.Context) {
	h.Logger.Info("DeleteReview handler is invoked")

	id := c.Param("id")
	if id == "" {
		handleError(c, h, nil, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	_, err := h.Review.DeleteReview(ctx, &pb.ID{Id: id})
	if err != nil {
		handleError(c, h, err, "error deleting review", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("DeleteReview handler is completed")
	c.JSON(http.StatusOK, "Review deleted successfully")
}

// FetchReviews godoc
// @Summary Fetches reviews
// @Description Fetches reviews
// @Tags review
// @Security ApiKeyAuth
// @Param page query int true "Page number"
// @Param limit query int true "Number of items per page"
// @Success 200 {object} reviews.ReviewsList
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /reviews/all [get]
func (h *Handler) FetchReviews(c *gin.Context) {
	h.Logger.Info("FetchReviews handler is invoked")

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

	resp, err := h.Review.ListReviews(ctx, &pb.Pagination{Page: page, Limit: limit})
	if err != nil {
		handleError(c, h, err, "error fetching reviews", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("FetchReviews handler is completed")
	c.JSON(http.StatusOK, resp)
}
