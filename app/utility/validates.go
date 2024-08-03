package utility

import (
	"errors"

	"github.com/fernandamelov/device-management-api/app/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateDevice(request *models.Device) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Brand, validation.Required),
	)
}

func ValidateID(id string) error {
	if id == "" {
		return errors.New("ID cannot be empty")
	}
	return nil
}

func ValidateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand cannot be empty")
	}
	return nil
}
