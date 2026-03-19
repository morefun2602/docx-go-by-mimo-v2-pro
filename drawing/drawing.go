package drawing

import (
	"github.com/docx-go/oxml"
	"github.com/docx-go/shared"
)

type Drawing struct {
	element *oxml.Element
	parent  interface{}
}

func NewDrawing(element *oxml.Element, parent interface{}) *Drawing {
	return &Drawing{
		element: element,
		parent:  parent,
	}
}

func (d *Drawing) Element() *oxml.Element {
	return d.element
}

type Inline struct {
	element *oxml.Element
	parent  *Drawing
}

func NewInline(element *oxml.Element, parent *Drawing) *Inline {
	return &Inline{
		element: element,
		parent:  parent,
	}
}

func (i *Inline) Element() *oxml.Element {
	return i.element
}

func (i *Inline) Width() *shared.Length {
	extent := i.element.Find("wp:extent")
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

func (i *Inline) SetWidth(width *shared.Length) {
	extent := i.element.Find("wp:extent")
	if extent == nil {
		extent = oxml.NewElement("wp:extent")
		i.element.AddChild(extent)
	}
	if width == nil {
		extent.RemoveAttr("cx")
	} else {
		extent.SetAttr("cx", formatInt(int(width.EMU())))
	}
}

func (i *Inline) Height() *shared.Length {
	extent := i.element.Find("wp:extent")
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

func (i *Inline) SetHeight(height *shared.Length) {
	extent := i.element.Find("wp:extent")
	if extent == nil {
		extent = oxml.NewElement("wp:extent")
		i.element.AddChild(extent)
	}
	if height == nil {
		extent.RemoveAttr("cy")
	} else {
		extent.SetAttr("cy", formatInt(int(height.EMU())))
	}
}

func (i *Inline) DocPr() *DocPr {
	docPr := i.element.Find("wp:docPr")
	if docPr == nil {
		return nil
	}
	return NewDocPr(docPr)
}

type DocPr struct {
	element *oxml.Element
}

func NewDocPr(element *oxml.Element) *DocPr {
	return &DocPr{element: element}
}

func (dp *DocPr) Element() *oxml.Element {
	return dp.element
}

func (dp *DocPr) ID() string {
	val, _ := dp.element.GetAttr("id")
	return val
}

func (dp *DocPr) SetID(id string) {
	dp.element.SetAttr("id", id)
}

func (dp *DocPr) Name() string {
	val, _ := dp.element.GetAttr("name")
	return val
}

func (dp *DocPr) SetName(name string) {
	dp.element.SetAttr("name", name)
}

func (dp *DocPr) Description() string {
	val, _ := dp.element.GetAttr("descr")
	return val
}

func (dp *DocPr) SetDescription(desc string) {
	dp.element.SetAttr("descr", desc)
}

type Graphic struct {
	element *oxml.Element
}

func NewGraphic(element *oxml.Element) *Graphic {
	return &Graphic{element: element}
}

func (g *Graphic) Element() *oxml.Element {
	return g.element
}

type GraphicData struct {
	element *oxml.Element
}

func NewGraphicData(element *oxml.Element) *GraphicData {
	return &GraphicData{element: element}
}

func (gd *GraphicData) Element() *oxml.Element {
	return gd.element
}

func (gd *GraphicData) URI() string {
	val, _ := gd.element.GetAttr("uri")
	return val
}

func (gd *GraphicData) SetURI(uri string) {
	gd.element.SetAttr("uri", uri)
}

type Picture struct {
	element *oxml.Element
}

func NewPicture(element *oxml.Element) *Picture {
	return &Picture{element: element}
}

func (p *Picture) Element() *oxml.Element {
	return p.element
}

type Blip struct {
	element *oxml.Element
}

func NewBlip(element *oxml.Element) *Blip {
	return &Blip{element: element}
}

func (b *Blip) Element() *oxml.Element {
	return b.element
}

func (b *Blip) Embed() string {
	val, _ := b.element.GetAttr("r:embed")
	return val
}

func (b *Blip) SetEmbed(rID string) {
	b.element.SetAttr("r:embed", rID)
}

func formatInt(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}
	return result
}
