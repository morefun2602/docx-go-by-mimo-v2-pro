package text

import (
	"fmt"
	"strings"

	"github.com/docx-go/oxml"
)

type BreakType int

const (
	BreakLine BreakType = iota
	BreakPage
	BreakColumn
	BreakLineClearLeft
	BreakLineClearRight
	BreakLineClearAll
)

type Paragraph struct {
	element *oxml.Element
	parent  interface{}
}

func NewParagraph(element *oxml.Element, parent interface{}) *Paragraph {
	return &Paragraph{
		element: element,
		parent:  parent,
	}
}

func (p *Paragraph) Element() *oxml.Element {
	return p.element
}

func (p *Paragraph) AddRun(text string, style string) *Run {
	rElem := oxml.NewElement("w:r")
	p.element.AddChild(rElem)

	run := NewRun(rElem, p)
	if text != "" {
		run.SetText(text)
	}
	if style != "" {
		run.SetStyle(style)
	}
	return run
}

func (p *Paragraph) Alignment() string {
	pPr := p.element.Find("w:pPr")
	if pPr == nil {
		return ""
	}
	jc := pPr.Find("w:jc")
	if jc == nil {
		return ""
	}
	val, _ := jc.GetAttr("w:val")
	return val
}

func (p *Paragraph) SetAlignment(alignment string) {
	pPr := p.getOrAddPPr()
	jc := pPr.Find("w:jc")
	if jc == nil {
		jc = oxml.NewElement("w:jc")
		pPr.AddChild(jc)
	}
	jc.SetAttr("w:val", alignment)
}

func (p *Paragraph) Clear() {
	p.element.RemoveAllChildren()
	return
}

func (p *Paragraph) Style() string {
	pPr := p.element.Find("w:pPr")
	if pPr == nil {
		return ""
	}
	style := pPr.Find("w:pStyle")
	if style == nil {
		return ""
	}
	val, _ := style.GetAttr("w:val")
	return val
}

func (p *Paragraph) SetStyle(style string) {
	pPr := p.getOrAddPPr()
	styleElem := pPr.Find("w:pStyle")
	if styleElem == nil {
		styleElem = oxml.NewElement("w:pStyle")
		pPr.AddChild(styleElem)
	}
	styleElem.SetAttr("w:val", style)
}

func (p *Paragraph) Text() string {
	var parts []string
	for _, child := range p.element.Children() {
		if child.Tag() == "w:r" {
			run := NewRun(child, p)
			parts = append(parts, run.Text())
		}
	}
	return strings.Join(parts, "")
}

func (p *Paragraph) SetText(text string) {
	p.element.RemoveAllChildren()
	p.AddRun(text, "")
}

func (p *Paragraph) Runs() []*Run {
	var runs []*Run
	for _, child := range p.element.Children() {
		if child.Tag() == "w:r" {
			runs = append(runs, NewRun(child, p))
		}
	}
	return runs
}

func (p *Paragraph) InsertBefore(text, style string) *Paragraph {
	newP := oxml.NewElement("w:p")
	p.element.Parent().InsertBefore(newP, p.element)
	para := NewParagraph(newP, p.parent)
	if text != "" {
		para.AddRun(text, style)
	}
	if style != "" {
		para.SetStyle(style)
	}
	return para
}

func (p *Paragraph) getOrAddPPr() *oxml.Element {
	pPr := p.element.Find("w:pPr")
	if pPr == nil {
		pPr = oxml.NewElement("w:pPr")
		children := p.element.Children()
		if len(children) > 0 {
			p.element.InsertBefore(pPr, children[0])
		} else {
			p.element.AddChild(pPr)
		}
	}
	return pPr
}

type Run struct {
	element *oxml.Element
	parent  *Paragraph
}

func NewRun(element *oxml.Element, parent *Paragraph) *Run {
	return &Run{
		element: element,
		parent:  parent,
	}
}

func (r *Run) Element() *oxml.Element {
	return r.element
}

