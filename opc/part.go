package opc

import (
	"encoding/xml"
	"fmt"
)

type Part struct {
	partname    PackURI
	contentType string
	blob        []byte
	pkg         *Package
	rels        *Relationships
}

func NewPart(partname PackURI, contentType string, blob []byte, pkg *Package) *Part {
	return &Part{
		partname:    partname,
		contentType: contentType,
		blob:        blob,
		pkg:         pkg,
		rels:        NewRelationships(partname.BaseURI()),
	}
}

func (p *Part) AfterUnmarshal() {}

func (p *Part) BeforeMarshal() {}

func (p *Part) Blob() []byte {
	if p.blob == nil {
		return []byte{}
	}
	return p.blob
}

func (p *Part) ContentType() string {
	return p.contentType
}

func (p *Part) DropRel(rID string) {
	if p.relRefCount(rID) < 2 {
		p.rels.Delete(rID)
	}
}

func (p *Part) Load(reltype string, target interface{}, rID string, isExternal bool) *Relationship {
	return p.rels.AddRelationship(reltype, target, rID, isExternal)
}

func (p *Part) Package() *Package {
	return p.pkg
}

func (p *Part) Partname() PackURI {
	return p.partname
}

func (p *Part) SetPartname(partname PackURI) error {
	p.partname = partname
	return nil
}

func (p *Part) PartRelatedBy(reltype string) (*Part, error) {
	return p.rels.PartWithReltype(reltype)
}

func (p *Part) RelateTo(target interface{}, reltype string, isExternal bool) string {
	if isExternal {
		return p.rels.GetOrAddExtRel(reltype, target.(string))
	}
	rel := p.rels.GetOrAdd(reltype, target.(*Part))
	return rel.rID
}

func (p *Part) RelatedParts() map[string]interface{} {
	return p.rels.RelatedParts()
}

func (p *Part) Rels() *Relationships {
	return p.rels
}

func (p *Part) TargetRef(rID string) (string, error) {
	rel, ok := p.rels.Get(rID)
	if !ok {
		return "", fmt.Errorf("relationship %s not found", rID)
	}
	return rel.TargetRef()
}

func (p *Part) relRefCount(rID string) int {
	return 0
}

type XmlPart struct {
	Part
	element interface{}
}

func NewXmlPart(partname PackURI, contentType string, element interface{}, pkg *Package) *XmlPart {
	return &XmlPart{
		Part: Part{
			partname:    partname,
			contentType: contentType,
			pkg:         pkg,
			rels:        NewRelationships(partname.BaseURI()),
		},
		element: element,
	}
}

func (xp *XmlPart) Blob() []byte {
	if xp.element == nil {
		return xp.Part.Blob()
	}
	data, err := xml.Marshal(xp.element)
	if err != nil {
		return xp.Part.Blob()
	}
	return data
}

func (xp *XmlPart) Element() interface{} {
	return xp.element
}

type PartFactory struct {
	partClassSelector func(contentType, reltype string) *Part
	partTypeFor       map[string]*Part
	defaultPartType   *Part
}

var globalPartFactory = &PartFactory{
	partTypeFor:     make(map[string]*Part),
	defaultPartType: &Part{},
}

func GetPartFactory() *PartFactory {
	return globalPartFactory
}

func (pf *PartFactory) SetPartClassSelector(fn func(contentType, reltype string) *Part) {
	pf.partClassSelector = fn
}

func (pf *PartFactory) RegisterPartType(contentType string, part *Part) {
	pf.partTypeFor[contentType] = part
}

func (pf *PartFactory) CreatePart(partname PackURI, contentType, reltype string, blob []byte, pkg *Package) *Part {
	if pf.partClassSelector != nil {
		if p := pf.partClassSelector(contentType, reltype); p != nil {
			return NewPart(partname, contentType, blob, pkg)
		}
	}
	if p, ok := pf.partTypeFor[contentType]; ok {
		_ = p
		return NewPart(partname, contentType, blob, pkg)
	}
	return NewPart(partname, contentType, blob, pkg)
}
