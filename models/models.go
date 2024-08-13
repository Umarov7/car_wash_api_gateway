package models

type UserUpdate struct {
	Email       string `json:"email" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type Location struct {
	Address   string  `json:"address" validate:"required"`
	City      string  `json:"city" validate:"required"`
	Country   string  `json:"country" validate:"required"`
	Latitude  float32 `json:"latitude" validate:"required"`
	Longitude float32 `json:"longitude" validate:"required"`
}

type ProviderCreate struct {
	CompanyName   string   `json:"company_name" validate:"required"`
	Description   string   `json:"description" validate:"required"`
	Services      []string `json:"services" validate:"required"`
	Availability  []string `json:"availability" validate:"required"`
	AverageRating float32  `json:"average_rating" validate:"required"`
	Location      Location `json:"location" validate:"required"`
}

type ServiceUpdate struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Duration    int32   `json:"duration"`
}

type BookingCreate struct {
	ProviderID    string   `json:"provider_id" validate:"required"`
	ServiceID     string   `json:"service_id" validate:"required"`
	Status        string   `json:"status" validate:"required"`
	ScheduledTime string   `json:"scheduled_time" validate:"required"`
	Location      Location `json:"location" validate:"required"`
	TotalPrice    float32  `json:"total_price" validate:"required"`
}

type BookingUpdate struct {
	Status        string   `json:"status"`
	ScheduledTime string   `json:"scheduled_time"`
	Location      Location `json:"location"`
	TotalPrice    float32  `json:"total_price"`
}

type ReviewCreate struct {
	BookingID  string `json:"booking_id" validate:"required"`
	ProviderID string `json:"provider_id" validate:"required"`
	Rating     int32  `json:"rating" validate:"required"`
	Comment    string `json:"comment" validate:"required"`
}

type ReviewUpdate struct {
	Rating  int32  `json:"rating"`
	Comment string `json:"comment"`
}

type ProviderUpdate struct {
	CompanyName   string   `json:"company_name"`
	Description   string   `json:"description"`
	Services      []string `json:"services"`
	Availability  []string `json:"availability"`
	AverageRating float32  `json:"average_rating"`
	Location      Location `json:"location"`
}
