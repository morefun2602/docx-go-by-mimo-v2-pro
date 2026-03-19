package oxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

const (
	NSW   = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"
	NSR   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"
	NSA   = "http://schemas.openxmlformats.org/drawingml/2006/main"
	NSWP  = "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
	NSPic = "http://schemas.openxmlformats.org/drawingml/2006/picture"
	NSCp  = "http://schemas.openxmlformats.org/package/2006/content-types"
)

func ParseXML(data []byte) (*Element, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.Strict = false

	var root *Element
	var stack []*Element

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("XML parse error: %w", err)
		}

		switch t := token.(type) {
		case xml.StartElement:
			elem := &Element{
				XMLName:  t.Name,
				Attr:     t.Attr,
				children: make([]*Element, 0),
			}
			if len(stack) == 0 {
				root = elem
			} else {
				parent := stack[len(stack)-1]
				parent.children = append(parent.children, elem)
				elem.parent = parent
			}
			stack = append(stack, elem)

		case xml.EndElement:
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

		case xml.CharData:
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.text = string(t)
			}
		}
	}

	return root, nil
}

type Element struct {
	XMLName  xml.Name
	Attr     []xml.Attr
	text     string
	children []*Element
	parent   *Element
}

func (e *Element) Name() string {
	return e.XMLName.Local
}

func (e *Element) Space() string {
	return e.XMLName.Space
}

func (e *Element) Tag() string {
	if e.XMLName.Space == "" {
		return e.XMLName.Local
	}
	prefix := namespaceToPrefix(e.XMLName.Space)
	if prefix == "" {
		return e.XMLName.Local
	}
	return prefix + ":" + e.XMLName.Local
}

func namespaceToPrefix(ns string) string {
	switch ns {
	case NSW:
		return "w"
	case NSR:
		return "r"
	case NSA:
		return "a"
	case NSWP:
		return "wp"
	case NSPic:
		return "pic"
	case NSCp:
		return "cp"
	default:
		return ""
	}
}

func prefixToNamespace(pfx string) string {
	switch pfx {
	case "w":
		return NSW
	case "r":
		return NSR
	case "a":
		return NSA
	case "wp":
		return NSWP
	case "pic":
		return NSPic
	case "cp":
		return NSCp
	default:
		return ""
	}
}

func (e *Element) GetAttr(name string) (string, bool) {
	for _, attr := range e.Attr {
		if attr.Name.Local == name {
			return attr.Value, true
		}
	}
	return "", false
}

func (e *Element) SetAttr(name, value string) {
	for i, attr := range e.Attr {
		if attr.Name.Local == name {
			e.Attr[i].Value = value
			return
		}
	}
	e.Attr = append(e.Attr, xml.Attr{
		Name:  xml.Name{Local: name},
		Value: value,
	})
}

func (e *Element) RemoveAttr(name string) {
	for i, attr := range e.Attr {
		if attr.Name.Local == name {
			e.Attr = append(e.Attr[:i], e.Attr[i+1:]...)
			return
		}
	}
}

func (e *Element) Text() string {
	return e.text
}

func (e *Element) SetText(text string) {
	e.text = text
}

func (e *Element) Children() []*Element {
	return e.children
}

func (e *Element) Parent() *Element {
	return e.parent
}

func (e *Element) Find(name string) *Element {
	for _, child := range e.children {
		if child.Tag() == name {
			return child
		}
	}
	return nil
}

func (e *Element) FindAll(name string) []*Element {
	var result []*Element
	for _, child := range e.children {
		if child.Tag() == name {
			result = append(result, child)
		}
	}
	return result
}

func (e *Element) FindRecursive(name string) []*Element {
	var result []*Element
	if e.Tag() == name {
		result = append(result, e)
	}
	for _, child := range e.children {
		result = append(result, child.FindRecursive(name)...)
	}
	return result
}

func (e *Element) AddChild(child *Element) {
	child.parent = e
	e.children = append(e.children, child)
}

func (e *Element) InsertBefore(child *Element, before *Element) {
	for i, c := range e.children {
		if c == before {
			child.parent = e
			e.children = append(e.children[:i], append([]*Element{child}, e.children[i:]...)...)
			return
		}
	}
	e.AddChild(child)
}

func (e *Element) InsertAfter(child *Element, after *Element) {
	for i, c := range e.children {
		if c == after {
			child.parent = e
			e.children = append(e.children[:i+1], append([]*Element{child}, e.children[i+1:]...)...)
			return
		}
	}
	e.AddChild(child)
}

