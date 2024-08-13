package api

import (
	_ "api-gateway/api/docs"
	"api-gateway/api/handler"
	"api-gateway/api/middleware"
	"api-gateway/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title On-Demand Car Wash Service
// @version 1.0
// @description API Gateway of On-Demand Car Wash Service
// @host localhost:8080
// @BasePath /car-wash
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(cfg *config.Config) *gin.Engine {
	h := handler.NewHandler(cfg)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/car-wash")
	api.Use(middleware.Check(cfg))

	u := api.Group("/users")
	{
		u.GET("/profile", h.GetProfile)
		u.PUT("/profile", h.UpdateProfile)
	}

	p := api.Group("/providers")
	{
		p.POST("/register", h.CreateProvider)
		p.GET("/:id", h.GetProvider)
		p.PUT("/:id", h.UpdateProvider)
		p.DELETE("/:id", h.DeleteProvider)
		p.GET("/all", h.FetchProviders)
		p.GET("/search", h.SearchProviders)
	}

	s := api.Group("/services")
	{
		s.POST("", h.CreateService)
		s.GET("/:id", h.GetService)
		s.PUT("/:id", h.UpdateService)
		s.DELETE("/:id", h.DeleteService)
		s.GET("/all", h.FetchServices)
		s.GET("/search", h.SearchServices)
		s.GET("/popular", h.GetPopularServices)
	}

	b := api.Group("/bookings")
	{
		b.POST("", h.CreateBooking)
		b.GET("/:id", h.GetBooking)
		b.PUT("/:id", h.UpdateBooking)
		b.DELETE("/:id", h.CancelBooking)
		b.GET("/all", h.FetchBookings)
	}

	pay := api.Group("/payments")
	{
		pay.POST("", h.CreatePayment)
		pay.GET("/:id", h.GetPayment)
		pay.GET("/all", h.FetchPayments)
	}

	r := api.Group("/reviews")
	{
		r.POST("", h.CreateReview)
		r.GET("/:id", h.GetReview)
		r.PUT("/:id", h.UpdateReview)
		r.DELETE("/:id", h.DeleteReview)
		r.GET("/all", h.FetchReviews)
	}

	n := api.Group("/notifications")
	{
		n.POST("", h.CreateNotification)
		n.GET("/:id", h.GetNotification)
	}

	return router
}
