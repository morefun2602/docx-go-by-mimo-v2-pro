package styles

import (
	"github.com/docx-go/oxml"
)

type StyleType int

const (
	StyleTypeParagraph StyleType = iota
	StyleTypeCharacter
	StyleTypeTable
	StyleTypeNumbering
)

func (t StyleType) String() string {
	switch t {
	case StyleTypeParagraph:
		return "paragraph"
	case StyleTypeCharacter:
		return "character"
	case StyleTypeTable:
		return "table"
	case StyleTypeNumbering:
		return "numbering"
	default:
		return "unknown"
	}
}

type Style struct {
	element        *oxml.Element
	styleID        string
	name           string
	styleType      StyleType
	basedOn        string
	next           string
	locked         bool
	hidden         bool
	uiPriority     int
	semiHidden     bool
	unhideWhenUsed bool
	qFormat        bool
}

func NewStyle(element *oxml.Element) *Style {
	s := &Style{element: element}
	s.parse()
	return s
}

func (s *Style) parse() {
	if s.element == nil {
		return
	}

	if val, ok := s.element.GetAttr("styleId"); ok {
		s.styleID = val
	}
	if val, ok := s.element.GetAttr("type"); ok {
		s.styleType = parseStyleType(val)
	}

	nameElem := s.element.Find("w:name")
	if nameElem != nil {
		if val, ok := nameElem.GetAttr("val"); ok {
			s.name = val
		}
	}

	basedOnElem := s.element.Find("w:basedOn")
	if basedOnElem != nil {
		if val, ok := basedOnElem.GetAttr("val"); ok {
			s.basedOn = val
		}
	}

	nextElem := s.element.Find("w:next")
	if nextElem != nil {
		if val, ok := nextElem.GetAttr("val"); ok {
			s.next = val
		}
	}

	uiPriorityElem := s.element.Find("w:uiPriority")
	if uiPriorityElem != nil {
		if val, ok := uiPriorityElem.GetAttr("val"); ok {
			s.uiPriority = parseInt(val)
		}
	}

	s.locked = s.element.Find("w:locked") != nil
	s.hidden = s.element.Find("w:hidden") != nil
	s.semiHidden = s.element.Find("w:semiHidden") != nil
	s.unhideWhenUsed = s.element.Find("w:unhideWhenUsed") != nil
	s.qFormat = s.element.Find("w:qFormat") != nil
}

func (s *Style) Element() *oxml.Element {
	return s.element
}

func (s *Style) StyleID() string {
	return s.styleID
}

func (s *Style) Name() string {
	return s.name
}

func (s *Style) SetName(name string) {
	s.name = name
	nameElem := s.element.Find("w:name")
	if nameElem == nil {
		nameElem = oxml.NewElement("w:name")
		s.element.AddChild(nameElem)
	}
	nameElem.SetAttr("val", name)
}

func (s *Style) Type() StyleType {
	return s.styleType
}

func (s *Style) BasedOn() string {
	return s.basedOn
}

func (s *Style) SetBasedOn(styleID string) {
	s.basedOn = styleID
	basedOnElem := s.element.Find("w:basedOn")
	if basedOnElem == nil {
		basedOnElem = oxml.NewElement("w:basedOn")
		s.element.AddChild(basedOnElem)
	}
	if styleID == "" {
		s.element.RemoveChild(basedOnElem)
	} else {
		basedOnElem.SetAttr("val", styleID)
	}
}

func (s *Style) Next() string {
	return s.next
}

func (s *Style) SetNext(styleID string) {
	s.next = styleID
	nextElem := s.element.Find("w:next")
	if nextElem == nil {
		nextElem = oxml.NewElement("w:next")
		s.element.AddChild(nextElem)
	}
	if styleID == "" {
		s.element.RemoveChild(nextElem)
	} else {
		nextElem.SetAttr("val", styleID)
	}
}

func (s *Style) Locked() bool {
	return s.locked
}

func (s *Style) SetLocked(locked bool) {
	s.locked = locked
	elem := s.element.Find("w:locked")
	if locked {
		if elem == nil {
			elem = oxml.NewElement("w:locked")
			s.element.AddChild(elem)
		}
	} else {
		if elem != nil {
			s.element.RemoveChild(elem)
		}
	}
}

func (s *Style) Hidden() bool {
	return s.hidden
}

func (s *Style) SetHidden(hidden bool) {
	s.hidden = hidden
	elem := s.element.Find("w:hidden")
	if hidden {
		if elem == nil {
			elem = oxml.NewElement("w:hidden")
			s.element.AddChild(elem)
		}
	} else {
		if elem != nil {
			s.element.RemoveChild(elem)
		}
	}
}

func (s *Style) UIPriority() int {
	return s.uiPriority
}

func (s *Style) SetUIPriority(priority int) {
	s.uiPriority = priority
	elem := s.element.Find("w:uiPriority")
	if elem == nil {
		elem = oxml.NewElement("w:uiPriority")
		s.element.AddChild(elem)
	}
	elem.SetAttr("val", formatInt(priority))
}

type Styles struct {
	element *oxml.Element
	parent  interface{}
	styles  []*Style
}

func NewStyles(element *oxml.Element, parent interface{}) *Styles {
	s := &Styles{
		element: element,
		parent:  parent,
	}
	s.parse()
	return s
}

func (s *Styles) parse() {
	s.styles = nil
	for _, child := range s.element.FindAll("w:style") {
		s.styles = append(s.styles, NewStyle(child))
	}
}

func (s *Styles) Element() *oxml.Element {
	return s.element
}

func (s *Styles) Len() int {
	return len(s.styles)
}

func (s *Styles) Get(index int) *Style {
	if index < 0 || index >= len(s.styles) {
		return nil
	}
	return s.styles[index]
}

func (s *Styles) GetByID(styleID string) *Style {
	for _, style := range s.styles {
		if style.StyleID() == styleID {
			return style
		}
	}
	return nil
}

func (s *Styles) GetByName(name string) *Style {
	for _, style := range s.styles {
		if style.Name() == name {
			return style
		}
	}
	return nil
}

func (s *Styles) GetByType(styleType StyleType) []*Style {
	var result []*Style
	for _, style := range s.styles {
		if style.Type() == styleType {
			result = append(result, style)
		}
	}
	return result
}

func (s *Styles) All() []*Style {
	return s.styles
}

func parseStyleType(val string) StyleType {
	switch val {
	case "paragraph":
		return StyleTypeParagraph
	case "character":
		return StyleTypeCharacter
	case "table":
		return StyleTypeTable
	case "numbering":
		return StyleTypeNumbering
	default:
		return StyleTypeParagraph
	}
}

func parseInt(s string) int {
	result := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		}
	}
	return result
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
