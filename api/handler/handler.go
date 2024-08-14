package handler

import (
	"api-gateway/config"
	pbb "api-gateway/genproto/bookings"
	pbn "api-gateway/genproto/notifications"
	pbpa "api-gateway/genproto/payments"
	pbp "api-gateway/genproto/providers"
	pbr "api-gateway/genproto/reviews"
	pbs "api-gateway/genproto/services"
	pbu "api-gateway/genproto/user"
	"api-gateway/kafka/producer"
	"api-gateway/pkg"
	"api-gateway/pkg/logger"
	"log/slog"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	User                     pbu.UserClient
	Provider                 pbp.ProvidersClient
	Service                  pbs.ServicesClient
	Booking                  pbb.BookingsClient
	Payment                  pbpa.PaymentsClient
	Review                   pbr.ReviewsClient
	Notification             pbn.NotificationsClient
	Logger                   *slog.Logger
	ContextTimeout           time.Duration
	KafkaProducer            producer.IKafkaProducer
	TopicBookingCreated      string
	TopicBookingUpdated      string
	TopicBookingCancelled    string
	TopicPaymentCreated      string
	TopicReviewCreated       string
	TopicNotificationCreated string
}

func NewHandler(cfg *config.Config) *Handler {
	kafkaBrokerAddress := cfg.KAFKA_HOST + ":" + cfg.KAFKA_PORT

	return &Handler{
		User:                     pkg.NewUserClient(cfg),
		Provider:                 pkg.NewProvidersClient(cfg),
		Service:                  pkg.NewServicesClient(cfg),
		Booking:                  pkg.NewBookingsClient(cfg),
		Payment:                  pkg.NewPaymentsClient(cfg),
		Review:                   pkg.NewReviewsClient(cfg),
		Notification:             pkg.NewNotificationClient(cfg),
		Logger:                   logger.NewLogger(),
		ContextTimeout:           time.Second * 10,
		KafkaProducer:            producer.NewKafkaProducer([]string{kafkaBrokerAddress}),
		TopicBookingCreated:      cfg.KAFKA_TOPIC_BOOKING_CREATED,
		TopicBookingUpdated:      cfg.KAFKA_TOPIC_BOOKING_UPDATED,
		TopicBookingCancelled:    cfg.KAFKA_TOPIC_BOOKING_CANCELLED,
		TopicPaymentCreated:      cfg.KAFKA_TOPIC_PAYMENT_CREATED,
		TopicReviewCreated:       cfg.KAFKA_TOPIC_REVIEW_CREATED,
		TopicNotificationCreated: cfg.KAFKA_TOPIC_NOTIFICATION_CREATED,
	}
}

func handleError(c *gin.Context, h *Handler, err error, msg string, code int) {
	er := errors.Wrap(err, msg).Error()
	c.AbortWithStatusJSON(code, gin.H{"error": er})
	h.Logger.Error(er)

}

func getUserID(c *gin.Context) (string, error) {
	id, ok := c.Get("user_id")
	if !ok {
		return "", errors.New("user id not found")
	}

	idStr, ok := id.(string)
	if !ok {
		return "", errors.New("invalid user id")
	}

	return idStr, nil
}

func parseIntQueryParam(queryParam string) (int32, error) {
	if queryParam == "" {
		return -1, errors.New("empty integer parameter")
	}

	value, err := strconv.Atoi(queryParam)
	if err != nil || value < 1 {
		return -1, errors.New("invalid integer parameter")
	}

	return int32(value), nil
}

func parseFloatQueryParam(queryParam string) (float32, error) {
	if queryParam == "" {
		return -1, errors.New("empty float parameter")
	}

	value, err := strconv.ParseFloat(queryParam, 32)
	if err != nil || value < 1 {
		return -1, errors.New("invalid float parameter")
	}

	return float32(value), nil
}
