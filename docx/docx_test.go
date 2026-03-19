package docx_test

import (
	"testing"

	"github.com/docx-go/docx"
	"github.com/docx-go/shared"
)

func TestNewDocument(t *testing.T) {
	doc := docx.New()
	if doc == nil {
		t.Fatal("New() returned nil")
	}
}

func TestAddParagraph(t *testing.T) {
	doc := docx.New()
	para := doc.AddParagraph("Hello, World!", "")
	if para == nil {
		t.Fatal("AddParagraph returned nil")
	}
	if para.Text() != "Hello, World!" {
		t.Errorf("expected 'Hello, World!', got '%s'", para.Text())
	}
}

func TestAddHeading(t *testing.T) {
	doc := docx.New()

	testCases := []struct {
		level int
		style string
	}{
		{0, "Title"},
		{1, "Heading 1"},
		{2, "Heading 2"},
		{9, "Heading 9"},
	}

	for _, tc := range testCases {
		para := doc.AddHeading("Test Heading", tc.level)
		if para == nil {
			t.Fatalf("AddHeading(%d) returned nil", tc.level)
		}
		if para.Style() != tc.style {
			t.Errorf("level %d: expected style '%s', got '%s'", tc.level, tc.style, para.Style())
		}
	}
}

func TestAddHeadingOutOfRange(t *testing.T) {
	doc := docx.New()

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for level -1")
		}
	}()

	doc.AddHeading("Test", -1)
}

func TestAddTable(t *testing.T) {
	doc := docx.New()
	table := doc.AddTable(2, 3, "")
	if table == nil {
		t.Fatal("AddTable returned nil")
	}

	if len(table.Rows()) != 2 {
		t.Errorf("expected 2 rows, got %d", len(table.Rows()))
	}

	if table.ColumnCount() != 3 {
		t.Errorf("expected 3 columns, got %d", table.ColumnCount())
	}
}

func TestAddSection(t *testing.T) {
	doc := docx.New()
	section := doc.AddSection(docx.SectionTypeNextPage)
	if section == nil {
		t.Fatal("AddSection returned nil")
	}
}

func TestParagraphText(t *testing.T) {
	doc := docx.New()
	para := doc.AddParagraph("", "")
	para.SetText("New text")

	if para.Text() != "New text" {
		t.Errorf("expected 'New text', got '%s'", para.Text())
	}
}

func TestRunBold(t *testing.T) {
	doc := docx.New()
	para := doc.AddParagraph("", "")
	run := para.AddRun("Bold text", "")
	run.SetBold(true)

	if !run.Bold() {
		t.Error("expected bold to be true")
	}
}

func TestRunItalic(t *testing.T) {
	doc := docx.New()
	para := doc.AddParagraph("", "")
	run := para.AddRun("Italic text", "")
	run.SetItalic(true)

	if !run.Italic() {
		t.Error("expected italic to be true")
	}
}

func TestTableCell(t *testing.T) {
	doc := docx.New()
	table := doc.AddTable(2, 2, "")

	cell := table.Cell(0, 0)
	if cell == nil {
		t.Fatal("Cell(0, 0) returned nil")
	}

	cell.SetText("Cell text")
	if cell.Text() != "Cell text" {
		t.Errorf("expected 'Cell text', got '%s'", cell.Text())
	}
}

func TestTableAddRow(t *testing.T) {
	doc := docx.New()
	table := doc.AddTable(1, 2, "")

	initialRows := len(table.Rows())
	table.AddRow()

	if len(table.Rows()) != initialRows+1 {
		t.Errorf("expected %d rows, got %d", initialRows+1, len(table.Rows()))
	}
}

func TestSectionMargins(t *testing.T) {
	doc := docx.New()
	section := doc.AddSection(docx.SectionTypeNextPage)

	leftMargin := section.LeftMargin()
	if leftMargin != nil {
		t.Error("expected nil for unset left margin")
	}
}

func TestSectionPageWidth(t *testing.T) {
	doc := docx.New()
	section := doc.AddSection(docx.SectionTypeNextPage)

	width := shared.Inches(8.5)
	section.SetPageWidth(&width)
	result := section.PageWidth()

	if result == nil {
		t.Fatal("PageWidth returned nil")
	}

	if result.Inches() != 8.5 {
		t.Errorf("expected 8.5 inches, got %f", result.Inches())
	}
}
