package docx

import (
	"github.com/docx-go/oxml"
	"github.com/docx-go/text"
)

type BlockItemContainer struct {
	element *oxml.Element
	parent  interface{}
}

func NewBlockItemContainer(element *oxml.Element, parent interface{}) *BlockItemContainer {
	return &BlockItemContainer{
		element: element,
		parent:  parent,
	}
}

func (bic *BlockItemContainer) Element() *oxml.Element {
	return bic.element
}

func (bic *BlockItemContainer) AddParagraph(content, style string) *text.Paragraph {
	p := oxml.NewElement("w:p")
	bic.element.AddChild(p)
	para := text.NewParagraph(p, bic)
	if content != "" {
		para.AddRun(content, "")
	}
	if style != "" {
		para.SetStyle(style)
	}
	return para
}

func (bic *BlockItemContainer) AddTable(rows, cols int) *Table {
	tbl := oxml.NewElement("w:tbl")
	bic.element.AddChild(tbl)

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

	return NewTable(tbl, bic)
}

func (bic *BlockItemContainer) Paragraphs() []*text.Paragraph {
	var paragraphs []*text.Paragraph
	for _, p := range bic.element.FindAll("w:p") {
		paragraphs = append(paragraphs, text.NewParagraph(p, bic))
	}
	return paragraphs
}

func (bic *BlockItemContainer) Tables() []*Table {
	var tables []*Table
	for _, tbl := range bic.element.FindAll("w:tbl") {
		tables = append(tables, NewTable(tbl, bic))
	}
	return tables
}

func (bic *BlockItemContainer) ClearContent() {
	var sectPr *oxml.Element
	for _, child := range bic.element.Children() {
		if child.Tag() == "w:sectPr" {
			sectPr = child
			break
		}
	}
	bic.element.RemoveAllChildren()
	if sectPr != nil {
		bic.element.AddChild(sectPr)
	}
}
