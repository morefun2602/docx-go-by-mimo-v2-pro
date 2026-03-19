package styles_test

import (
	"testing"

	"github.com/docx-go/oxml"
	"github.com/docx-go/styles"
)

func TestStyleTypeString(t *testing.T) {
	tests := []struct {
		stype    styles.StyleType
		expected string
	}{
		{styles.StyleTypeParagraph, "paragraph"},
		{styles.StyleTypeCharacter, "character"},
		{styles.StyleTypeTable, "table"},
		{styles.StyleTypeNumbering, "numbering"},
	}

	for _, tt := range tests {
		if got := tt.stype.String(); got != tt.expected {
			t.Errorf("StyleType(%d).String() = %s, want %s", tt.stype, got, tt.expected)
		}
	}
}

func TestNewStyle(t *testing.T) {
	elem := oxml.NewElement("w:style")
	elem.SetAttr("styleId", "TestStyle")
	elem.SetAttr("type", "paragraph")

	nameElem := oxml.NewElement("w:name")
	nameElem.SetAttr("val", "Test Style")
	elem.AddChild(nameElem)

	style := styles.NewStyle(elem)
	if style.StyleID() != "TestStyle" {
		t.Errorf("StyleID() = %s, want TestStyle", style.StyleID())
	}
	if style.Name() != "Test Style" {
		t.Errorf("Name() = %s, want Test Style", style.Name())
	}
	if style.Type() != styles.StyleTypeParagraph {
		t.Errorf("Type() = %v, want StyleTypeParagraph", style.Type())
	}
}

func TestStyleSetName(t *testing.T) {
	elem := oxml.NewElement("w:style")
	style := styles.NewStyle(elem)

	style.SetName("New Name")
	if style.Name() != "New Name" {
		t.Errorf("Name() = %s, want New Name", style.Name())
	}
}

func TestStyleBasedOn(t *testing.T) {
	elem := oxml.NewElement("w:style")
	basedOnElem := oxml.NewElement("w:basedOn")
	basedOnElem.SetAttr("val", "Normal")
	elem.AddChild(basedOnElem)

	style := styles.NewStyle(elem)
	if style.BasedOn() != "Normal" {
		t.Errorf("BasedOn() = %s, want Normal", style.BasedOn())
	}
}

func TestStyleSetBasedOn(t *testing.T) {
	elem := oxml.NewElement("w:style")
	style := styles.NewStyle(elem)

	style.SetBasedOn("Heading1")
	if style.BasedOn() != "Heading1" {
		t.Errorf("BasedOn() = %s, want Heading1", style.BasedOn())
	}
}

func TestStyleUIPriority(t *testing.T) {
	elem := oxml.NewElement("w:style")
	uiPriorityElem := oxml.NewElement("w:uiPriority")
	uiPriorityElem.SetAttr("val", "1")
	elem.AddChild(uiPriorityElem)

	style := styles.NewStyle(elem)
	if style.UIPriority() != 1 {
		t.Errorf("UIPriority() = %d, want 1", style.UIPriority())
	}
}

func TestStyleHidden(t *testing.T) {
	elem := oxml.NewElement("w:style")
	hiddenElem := oxml.NewElement("w:hidden")
	elem.AddChild(hiddenElem)

	style := styles.NewStyle(elem)
	if !style.Hidden() {
		t.Error("Hidden() = false, want true")
	}
}

func TestStyleLocked(t *testing.T) {
	elem := oxml.NewElement("w:style")
	lockedElem := oxml.NewElement("w:locked")
	elem.AddChild(lockedElem)

	style := styles.NewStyle(elem)
	if !style.Locked() {
		t.Error("Locked() = false, want true")
	}
}

func TestStylesCollection(t *testing.T) {
	elem := oxml.NewElement("w:styles")

	style1 := oxml.NewElement("w:style")
	style1.SetAttr("styleId", "Style1")
	style1.SetAttr("type", "paragraph")
	name1 := oxml.NewElement("w:name")
	name1.SetAttr("val", "Style 1")
	style1.AddChild(name1)
	elem.AddChild(style1)

	style2 := oxml.NewElement("w:style")
	style2.SetAttr("styleId", "Style2")
	style2.SetAttr("type", "character")
	name2 := oxml.NewElement("w:name")
	name2.SetAttr("val", "Style 2")
	style2.AddChild(name2)
	elem.AddChild(style2)

	stylesCol := styles.NewStyles(elem, nil)

	if stylesCol.Len() != 2 {
		t.Errorf("Len() = %d, want 2", stylesCol.Len())
	}

	s := stylesCol.GetByID("Style1")
	if s == nil {
		t.Fatal("GetByID(Style1) returned nil")
	}
	if s.Name() != "Style 1" {
		t.Errorf("GetByID(Style1).Name() = %s, want Style 1", s.Name())
	}

	s2 := stylesCol.GetByName("Style 2")
	if s2 == nil {
		t.Fatal("GetByName(Style 2) returned nil")
	}
	if s2.StyleID() != "Style2" {
		t.Errorf("GetByName(Style 2).StyleID() = %s, want Style2", s2.StyleID())
	}
}

func TestStylesByType(t *testing.T) {
	elem := oxml.NewElement("w:styles")

	style1 := oxml.NewElement("w:style")
	style1.SetAttr("styleId", "ParaStyle")
	style1.SetAttr("type", "paragraph")
	elem.AddChild(style1)

	style2 := oxml.NewElement("w:style")
	style2.SetAttr("styleId", "CharStyle")
	style2.SetAttr("type", "character")
	elem.AddChild(style2)

	stylesCol := styles.NewStyles(elem, nil)

	paraStyles := stylesCol.GetByType(styles.StyleTypeParagraph)
	if len(paraStyles) != 1 {
		t.Errorf("GetByType(Paragraph) returned %d styles, want 1", len(paraStyles))
	}

	charStyles := stylesCol.GetByType(styles.StyleTypeCharacter)
	if len(charStyles) != 1 {
		t.Errorf("GetByType(Character) returned %d styles, want 1", len(charStyles))
	}
}
