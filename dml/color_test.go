package dml_test

import (
	"testing"

	"github.com/docx-go/dml"
	"github.com/docx-go/oxml"
)

func TestNewColor(t *testing.T) {
	elem := oxml.NewElement("w:color")
	elem.SetAttr("val", "FF0000")

	color := dml.NewColor(elem)
	if color.Val() != "FF0000" {
		t.Errorf("Val() = %s, want FF0000", color.Val())
	}
}

func TestNewColorFromHex(t *testing.T) {
	color := dml.NewColorFromHex("#00FF00")
	if color.Val() != "00FF00" {
		t.Errorf("Val() = %s, want 00FF00", color.Val())
	}
}

func TestColorSetVal(t *testing.T) {
	elem := oxml.NewElement("w:color")
	color := dml.NewColor(elem)

	color.SetVal("0000FF")
	if color.Val() != "0000FF" {
		t.Errorf("Val() = %s, want 0000FF", color.Val())
	}
}

func TestColorRGB(t *testing.T) {
	elem := oxml.NewElement("w:color")
	elem.SetAttr("val", "FF8000")

	color := dml.NewColor(elem)
	r, g, b, err := color.RGB()
	if err != nil {
		t.Fatalf("RGB() error = %v", err)
	}
	if r != 0xFF {
		t.Errorf("R = %d, want 255", r)
	}
	if g != 0x80 {
		t.Errorf("G = %d, want 128", g)
	}
	if b != 0x00 {
		t.Errorf("B = %d, want 0", b)
	}
}

func TestColorSetRGB(t *testing.T) {
	elem := oxml.NewElement("w:color")
	color := dml.NewColor(elem)

	color.SetRGB(0xFF, 0x80, 0x00)
	if color.Val() != "ff8000" {
		t.Errorf("Val() = %s, want ff8000", color.Val())
	}
}

func TestColorRGBInvalid(t *testing.T) {
	elem := oxml.NewElement("w:color")
	elem.SetAttr("val", "XYZ")

	color := dml.NewColor(elem)
	_, _, _, err := color.RGB()
	if err == nil {
		t.Error("RGB() expected error for invalid hex")
	}
}

func TestNewNoFill(t *testing.T) {
	nf := dml.NewNoFill()
	if nf.Element().Tag() != "a:noFill" {
		t.Errorf("Tag() = %s, want a:noFill", nf.Element().Tag())
	}
}

func TestNewSolidFill(t *testing.T) {
	colorElem := oxml.NewElement("a:srgbClr")
	colorElem.SetAttr("val", "FF0000")
	color := dml.NewColor(colorElem)

	sf := dml.NewSolidFill(color)
	if sf.Element().Tag() != "a:solidFill" {
		t.Errorf("Tag() = %s, want a:solidFill", sf.Element().Tag())
	}
	if sf.Color() != color {
		t.Error("Color() returned different color")
	}
}

func TestThemeColor(t *testing.T) {
	elem := oxml.NewElement("a:themeColor")
	elem.SetAttr("val", "accent1")

	tc := dml.NewThemeColor(elem)
	if tc.Val() != "accent1" {
		t.Errorf("Val() = %s, want accent1", tc.Val())
	}

	tc.SetVal("accent2")
	if tc.Val() != "accent2" {
		t.Errorf("Val() = %s, want accent2", tc.Val())
	}
}
