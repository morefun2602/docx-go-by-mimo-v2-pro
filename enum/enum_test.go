package enum_test

import (
	"testing"

	"github.com/docx-go/enum"
)

func TestParagraphAlignmentXmlValue(t *testing.T) {
	tests := []struct {
		align    enum.ParagraphAlignment
		expected string
	}{
		{enum.ParagraphAlignmentLeft, "left"},
		{enum.ParagraphAlignmentCenter, "center"},
		{enum.ParagraphAlignmentRight, "right"},
		{enum.ParagraphAlignmentJustify, "both"},
		{enum.ParagraphAlignmentDistribute, "distribute"},
	}

	for _, tt := range tests {
		if got := tt.align.XmlValue(); got != tt.expected {
			t.Errorf("ParagraphAlignment(%d).XmlValue() = %s, want %s", tt.align, got, tt.expected)
		}
	}
}

func TestColorIndexXmlValue(t *testing.T) {
	tests := []struct {
		color    enum.ColorIndex
		expected string
	}{
		{enum.ColorIndexAuto, "default"},
		{enum.ColorIndexBlack, "black"},
		{enum.ColorIndexBlue, "blue"},
		{enum.ColorIndexRed, "red"},
		{enum.ColorIndexGreen, "darkGreen"},
	}

	for _, tt := range tests {
		if got := tt.color.XmlValue(); got != tt.expected {
			t.Errorf("ColorIndex(%d).XmlValue() = %s, want %s", tt.color, got, tt.expected)
		}
	}
}

func TestUnderlineXmlValue(t *testing.T) {
	tests := []struct {
		underline enum.Underline
		expected  string
	}{
		{enum.UnderlineNone, "none"},
		{enum.UnderlineSingle, "single"},
		{enum.UnderlineDouble, "double"},
		{enum.UnderlineDotted, "dotted"},
		{enum.UnderlineWavy, "wave"},
	}

	for _, tt := range tests {
		if got := tt.underline.XmlValue(); got != tt.expected {
			t.Errorf("Underline(%d).XmlValue() = %s, want %s", tt.underline, got, tt.expected)
		}
	}
}

func TestOrientationXmlValue(t *testing.T) {
	tests := []struct {
		orient   enum.Orientation
		expected string
	}{
		{enum.OrientationPortrait, "portrait"},
		{enum.OrientationLandscape, "landscape"},
	}

	for _, tt := range tests {
		if got := tt.orient.XmlValue(); got != tt.expected {
			t.Errorf("Orientation(%d).XmlValue() = %s, want %s", tt.orient, got, tt.expected)
		}
	}
}

func TestSectionStartXmlValue(t *testing.T) {
	tests := []struct {
		start    enum.SectionStart
		expected string
	}{
		{enum.SectionStartContinuous, "continuous"},
		{enum.SectionStartNewPage, "nextPage"},
		{enum.SectionStartEvenPage, "evenPage"},
		{enum.SectionStartOddPage, "oddPage"},
	}

	for _, tt := range tests {
		if got := tt.start.XmlValue(); got != tt.expected {
			t.Errorf("SectionStart(%d).XmlValue() = %s, want %s", tt.start, got, tt.expected)
		}
	}
}

func TestCellVerticalAlignmentXmlValue(t *testing.T) {
	tests := []struct {
		align    enum.CellVerticalAlignment
		expected string
	}{
		{enum.CellVerticalAlignmentTop, "top"},
		{enum.CellVerticalAlignmentCenter, "center"},
		{enum.CellVerticalAlignmentBottom, "bottom"},
	}

	for _, tt := range tests {
		if got := tt.align.XmlValue(); got != tt.expected {
			t.Errorf("CellVerticalAlignment(%d).XmlValue() = %s, want %s", tt.align, got, tt.expected)
		}
	}
}

func TestTableAlignmentXmlValue(t *testing.T) {
	tests := []struct {
		align    enum.TableAlignment
		expected string
	}{
		{enum.TableAlignmentLeft, "left"},
		{enum.TableAlignmentCenter, "center"},
		{enum.TableAlignmentRight, "right"},
	}

	for _, tt := range tests {
		if got := tt.align.XmlValue(); got != tt.expected {
			t.Errorf("TableAlignment(%d).XmlValue() = %s, want %s", tt.align, got, tt.expected)
		}
	}
}

func TestTabAlignmentXmlValue(t *testing.T) {
	tests := []struct {
		tab      enum.TabAlignment
		expected string
	}{
		{enum.TabAlignmentLeft, "left"},
		{enum.TabAlignmentCenter, "center"},
		{enum.TabAlignmentRight, "right"},
		{enum.TabAlignmentDecimal, "decimal"},
	}

	for _, tt := range tests {
		if got := tt.tab.XmlValue(); got != tt.expected {
			t.Errorf("TabAlignment(%d).XmlValue() = %s, want %s", tt.tab, got, tt.expected)
		}
	}
}

func TestTabLeaderXmlValue(t *testing.T) {
	tests := []struct {
		leader   enum.TabLeader
		expected string
	}{
		{enum.TabLeaderSpaces, "none"},
		{enum.TabLeaderDots, "dot"},
		{enum.TabLeaderDashes, "hyphen"},
		{enum.TabLeaderLines, "underscore"},
	}

	for _, tt := range tests {
		if got := tt.leader.XmlValue(); got != tt.expected {
			t.Errorf("TabLeader(%d).XmlValue() = %s, want %s", tt.leader, got, tt.expected)
		}
	}
}

func TestLineSpacingXmlValue(t *testing.T) {
	tests := []struct {
		spacing  enum.LineSpacing
		expected string
	}{
		{enum.LineSpacingAtLeast, "atLeast"},
		{enum.LineSpacingExactly, "exact"},
		{enum.LineSpacingMultiple, "auto"},
	}

	for _, tt := range tests {
		if got := tt.spacing.XmlValue(); got != tt.expected {
			t.Errorf("LineSpacing(%d).XmlValue() = %s, want %s", tt.spacing, got, tt.expected)
		}
	}
}
