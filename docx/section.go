package docx

import (
	"fmt"

	"github.com/docx-go/oxml"
	"github.com/docx-go/shared"
	"github.com/docx-go/text"
)

type SectionType int

const (
	SectionTypeNextPage SectionType = iota
	SectionTypeContinuous
	SectionTypeEvenPage
	SectionTypeOddPage
)

type Orientation int

const (
	OrientationPortrait Orientation = iota
	OrientationLandscape
)

type Section struct {
	element *oxml.Element
	parent  interface{}
}

func NewSection(element *oxml.Element, parent interface{}) *Section {
	return &Section{
		element: element,
		parent:  parent,
	}
}

func (s *Section) Element() *oxml.Element {
	return s.element
}

func (s *Section) PageWidth() *shared.Length {
	pgSz := s.element.Find("w:pgSz")
	if pgSz == nil {
		return nil
	}
	val, exists := pgSz.GetAttr("w:w")
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

func (s *Section) SetPageWidth(width *shared.Length) {
	pgSz := s.getOrAddPgSz()
	if width == nil {
		pgSz.RemoveAttr("w:w")
	} else {
		pgSz.SetAttr("w:w", fmt.Sprintf("%d", width.EMU()))
	}
}

func (s *Section) PageHeight() *shared.Length {
	pgSz := s.element.Find("w:pgSz")
	if pgSz == nil {
		return nil
	}
	val, exists := pgSz.GetAttr("w:h")
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

func (s *Section) SetPageHeight(height *shared.Length) {
	pgSz := s.getOrAddPgSz()
	if height == nil {
		pgSz.RemoveAttr("w:h")
	} else {
		pgSz.SetAttr("w:h", fmt.Sprintf("%d", height.EMU()))
	}
}

func (s *Section) Orientation() Orientation {
	pgSz := s.element.Find("w:pgSz")
	if pgSz == nil {
		return OrientationPortrait
	}
	val, exists := pgSz.GetAttr("w:orient")
	if !exists {
		return OrientationPortrait
	}
	if val == "landscape" {
		return OrientationLandscape
	}
	return OrientationPortrait
}

func (s *Section) SetOrientation(orientation Orientation) {
	pgSz := s.getOrAddPgSz()
	if orientation == OrientationLandscape {
		pgSz.SetAttr("w:orient", "landscape")
	} else {
		pgSz.SetAttr("w:orient", "portrait")
	}
}

func (s *Section) LeftMargin() *shared.Length {
	return s.getMargin("w:left")
}

func (s *Section) SetLeftMargin(margin *shared.Length) {
	s.setMargin("w:left", margin)
}

func (s *Section) RightMargin() *shared.Length {
	return s.getMargin("w:right")
}

func (s *Section) SetRightMargin(margin *shared.Length) {
	s.setMargin("w:right", margin)
}

func (s *Section) TopMargin() *shared.Length {
	return s.getMargin("w:top")
}

func (s *Section) SetTopMargin(margin *shared.Length) {
	s.setMargin("w:top", margin)
}

func (s *Section) BottomMargin() *shared.Length {
	return s.getMargin("w:bottom")
}

func (s *Section) SetBottomMargin(margin *shared.Length) {
	s.setMargin("w:bottom", margin)
}

func (s *Section) HeaderDistance() *shared.Length {
	return s.getMargin("w:header")
}

func (s *Section) SetHeaderDistance(distance *shared.Length) {
	s.setMargin("w:header", distance)
}

func (s *Section) FooterDistance() *shared.Length {
	return s.getMargin("w:footer")
}

func (s *Section) SetFooterDistance(distance *shared.Length) {
	s.setMargin("w:footer", distance)
}

func (s *Section) Gutter() *shared.Length {
	return s.getMargin("w:gutter")
}

func (s *Section) SetGutter(gutter *shared.Length) {
	s.setMargin("w:gutter", gutter)
}

func (s *Section) StartType() SectionType {
	typeElem := s.element.Find("w:type")
	if typeElem == nil {
		return SectionTypeNextPage
	}
	val, exists := typeElem.GetAttr("w:val")
	if !exists {
		return SectionTypeNextPage
	}
	switch val {
	case "continuous":
		return SectionTypeContinuous
	case "evenPage":
		return SectionTypeEvenPage
	case "oddPage":
		return SectionTypeOddPage
	default:
		return SectionTypeNextPage
	}
}

func (s *Section) SetStartType(startType SectionType) {
	typeElem := s.element.Find("w:type")
	if typeElem == nil {
		typeElem = oxml.NewElement("w:type")
		s.element.AddChild(typeElem)
	}
	switch startType {
	case SectionTypeContinuous:
		typeElem.SetAttr("w:val", "continuous")
	case SectionTypeEvenPage:
		typeElem.SetAttr("w:val", "evenPage")
	case SectionTypeOddPage:
		typeElem.SetAttr("w:val", "oddPage")
	default:
		typeElem.SetAttr("w:val", "nextPage")
	}
}

func (s *Section) Header() *Header {
	return NewHeader(s.element, s.parent)
}

func (s *Section) Footer() *Footer {
	return NewFooter(s.element, s.parent)
}

func (s *Section) FirstPageHeader() *Header {
	return NewFirstPageHeader(s.element, s.parent)
}

func (s *Section) FirstPageFooter() *Footer {
	return NewFirstPageFooter(s.element, s.parent)
}

func (s *Section) EvenPageHeader() *Header {
	return NewEvenPageHeader(s.element, s.parent)
}

func (s *Section) EvenPageFooter() *Footer {
	return NewEvenPageFooter(s.element, s.parent)
}

func (s *Section) DifferentFirstPageHeaderFooter() bool {
	titlePg := s.element.Find("w:titlePg")
	if titlePg == nil {
		return false
	}
	val, exists := titlePg.GetAttr("w:val")
	if !exists {
		return true
	}
	return val == "true" || val == "1"
}

func (s *Section) SetDifferentFirstPageHeaderFooter(value bool) {
	titlePg := s.element.Find("w:titlePg")
	if titlePg == nil {
		titlePg = oxml.NewElement("w:titlePg")
		s.element.AddChild(titlePg)
	}
	if value {
		titlePg.SetAttr("w:val", "true")
	} else {
		titlePg.SetAttr("w:val", "false")
	}
}

func (s *Section) getMargin(name string) *shared.Length {
	pgMar := s.element.Find("w:pgMar")
	if pgMar == nil {
		return nil
	}
	val, exists := pgMar.GetAttr(name)
	if !exists {
		return nil
	}
	margin := 0
	for _, c := range val {
		margin = margin*10 + int(c-'0')
	}
	l := shared.Length(margin)
	return &l
}

func (s *Section) setMargin(name string, margin *shared.Length) {
	pgMar := s.getOrAddPgMar()
	if margin == nil {
		pgMar.RemoveAttr(name)
	} else {
		pgMar.SetAttr(name, fmt.Sprintf("%d", margin.EMU()))
	}
}

func (s *Section) getOrAddPgSz() *oxml.Element {
	pgSz := s.element.Find("w:pgSz")
	if pgSz == nil {
		pgSz = oxml.NewElement("w:pgSz")
		s.element.AddChild(pgSz)
	}
	return pgSz
}

func (s *Section) getOrAddPgMar() *oxml.Element {
	pgMar := s.element.Find("w:pgMar")
	if pgMar == nil {
		pgMar = oxml.NewElement("w:pgMar")
		s.element.AddChild(pgMar)
	}
	return pgMar
}

type Sections struct {
	element *oxml.Element
	parent  interface{}
}

func NewSections(element *oxml.Element, parent interface{}) *Sections {
	return &Sections{
		element: element,
		parent:  parent,
	}
}

func (s *Sections) Get(index int) *Section {
	sectPrList := s.element.FindAll("w:sectPr")
	if index < 0 || index >= len(sectPrList) {
		return nil
	}
	return NewSection(sectPrList[index], s.parent)
}

func (s *Sections) Len() int {
	return len(s.element.FindAll("w:sectPr"))
}

func (s *Sections) All() []*Section {
	var sections []*Section
	for _, sectPr := range s.element.FindAll("w:sectPr") {
		sections = append(sections, NewSection(sectPr, s.parent))
	}
	return sections
}

type Header struct {
	element *oxml.Element
	parent  interface{}
}

func NewHeader(element *oxml.Element, parent interface{}) *Header {
	return &Header{
		element: element,
		parent:  parent,
	}
}

func NewFirstPageHeader(element *oxml.Element, parent interface{}) *Header {
	return &Header{
		element: element,
		parent:  parent,
	}
}

func NewEvenPageHeader(element *oxml.Element, parent interface{}) *Header {
	return &Header{
		element: element,
		parent:  parent,
	}
}

func (h *Header) Element() *oxml.Element {
	return h.element
}

func (h *Header) AddParagraph(content, style string) *text.Paragraph {
	p := oxml.NewElement("w:p")
	h.element.AddChild(p)
	para := text.NewParagraph(p, h)
	if content != "" {
		para.AddRun(content, style)
	}
	if style != "" {
		para.SetStyle(style)
	}
	return para
}

func (h *Header) Paragraphs() []*text.Paragraph {
	var paragraphs []*text.Paragraph
	for _, p := range h.element.FindAll("w:p") {
		paragraphs = append(paragraphs, text.NewParagraph(p, h))
	}
	return paragraphs
}

type Footer struct {
	element *oxml.Element
	parent  interface{}
}

func NewFooter(element *oxml.Element, parent interface{}) *Footer {
	return &Footer{
		element: element,
		parent:  parent,
	}
}

func NewFirstPageFooter(element *oxml.Element, parent interface{}) *Footer {
	return &Footer{
		element: element,
		parent:  parent,
	}
}

func NewEvenPageFooter(element *oxml.Element, parent interface{}) *Footer {
	return &Footer{
		element: element,
		parent:  parent,
	}
}

func (f *Footer) Element() *oxml.Element {
	return f.element
}

func (f *Footer) AddParagraph(content, style string) *text.Paragraph {
	p := oxml.NewElement("w:p")
	f.element.AddChild(p)
	para := text.NewParagraph(p, f)
	if content != "" {
		para.AddRun(content, style)
	}
	if style != "" {
		para.SetStyle(style)
	}
	return para
}

func (f *Footer) Paragraphs() []*text.Paragraph {
	var paragraphs []*text.Paragraph
	for _, p := range f.element.FindAll("w:p") {
		paragraphs = append(paragraphs, text.NewParagraph(p, f))
	}
	return paragraphs
}
