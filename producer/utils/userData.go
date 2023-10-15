package utils

import (
	"github.com/ankitdmon/producer/models"
	"gorm.io/gorm"
)

func InsertInitialUserData(db *gorm.DB) {
	users := []models.User{
		{Name: "John Doe", Mobile: "123-456-7890", Latitude: 37.7749, Longitude: 122.4194},
		{Name: "Jane Smith", Mobile: "555-555-5555", Latitude: 40.7128, Longitude: 74.0060},
		{Name: "Bob Johnson", Mobile: "987-654-3210", Latitude: 28.6139, Longitude: 77.2090},
		{Name: "Alice Brown", Mobile: "111-222-3333", Latitude: 19.0760, Longitude: 72.8777},
		{Name: "Charlie Williams", Mobile: "777-888-9999", Latitude: 12.9716, Longitude: 77.5946},
		{Name: "Eve Davis", Mobile: "555-123-4567", Latitude: 22.5726, Longitude: 88.3639},
		{Name: "Grace Taylor", Mobile: "999-888-7777", Latitude: 26.9124, Longitude: 75.7873},
		{Name: "Frank Wilson", Mobile: "555-999-5555", Latitude: 17.3850, Longitude: 78.4867},
		{Name: "Hannah Lee", Mobile: "123-987-3210", Latitude: 18.5204, Longitude: 73.8567},
		{Name: "Isaac Harris", Mobile: "777-333-1111", Latitude: 23.2599, Longitude: 77.4126},
		{Name: "Sarah White", Mobile: "555-111-3333", Latitude: 41.8781, Longitude: 87.6298},
		{Name: "Michael Brown", Mobile: "123-456-7890", Latitude: 34.0522, Longitude: 118.2437},
		{Name: "Emily Davis", Mobile: "555-777-9999", Latitude: 51.5074, Longitude: 0.1278},
		{Name: "Daniel Wilson", Mobile: "555-123-4567", Latitude: 48.8566, Longitude: 2.3522},
		{Name: "Olivia Harris", Mobile: "123-987-3210", Latitude: 55.7558, Longitude: 37.6176},
		{Name: "William Johnson", Mobile: "777-333-1111", Latitude: 35.682839, Longitude: 139.759455},
		{Name: "Mia Smith", Mobile: "123-123-1234", Latitude: 25.276987, Longitude: 51.518133},
		{Name: "Ethan Martinez", Mobile: "321-654-9876", Latitude: 29.760427, Longitude: -95.369804},
		{Name: "Ava Johnson", Mobile: "444-777-2222", Latitude: 40.7128, Longitude: -74.0060},
		{Name: "Sophia Lee", Mobile: "555-555-1234", Latitude: 34.0522, Longitude: -118.2437},
	}

	for _, user := range users {
		db.Create(&user)
	}
}
