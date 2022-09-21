package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/go-playground/validator"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" validate:"required"`
	DateOfBirth string    `json:"date_of_birth" validate:"required"`
	Address     string    `json:"address" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	return nil
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateUserStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	validation := validator.New()
	err := validation.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