func (e *Element) RemoveChild(child *Element) {
	for i, c := range e.children {
		if c == child {
			e.children = append(e.children[:i], e.children[i+1:]...)
			child.parent = nil
			return
		}
	}
}

func (e *Element) RemoveAllChildren() {
	for _, child := range e.children {
		child.parent = nil
	}
	e.children = nil
}

func (e *Element) MarshalXML() ([]byte, error) {
	var buf bytes.Buffer
	if err := e.writeXML(&buf, ""); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *Element) writeXML(buf *bytes.Buffer, indent string) error {
	buf.WriteString(indent)
	buf.WriteString("<")
	buf.WriteString(e.Tag())

	for _, attr := range e.Attr {
		buf.WriteString(" ")
		buf.WriteString(attr.Name.Local)
		buf.WriteString(`="`)
		buf.WriteString(escapeXML(attr.Value))
		buf.WriteString(`"`)

	}

	if len(e.children) == 0 && e.text == "" {
		buf.WriteString("/>\n")
		return nil
	}

	buf.WriteString(">\n")

	if e.text != "" {
		buf.WriteString(indent + "  ")
		buf.WriteString(escapeXML(e.text))
		buf.WriteString("\n")
	}

	for _, child := range e.children {
		if err := child.writeXML(buf, indent+"  "); err != nil {
			return err
		}
	}

	buf.WriteString(indent)
	buf.WriteString("</")
	buf.WriteString(e.Tag())
	buf.WriteString(">\n")

	return nil
}

func escapeXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, `"`, "&quot;")
	return s
}

func NewElement(tag string) *Element {
	parts := strings.SplitN(tag, ":", 2)
	if len(parts) == 2 {
		ns := prefixToNamespace(parts[0])
		return &Element{
			XMLName:  xml.Name{Space: ns, Local: parts[1]},
			children: make([]*Element, 0),
		}
	}
	return &Element{
		XMLName:  xml.Name{Local: tag},
		children: make([]*Element, 0),
	}
}

func NewElementWithText(tag, text string) *Element {
	elem := NewElement(tag)
	elem.text = text
	return elem
}

type CTOnOff struct {
	Val *bool
}

func (c *CTOnOff) IsOn() bool {
	return c.Val == nil || *c.Val
}

func (c *CTOnOff) IsOff() bool {
	return c.Val != nil && !*c.Val
}

type CTString struct {
	Val string
}

type CTDecimalNumber struct {
	Val int
}

type CTHpsMeasure struct {
	Val int
}

type CTColor struct {
	Val        string
	ThemeColor *string
}

type CTFonts struct {
	Ascii    *string
	HAnsi    *string
	EastAsia *string
	Cs       *string
}

type CTUnderline struct {
	Val *string
}

type CTRPr struct {
	Bold      *CTOnOff
	Italic    *CTOnOff
	Color     *CTColor
	Fonts     *CTFonts
	FontSize  *CTHpsMeasure
	Underline *CTUnderline
	Style     *CTString
}

type CTPPr struct {
	Style           *CTString
	Alignment       *CTJc
	Spacing         *CTSpacing
	Indent          *CTInd
	KeepNext        *CTOnOff
	KeepLines       *CTOnOff
	PageBreakBefore *CTOnOff
	OutlineLevel    *CTDecimalNumber
	WidowControl    *CTOnOff
	Tabs            *CTTabStops
}

type CTJc struct {
	Val string
}

type CTSpacing struct {
	Before   *int
	After    *int
	Line     *int
	LineRule *string
}

type CTInd struct {
	Left      *int
	Right     *int
	FirstLine *int
	Hanging   *int
}

type CTTabStop struct {
	Val    string
	Pos    *int
	Leader *string
}

type CTTabStops struct {
	Tab []*CTTabStop
}

type CTDocument struct {
	Body *CTBody
}

type CTBody struct {
	SectPr   []*CTSectPr
	Elements []interface{}
}

type CTSectPr struct {
	Type       *CTSectType
	PageSize   *CTPageSz
	PageMargin *CTPageMar
	TitlePg    *CTOnOff
}

type CTSectType struct {
	Val *string
}

type CTPageSz struct {
	W      *int
	H      *int
	Orient *string
}

type CTPageMar struct {
	Top    *int
	Right  *int
	Bottom *int
	Left   *int
	Header *int
	Footer *int
	Gutter *int
}

type CTR struct {
	RPr      *CTRPr
	Elements []interface{}
}

type CTP struct {
	PPr      *CTPPr
	Elements []interface{}
}

type CTText struct {
	Space string
	Text  string
}

type CTBr struct {
	Type  *string
	Clear *string
}
