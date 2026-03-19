package docx

import (
	"github.com/docx-go/oxml"
)

type DocSettings struct {
	element *oxml.Element
	parent  interface{}
}

func NewDocSettings(element *oxml.Element, parent interface{}) *DocSettings {
	return &DocSettings{
		element: element,
		parent:  parent,
	}
}

func (ds *DocSettings) Element() *oxml.Element {
	return ds.element
}

func (ds *DocSettings) getOrAddElement(tag string) *oxml.Element {
	elem := ds.element.Find(tag)
	if elem == nil {
		elem = oxml.NewElement(tag)
		ds.element.AddChild(elem)
	}
	return elem
}

func (ds *DocSettings) UpdateFieldsOnOpen() bool {
	elem := ds.element.Find("w:updateFields")
	if elem == nil {
		return false
	}
	val, exists := elem.GetAttr("val")
	if !exists {
		return true
	}
	return val == "true" || val == "1"
}

func (ds *DocSettings) SetUpdateFieldsOnOpen(update bool) {
	elem := ds.getOrAddElement("w:updateFields")
	if update {
		elem.SetAttr("val", "true")
	} else {
		elem.SetAttr("val", "false")
	}
}

func (ds *DocSettings) DefaultTabStop() *int {
	elem := ds.element.Find("w:defaultTabStop")
	if elem == nil {
		return nil
	}
	val, exists := elem.GetAttr("val")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (ds *DocSettings) SetDefaultTabStop(stop *int) {
	if stop == nil {
		elem := ds.element.Find("w:defaultTabStop")
		if elem != nil {
			ds.element.RemoveChild(elem)
		}
		return
	}
	elem := ds.getOrAddElement("w:defaultTabStop")
	elem.SetAttr("val", formatInt(*stop))
}

func (ds *DocSettings) EvenAndOddHeaders() bool {
	elem := ds.element.Find("w:evenAndOddHeaders")
	return elem != nil
}

func (ds *DocSettings) SetEvenAndOddHeaders(value bool) {
	elem := ds.element.Find("w:evenAndOddHeaders")
	if value {
		if elem == nil {
			elem = oxml.NewElement("w:evenAndOddHeaders")
			ds.element.AddChild(elem)
		}
	} else {
		if elem != nil {
			ds.element.RemoveChild(elem)
		}
	}
}

func (ds *DocSettings) DefaultTableStyle() string {
	elem := ds.element.Find("w:defaultTableStyle")
	if elem == nil {
		return ""
	}
	val, _ := elem.GetAttr("val")
	return val
}

func (ds *DocSettings) SetDefaultTableStyle(style string) {
	elem := ds.getOrAddElement("w:defaultTableStyle")
	elem.SetAttr("val", style)
}

func (ds *DocSettings) HideGrammaticalErrors() bool {
	elem := ds.element.Find("w:hideGrammaticalErrors")
	return elem != nil
}

func (ds *DocSettings) SetHideGrammaticalErrors(hide bool) {
	elem := ds.element.Find("w:hideGrammaticalErrors")
	if hide {
		if elem == nil {
			elem = oxml.NewElement("w:hideGrammaticalErrors")
			ds.element.AddChild(elem)
		}
	} else {
		if elem != nil {
			ds.element.RemoveChild(elem)
		}
	}
}

func (ds *DocSettings) HideSpellingErrors() bool {
	elem := ds.element.Find("w:hideSpellingErrors")
	return elem != nil
}

func (ds *DocSettings) SetHideSpellingErrors(hide bool) {
	elem := ds.element.Find("w:hideSpellingErrors")
	if hide {
		if elem == nil {
			elem = oxml.NewElement("w:hideSpellingErrors")
			ds.element.AddChild(elem)
		}
	} else {
		if elem != nil {
			ds.element.RemoveChild(elem)
		}
	}
}

func (ds *DocSettings) NoPunctuationKerning() bool {
	elem := ds.element.Find("w:noPunctuationKerning")
	return elem != nil
}

func (ds *DocSettings) SetNoPunctuationKerning(value bool) {
	elem := ds.element.Find("w:noPunctuationKerning")
	if value {
		if elem == nil {
			elem = oxml.NewElement("w:noPunctuationKerning")
			ds.element.AddChild(elem)
		}
	} else {
		if elem != nil {
			ds.element.RemoveChild(elem)
		}
	}
}

func (ds *DocSettings) PrintTwoOnOne() bool {
	elem := ds.element.Find("w:printTwoOnOne")
	return elem != nil
}

func (ds *DocSettings) SetPrintTwoOnOne(value bool) {
	elem := ds.element.Find("w:printTwoOnOne")
	if value {
		if elem == nil {
			elem = oxml.NewElement("w:printTwoOnOne")
			ds.element.AddChild(elem)
		}
	} else {
		if elem != nil {
			ds.element.RemoveChild(elem)
		}
	}
}

func (ds *DocSettings) BookFoldPrinting() bool {
	elem := ds.element.Find("w:bookFoldPrinting")
	return elem != nil
}

func (ds *DocSettings) SetBookFoldPrinting(value bool) {
	elem := ds.element.Find("w:bookFoldPrinting")
	if value {
		if elem == nil {
			elem = oxml.NewElement("w:bookFoldPrinting")
			ds.element.AddChild(elem)
		}
	} else {
		if elem != nil {
			ds.element.RemoveChild(elem)
		}
	}
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
