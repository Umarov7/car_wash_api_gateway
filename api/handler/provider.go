package handler

import (
	pb "api-gateway/genproto/providers"
	"api-gateway/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProvider godoc
// @Summary Creates provider
// @Description Adds a new provider
// @Tags provider
// @Security ApiKeyAuth
// @Param data body models.ProviderCreate true "New provider"
// @Success 201 {object} providers.CreateResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /providers [post]
func (h *Handler) CreateProvider(c *gin.Context) {
	h.Logger.Info("CreateProvider handler is invoked")

	id, err := getUserID(c)
	if err != nil {
		handleError(c, h, err, "invalid user", http.StatusUnauthorized)
		return
	}

	var req models.ProviderCreate
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Provider.CreateProvider(ctx, &pb.NewProvider{
		UserId:        id,
		CompanyName:   req.CompanyName,
		Description:   req.Description,
		Services:      req.Services,
		Availability:  req.Availability,
		AverageRating: req.AverageRating,
		Location: &pb.Location{
			Address:   req.Location.Address,
			City:      req.Location.City,
			Country:   req.Location.Country,
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		},
	})
	if err != nil {
		handleError(c, h, err, "error creating provider", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("CreateProvider handler is completed")
	c.JSON(http.StatusCreated, resp)
}

// SearchProviders godoc
// @Summary Searches providers
// @Description Searches providers
// @Tags provider
// @Security ApiKeyAuth
// @Param data body providers.Filter true "Search data"
// @Success 200 {object} providers.SearchResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /providers/search [get]
func (h *Handler) SearchProviders(c *gin.Context) {
	h.Logger.Info("SearchProviders handler is invoked")

	var req pb.Filter
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Provider.SearchProviders(ctx, &req)
	if err != nil {
		handleError(c, h, err, "error finding providers", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("SearchProviders handler is completed")
	c.JSON(http.StatusOK, resp)
}
