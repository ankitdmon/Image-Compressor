package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ankitdmon/producer/routes"
	"github.com/gin-gonic/gin"
)

func TestCreateProduct(t *testing.T) {
	r := gin.New()
	routes.SetupProductRoutes(r)

	product := map[string]interface{}{
		"user_id":                   1,
		"product_name":              "Sample Product",
		"product_description":       "This is a sample product description.",
		"product_images":            []string{"image1.jpg", "image2.jpg"},
		"product_price":             19.99,
		"compressed_product_images": []string{"compressed1.jpg", "compressed2.jpg"},
	}

	productJSON, _ := json.Marshal(product)

	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(productJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestGetProducts(t *testing.T) {
	r := gin.New()
	routes.SetupProductRoutes(r)

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}
