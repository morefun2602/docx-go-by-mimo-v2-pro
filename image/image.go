package image

import (
	"crypto/sha1"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ImageFormat int

const (
	FormatUnknown ImageFormat = iota
	FormatBMP
	FormatGIF
	FormatJPEG
	FormatPNG
	FormatTIFF
)

func (f ImageFormat) String() string {
	switch f {
	case FormatBMP:
		return "bmp"
	case FormatGIF:
		return "gif"
	case FormatJPEG:
		return "jpeg"
	case FormatPNG:
		return "png"
	case FormatTIFF:
		return "tiff"
	default:
		return "unknown"
	}
}

func (f ImageFormat) ContentType() string {
	switch f {
	case FormatBMP:
		return "image/bmp"
	case FormatGIF:
		return "image/gif"
	case FormatJPEG:
		return "image/jpeg"
	case FormatPNG:
		return "image/png"
	case FormatTIFF:
		return "image/tiff"
	default:
		return ""
	}
}

func (f ImageFormat) Extension() string {
	switch f {
	case FormatBMP:
		return "bmp"
	case FormatGIF:
		return "gif"
	case FormatJPEG:
		return "jpg"
	case FormatPNG:
		return "png"
	case FormatTIFF:
		return "tiff"
	default:
		return ""
	}
}

type Image struct {
	content  []byte
	format   ImageFormat
	sha1Hash string
	width    int
	height   int
	dpi      int
}

func New(content []byte) (*Image, error) {
	format := detectFormat(content)
	if format == FormatUnknown {
		return nil, fmt.Errorf("unknown image format")
	}

	img := &Image{
		content: content,
		format:  format,
		dpi:     72,
	}

	hash := sha1.Sum(content)
	img.sha1Hash = fmt.Sprintf("%x", hash)

	if err := img.parseDimensions(); err != nil {
		return nil, err
	}

	return img, nil
}

func FromFile(path string) (*Image, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return New(data)
}

func FromReader(r io.Reader) (*Image, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return New(data)
}

func (img *Image) Content() []byte {
	return img.content
}

func (img *Image) Format() ImageFormat {
	return img.format
}

func (img *Image) ContentType() string {
	return img.format.ContentType()
}

func (img *Image) Extension() string {
	return img.format.Extension()
}

func (img *Image) SHA1() string {
	return img.sha1Hash
}

func (img *Image) Width() int {
	return img.width
}

func (img *Image) Height() int {
	return img.height
}

func (img *Image) DPI() int {
	return img.dpi
}

func (img *Image) WidthEMU() int {
	return int(float64(img.width) / float64(img.dpi) * 914400)
}

func (img *Image) HeightEMU() int {
	return int(float64(img.height) / float64(img.dpi) * 914400)
}

func (img *Image) Filename() string {
	return fmt.Sprintf("image.%s", img.Extension())
}

func (img *Image) parseDimensions() error {
	cfg, _, err := image.DecodeConfig(strings.NewReader(string(img.content)))
	if err != nil {
		return nil
	}
	img.width = cfg.Width
	img.height = cfg.Height
	return nil
}

func detectFormat(data []byte) ImageFormat {
	if len(data) < 4 {
		return FormatUnknown
	}

	if isBMP(data) {
		return FormatBMP
	}
	if isGIF(data) {
		return FormatGIF
	}
	if isJPEG(data) {
		return FormatJPEG
	}
	if isPNG(data) {
		return FormatPNG
	}
	if isTIFF(data) {
		return FormatTIFF
	}

	return FormatUnknown
}

func isBMP(data []byte) bool {
	return len(data) >= 2 && data[0] == 'B' && data[1] == 'M'
}

func isGIF(data []byte) bool {
	if len(data) < 6 {
		return false
	}
	return string(data[:6]) == "GIF87a" || string(data[:6]) == "GIF89a"
}

func isJPEG(data []byte) bool {
	if len(data) < 3 {
		return false
	}
	return data[0] == 0xFF && data[1] == 0xD8
}

func isPNG(data []byte) bool {
	return len(data) >= 8 && data[0] == 0x89 && data[1] == 'P' && data[2] == 'N' && data[3] == 'G'
}

func isTIFF(data []byte) bool {
	if len(data) < 4 {
		return false
	}
	return (data[0] == 'I' && data[1] == 'I' && data[2] == 42 && data[3] == 0) ||
		(data[0] == 'M' && data[1] == 'M' && data[2] == 0 && data[3] == 42)
}

func ContentTypeName(format ImageFormat) string {
	return format.ContentType()
}

func ExtensionForFormat(format ImageFormat) string {
	return format.Extension()
}

func FormatFromExtension(ext string) ImageFormat {
	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	switch ext {
	case "bmp":
		return FormatBMP
	case "gif":
		return FormatGIF
	case "jpg", "jpeg":
		return FormatJPEG
	case "png":
		return FormatPNG
	case "tif", "tiff":
		return FormatTIFF
	default:
		return FormatUnknown
	}
}

func FormatFromContentType(ct string) ImageFormat {
	ct = strings.ToLower(ct)
	switch {
	case strings.Contains(ct, "bmp"):
		return FormatBMP
	case strings.Contains(ct, "gif"):
		return FormatGIF
	case strings.Contains(ct, "jpeg"):
		return FormatJPEG
	case strings.Contains(ct, "png"):
		return FormatPNG
	case strings.Contains(ct, "tiff"):
		return FormatTIFF
	default:
		return FormatUnknown
	}
}

func ImageFilename(base string, format ImageFormat) string {
	ext := filepath.Ext(base)
	if ext == "" {
		return fmt.Sprintf("%s.%s", base, format.Extension())
	}
	return base
}
