package image

const (
	EmusPerInch = 914400
	EmusPerCm   = 360000
	EmusPerPx   = 9525
)

var (
	MimeBMP  = "image/bmp"
	MimeGIF  = "image/gif"
	MimeJPEG = "image/jpeg"
	MimePNG  = "image/png"
	MimeTIFF = "image/tiff"
	MimeEMF  = "image/x-emf"
	MimeWMF  = "image/x-wmf"
	MimeICO  = "image/x-icon"
)

var ImageSignatures = map[string][]byte{
	"bmp":   {'B', 'M'},
	"gif":   []byte("GIF"),
	"jpeg":  {0xFF, 0xD8},
	"png":   {0x89, 0x50, 0x4E, 0x47},
	"tiff":  {'I', 'I'},
	"tiff2": {'M', 'M'},
}

var ContentTypeToExt = map[string]string{
	"image/bmp":    "bmp",
	"image/gif":    "gif",
	"image/jpeg":   "jpg",
	"image/png":    "png",
	"image/tiff":   "tif",
	"image/x-emf":  "emf",
	"image/x-wmf":  "wmf",
	"image/x-icon": "ico",
}

var ExtToContentType = map[string]string{
	"bmp":  "image/bmp",
	"gif":  "image/gif",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"tif":  "image/tiff",
	"tiff": "image/tiff",
	"emf":  "image/x-emf",
	"wmf":  "image/x-wmf",
	"ico":  "image/x-icon",
}
