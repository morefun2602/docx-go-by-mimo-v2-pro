package docx

import (
	"github.com/docx-go/image"
	"github.com/docx-go/oxml"
	"github.com/docx-go/shared"
)

type InlineShape struct {
	element *oxml.Element
	parent  interface{}
}

func NewInlineShape(element *oxml.Element, parent interface{}) *InlineShape {
	return &InlineShape{
		element: element,
		parent:  parent,
	}
}

func (is *InlineShape) Element() *oxml.Element {
	return is.element
}

func (is *InlineShape) Width() *shared.Length {
	extent := is.element.Find("wp:extent")
	if extent == nil {
		return nil
	}
	val, exists := extent.GetAttr("cx")
	if !exists {
		return nil
	}
	w := 0
	for _, c := range val {
		w = w*10 + int(c-'0')
	}
	l := shared.Length(w)
	return &l
}

func (is *InlineShape) SetWidth(width *shared.Length) {
	extent := is.element.Find("wp:extent")
	if extent == nil {
		extent = oxml.NewElement("wp:extent")
		is.element.AddChild(extent)
	}
	if width != nil {
		extent.SetAttr("cx", formatInt(int(width.EMU())))
	}
}

func (is *InlineShape) Height() *shared.Length {
	extent := is.element.Find("wp:extent")
	if extent == nil {
		return nil
	}
	val, exists := extent.GetAttr("cy")
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

func (is *InlineShape) SetHeight(height *shared.Length) {
	extent := is.element.Find("wp:extent")
	if extent == nil {
		extent = oxml.NewElement("wp:extent")
		is.element.AddChild(extent)
	}
	if height != nil {
		extent.SetAttr("cy", formatInt(int(height.EMU())))
	}
}

func (is *InlineShape) ProportionalLock() bool {
	cNvPr := is.element.Find("wp:docPr")
	if cNvPr == nil {
		return false
	}
	return cNvPr.Find("a:graphicFrameLocks") != nil
}

type Anchor struct {
	element *oxml.Element
	parent  interface{}
}

func NewAnchor(element *oxml.Element, parent interface{}) *Anchor {
	return &Anchor{
		element: element,
		parent:  parent,
	}
}

func (a *Anchor) Element() *oxml.Element {
	return a.element
}

func (a *Anchor) Width() *shared.Length {
	extent := a.element.Find("wp:extent")
	if extent == nil {
		return nil
	}
	val, exists := extent.GetAttr("cx")
	if !exists {
		return nil
	}
	w := 0
	for _, c := range val {
		w = w*10 + int(c-'0')
	}
	l := shared.Length(w)
	return &l
}

func (a *Anchor) SetWidth(width *shared.Length) {
	extent := a.element.Find("wp:extent")
	if extent == nil {
		extent = oxml.NewElement("wp:extent")
		a.element.AddChild(extent)
	}
	if width != nil {
		extent.SetAttr("cx", formatInt(int(width.EMU())))
	}
}

func (a *Anchor) Height() *shared.Length {
	extent := a.element.Find("wp:extent")
	if extent == nil {
		return nil
	}
	val, exists := extent.GetAttr("cy")
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

func (a *Anchor) SetHeight(height *shared.Length) {
	extent := a.element.Find("wp:extent")
	if extent == nil {
		extent = oxml.NewElement("wp:extent")
		a.element.AddChild(extent)
	}
	if height != nil {
		extent.SetAttr("cy", formatInt(int(height.EMU())))
	}
}

func (a *Anchor) PositionH() *PositionH {
	posH := a.element.Find("wp:positionH")
	if posH == nil {
		return nil
	}
	return NewPositionH(posH)
}

func (a *Anchor) PositionV() *PositionV {
	posV := a.element.Find("wp:positionV")
	if posV == nil {
		return nil
	}
	return NewPositionV(posV)
}

type PositionH struct {
	element *oxml.Element
}

func NewPositionH(element *oxml.Element) *PositionH {
	return &PositionH{element: element}
}

func (p *PositionH) Element() *oxml.Element {
	return p.element
}

func (p *PositionH) RelativeFrom() string {
	val, _ := p.element.GetAttr("relativeFrom")
	return val
}

func (p *PositionH) SetRelativeFrom(relativeFrom string) {
	p.element.SetAttr("relativeFrom", relativeFrom)
}

func (p *PositionH) Align() string {
	align := p.element.Find("wp:align")
	if align == nil {
		return ""
	}
	val, _ := align.GetAttr("val")
	return val
}

func (p *PositionH) SetAlign(align string) {
	alignElem := p.element.Find("wp:align")
	if alignElem == nil {
		alignElem = oxml.NewElement("wp:align")
		p.element.AddChild(alignElem)
	}
	alignElem.SetAttr("val", align)
}

func (p *PositionH) PosOffset() *int {
	posOffset := p.element.Find("wp:posOffset")
	if posOffset == nil {
		return nil
	}
	val := posOffset.Text()
	if val == "" {
		return nil
	}
	offset := 0
	for _, c := range val {
		offset = offset*10 + int(c-'0')
	}
	return &offset
}

func (p *PositionH) SetPosOffset(offset *int) {
	posOffset := p.element.Find("wp:posOffset")
	if offset == nil {
		if posOffset != nil {
			p.element.RemoveChild(posOffset)
		}
		return
	}
	if posOffset == nil {
		posOffset = oxml.NewElement("wp:posOffset")
		p.element.AddChild(posOffset)
	}
	posOffset.SetText(formatInt(*offset))
}

type PositionV struct {
	element *oxml.Element
}

func NewPositionV(element *oxml.Element) *PositionV {
	return &PositionV{element: element}
}

func (p *PositionV) Element() *oxml.Element {
	return p.element
}

func (p *PositionV) RelativeFrom() string {
	val, _ := p.element.GetAttr("relativeFrom")
	return val
}

func (p *PositionV) SetRelativeFrom(relativeFrom string) {
	p.element.SetAttr("relativeFrom", relativeFrom)
}

func (p *PositionV) Align() string {
	align := p.element.Find("wp:align")
	if align == nil {
		return ""
	}
	val, _ := align.GetAttr("val")
	return val
}

func (p *PositionV) SetAlign(align string) {
	alignElem := p.element.Find("wp:align")
	if alignElem == nil {
		alignElem = oxml.NewElement("wp:align")
		p.element.AddChild(alignElem)
	}
	alignElem.SetAttr("val", align)
}

func (p *PositionV) PosOffset() *int {
	posOffset := p.element.Find("wp:posOffset")
	if posOffset == nil {
		return nil
	}
	val := posOffset.Text()
	if val == "" {
		return nil
	}
	offset := 0
	for _, c := range val {
		offset = offset*10 + int(c-'0')
	}
	return &offset
}

func (p *PositionV) SetPosOffset(offset *int) {
	posOffset := p.element.Find("wp:posOffset")
	if offset == nil {
		if posOffset != nil {
			p.element.RemoveChild(posOffset)
		}
		return
	}
	if posOffset == nil {
		posOffset = oxml.NewElement("wp:posOffset")
		p.element.AddChild(posOffset)
	}
	posOffset.SetText(formatInt(*offset))
}

type InlineShapes struct {
	element *oxml.Element
	parent  interface{}
}

func NewInlineShapes(element *oxml.Element, parent interface{}) *InlineShapes {
	return &InlineShapes{
		element: element,
		parent:  parent,
	}
}

func (is *InlineShapes) Element() *oxml.Element {
	return is.element
}

func (is *InlineShapes) Len() int {
	return len(is.element.FindAll("wp:inline"))
}

func (is *InlineShapes) Get(index int) *InlineShape {
	inlineShapes := is.element.FindAll("wp:inline")
	if index < 0 || index >= len(inlineShapes) {
		return nil
	}
	return NewInlineShape(inlineShapes[index], is)
}

func (is *InlineShapes) AddPicture(img *image.Image) *InlineShape {
	drawing := oxml.NewElement("w:drawing")
	inline := oxml.NewElement("wp:inline")
	drawing.AddChild(inline)

	extent := oxml.NewElement("wp:extent")
	extent.SetAttr("cx", formatInt(img.WidthEMU()))
	extent.SetAttr("cy", formatInt(img.HeightEMU()))
	inline.AddChild(extent)

	docPr := oxml.NewElement("wp:docPr")
	docPr.SetAttr("id", "1")
	docPr.SetAttr("name", "Picture")
	inline.AddChild(docPr)

	graphic := oxml.NewElement("a:graphic")
	inline.AddChild(graphic)

	graphicData := oxml.NewElement("a:graphicData")
	graphicData.SetAttr("uri", "http://schemas.openxmlformats.org/drawingml/2006/picture")
	graphic.AddChild(graphicData)

	pic := oxml.NewElement("pic:pic")
	graphicData.AddChild(pic)

	return NewInlineShape(inline, is)
}
