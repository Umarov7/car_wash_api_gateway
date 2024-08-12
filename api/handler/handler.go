package handler

import (
	"api-gateway/config"
	pbb "api-gateway/genproto/bookings"
	pbpa "api-gateway/genproto/payments"
	pbp "api-gateway/genproto/providers"
	pbr "api-gateway/genproto/reviews"
	pbs "api-gateway/genproto/services"
	pbu "api-gateway/genproto/user"
	"api-gateway/pkg"
	"api-gateway/pkg/logger"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	User           pbu.UserClient
	Provider       pbp.ProvidersClient
	Service        pbs.ServicesClient
	Booking        pbb.BookingsClient
	Payment        pbpa.PaymentsClient
	Review         pbr.ReviewsClient
	Logger         *slog.Logger
	ContextTimeout time.Duration
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		User:           pkg.NewUserClient(cfg),
		Provider:       pkg.NewProvidersClient(cfg),
		Service:        pkg.NewServicesClient(cfg),
		Booking:        pkg.NewBookingsClient(cfg),
		Payment:        pkg.NewPaymentsClient(cfg),
		Review:         pkg.NewReviewsClient(cfg),
		Logger:         logger.NewLogger(),
		ContextTimeout: time.Second * 5,
	}
}

func handleError(c *gin.Context, h *Handler, err error, msg string, code int) {
	er := errors.Wrap(err, msg).Error()
	c.AbortWithStatusJSON(code, gin.H{"error": er})
	h.Logger.Error(er)
}
