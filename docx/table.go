package docx

import (
	"fmt"

	"github.com/docx-go/oxml"
	"github.com/docx-go/shared"
)

type Table struct {
	element *oxml.Element
	parent  interface{}
}

func NewTable(element *oxml.Element, parent interface{}) *Table {
	return &Table{
		element: element,
		parent:  parent,
	}
}

func (t *Table) Element() *oxml.Element {
	return t.element
}

func (t *Table) AddRow() *Row {
	tr := oxml.NewElement("w:tr")
	t.element.AddChild(tr)

	tblGrid := t.element.Find("w:tblGrid")
	if tblGrid != nil {
		for _, gridCol := range tblGrid.FindAll("w:gridCol") {
			tc := oxml.NewElement("w:tc")
			tr.AddChild(tc)
			w, exists := gridCol.GetAttr("w:w")
			if exists {
				tcPr := oxml.NewElement("w:tcPr")
				tc.AddChild(tcPr)
				tcW := oxml.NewElement("w:tcW")
				tcPr.AddChild(tcW)
				tcW.SetAttr("w:w", w)
				tcW.SetAttr("w:type", "dxa")
			}
		}
	}

	return NewRow(tr, t)
}

func (t *Table) AddColumn(width shared.Length) *Column {
	tblGrid := t.element.Find("w:tblGrid")
	if tblGrid == nil {
		tblGrid = oxml.NewElement("w:tblGrid")
		t.element.InsertBefore(tblGrid, t.element.Children()[0])
	}

	gridCol := oxml.NewElement("w:gridCol")
	gridCol.SetAttr("w:w", fmt.Sprintf("%d", width.EMU()))
	tblGrid.AddChild(gridCol)

	for _, tr := range t.element.FindAll("w:tr") {
		tc := oxml.NewElement("w:tc")
		tr.AddChild(tc)
		tcPr := oxml.NewElement("w:tcPr")
		tc.AddChild(tcPr)
		tcW := oxml.NewElement("w:tcW")
		tcPr.AddChild(tcW)
		tcW.SetAttr("w:w", fmt.Sprintf("%d", width.EMU()))
		tcW.SetAttr("w:type", "dxa")
	}

	return NewColumn(gridCol, t)
}

func (t *Table) Alignment() string {
	tblPr := t.element.Find("w:tblPr")
	if tblPr == nil {
		return ""
	}
	jc := tblPr.Find("w:jc")
	if jc == nil {
		return ""
	}
	val, _ := jc.GetAttr("w:val")
	return val
}

func (t *Table) SetAlignment(alignment string) {
	tblPr := t.getOrAddTblPr()
	jc := tblPr.Find("w:jc")
	if jc == nil {
		jc = oxml.NewElement("w:jc")
		tblPr.AddChild(jc)
	}
	jc.SetAttr("w:val", alignment)
}

func (t *Table) Autofit() bool {
	tblPr := t.element.Find("w:tblPr")
	if tblPr == nil {
		return true
	}
	layout := tblPr.Find("w:tblLayout")
	if layout == nil {
		return true
	}
	val, _ := layout.GetAttr("w:val")
	return val != "fixed"
}

func (t *Table) SetAutofit(autofit bool) {
	tblPr := t.getOrAddTblPr()
	layout := tblPr.Find("w:tblLayout")
	if layout == nil {
		layout = oxml.NewElement("w:tblLayout")
		tblPr.AddChild(layout)
	}
	if autofit {
		layout.SetAttr("w:val", "autofit")
	} else {
		layout.SetAttr("w:val", "fixed")
	}
}

func (t *Table) Cell(rowIdx, colIdx int) *Cell {
	cells := t.Cells()
	columnCount := t.ColumnCount()
	if columnCount == 0 {
		return nil
	}
	cellIdx := colIdx + (rowIdx * columnCount)
	if cellIdx >= len(cells) {
		return nil
	}
	return cells[cellIdx]
}

func (t *Table) ColumnCells(columnIdx int) []*Cell {
	cells := t.Cells()
	columnCount := t.ColumnCount()
	var result []*Cell
	for i := columnIdx; i < len(cells); i += columnCount {
		result = append(result, cells[i])
	}
	return result
}

func (t *Table) RowCells(rowIdx int) []*Cell {
	cells := t.Cells()
	columnCount := t.ColumnCount()
	start := rowIdx * columnCount
	end := start + columnCount
	if start >= len(cells) {
		return nil
	}
	if end > len(cells) {
		end = len(cells)
	}
	return cells[start:end]
}

func (t *Table) Rows() []*Row {
	var rows []*Row
	for _, tr := range t.element.FindAll("w:tr") {
		rows = append(rows, NewRow(tr, t))
	}
	return rows
}

