package oxml_test

import (
	"testing"

	"github.com/docx-go/oxml"
)

func TestParseXML(t *testing.T) {
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:body>
    <w:p>
      <w:r>
        <w:t>Hello</w:t>
      </w:r>
    </w:p>
  </w:body>
</w:document>`

	element, err := oxml.ParseXML([]byte(xml))
	if err != nil {
		t.Fatalf("ParseXML failed: %v", err)
	}

	if element.Tag() != "w:document" {
		t.Errorf("expected 'w:document', got '%s'", element.Tag())
	}
}

func TestElementFind(t *testing.T) {
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:body>
    <w:p>
      <w:r>
        <w:t>Hello</w:t>
      </w:r>
    </w:p>
  </w:body>
</w:document>`

	element, err := oxml.ParseXML([]byte(xml))
	if err != nil {
		t.Fatalf("ParseXML failed: %v", err)
	}

	body := element.Find("w:body")
	if body == nil {
		t.Fatal("body not found")
	}

	p := body.Find("w:p")
	if p == nil {
		t.Fatal("p not found")
	}

	r := p.Find("w:r")
	if r == nil {
		t.Fatal("r not found")
	}

	tElem := r.Find("w:t")
	if tElem == nil {
		t.Fatal("t not found")
	}

	if tElem.Text() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", tElem.Text())
	}
}

func TestElementGetAttr(t *testing.T) {
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<w:p xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" w:val="test">
</w:p>`

	element, err := oxml.ParseXML([]byte(xml))
	if err != nil {
		t.Fatalf("ParseXML failed: %v", err)
	}

	val, exists := element.GetAttr("val")
	if !exists {
		t.Error("attribute val not found")
	}
	if val != "test" {
		t.Errorf("expected 'test', got '%s'", val)
	}
}

func TestElementSetAttr(t *testing.T) {
	element := oxml.NewElement("w:p")
	element.SetAttr("w:val", "test")

	val, exists := element.GetAttr("w:val")
	if !exists {
		t.Error("attribute w:val not found")
	}
	if val != "test" {
		t.Errorf("expected 'test', got '%s'", val)
	}
}

func TestElementRemoveAttr(t *testing.T) {
	element := oxml.NewElement("w:p")
	element.SetAttr("w:val", "test")
	element.RemoveAttr("w:val")

	_, exists := element.GetAttr("w:val")
	if exists {
		t.Error("attribute w:val should have been removed")
	}
}

func TestElementSetText(t *testing.T) {
	element := oxml.NewElement("w:t")
	element.SetText("Hello")

	if element.Text() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", element.Text())
	}
}

func TestElementAddChild(t *testing.T) {
	parent := oxml.NewElement("w:p")
	child := oxml.NewElement("w:r")
	parent.AddChild(child)

	if len(parent.Children()) != 1 {
		t.Errorf("expected 1 child, got %d", len(parent.Children()))
	}
	if parent.Children()[0] != child {
		t.Error("child not found")
	}
}

func TestElementRemoveChild(t *testing.T) {
	parent := oxml.NewElement("w:p")
	child := oxml.NewElement("w:r")
	parent.AddChild(child)
	parent.RemoveChild(child)

	if len(parent.Children()) != 0 {
		t.Errorf("expected 0 children, got %d", len(parent.Children()))
	}
}

func TestElementRemoveAllChildren(t *testing.T) {
	parent := oxml.NewElement("w:p")
	child1 := oxml.NewElement("w:r")
	child2 := oxml.NewElement("w:r")
	parent.AddChild(child1)
	parent.AddChild(child2)
	parent.RemoveAllChildren()

	if len(parent.Children()) != 0 {
		t.Errorf("expected 0 children, got %d", len(parent.Children()))
	}
}

func TestElementMarshalXML(t *testing.T) {
	element := oxml.NewElement("w:p")
	child := oxml.NewElement("w:r")
	element.AddChild(child)

	data, err := element.MarshalXML()
	if err != nil {
		t.Fatalf("MarshalXML failed: %v", err)
	}

	if len(data) == 0 {
		t.Error("expected non-empty XML")
	}
}
