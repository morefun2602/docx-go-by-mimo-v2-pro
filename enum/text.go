package enum

type ParagraphAlignment int

const (
	ParagraphAlignmentLeft        ParagraphAlignment = 0
	ParagraphAlignmentCenter      ParagraphAlignment = 1
	ParagraphAlignmentRight       ParagraphAlignment = 2
	ParagraphAlignmentJustify     ParagraphAlignment = 3
	ParagraphAlignmentDistribute  ParagraphAlignment = 4
	ParagraphAlignmentJustifyMed  ParagraphAlignment = 5
	ParagraphAlignmentJustifyHi   ParagraphAlignment = 7
	ParagraphAlignmentJustifyLow  ParagraphAlignment = 8
	ParagraphAlignmentThaiJustify ParagraphAlignment = 9
)

func (a ParagraphAlignment) XmlValue() string {
	switch a {
	case ParagraphAlignmentLeft:
		return "left"
	case ParagraphAlignmentCenter:
		return "center"
	case ParagraphAlignmentRight:
		return "right"
	case ParagraphAlignmentJustify:
		return "both"
	case ParagraphAlignmentDistribute:
		return "distribute"
	case ParagraphAlignmentJustifyMed:
		return "mediumKashida"
	case ParagraphAlignmentJustifyHi:
		return "highKashida"
	case ParagraphAlignmentJustifyLow:
		return "lowKashida"
	case ParagraphAlignmentThaiJustify:
		return "thaiDistribute"
	default:
		return ""
	}
}

type BreakType int

const (
	BreakTypeColumn            BreakType = 8
	BreakTypeLine              BreakType = 6
	BreakTypeLineClearLeft     BreakType = 9
	BreakTypeLineClearRight    BreakType = 10
	BreakTypeLineClearAll      BreakType = 11
	BreakTypePage              BreakType = 7
	BreakTypeSectionContinuous BreakType = 3
	BreakTypeSectionEvenPage   BreakType = 4
	BreakTypeSectionNextPage   BreakType = 2
	BreakTypeSectionOddPage    BreakType = 5
	BreakTypeTextWrapping      BreakType = 11
)

type ColorIndex int

const (
	ColorIndexInherited   ColorIndex = -1
	ColorIndexAuto        ColorIndex = 0
	ColorIndexBlack       ColorIndex = 1
	ColorIndexBlue        ColorIndex = 2
	ColorIndexBrightGreen ColorIndex = 4
	ColorIndexDarkBlue    ColorIndex = 9
	ColorIndexDarkRed     ColorIndex = 13
	ColorIndexDarkYellow  ColorIndex = 14
	ColorIndexGray25      ColorIndex = 16
	ColorIndexGray50      ColorIndex = 15
	ColorIndexGreen       ColorIndex = 11
	ColorIndexPink        ColorIndex = 5
	ColorIndexRed         ColorIndex = 6
	ColorIndexTeal        ColorIndex = 10
	ColorIndexTurquoise   ColorIndex = 3
	ColorIndexViolet      ColorIndex = 12
	ColorIndexWhite       ColorIndex = 8
	ColorIndexYellow      ColorIndex = 7
)

func (c ColorIndex) XmlValue() string {
	switch c {
	case ColorIndexInherited:
		return ""
	case ColorIndexAuto:
		return "default"
	case ColorIndexBlack:
		return "black"
	case ColorIndexBlue:
		return "blue"
	case ColorIndexBrightGreen:
		return "green"
	case ColorIndexDarkBlue:
		return "darkBlue"
	case ColorIndexDarkRed:
		return "darkRed"
	case ColorIndexDarkYellow:
		return "darkYellow"
	case ColorIndexGray25:
		return "lightGray"
	case ColorIndexGray50:
		return "darkGray"
	case ColorIndexGreen:
		return "darkGreen"
	case ColorIndexPink:
		return "magenta"
	case ColorIndexRed:
		return "red"
	case ColorIndexTeal:
		return "darkCyan"
	case ColorIndexTurquoise:
		return "cyan"
	case ColorIndexViolet:
		return "darkMagenta"
	case ColorIndexWhite:
		return "white"
	case ColorIndexYellow:
		return "yellow"
	default:
		return ""
	}
}

type LineSpacing int

const (
	LineSpacingSingle       LineSpacing = 0
	LineSpacingOnePointFive LineSpacing = 1
	LineSpacingDouble       LineSpacing = 2
	LineSpacingAtLeast      LineSpacing = 3
	LineSpacingExactly      LineSpacing = 4
	LineSpacingMultiple     LineSpacing = 5
)

