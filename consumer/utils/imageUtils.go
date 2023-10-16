package imageUtils

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"path/filepath"

	"strings"

	"github.com/ankitdmon/producer/models"
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
)

func DownloadImage(imgURL string) (image.Image, error) {

	resp, err := http.Get(imgURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get image from URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image, status code: %d", resp.StatusCode)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	logrus.Info("Image downloaded successfully from URL")
	return img, nil
}

func ResizeAndCompressImage(img image.Image, quality int) ([]byte, error) {
	imgResized := resize.Resize(1024, 0, img, resize.Lanczos3)

	buffer := new(strings.Builder)
	err := jpeg.Encode(buffer, imgResized, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, fmt.Errorf("failed to resize and compress image: %v", err)
	}
	logrus.Info("Image resized and compressed successfully")
	return []byte(buffer.String()), nil
}

func SaveImageToLocal(filename string, data []byte, dir string) (string, error) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	filePath := filepath.Join(dir, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create image file: %v", err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return "", fmt.Errorf("failed to write data to image file: %v", err)
	}

	logrus.Infof("Image saved to file %s", filePath)
	return filePath, nil
}

func CompressAndUpdateImageInDB(productID int, quality int) error {
	images, err := models.GetProductImagesByProductID(productID)
	if err != nil {
		return fmt.Errorf("failed to get product images from db: %v", err)
	}

	for i, imgURL := range images {

		img, err := DownloadImage(imgURL)
		if err != nil {
			return err
		}

		compressedImage, err := ResizeAndCompressImage(img, quality)
		if err != nil {
			return err
		}

		filename := fmt.Sprintf("compressed_image_%d_%d.jpg", productID, i+1)
		filepath, err := SaveImageToLocal(filename, compressedImage, "compressedImages")
		if err != nil {
			return fmt.Errorf("failed to save compressed image to file: %v", err)
		}

		err = models.UpdateProductImage(productID, filepath)
		if err != nil {
			return err
		}

	}

	return nil
}