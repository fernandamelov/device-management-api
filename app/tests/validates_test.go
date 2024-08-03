package tests

import (
	"testing"

	"github.com/fernandamelov/device-management-api/app/models"
	"github.com/fernandamelov/device-management-api/app/utility"
	"github.com/stretchr/testify/assert"
)

func TestValidateDevice(t *testing.T) {
	tests := []struct {
		name    string
		request *models.Device
		wantErr bool
	}{
		{
			name: "Valid Device",
			request: &models.Device{
				Name:  "Test Device",
				Brand: "Test Brand",
			},
			wantErr: false,
		},
		{
			name: "Missing Name",
			request: &models.Device{
				Brand: "Test Brand",
			},
			wantErr: true,
		},
		{
			name: "Missing Brand",
			request: &models.Device{
				Name: "Test Device",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utility.ValidateDevice(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateBrand(t *testing.T) {
	tests := []struct {
		name    string
		brand   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid Brand",
			brand:   "Test Brand",
			wantErr: false,
		},
		{
			name:    "Empty Brand",
			brand:   "",
			wantErr: true,
			errMsg:  "brand cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utility.ValidateBrand(tt.brand)
			if tt.wantErr {
				assert.EqualError(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateID(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Valid ID",
			id:      "123",
			wantErr: false,
		},
		{
			name:    "Empty ID",
			id:      "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utility.ValidateID(tt.id)
			if tt.wantErr {
				assert.EqualError(t, err, "ID cannot be empty")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