func (l LineSpacing) XmlValue() string {
	switch l {
	case LineSpacingSingle:
		return "UNMAPPED"
	case LineSpacingOnePointFive:
		return "UNMAPPED"
	case LineSpacingDouble:
		return "UNMAPPED"
	case LineSpacingAtLeast:
		return "atLeast"
	case LineSpacingExactly:
		return "exact"
	case LineSpacingMultiple:
		return "auto"
	default:
		return ""
	}
}

type TabAlignment int

const (
	TabAlignmentLeft    TabAlignment = 0
	TabAlignmentCenter  TabAlignment = 1
	TabAlignmentRight   TabAlignment = 2
	TabAlignmentDecimal TabAlignment = 3
	TabAlignmentBar     TabAlignment = 4
	TabAlignmentList    TabAlignment = 6
	TabAlignmentClear   TabAlignment = 101
	TabAlignmentEnd     TabAlignment = 102
	TabAlignmentNum     TabAlignment = 103
	TabAlignmentStart   TabAlignment = 104
)

func (t TabAlignment) XmlValue() string {
	switch t {
	case TabAlignmentLeft:
		return "left"
	case TabAlignmentCenter:
		return "center"
	case TabAlignmentRight:
		return "right"
	case TabAlignmentDecimal:
		return "decimal"
	case TabAlignmentBar:
		return "bar"
	case TabAlignmentList:
		return "list"
	case TabAlignmentClear:
		return "clear"
	case TabAlignmentEnd:
		return "end"
	case TabAlignmentNum:
		return "num"
	case TabAlignmentStart:
		return "start"
	default:
		return ""
	}
}

type TabLeader int

const (
	TabLeaderSpaces    TabLeader = 0
	TabLeaderDots      TabLeader = 1
	TabLeaderDashes    TabLeader = 2
	TabLeaderLines     TabLeader = 3
	TabLeaderHeavy     TabLeader = 4
	TabLeaderMiddleDot TabLeader = 5
)

func (t TabLeader) XmlValue() string {
	switch t {
	case TabLeaderSpaces:
		return "none"
	case TabLeaderDots:
		return "dot"
	case TabLeaderDashes:
		return "hyphen"
	case TabLeaderLines:
		return "underscore"
	case TabLeaderHeavy:
		return "heavy"
	case TabLeaderMiddleDot:
		return "middleDot"
	default:
		return ""
	}
}

type Underline int

const (
	UnderlineInherited       Underline = -1
	UnderlineNone            Underline = 0
	UnderlineSingle          Underline = 1
	UnderlineWords           Underline = 2
	UnderlineDouble          Underline = 3
	UnderlineDotted          Underline = 4
	UnderlineThick           Underline = 6
	UnderlineDash            Underline = 7
	UnderlineDotDash         Underline = 9
	UnderlineDotDotDash      Underline = 10
	UnderlineWavy            Underline = 11
	UnderlineDottedHeavy     Underline = 20
	UnderlineDashHeavy       Underline = 23
	UnderlineDotDashHeavy    Underline = 25
	UnderlineDotDotDashHeavy Underline = 26
	UnderlineWavyHeavy       Underline = 27
	UnderlineDashLong        Underline = 39
	UnderlineWavyDouble      Underline = 43
	UnderlineDashLongHeavy   Underline = 55
)

func (u Underline) XmlValue() string {
	switch u {
	case UnderlineInherited:
		return ""
	case UnderlineNone:
		return "none"
	case UnderlineSingle:
		return "single"
	case UnderlineWords:
		return "words"
	case UnderlineDouble:
		return "double"
	case UnderlineDotted:
		return "dotted"
	case UnderlineThick:
		return "thick"
	case UnderlineDash:
		return "dash"
	case UnderlineDotDash:
		return "dotDash"
	case UnderlineDotDotDash:
		return "dotDotDash"
	case UnderlineWavy:
		return "wave"
	case UnderlineDottedHeavy:
		return "dottedHeavy"
	case UnderlineDashHeavy:
		return "dashedHeavy"
	case UnderlineDotDashHeavy:
		return "dashDotHeavy"
	case UnderlineDotDotDashHeavy:
		return "dashDotDotHeavy"
	case UnderlineWavyHeavy:
		return "wavyHeavy"
	case UnderlineDashLong:
		return "dashLong"
	case UnderlineWavyDouble:
		return "wavyDouble"
	case UnderlineDashLongHeavy:
		return "dashLongHeavy"
	default:
		return ""
	}
}