func (t *Table) Columns() []*Column {
	tblGrid := t.element.Find("w:tblGrid")
	if tblGrid == nil {
		return nil
	}
	var columns []*Column
	for _, gridCol := range tblGrid.FindAll("w:gridCol") {
		columns = append(columns, NewColumn(gridCol, t))
	}
	return columns
}

func (t *Table) Style() string {
	tblPr := t.element.Find("w:tblPr")
	if tblPr == nil {
		return ""
	}
	style := tblPr.Find("w:tblStyle")
	if style == nil {
		return ""
	}
	val, _ := style.GetAttr("w:val")
	return val
}

func (t *Table) SetStyle(style string) {
	tblPr := t.getOrAddTblPr()
	styleElem := tblPr.Find("w:tblStyle")
	if styleElem == nil {
		styleElem = oxml.NewElement("w:tblStyle")
		tblPr.AddChild(styleElem)
	}
	styleElem.SetAttr("w:val", style)
}

func (t *Table) Cells() []*Cell {
	var cells []*Cell
	for _, tc := range t.element.FindRecursive("w:tc") {
		cells = append(cells, NewCell(tc, t))
	}
	return cells
}

func (t *Table) ColumnCount() int {
	tblGrid := t.element.Find("w:tblGrid")
	if tblGrid == nil {
		return 0
	}
	return len(tblGrid.FindAll("w:gridCol"))
}

func (t *Table) getOrAddTblPr() *oxml.Element {
	tblPr := t.element.Find("w:tblPr")
	if tblPr == nil {
		tblPr = oxml.NewElement("w:tblPr")
		t.element.InsertBefore(tblPr, t.element.Children()[0])
	}
	return tblPr
}

type Row struct {
	element *oxml.Element
	parent  *Table
}

func NewRow(element *oxml.Element, parent *Table) *Row {
	return &Row{
		element: element,
		parent:  parent,
	}
}

func (r *Row) Element() *oxml.Element {
	return r.element
}

func (r *Row) Cells() []*Cell {
	var cells []*Cell
	for _, tc := range r.element.FindAll("w:tc") {
		cells = append(cells, NewCell(tc, r.parent))
	}
	return cells
}

func (r *Row) Height() *shared.Length {
	trPr := r.element.Find("w:trPr")
	if trPr == nil {
		return nil
	}
	trHeight := trPr.Find("w:trHeight")
	if trHeight == nil {
		return nil
	}
	val, exists := trHeight.GetAttr("w:val")
	if !exists {
		return nil
	}
	h := 0
	for _, c := range val {
		h = h*10 + int(c-'0')
	}
	l := shared.Length(h)
	return &l
}

func (r *Row) SetHeight(height *shared.Length) {
	trPr := r.getOrAddTrPr()
	trHeight := trPr.Find("w:trHeight")
	if trHeight == nil {
		trHeight = oxml.NewElement("w:trHeight")
		trPr.AddChild(trHeight)
	}
	if height == nil {
		trPr.RemoveChild(trHeight)
	} else {
		trHeight.SetAttr("w:val", fmt.Sprintf("%d", height.EMU()))
		trHeight.SetAttr("w:hRule", "atLeast")
	}
}

func (r *Row) getOrAddTrPr() *oxml.Element {
	trPr := r.element.Find("w:trPr")
	if trPr == nil {
		trPr = oxml.NewElement("w:trPr")
		r.element.InsertBefore(trPr, r.element.Children()[0])
	}
	return trPr
}

type Column struct {
	element *oxml.Element
	parent  *Table
}

func NewColumn(element *oxml.Element, parent *Table) *Column {
	return &Column{
		element: element,
		parent:  parent,
	}
}

func (c *Column) Element() *oxml.Element {
	return c.element
}

func (c *Column) Width() *shared.Length {
	val, exists := c.element.GetAttr("w:w")
	if !exists {
		return nil
	}
	w := 0
	for _, ch := range val {
		w = w*10 + int(ch-'0')
	}
	l := shared.Length(w)
	return &l
}

func (c *Column) SetWidth(width *shared.Length) {
	if width == nil {
		c.element.RemoveAttr("w:w")
	} else {
		c.element.SetAttr("w:w", fmt.Sprintf("%d", width.EMU()))
	}
}

type Cell struct {
	element *oxml.Element
	parent  *Table
}

func NewCell(element *oxml.Element, parent *Table) *Cell {
	return &Cell{
		element: element,
		parent:  parent,
	}
}

func (c *Cell) Element() *oxml.Element {
	return c.element
}

func (c *Cell) AddParagraph(text, style string) {
	p := oxml.NewElement("w:p")
	c.element.InsertBefore(p, c.element.Children()[len(c.element.Children())-1])
	if style != "" {
		pPr := oxml.NewElement("w:pPr")
		p.AddChild(pPr)
		pStyle := oxml.NewElement("w:pStyle")
		pPr.AddChild(pStyle)
		pStyle.SetAttr("w:val", style)
	}
	if text != "" {
		r := oxml.NewElement("w:r")
		p.AddChild(r)
		t := oxml.NewElement("w:t")
		r.AddChild(t)
		t.SetText(text)
	}
}

