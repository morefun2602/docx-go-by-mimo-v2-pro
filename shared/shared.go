package shared

const (
	emusPerInch = 914400
	emusPerCM   = 360000
	emusPerMM   = 36000
	emusPerPt   = 12700
	emusPerTwip = 635
)

type Length int64

func (l Length) CM() float64 {
	return float64(l) / float64(emusPerCM)
}

func (l Length) EMU() int64 {
	return int64(l)
}

func (l Length) Inches() float64 {
	return float64(l) / float64(emusPerInch)
}

func (l Length) MM() float64 {
	return float64(l) / float64(emusPerMM)
}

func (l Length) Pt() float64 {
	return float64(l) / float64(emusPerPt)
}

func (l Length) Twips() int64 {
	return int64(float64(l)/float64(emusPerTwip) + 0.5)
}

func Inches(inches float64) Length {
	return Length(int64(inches * float64(emusPerInch)))
}

func CM(cm float64) Length {
	return Length(int64(cm * float64(emusPerCM)))
}

func EMU(emu int64) Length {
	return Length(emu)
}

func MM(mm float64) Length {
	return Length(int64(mm * float64(emusPerMM)))
}

func Pt(points float64) Length {
	return Length(int64(points * float64(emusPerPt)))
}

func Twips(twips float64) Length {
	return Length(int64(twips * float64(emusPerTwip)))
}

type RGBColor struct {
	R, G, B uint8
}

func NewRGBColor(r, g, b uint8) RGBColor {
	return RGBColor{R: r, G: g, B: b}
}

func (c RGBColor) String() string {
	return string([]byte{
		hexChar(c.R >> 4), hexChar(c.R & 0xf),
		hexChar(c.G >> 4), hexChar(c.G & 0xf),
		hexChar(c.B >> 4), hexChar(c.B & 0xf),
	})
}

func hexChar(b byte) byte {
	if b < 10 {
		return '0' + b
	}
	return 'a' + (b - 10)
}

func RGBColorFromString(hex string) (RGBColor, error) {
	if len(hex) != 6 {
		return RGBColor{}, ErrInvalidRGBColor(hex)
	}
	r, err := parseHexByte(hex[0:2])
	if err != nil {
		return RGBColor{}, err
	}
	g, err := parseHexByte(hex[2:4])
	if err != nil {
		return RGBColor{}, err
	}
	b, err := parseHexByte(hex[4:6])
	if err != nil {
		return RGBColor{}, err
	}
	return RGBColor{R: r, G: g, B: b}, nil
}

func parseHexByte(s string) (byte, error) {
	var result byte
	for i := 0; i < 2; i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			result = result*16 + (c - '0')
		case c >= 'a' && c <= 'f':
			result = result*16 + (c - 'a' + 10)
		case c >= 'A' && c <= 'F':
			result = result*16 + (c - 'A' + 10)
		default:
			return 0, ErrInvalidRGBColor(s)
		}
	}
	return result, nil
}

type ErrInvalidRGBColor string

func (e ErrInvalidRGBColor) Error() string {
	return "invalid RGB color: " + string(e)
}
