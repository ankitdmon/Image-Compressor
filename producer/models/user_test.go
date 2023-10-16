package models

import (
	"testing"
)

func TestUserModelValidation(t *testing.T) {
	user := User{
		Name:      "John Doe",
		Mobile:    "123-456-7890",
		Latitude:  37.7749,
		Longitude: 122.4194,
	}

	if user.ID != 0 {
		t.Errorf("Expected ID to be 0, but got %d", user.ID)
	}

	if user.Name != "John Doe" {
		t.Errorf("Expected Name to be 'John Doe', but got '%s'", user.Name)
	}

	if user.Mobile != "123-456-7890" {
		t.Errorf("Expected Mobile to be '123-456-7890', but got '%s'", user.Mobile)
	}

	if user.Latitude != 37.7749 {
		t.Errorf("Expected Latitude to be 37.7749, but got %f", user.Latitude)
	}

	if user.Longitude != 122.4194 {
		t.Errorf("Expected Longitude to be 122.4194, but got %f", user.Longitude)
	}

	if !user.CreatedAt.IsZero() {
		t.Errorf("Expected CreatedAt to be zero, but got a non-zero timestamp")
	}

	if !user.UpdatedAt.IsZero() {
		t.Errorf("Expected UpdatedAt to be zero, but got a non-zero timestamp")
	}
}
