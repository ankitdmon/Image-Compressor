package models

import (
	"testing"

	"github.com/ankitdmon/producer/models"
)

func TestProductModel(t *testing.T) {
	product := models.Product    {
		UserID:             1,
		ProductName:        "Sample Product",
		ProductDescription: "This is a sample product description.",
		ProductImages:      []string{"image1.jpg", "image2.jpg"},
		ProductPrice:       19.99,
		CompressedImages:   []string{"compressed1.jpg", "compressed2.jpg"},
	}

	if product.ProductID != 0 {
		t.Errorf("Expected ProductID to be 0, but got %d", product.ProductID)
	}

	if product.ProductName != "Sample Product" {
		t.Errorf("Expected ProductName to be 'Sample Product', but got '%s'", product.ProductName)
	}

	if product.ProductDescription != "This is a sample product description." {
		t.Errorf("Expected ProductDescription to be 'This is a sample product description.', but got '%s'", product.ProductDescription)
	}

	if product.ProductPrice != 19.99 {
		t.Errorf("Expected ProductPrice to be 19.99, but got %f", product.ProductPrice)
	}

	if !product.CreatedAt.IsZero() {
		t.Errorf("Expected CreatedAt to be zero, but got a non-zero timestamp")
	}
	if !product.UpdatedAt.IsZero() {
		t.Errorf("Expected UpdatedAt to be zero, but got a non-zero timestamp")
	}

	if len(product.ProductImages) != 2 {
		t.Errorf("Expected 2 images in ProductImages, but got %d", len(product.ProductImages))
	}

	if len(product.CompressedImages) != 2 {
		t.Errorf("Expected 2 images in CompressedImages, but got %d", len(product.CompressedImages))
	}
}
