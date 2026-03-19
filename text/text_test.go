package text_test

import (
	"testing"

	"github.com/docx-go/oxml"
	"github.com/docx-go/text"
)

func TestParagraphAddRun(t *testing.T) {
	p := oxml.NewElement("w:p")
	para := text.NewParagraph(p, nil)

	run := para.AddRun("Hello", "")
	if run == nil {
		t.Fatal("AddRun returned nil")
	}

	if run.Text() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", run.Text())
	}
}

func TestParagraphStyle(t *testing.T) {
	p := oxml.NewElement("w:p")
	para := text.NewParagraph(p, nil)

	para.SetStyle("Heading 1")
	if para.Style() != "Heading 1" {
		t.Errorf("expected 'Heading 1', got '%s'", para.Style())
	}
}

func TestParagraphAlignment(t *testing.T) {
	p := oxml.NewElement("w:p")
	para := text.NewParagraph(p, nil)

	para.SetAlignment("center")
	if para.Alignment() != "center" {
		t.Errorf("expected 'center', got '%s'", para.Alignment())
	}
}

func TestRunBold(t *testing.T) {
	r := oxml.NewElement("w:r")
	run := text.NewRun(r, nil)

	run.SetBold(true)
	if !run.Bold() {
		t.Error("expected bold to be true")
	}

	run.SetBold(false)
	if run.Bold() {
		t.Error("expected bold to be false")
	}
}

func TestRunItalic(t *testing.T) {
	r := oxml.NewElement("w:r")
	run := text.NewRun(r, nil)

	run.SetItalic(true)
	if !run.Italic() {
		t.Error("expected italic to be true")
	}

	run.SetItalic(false)
	if run.Italic() {
		t.Error("expected italic to be false")
	}
}

func TestRunText(t *testing.T) {
	r := oxml.NewElement("w:r")
	run := text.NewRun(r, nil)

	run.SetText("Hello World")
	if run.Text() != "Hello World" {
		t.Errorf("expected 'Hello World', got '%s'", run.Text())
	}
}

func TestRunStyle(t *testing.T) {
	r := oxml.NewElement("w:r")
	run := text.NewRun(r, nil)

	run.SetStyle("Strong")
	if run.Style() != "Strong" {
		t.Errorf("expected 'Strong', got '%s'", run.Style())
	}
}

func TestRunUnderline(t *testing.T) {
	r := oxml.NewElement("w:r")
	run := text.NewRun(r, nil)

	run.SetUnderline("single")
	if run.Underline() != "single" {
		t.Errorf("expected 'single', got '%s'", run.Underline())
	}
}

func TestRunAddBreak(t *testing.T) {
	r := oxml.NewElement("w:r")
	run := text.NewRun(r, nil)

	run.AddBreak(text.BreakPage)
	if len(r.Children()) != 1 {
		t.Errorf("expected 1 child, got %d", len(r.Children()))
	}
}

func TestRunAddTab(t *testing.T) {
	r := oxml.NewElement("w:r")
	run := text.NewRun(r, nil)

	run.AddTab()
	if len(r.Children()) != 1 {
		t.Errorf("expected 1 child, got %d", len(r.Children()))
	}
}

func TestFontSize(t *testing.T) {
	rPr := oxml.NewElement("w:rPr")
	font := text.NewFont(rPr)

	font.SetSize(12)
	if font.Size() != 24 {
		t.Errorf("expected 24, got %d", font.Size())
	}
}

func TestFontName(t *testing.T) {
	rPr := oxml.NewElement("w:rPr")
	font := text.NewFont(rPr)

	font.SetName("Arial")
	if font.Name() != "Arial" {
		t.Errorf("expected 'Arial', got '%s'", font.Name())
	}
}

func TestFontBold(t *testing.T) {
	rPr := oxml.NewElement("w:rPr")
	font := text.NewFont(rPr)

	font.SetBold(true)
	if !font.Bold() {
		t.Error("expected bold to be true")
	}
}

func TestFontItalic(t *testing.T) {
	rPr := oxml.NewElement("w:rPr")
	font := text.NewFont(rPr)

	font.SetItalic(true)
	if !font.Italic() {
		t.Error("expected italic to be true")
	}
}

func TestFontColor(t *testing.T) {
	rPr := oxml.NewElement("w:rPr")
	font := text.NewFont(rPr)

	font.SetColor("FF0000")
	if font.Color() != "FF0000" {
		t.Errorf("expected 'FF0000', got '%s'", font.Color())
	}
}
