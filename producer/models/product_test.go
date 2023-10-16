package models

import (
	"testing"
)

func TestProductModel(t *testing.T) {
	product := Product{
		UserID:                  1,
		ProductName:             "Sample Product",
		ProductDescription:      "This is a sample product description.",
		ProductImages:           "test image",
		ProductPrice:            19.99,
		CompressedProductImages: "compressed1.jpg,compressed2.jpg",
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

	if product.ProductImages != "test image" {
		t.Errorf("Expected ProductImages to be 'test image', got '%s'", product.ProductImages)
	}
}

func BenchmarkCreateProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CreateProduct(1, "test product", "test description", "test image", 10.0)
		if err != nil {
			b.Errorf("Error creating product: %s", err.Error())
		}
	}
}

func BenchmarkGetProductId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetProductId()
		if err != nil {
			b.Errorf("Error getting product ID: %s", err.Error())
		}
	}
}

func BenchmarkGetProductImagesByProductID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetProductImagesByProductID(7)
		if err != nil {
			b.Errorf("Error getting product images: %s", err.Error())
		}
	}
}
