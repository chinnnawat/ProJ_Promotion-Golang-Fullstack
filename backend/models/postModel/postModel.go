package postmodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostUser struct {
	gorm.Model
	// UserID         string from keycloak
	PostID         uuid.UUID `gorm:"type:uuid;primaryKey" validate:"required"`
	Title          string
	Description    string
	NamePost       string `validate:"required,min=3"`
	RestaurantName string `validate:"required,min=3"`
	Category       string `validate:"required"`
	Period         string `validate:"required"`
	UrlPost        string `validate:"required"`
	MapUrl         string `validate:"required"`
	Location       string `validate:"required"`
	Participants   int8   `validate:"required"`
	ReviewPost     []reviewPost
}

type reviewPost struct {
	Rating  *int8
	Comment *string
}
