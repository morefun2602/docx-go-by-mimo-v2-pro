package enum

type WdColor int

const (
	WdColorAutomatic   WdColor = 0
	WdColorBlack       WdColor = 0
	WdColorBlue        WdColor = 0xFF0000
	WdColorBrightGreen WdColor = 0x00FF00
	WdColorDarkBlue    WdColor = 0x800000
	WdColorDarkRed     WdColor = 0x000080
	WdColorDarkYellow  WdColor = 0x008080
	WdColorGray25      WdColor = 0xC0C0C0
	WdColorGray50      WdColor = 0x808080
	WdColorGreen       WdColor = 0x008000
	WdColorPink        WdColor = 0xFF00FF
	WdColorRed         WdColor = 0x0000FF
	WdColorTeal        WdColor = 0x808000
	WdColorTurquoise   WdColor = 0xFFFF00
	WdColorViolet      WdColor = 0x800080
	WdColorWhite       WdColor = 0xFFFFFF
	WdColorYellow      WdColor = 0x00FFFF
)

type PictureType int

const (
	PictureTypeBMP PictureType = iota
	PictureTypeGIF
	PictureTypeJPEG
	PictureTypePNG
	PictureTypeTIFF
	PictureTypeEMF
	PictureTypeWMF
	PictureTypeICO
)

func (p PictureType) ContentType() string {
	switch p {
	case PictureTypeBMP:
		return "image/bmp"
	case PictureTypeGIF:
		return "image/gif"
	case PictureTypeJPEG:
		return "image/jpeg"
	case PictureTypePNG:
		return "image/png"
	case PictureTypeTIFF:
		return "image/tiff"
	case PictureTypeEMF:
		return "image/x-emf"
	case PictureTypeWMF:
		return "image/x-wmf"
	case PictureTypeICO:
		return "image/x-icon"
	default:
		return ""
	}
}

func (p PictureType) Extension() string {
	switch p {
	case PictureTypeBMP:
		return ".bmp"
	case PictureTypeGIF:
		return ".gif"
	case PictureTypeJPEG:
		return ".jpg"
	case PictureTypePNG:
		return ".png"
	case PictureTypeTIFF:
		return ".tif"
	case PictureTypeEMF:
		return ".emf"
	case PictureTypeWMF:
		return ".wmf"
	case PictureTypeICO:
		return ".ico"
	default:
		return ""
	}
}

type ShapeType int

const (
	ShapeTypeInline ShapeType = iota
	ShapeTypeAnchor
)

type HorizontalAlignment int

const (
	HorizontalAlignmentLeft HorizontalAlignment = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
	HorizontalAlignmentInside
	HorizontalAlignmentOutside
)

func (h HorizontalAlignment) XmlValue() string {
	switch h {
	case HorizontalAlignmentLeft:
		return "left"
	case HorizontalAlignmentCenter:
		return "center"
	case HorizontalAlignmentRight:
		return "right"
	case HorizontalAlignmentInside:
		return "inside"
	case HorizontalAlignmentOutside:
		return "outside"
	default:
		return ""
	}
}

type VerticalAlignment int

const (
	VerticalAlignmentTop VerticalAlignment = iota
	VerticalAlignmentCenter
	VerticalAlignmentBottom
	VerticalAlignmentInside
	VerticalAlignmentOutside
)

func (v VerticalAlignment) XmlValue() string {
	switch v {
	case VerticalAlignmentTop:
		return "top"
	case VerticalAlignmentCenter:
		return "center"
	case VerticalAlignmentBottom:
		return "bottom"
	case VerticalAlignmentInside:
		return "inside"
	case VerticalAlignmentOutside:
		return "outside"
	default:
		return ""
	}
}

type WrapType int

const (
	WrapTypeNone WrapType = iota
	WrapTypeSquare
	WrapTypeTight
	WrapTypeThrough
	WrapTypeTopAndBottom
)

func (w WrapType) XmlValue() string {
	switch w {
	case WrapTypeNone:
		return "none"
	case WrapTypeSquare:
		return "square"
	case WrapTypeTight:
		return "tight"
	case WrapTypeThrough:
		return "through"
	case WrapTypeTopAndBottom:
		return "topAndBottom"
	default:
		return ""
	}
}

type WdOLEType int

const (
	WdOLETypeEmbed   WdOLEType = 1
	WdOLETypeLink    WdOLEType = 0
	WdOLETypePicture WdOLEType = 4
	WdOLETypeControl WdOLEType = 5
	WdOLETypeAuto    WdOLEType = 2
)

type WdOLEVerb int

const (
	WdOLEVerbPrimary          WdOLEVerb = 0
	WdOLEVerbShow             WdOLEVerb = -1
	WdOLEVerbOpen             WdOLEVerb = -2
	WdOLEVerbHide             WdOLEVerb = -3
	WdOLEVerbUIActivate       WdOLEVerb = -4
	WdOLEVerbInPlaceActivate  WdOLEVerb = -5
	WdOLEVerbDiscardUndoState WdOLEVerb = -6
)
