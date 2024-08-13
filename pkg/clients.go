package pkg

import (
	"api-gateway/config"
	pbb "api-gateway/genproto/bookings"
	pbn "api-gateway/genproto/notifications"
	pbpa "api-gateway/genproto/payments"
	pbp "api-gateway/genproto/providers"
	pbr "api-gateway/genproto/reviews"
	pbs "api-gateway/genproto/services"
	pbu "api-gateway/genproto/user"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(cfg *config.Config) pbu.UserClient {
	conn, err := grpc.NewClient(cfg.AUTH_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbu.NewUserClient(conn)
}

func NewProvidersClient(cfg *config.Config) pbp.ProvidersClient {
	conn, err := grpc.NewClient(cfg.BOOKING_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbp.NewProvidersClient(conn)
}

func NewServicesClient(cfg *config.Config) pbs.ServicesClient {
	conn, err := grpc.NewClient(cfg.BOOKING_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbs.NewServicesClient(conn)
}

func NewBookingsClient(cfg *config.Config) pbb.BookingsClient {
	conn, err := grpc.NewClient(cfg.BOOKING_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbb.NewBookingsClient(conn)
}

func NewPaymentsClient(cfg *config.Config) pbpa.PaymentsClient {
	conn, err := grpc.NewClient(cfg.BOOKING_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbpa.NewPaymentsClient(conn)
}

func NewReviewsClient(cfg *config.Config) pbr.ReviewsClient {
	conn, err := grpc.NewClient(cfg.BOOKING_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbr.NewReviewsClient(conn)
}

func NewNotificationClient(cfg *config.Config) pbn.NotificationsClient {
	conn, err := grpc.NewClient(cfg.BOOKING_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbn.NewNotificationsClient(conn)
}