func (r *Run) AddBreak(breakType BreakType) {
	br := oxml.NewElement("w:br")
	switch breakType {
	case BreakPage:
		br.SetAttr("w:type", "page")
	case BreakColumn:
		br.SetAttr("w:type", "column")
	case BreakLineClearLeft:
		br.SetAttr("w:type", "textWrapping")
		br.SetAttr("w:clear", "left")
	case BreakLineClearRight:
		br.SetAttr("w:type", "textWrapping")
		br.SetAttr("w:clear", "right")
	case BreakLineClearAll:
		br.SetAttr("w:type", "textWrapping")
		br.SetAttr("w:clear", "all")
	}
	r.element.AddChild(br)
}

func (r *Run) AddText(text string) {
	t := oxml.NewElementWithText("w:t", text)
	r.element.AddChild(t)
}

func (r *Run) Clear() {
	r.element.RemoveAllChildren()
	return
}

func (r *Run) Bold() bool {
	rPr := r.element.Find("w:rPr")
	if rPr == nil {
		return false
	}
	b := rPr.Find("w:b")
	if b == nil {
		return false
	}
	val, exists := b.GetAttr("w:val")
	if !exists {
		return true
	}
	return val == "true" || val == "1"
}

func (r *Run) SetBold(bold bool) {
	rPr := r.getOrAddRPr()
	b := rPr.Find("w:b")
	if b == nil {
		b = oxml.NewElement("w:b")
		rPr.AddChild(b)
	}
	if bold {
		b.SetAttr("w:val", "true")
	} else {
		b.SetAttr("w:val", "false")
	}
}

func (r *Run) Italic() bool {
	rPr := r.element.Find("w:rPr")
	if rPr == nil {
		return false
	}
	i := rPr.Find("w:i")
	if i == nil {
		return false
	}
	val, exists := i.GetAttr("w:val")
	if !exists {
		return true
	}
	return val == "true" || val == "1"
}

func (r *Run) SetItalic(italic bool) {
	rPr := r.getOrAddRPr()
	i := rPr.Find("w:i")
	if i == nil {
		i = oxml.NewElement("w:i")
		rPr.AddChild(i)
	}
	if italic {
		i.SetAttr("w:val", "true")
	} else {
		i.SetAttr("w:val", "false")
	}
}

func (r *Run) Style() string {
	rPr := r.element.Find("w:rPr")
	if rPr == nil {
		return ""
	}
	style := rPr.Find("w:rStyle")
	if style == nil {
		return ""
	}
	val, _ := style.GetAttr("w:val")
	return val
}

func (r *Run) SetStyle(style string) {
	rPr := r.getOrAddRPr()
	styleElem := rPr.Find("w:rStyle")
	if styleElem == nil {
		styleElem = oxml.NewElement("w:rStyle")
		rPr.AddChild(styleElem)
	}
	styleElem.SetAttr("w:val", style)
}

func (r *Run) Text() string {
	var parts []string
	for _, child := range r.element.Children() {
		switch child.Tag() {
		case "w:t":
			parts = append(parts, child.Text())
		case "w:tab":
			parts = append(parts, "\t")
		case "w:br":
			val, exists := child.GetAttr("w:type")
			if !exists || val == "" {
				parts = append(parts, "\n")
			}
		case "w:cr":
			parts = append(parts, "\n")
		}
	}
	return strings.Join(parts, "")
}

func (r *Run) SetText(text string) {
	r.element.RemoveAllChildren()
	for _, c := range text {
		switch c {
		case '\t':
			tab := oxml.NewElement("w:tab")
			r.element.AddChild(tab)
		case '\n', '\r':
			cr := oxml.NewElement("w:cr")
			r.element.AddChild(cr)
		default:
			t := r.element.Find("w:t")
			if t == nil {
				t = oxml.NewElement("w:t")
				r.element.AddChild(t)
			}
			t.SetText(t.Text() + string(c))
		}
	}
}

func (r *Run) Font() *Font {
	return NewFont(r.getOrAddRPr())
}

