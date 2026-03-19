package dml

import (
	"fmt"
	"strings"

	"github.com/docx-go/oxml"
)

type Color struct {
	element *oxml.Element
	val     string
}

func NewColor(element *oxml.Element) *Color {
	c := &Color{element: element}
	c.parse()
	return c
}

func (c *Color) parse() {
	if c.element == nil {
		return
	}
	if val, ok := c.element.GetAttr("val"); ok {
		c.val = val
	}
}

func (c *Color) Element() *oxml.Element {
	return c.element
}

func (c *Color) Val() string {
	return c.val
}

func (c *Color) SetVal(val string) {
	c.val = val
	c.element.SetAttr("val", val)
}

func NewColorFromHex(hex string) *Color {
	element := oxml.NewElement("w:color")
	hex = strings.TrimPrefix(hex, "#")
	element.SetAttr("val", hex)
	return NewColor(element)
}

func (c *Color) RGB() (r, g, b uint8, err error) {
	hex := strings.TrimPrefix(c.val, "#")
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}

	r, err = parseHexByte(hex[0:2])
	if err != nil {
		return
	}
	g, err = parseHexByte(hex[2:4])
	if err != nil {
		return
	}
	b, err = parseHexByte(hex[4:6])
	return
}

func (c *Color) SetRGB(r, g, b uint8) {
	c.val = fmt.Sprintf("%02x%02x%02x", r, g, b)
	c.element.SetAttr("val", c.val)
}

func parseHexByte(s string) (byte, error) {
	var result byte
	for i := 0; i < 2; i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			result = result*16 + (c - '0')
		case c >= 'a' && c <= 'f':
			result = result*16 + (c - 'a' + 10)
		case c >= 'A' && c <= 'F':
			result = result*16 + (c - 'A' + 10)
		default:
			return 0, fmt.Errorf("invalid hex char: %c", c)
		}
	}
	return result, nil
}

type ThemeColor struct {
	element *oxml.Element
	val     string
}

func NewThemeColor(element *oxml.Element) *ThemeColor {
	tc := &ThemeColor{element: element}
	tc.parse()
	return tc
}

func (tc *ThemeColor) parse() {
	if tc.element == nil {
		return
	}
	if val, ok := tc.element.GetAttr("val"); ok {
		tc.val = val
	}
}

func (tc *ThemeColor) Element() *oxml.Element {
	return tc.element
}

func (tc *ThemeColor) Val() string {
	return tc.val
}

func (tc *ThemeColor) SetVal(val string) {
	tc.val = val
	tc.element.SetAttr("val", val)
}

type Fill struct {
	element *oxml.Element
}

func NewFill(element *oxml.Element) *Fill {
	return &Fill{element: element}
}

func (f *Fill) Element() *oxml.Element {
	return f.element
}

type NoFill struct {
	element *oxml.Element
}

func NewNoFill() *NoFill {
	return &NoFill{element: oxml.NewElement("a:noFill")}
}

func (nf *NoFill) Element() *oxml.Element {
	return nf.element
}

type SolidFill struct {
	element *oxml.Element
	color   *Color
}

func NewSolidFill(color *Color) *SolidFill {
	element := oxml.NewElement("a:solidFill")
	element.AddChild(color.Element())
	return &SolidFill{
		element: element,
		color:   color,
	}
}

func (sf *SolidFill) Element() *oxml.Element {
	return sf.element
}

func (sf *SolidFill) Color() *Color {
	return sf.color
}
