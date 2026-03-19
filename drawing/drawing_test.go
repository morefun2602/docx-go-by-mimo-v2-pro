package drawing_test

import (
	"testing"

	"github.com/docx-go/drawing"
	"github.com/docx-go/oxml"
	"github.com/docx-go/shared"
)

func TestNewDrawing(t *testing.T) {
	elem := oxml.NewElement("w:drawing")
	d := drawing.NewDrawing(elem, nil)
	if d.Element() != elem {
		t.Error("Element() returned different element")
	}
}

func TestNewInline(t *testing.T) {
	drawingElem := oxml.NewElement("w:drawing")
	inlineElem := oxml.NewElement("wp:inline")
	drawingElem.AddChild(inlineElem)

	inline := drawing.NewInline(inlineElem, nil)
	if inline.Element() != inlineElem {
		t.Error("Element() returned different element")
	}
}

func TestInlineWidth(t *testing.T) {
	inlineElem := oxml.NewElement("wp:inline")
	extentElem := oxml.NewElement("wp:extent")
	extentElem.SetAttr("cx", "914400")
	extentElem.SetAttr("cy", "914400")
	inlineElem.AddChild(extentElem)

	inline := drawing.NewInline(inlineElem, nil)
	width := inline.Width()
	if width == nil {
		t.Fatal("Width() returned nil")
	}
	if width.EMU() != 914400 {
		t.Errorf("Width().EMU() = %d, want 914400", width.EMU())
	}
}

func TestInlineSetWidth(t *testing.T) {
	inlineElem := oxml.NewElement("wp:inline")
	inline := drawing.NewInline(inlineElem, nil)

	w := shared.Inches(2.0)
	inline.SetWidth(&w)
	width := inline.Width()
	if width == nil {
		t.Fatal("Width() returned nil")
	}
	if width.Inches() != 2.0 {
		t.Errorf("Width().Inches() = %f, want 2.0", width.Inches())
	}
}

func TestInlineHeight(t *testing.T) {
	inlineElem := oxml.NewElement("wp:inline")
	extentElem := oxml.NewElement("wp:extent")
	extentElem.SetAttr("cx", "914400")
	extentElem.SetAttr("cy", "1828800")
	inlineElem.AddChild(extentElem)

	inline := drawing.NewInline(inlineElem, nil)
	height := inline.Height()
	if height == nil {
		t.Fatal("Height() returned nil")
	}
	if height.EMU() != 1828800 {
		t.Errorf("Height().EMU() = %d, want 1828800", height.EMU())
	}
}

func TestInlineSetHeight(t *testing.T) {
	inlineElem := oxml.NewElement("wp:inline")
	inline := drawing.NewInline(inlineElem, nil)

	h := shared.Inches(1.5)
	inline.SetHeight(&h)
	height := inline.Height()
	if height == nil {
		t.Fatal("Height() returned nil")
	}
	if height.Inches() != 1.5 {
		t.Errorf("Height().Inches() = %f, want 1.5", height.Inches())
	}
}

func TestDocPr(t *testing.T) {
	elem := oxml.NewElement("wp:docPr")
	elem.SetAttr("id", "1")
	elem.SetAttr("name", "Picture 1")
	elem.SetAttr("descr", "A picture")

	dp := drawing.NewDocPr(elem)
	if dp.ID() != "1" {
		t.Errorf("ID() = %s, want 1", dp.ID())
	}
	if dp.Name() != "Picture 1" {
		t.Errorf("Name() = %s, want Picture 1", dp.Name())
	}
	if dp.Description() != "A picture" {
		t.Errorf("Description() = %s, want A picture", dp.Description())
	}
}

func TestDocPrSetters(t *testing.T) {
	elem := oxml.NewElement("wp:docPr")
	dp := drawing.NewDocPr(elem)

	dp.SetID("2")
	dp.SetName("New Name")
	dp.SetDescription("New description")

	if dp.ID() != "2" {
		t.Errorf("ID() = %s, want 2", dp.ID())
	}
	if dp.Name() != "New Name" {
		t.Errorf("Name() = %s, want New Name", dp.Name())
	}
	if dp.Description() != "New description" {
		t.Errorf("Description() = %s, want New description", dp.Description())
	}
}

func TestBlip(t *testing.T) {
	elem := oxml.NewElement("a:blip")
	elem.SetAttr("r:embed", "rId1")

	blip := drawing.NewBlip(elem)
	if blip.Embed() != "rId1" {
		t.Errorf("Embed() = %s, want rId1", blip.Embed())
	}

	blip.SetEmbed("rId2")
	if blip.Embed() != "rId2" {
		t.Errorf("Embed() = %s, want rId2", blip.Embed())
	}
}

func TestGraphic(t *testing.T) {
	elem := oxml.NewElement("a:graphic")
	g := drawing.NewGraphic(elem)
	if g.Element() != elem {
		t.Error("Element() returned different element")
	}
}

func TestGraphicData(t *testing.T) {
	elem := oxml.NewElement("a:graphicData")
	elem.SetAttr("uri", "http://schemas.openxmlformats.org/drawingml/2006/picture")

	gd := drawing.NewGraphicData(elem)
	if gd.URI() != "http://schemas.openxmlformats.org/drawingml/2006/picture" {
		t.Errorf("URI() = %s, want picture URI", gd.URI())
	}
}

func TestPicture(t *testing.T) {
	elem := oxml.NewElement("pic:pic")
	p := drawing.NewPicture(elem)
	if p.Element() != elem {
		t.Error("Element() returned different element")
	}
}
