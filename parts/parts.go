package parts

import (
	"github.com/docx-go/opc"
	"github.com/docx-go/oxml"
)

type DocumentPart struct {
	*opc.Part
	element *oxml.Element
}

func NewDocumentPart(partname opc.PackURI, contentType string, blob []byte, pkg *opc.Package) *DocumentPart {
	element, err := oxml.ParseXML(blob)
	if err != nil {
		element = oxml.NewElement("w:document")
	}
	return &DocumentPart{
		Part:    opc.NewPart(partname, contentType, blob, pkg),
		element: element,
	}
}

func (dp *DocumentPart) Element() *oxml.Element {
	return dp.element
}

func (dp *DocumentPart) Blob() []byte {
	data, err := dp.element.MarshalXML()
	if err != nil {
		return dp.Part.Blob()
	}
	return data
}

type HeaderPart struct {
	*opc.Part
	element *oxml.Element
}

func NewHeaderPart(partname opc.PackURI, contentType string, blob []byte, pkg *opc.Package) *HeaderPart {
	element, err := oxml.ParseXML(blob)
	if err != nil {
		element = oxml.NewElement("w:hdr")
	}
	return &HeaderPart{
		Part:    opc.NewPart(partname, contentType, blob, pkg),
		element: element,
	}
}

func (hp *HeaderPart) Element() *oxml.Element {
	return hp.element
}

func (hp *HeaderPart) Blob() []byte {
	data, err := hp.element.MarshalXML()
	if err != nil {
		return hp.Part.Blob()
	}
	return data
}

type FooterPart struct {
	*opc.Part
	element *oxml.Element
}

func NewFooterPart(partname opc.PackURI, contentType string, blob []byte, pkg *opc.Package) *FooterPart {
	element, err := oxml.ParseXML(blob)
	if err != nil {
		element = oxml.NewElement("w:ftr")
	}
	return &FooterPart{
		Part:    opc.NewPart(partname, contentType, blob, pkg),
		element: element,
	}
}

func (fp *FooterPart) Element() *oxml.Element {
	return fp.element
}

func (fp *FooterPart) Blob() []byte {
	data, err := fp.element.MarshalXML()
	if err != nil {
		return fp.Part.Blob()
	}
	return data
}

type StylesPart struct {
	*opc.Part
	element *oxml.Element
}

func NewStylesPart(partname opc.PackURI, contentType string, blob []byte, pkg *opc.Package) *StylesPart {
	element, err := oxml.ParseXML(blob)
	if err != nil {
		element = oxml.NewElement("w:styles")
	}
	return &StylesPart{
		Part:    opc.NewPart(partname, contentType, blob, pkg),
		element: element,
	}
}

func (sp *StylesPart) Element() *oxml.Element {
	return sp.element
}

func (sp *StylesPart) Blob() []byte {
	data, err := sp.element.MarshalXML()
	if err != nil {
		return sp.Part.Blob()
	}
	return data
}

type NumberingPart struct {
	*opc.Part
	element *oxml.Element
}

func NewNumberingPart(partname opc.PackURI, contentType string, blob []byte, pkg *opc.Package) *NumberingPart {
	element, err := oxml.ParseXML(blob)
	if err != nil {
		element = oxml.NewElement("w:numbering")
	}
	return &NumberingPart{
		Part:    opc.NewPart(partname, contentType, blob, pkg),
		element: element,
	}
}

func (np *NumberingPart) Element() *oxml.Element {
	return np.element
}

func (np *NumberingPart) Blob() []byte {
	data, err := np.element.MarshalXML()
	if err != nil {
		return np.Part.Blob()
	}
	return data
}

type SettingsPart struct {
	*opc.Part
	element *oxml.Element
}

func NewSettingsPart(partname opc.PackURI, contentType string, blob []byte, pkg *opc.Package) *SettingsPart {
	element, err := oxml.ParseXML(blob)
	if err != nil {
		element = oxml.NewElement("w:settings")
	}
	return &SettingsPart{
		Part:    opc.NewPart(partname, contentType, blob, pkg),
		element: element,
	}
}

func (sp *SettingsPart) Element() *oxml.Element {
	return sp.element
}

func (sp *SettingsPart) Blob() []byte {
	data, err := sp.element.MarshalXML()
	if err != nil {
		return sp.Part.Blob()
	}
	return data
}

type CommentsPart struct {
	*opc.Part
	element *oxml.Element
}

func NewCommentsPart(partname opc.PackURI, contentType string, blob []byte, pkg *opc.Package) *CommentsPart {
	element, err := oxml.ParseXML(blob)
	if err != nil {
		element = oxml.NewElement("w:comments")
	}
	return &CommentsPart{
		Part:    opc.NewPart(partname, contentType, blob, pkg),
		element: element,
	}
}

func (cp *CommentsPart) Element() *oxml.Element {
	return cp.element
}

func (cp *CommentsPart) Blob() []byte {
	data, err := cp.element.MarshalXML()
	if err != nil {
		return cp.Part.Blob()
	}
	return data
}

type ImagePart struct {
	*opc.Part
	sha1 string
}

func NewImagePart(partname opc.PackURI, contentType string, blob []byte, pkg *opc.Package) *ImagePart {
	return &ImagePart{
		Part: opc.NewPart(partname, contentType, blob, pkg),
	}
}

func (ip *ImagePart) SHA1() string {
	return ip.sha1
}

func RegisterPartTypes() {
	pf := opc.GetPartFactory()
	pf.RegisterPartType(opc.ContentTypeWMLDocumentMain, &opc.Part{})
	pf.RegisterPartType(opc.ContentTypeWMLStyles, &opc.Part{})
	pf.RegisterPartType(opc.ContentTypeWMLNumbering, &opc.Part{})
	pf.RegisterPartType(opc.ContentTypeWMLSettings, &opc.Part{})
	pf.RegisterPartType(opc.ContentTypeWMLHeader, &opc.Part{})
	pf.RegisterPartType(opc.ContentTypeWMLFooter, &opc.Part{})
	pf.RegisterPartType(opc.ContentTypeWMLComments, &opc.Part{})
}
