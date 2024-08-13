package handler

import (
	pb "api-gateway/genproto/services"
	"api-gateway/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateService godoc
// @Summary Creates service
// @Description Adds a new service
// @Tags service
// @Security ApiKeyAuth
// @Param data body services.NewService true "New service"
// @Success 201 {object} services.CreateResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /services [post]
func (h *Handler) CreateService(c *gin.Context) {
	h.Logger.Info("CreateService handler is invoked")

	var req pb.NewService
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Service.CreateService(ctx, &req)
	if err != nil {
		handleError(c, h, err, "error creating service", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("CreateService handler is completed")
	c.JSON(http.StatusCreated, resp)
}

// UpdateService godoc
// @Summary Updates service
// @Description Updates service
// @Tags service
// @Security ApiKeyAuth
// @Param id path string true "Service ID"
// @Param data body models.ServiceUpdate true "New service data"
// @Success 200 {object} services.UpdateResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /services/{id} [put]
func (h *Handler) UpdateService(c *gin.Context) {
	h.Logger.Info("UpdateService handler is invoked")

	id := c.Param("id")
	if id == "" {
		handleError(c, h, nil, "invalid data format", http.StatusBadRequest)
		return
	}

	var req models.ServiceUpdate
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Service.UpdateService(ctx, &pb.NewData{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Duration:    req.Duration,
	})
	if err != nil {
		handleError(c, h, err, "error updating service", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("UpdateService handler is completed")
	c.JSON(http.StatusOK, resp)
}

// DeleteService godoc
// @Summary Deletes service
// @Description Deletes service
// @Tags service
// @Security ApiKeyAuth
// @Param id path string true "Service ID"
// @Success 200 {object} string "Service deleted successfully"
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /services/{id} [delete]
func (h *Handler) DeleteService(c *gin.Context) {
	h.Logger.Info("DeleteService handler is invoked")

	id := c.Param("id")
	if id == "" {
		handleError(c, h, nil, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	_, err := h.Service.DeleteService(ctx, &pb.ID{Id: id})
	if err != nil {
		handleError(c, h, err, "error deleting service", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("DeleteService handler is completed")
	c.JSON(http.StatusOK, "Service deleted successfully")
}

// FetchServices godoc
// @Summary Fetches services
// @Description Fetches services
// @Tags service
// @Security ApiKeyAuth
// @Param page query int true "Page number"
// @Param limit query int true "Number of items per page"
// @Success 200 {object} services.ServicesList
// @Failure 400 {object} string "Invalid pagination parameter"
// @Failure 500 {object} string "Server error while processing request"
// @Router /services/all [get]
func (h *Handler) FetchServices(c *gin.Context) {
	h.Logger.Info("FetchServices handler is invoked")

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

	resp, err := h.Service.ListServices(ctx, &pb.Pagination{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		handleError(c, h, err, "error fetching services", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("FetchServices handler is completed")
	c.JSON(http.StatusOK, resp)
}

// SearchServices godoc
// @Summary Searches services
// @Description Searches services
// @Tags service
// @Security ApiKeyAuth
// @Param data body services.Filter true "Search data"
// @Success 200 {object} services.SearchResp
// @Failure 400 {object} string "Invalid data format"
// @Failure 500 {object} string "Server error while processing request"
// @Router /services/search [post]
func (h *Handler) SearchServices(c *gin.Context) {
	h.Logger.Info("SearchServices handler is invoked")

	var req pb.Filter
	if err := c.ShouldBind(&req); err != nil {
		handleError(c, h, err, "invalid data format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	resp, err := h.Service.SearchServices(ctx, &req)
	if err != nil {
		handleError(c, h, err, "error searching services", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("SearchServices handler is completed")
	c.JSON(http.StatusOK, resp)
}
