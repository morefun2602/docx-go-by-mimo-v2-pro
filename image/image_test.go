package image_test

import (
	"testing"

	"github.com/docx-go/image"
)

func TestFormatFromExtension(t *testing.T) {
	tests := []struct {
		ext      string
		expected image.ImageFormat
	}{
		{"bmp", image.FormatBMP},
		{"gif", image.FormatGIF},
		{"jpg", image.FormatJPEG},
		{"jpeg", image.FormatJPEG},
		{"png", image.FormatPNG},
		{"tif", image.FormatTIFF},
		{"tiff", image.FormatTIFF},
		{"xyz", image.FormatUnknown},
	}

	for _, tt := range tests {
		if got := image.FormatFromExtension(tt.ext); got != tt.expected {
			t.Errorf("FormatFromExtension(%s) = %v, want %v", tt.ext, got, tt.expected)
		}
	}
}

func TestFormatFromContentType(t *testing.T) {
	tests := []struct {
		ct       string
		expected image.ImageFormat
	}{
		{"image/bmp", image.FormatBMP},
		{"image/gif", image.FormatGIF},
		{"image/jpeg", image.FormatJPEG},
		{"image/png", image.FormatPNG},
		{"image/tiff", image.FormatTIFF},
	}

	for _, tt := range tests {
		if got := image.FormatFromContentType(tt.ct); got != tt.expected {
			t.Errorf("FormatFromContentType(%s) = %v, want %v", tt.ct, got, tt.expected)
		}
	}
}

func TestImageFormatString(t *testing.T) {
	tests := []struct {
		format   image.ImageFormat
		expected string
	}{
		{image.FormatBMP, "bmp"},
		{image.FormatGIF, "gif"},
		{image.FormatJPEG, "jpeg"},
		{image.FormatPNG, "png"},
		{image.FormatTIFF, "tiff"},
		{image.FormatUnknown, "unknown"},
	}

	for _, tt := range tests {
		if got := tt.format.String(); got != tt.expected {
			t.Errorf("ImageFormat(%d).String() = %s, want %s", tt.format, got, tt.expected)
		}
	}
}

func TestImageFormatContentType(t *testing.T) {
	tests := []struct {
		format   image.ImageFormat
		expected string
	}{
		{image.FormatBMP, "image/bmp"},
		{image.FormatGIF, "image/gif"},
		{image.FormatJPEG, "image/jpeg"},
		{image.FormatPNG, "image/png"},
		{image.FormatTIFF, "image/tiff"},
	}

	for _, tt := range tests {
		if got := tt.format.ContentType(); got != tt.expected {
			t.Errorf("ImageFormat(%d).ContentType() = %s, want %s", tt.format, got, tt.expected)
		}
	}
}

func TestImageFormatExtension(t *testing.T) {
	tests := []struct {
		format   image.ImageFormat
		expected string
	}{
		{image.FormatBMP, "bmp"},
		{image.FormatGIF, "gif"},
		{image.FormatJPEG, "jpg"},
		{image.FormatPNG, "png"},
		{image.FormatTIFF, "tiff"},
	}

	for _, tt := range tests {
		if got := tt.format.Extension(); got != tt.expected {
			t.Errorf("ImageFormat(%d).Extension() = %s, want %s", tt.format, got, tt.expected)
		}
	}
}

func TestImageFilename(t *testing.T) {
	tests := []struct {
		base     string
		format   image.ImageFormat
		expected string
	}{
		{"image", image.FormatPNG, "image.png"},
		{"photo", image.FormatJPEG, "photo.jpg"},
		{"icon.gif", image.FormatGIF, "icon.gif"},
	}

	for _, tt := range tests {
		if got := image.ImageFilename(tt.base, tt.format); got != tt.expected {
			t.Errorf("ImageFilename(%s, %v) = %s, want %s", tt.base, tt.format, got, tt.expected)
		}
	}
}

func TestNewImageFromPNG(t *testing.T) {
	pngHeader := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	img, err := image.New(pngHeader)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if img.Format() != image.FormatPNG {
		t.Errorf("Format() = %v, want %v", img.Format(), image.FormatPNG)
	}
}

func TestNewImageFromJPEG(t *testing.T) {
	jpegHeader := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	img, err := image.New(jpegHeader)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if img.Format() != image.FormatJPEG {
		t.Errorf("Format() = %v, want %v", img.Format(), image.FormatJPEG)
	}
}

func TestNewImageFromGIF(t *testing.T) {
	gifHeader := []byte("GIF89a")
	img, err := image.New(gifHeader)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if img.Format() != image.FormatGIF {
		t.Errorf("Format() = %v, want %v", img.Format(), image.FormatGIF)
	}
}

func TestNewImageFromBMP(t *testing.T) {
	bmpHeader := []byte{'B', 'M', 0, 0, 0, 0}
	img, err := image.New(bmpHeader)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if img.Format() != image.FormatBMP {
		t.Errorf("Format() = %v, want %v", img.Format(), image.FormatBMP)
	}
}

func TestNewImageFromTIFF(t *testing.T) {
	tiffHeader := []byte{'I', 'I', 42, 0, 0, 0, 0, 0}
	img, err := image.New(tiffHeader)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if img.Format() != image.FormatTIFF {
		t.Errorf("Format() = %v, want %v", img.Format(), image.FormatTIFF)
	}
}

func TestNewImageUnknownFormat(t *testing.T) {
	_, err := image.New([]byte{0x00, 0x01, 0x02})
	if err == nil {
		t.Error("New() expected error for unknown format")
	}
}

func TestImageSHA1(t *testing.T) {
	pngHeader := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	img, _ := image.New(pngHeader)
	if img.SHA1() == "" {
		t.Error("SHA1() returned empty string")
	}
}
