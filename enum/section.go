package enum

type HeaderFooterIndex int

const (
	HeaderFooterIndexPrimary   HeaderFooterIndex = 1
	HeaderFooterIndexFirstPage HeaderFooterIndex = 2
	HeaderFooterIndexEvenPage  HeaderFooterIndex = 3
)

func (h HeaderFooterIndex) XmlValue() string {
	switch h {
	case HeaderFooterIndexPrimary:
		return "default"
	case HeaderFooterIndexFirstPage:
		return "first"
	case HeaderFooterIndexEvenPage:
		return "even"
	default:
		return ""
	}
}

type Orientation int

const (
	OrientationPortrait  Orientation = 0
	OrientationLandscape Orientation = 1
)

func (o Orientation) XmlValue() string {
	switch o {
	case OrientationPortrait:
		return "portrait"
	case OrientationLandscape:
		return "landscape"
	default:
		return ""
	}
}

type SectionStart int

const (
	SectionStartContinuous SectionStart = 0
	SectionStartNewColumn  SectionStart = 1
	SectionStartNewPage    SectionStart = 2
	SectionStartEvenPage   SectionStart = 3
	SectionStartOddPage    SectionStart = 4
)

func (s SectionStart) XmlValue() string {
	switch s {
	case SectionStartContinuous:
		return "continuous"
	case SectionStartNewColumn:
		return "nextColumn"
	case SectionStartNewPage:
		return "nextPage"
	case SectionStartEvenPage:
		return "evenPage"
	case SectionStartOddPage:
		return "oddPage"
	default:
		return ""
	}
}
