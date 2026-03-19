package docx

import (
	"fmt"
	"os"

	"github.com/docx-go/opc"
	"github.com/docx-go/oxml"
	"github.com/docx-go/text"
)

type Document struct {
	element *oxml.Element
	part    interface{}
	body    *Body
}

func NewDocument(element *oxml.Element, part interface{}) *Document {
	doc := &Document{
		element: element,
		part:    part,
	}
	body := element.Find("w:body")
	if body != nil {
		doc.body = NewBody(body, doc)
	}
	return doc
}

func (d *Document) Element() *oxml.Element {
	return d.element
}

func (d *Document) AddParagraph(content, style string) *text.Paragraph {
	return d.body.AddParagraph(content, style)
}

func (d *Document) AddHeading(content string, level int) *text.Paragraph {
	if level < 0 || level > 9 {
		panic(fmt.Sprintf("level must be in range 0-9, got %d", level))
	}
	style := "Heading 1"
	if level == 0 {
		style = "Title"
	} else if level > 1 {
		style = fmt.Sprintf("Heading %d", level)
	}
	return d.AddParagraph(content, style)
}

func (d *Document) AddPageBreak() *text.Paragraph {
	para := d.AddParagraph("", "")
	run := para.AddRun("", "")
	run.AddBreak(text.BreakPage)
	return para
}

func (d *Document) AddTable(rows, cols int, style string) *Table {
	tbl := d.body.AddTable(rows, cols)
	if style != "" {
		tbl.SetStyle(style)
	}
	return tbl
}

func (d *Document) AddSection(startType SectionType) *Section {
	sectPr := oxml.NewElement("w:sectPr")
	d.element.AddChild(sectPr)
	section := NewSection(sectPr, d)
	section.SetStartType(startType)
	return section
}

func (d *Document) Paragraphs() []*text.Paragraph {
	return d.body.Paragraphs()
}

func (d *Document) Tables() []*Table {
	return d.body.Tables()
}

func (d *Document) Sections() *Sections {
	return NewSections(d.element, d)
}

func (d *Document) Styles() *Styles {
	return NewStyles(d.element, d)
}

func (d *Document) Settings() *Settings {
	return NewSettings(d.element, d)
}

func (d *Document) Save(pathOrStream interface{}) error {
	switch v := pathOrStream.(type) {
	case string:
		return d.saveToFile(v)
	default:
		return fmt.Errorf("unsupported save target")
	}
}

func (d *Document) saveToFile(path string) error {
	data, err := d.element.MarshalXML()
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

type Body struct {
	element *oxml.Element
	parent  interface{}
}

func NewBody(element *oxml.Element, parent interface{}) *Body {
	return &Body{
		element: element,
		parent:  parent,
	}
}

func (b *Body) Element() *oxml.Element {
	return b.element
}

func (b *Body) AddParagraph(content, style string) *text.Paragraph {
	p := oxml.NewElement("w:p")
	b.element.AddChild(p)
	para := text.NewParagraph(p, b)
	if content != "" {
		para.AddRun(content, "")
	}
	if style != "" {
		para.SetStyle(style)
	}
	return para
}

func (b *Body) AddTable(rows, cols int) *Table {
	tbl := oxml.NewElement("w:tbl")
	b.element.AddChild(tbl)

	tblPr := oxml.NewElement("w:tblPr")
	tbl.AddChild(tblPr)

	tblGrid := oxml.NewElement("w:tblGrid")
	tbl.AddChild(tblGrid)
	for i := 0; i < cols; i++ {
		gridCol := oxml.NewElement("w:gridCol")
		gridCol.SetAttr("w:w", "5000")
		tblGrid.AddChild(gridCol)
	}

	for i := 0; i < rows; i++ {
		tr := oxml.NewElement("w:tr")
		tbl.AddChild(tr)
		for j := 0; j < cols; j++ {
			tc := oxml.NewElement("w:tc")
			tr.AddChild(tc)
			p := oxml.NewElement("w:p")
			tc.AddChild(p)
		}
	}

	return NewTable(tbl, b)
}

func (b *Body) Paragraphs() []*text.Paragraph {
	var paragraphs []*text.Paragraph
	for _, p := range b.element.FindAll("w:p") {
		paragraphs = append(paragraphs, text.NewParagraph(p, b))
	}
	return paragraphs
}

func (b *Body) Tables() []*Table {
	var tables []*Table
	for _, tbl := range b.element.FindAll("w:tbl") {
		tables = append(tables, NewTable(tbl, b))
	}
	return tables
}

type Styles struct {
	element *oxml.Element
	parent  interface{}
}

func NewStyles(element *oxml.Element, parent interface{}) *Styles {
	return &Styles{
		element: element,
		parent:  parent,
	}
}

func (s *Styles) Element() *oxml.Element {
	return s.element
}

type Settings struct {
	element *oxml.Element
	parent  interface{}
}

func NewSettings(element *oxml.Element, parent interface{}) *Settings {
	return &Settings{
		element: element,
		parent:  parent,
	}
}

func (s *Settings) Element() *oxml.Element {
	return s.element
}

func DocumentFromFile(path string) (*Document, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	element, err := oxml.ParseXML(data)
	if err != nil {
		return nil, err
	}

	return NewDocument(element, nil), nil
}

func NewDocumentFromPackage(pkg *opc.Package) (*Document, error) {
	mainPart, err := pkg.MainDocumentPart()
	if err != nil {
		return nil, err
	}
	if mainPart.ContentType() != opc.ContentTypeWMLDocumentMain {
		return nil, fmt.Errorf("not a Word file, content type is '%s'", mainPart.ContentType())
	}

	data := mainPart.Blob()
	element, err := oxml.ParseXML(data)
	if err != nil {
		return nil, err
	}

	return NewDocument(element, mainPart), nil
}

func New() *Document {
	doc := &Document{
		element: oxml.NewElement("w:document"),
	}
	doc.element.SetAttr("xmlns:w", oxml.NSW)

	body := oxml.NewElement("w:body")
	doc.element.AddChild(body)
	doc.body = NewBody(body, doc)

	sectPr := oxml.NewElement("w:sectPr")
	body.AddChild(sectPr)

	return doc
}

func (d *Document) Body() *Body {
	return d.body
}
