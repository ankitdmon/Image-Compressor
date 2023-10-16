package utils

import (
	"fmt"

	"net/http"
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