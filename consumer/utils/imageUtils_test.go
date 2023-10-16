package imageUtils_test

import (
	"testing"

	imageUtils "github.com/ankitdmon/consumer/utils"
)

const (
	imgURL = "https://c4.wallpaperflare.com/wallpaper/41/681/303/pc-hd-1080p-nature-1920x1080-wallpaper-preview.jpg"
)

func BenchmarkDownloadImage(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := imageUtils.DownloadImage(imgURL)
		if err != nil {
			b.Fatalf("Error downloading image: %v", err)
		}
	}
}

func BenchmarkResizeAndCompressImage(b *testing.B) {
	img, err := imageUtils.DownloadImage(imgURL)
	if err != nil {
		b.Fatalf("Error downloading image: %v", err)
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := imageUtils.ResizeAndCompressImage(img, 80)
		if err != nil {
			b.Fatalf("Error resizing and compressing image: %v", err)
		}
	}
}

func BenchmarkSaveImageToLocal(b *testing.B) {
	img, err := imageUtils.DownloadImage(imgURL)
	if err != nil {
		b.Fatalf("Error downloading image: %v", err)
	}
	compressedImage, err := imageUtils.ResizeAndCompressImage(img, 80)
	if err != nil {
		b.Fatalf("Error resizing and compressing image: %v", err)
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := imageUtils.SaveImageToLocal("test_image.jpg", compressedImage, "test_dir")
		if err != nil {
			b.Fatalf("Error saving image to local: %v", err)
		}
	}
}