func (r *Run) Underline() string {
	rPr := r.element.Find("w:rPr")
	if rPr == nil {
		return ""
	}
	u := rPr.Find("w:u")
	if u == nil {
		return ""
	}
	val, _ := u.GetAttr("w:val")
	return val
}

func (r *Run) SetUnderline(underline string) {
	rPr := r.getOrAddRPr()
	u := rPr.Find("w:u")
	if u == nil {
		u = oxml.NewElement("w:u")
		rPr.AddChild(u)
	}
	if underline == "" {
		rPr.RemoveChild(u)
	} else {
		u.SetAttr("w:val", underline)
	}
}

func (r *Run) AddTab() {
	tab := oxml.NewElement("w:tab")
	r.element.AddChild(tab)
}

func (r *Run) getOrAddRPr() *oxml.Element {
	rPr := r.element.Find("w:rPr")
	if rPr == nil {
		rPr = oxml.NewElement("w:rPr")
		children := r.element.Children()
		if len(children) > 0 {
			r.element.InsertBefore(rPr, children[0])
		} else {
			r.element.AddChild(rPr)
		}
	}
	return rPr
}

type Font struct {
	rPr *oxml.Element
}

func NewFont(rPr *oxml.Element) *Font {
	return &Font{rPr: rPr}
}

func (f *Font) Size() int {
	sz := f.rPr.Find("w:sz")
	if sz == nil {
		return 0
	}
	val, _ := sz.GetAttr("w:val")
	if val == "" {
		return 0
	}
	size := 0
	fmt.Sscanf(val, "%d", &size)
	return size
}

func (f *Font) SetSize(size int) {
	sz := f.rPr.Find("w:sz")
	if sz == nil {
		sz = oxml.NewElement("w:sz")
		f.rPr.AddChild(sz)
	}
	sz.SetAttr("w:val", fmt.Sprintf("%d", size*2))
}

func (f *Font) Name() string {
	rFonts := f.rPr.Find("w:rFonts")
	if rFonts == nil {
		return ""
	}
	val, _ := rFonts.GetAttr("w:ascii")
	return val
}

func (f *Font) SetName(name string) {
	rFonts := f.rPr.Find("w:rFonts")
	if rFonts == nil {
		rFonts = oxml.NewElement("w:rFonts")
		f.rPr.AddChild(rFonts)
	}
	rFonts.SetAttr("w:ascii", name)
}

func (f *Font) Color() string {
	color := f.rPr.Find("w:color")
	if color == nil {
		return ""
	}
	val, _ := color.GetAttr("w:val")
	return val
}

func (f *Font) SetColor(color string) {
	colorElem := f.rPr.Find("w:color")
	if colorElem == nil {
		colorElem = oxml.NewElement("w:color")
		f.rPr.AddChild(colorElem)
	}
	colorElem.SetAttr("w:val", color)
}

func (f *Font) Bold() bool {
	b := f.rPr.Find("w:b")
	if b == nil {
		return false
	}
	val, exists := b.GetAttr("w:val")
	if !exists {
		return true
	}
	return val == "true" || val == "1"
}

func (f *Font) SetBold(bold bool) {
	b := f.rPr.Find("w:b")
	if b == nil {
		b = oxml.NewElement("w:b")
		f.rPr.AddChild(b)
	}
	if bold {
		b.SetAttr("w:val", "true")
	} else {
		b.SetAttr("w:val", "false")
	}
}

func (f *Font) Italic() bool {
	i := f.rPr.Find("w:i")
	if i == nil {
		return false
	}
	val, exists := i.GetAttr("w:val")
	if !exists {
		return true
	}
	return val == "true" || val == "1"
}

func (f *Font) SetItalic(italic bool) {
	i := f.rPr.Find("w:i")
	if i == nil {
		i = oxml.NewElement("w:i")
		f.rPr.AddChild(i)
	}
	if italic {
		i.SetAttr("w:val", "true")
	} else {
		i.SetAttr("w:val", "false")
	}
}
