package shared_test

import (
	"testing"

	"github.com/docx-go/shared"
)

func TestLengthInches(t *testing.T) {
	length := shared.Inches(1.0)
	if length != 914400 {
		t.Errorf("expected 914400, got %d", length)
	}
	if length.Inches() != 1.0 {
		t.Errorf("expected 1.0 inches, got %f", length.Inches())
	}
}

func TestLengthCM(t *testing.T) {
	length := shared.CM(1.0)
	if length != 360000 {
		t.Errorf("expected 360000, got %d", length)
	}
	if length.CM() != 1.0 {
		t.Errorf("expected 1.0 cm, got %f", length.CM())
	}
}

func TestLengthMM(t *testing.T) {
	length := shared.MM(10.0)
	if length != 360000 {
		t.Errorf("expected 360000, got %d", length)
	}
	if length.MM() != 10.0 {
		t.Errorf("expected 10.0 mm, got %f", length.MM())
	}
}

func TestLengthPt(t *testing.T) {
	length := shared.Pt(72.0)
	if length != 914400 {
		t.Errorf("expected 914400, got %d", length)
	}
	if length.Pt() != 72.0 {
		t.Errorf("expected 72.0 pt, got %f", length.Pt())
	}
}

func TestLengthTwips(t *testing.T) {
	length := shared.Twips(1440.0)
	if length != 914400 {
		t.Errorf("expected 914400, got %d", length)
	}
	if length.Twips() != 1440 {
		t.Errorf("expected 1440 twips, got %d", length.Twips())
	}
}

func TestRGBColorString(t *testing.T) {
	color := shared.NewRGBColor(0xFF, 0x80, 0x00)
	if color.String() != "ff8000" {
		t.Errorf("expected 'ff8000', got '%s'", color.String())
	}
}

func TestRGBColorFromString(t *testing.T) {
	color, err := shared.RGBColorFromString("ff8000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if color.R != 0xFF {
		t.Errorf("expected R=255, got %d", color.R)
	}
	if color.G != 0x80 {
		t.Errorf("expected G=128, got %d", color.G)
	}
	if color.B != 0x00 {
		t.Errorf("expected B=0, got %d", color.B)
	}
}

func TestRGBColorFromStringInvalid(t *testing.T) {
	_, err := shared.RGBColorFromString("xyz")
	if err == nil {
		t.Error("expected error for invalid hex string")
	}
}