func (c *Cell) AddTable(rows, cols int) *Table {
	tbl := oxml.NewElement("w:tbl")
	c.element.AddChild(tbl)

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

	p := oxml.NewElement("w:p")
	c.element.AddChild(p)

	return NewTable(tbl, c)
}

func (c *Cell) Text() string {
	var parts []string
	for _, p := range c.element.FindAll("w:p") {
		for _, r := range p.FindAll("w:r") {
			for _, t := range r.FindAll("w:t") {
				parts = append(parts, t.Text())
			}
		}
	}
	return joinStrings(parts, "\n")
}

func (c *Cell) SetText(text string) {
	c.element.RemoveAllChildren()
	p := oxml.NewElement("w:p")
	c.element.AddChild(p)
	r := oxml.NewElement("w:r")
	p.AddChild(r)
	t := oxml.NewElement("w:t")
	r.AddChild(t)
	t.SetText(text)
}

func (c *Cell) VerticalAlignment() string {
	tcPr := c.element.Find("w:tcPr")
	if tcPr == nil {
		return ""
	}
	vAlign := tcPr.Find("w:vAlign")
	if vAlign == nil {
		return ""
	}
	val, _ := vAlign.GetAttr("w:val")
	return val
}

func (c *Cell) SetVerticalAlignment(alignment string) {
	tcPr := c.getOrAddTcPr()
	vAlign := tcPr.Find("w:vAlign")
	if vAlign == nil {
		vAlign = oxml.NewElement("w:vAlign")
		tcPr.AddChild(vAlign)
	}
	if alignment == "" {
		tcPr.RemoveChild(vAlign)
	} else {
		vAlign.SetAttr("w:val", alignment)
	}
}

func (c *Cell) Width() *shared.Length {
	tcPr := c.element.Find("w:tcPr")
	if tcPr == nil {
		return nil
	}
	tcW := tcPr.Find("w:tcW")
	if tcW == nil {
		return nil
	}
	val, exists := tcW.GetAttr("w:w")
	if !exists {
		return nil
	}
	w := 0
	for _, ch := range val {
		w = w*10 + int(ch-'0')
	}
	l := shared.Length(w)
	return &l
}

func (c *Cell) SetWidth(width *shared.Length) {
	tcPr := c.getOrAddTcPr()
	tcW := tcPr.Find("w:tcW")
	if tcW == nil {
		tcW = oxml.NewElement("w:tcW")
		tcPr.AddChild(tcW)
	}
	if width == nil {
		tcPr.RemoveChild(tcW)
	} else {
		tcW.SetAttr("w:w", fmt.Sprintf("%d", width.EMU()))
		tcW.SetAttr("w:type", "dxa")
	}
}

func (c *Cell) Merge(other *Cell) *Cell {
	gridSpan1 := c.GridSpan()
	gridSpan2 := other.GridSpan()

	tcPr := c.getOrAddTcPr()
	mergeElem := tcPr.Find("w:vMerge")
	if mergeElem == nil {
		mergeElem = oxml.NewElement("w:vMerge")
		tcPr.AddChild(mergeElem)
	}
	mergeElem.SetAttr("w:val", "restart")

	tcPr2 := other.getOrAddTcPr()
	mergeElem2 := tcPr2.Find("w:vMerge")
	if mergeElem2 == nil {
		mergeElem2 = oxml.NewElement("w:vMerge")
		tcPr2.AddChild(mergeElem2)
	}
	mergeElem2.SetAttr("w:val", "continue")

	if gridSpan1 > 1 {
		gridSpan := tcPr.Find("w:gridSpan")
		if gridSpan == nil {
			gridSpan = oxml.NewElement("w:gridSpan")
			tcPr.AddChild(gridSpan)
		}
		gridSpan.SetAttr("w:val", fmt.Sprintf("%d", gridSpan1+gridSpan2))
	}

	return c
}

func (c *Cell) GridSpan() int {
	tcPr := c.element.Find("w:tcPr")
	if tcPr == nil {
		return 1
	}
	gridSpan := tcPr.Find("w:gridSpan")
	if gridSpan == nil {
		return 1
	}
	val, _ := gridSpan.GetAttr("w:val")
	if val == "" {
		return 1
	}
	span := 0
	for _, ch := range val {
		span = span*10 + int(ch-'0')
	}
	return span
}

func (c *Cell) getOrAddTcPr() *oxml.Element {
	tcPr := c.element.Find("w:tcPr")
	if tcPr == nil {
		tcPr = oxml.NewElement("w:tcPr")
		c.element.InsertBefore(tcPr, c.element.Children()[0])
	}
	return tcPr
}

func joinStrings(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for _, p := range parts[1:] {
		result += sep + p
	}
	return result
}
