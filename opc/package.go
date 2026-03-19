package opc

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Package struct {
	rels     *Relationships
	parts    map[PackURI]*Part
	partname PackURI
}

func NewPackage() *Package {
	return &Package{
		rels:     NewRelationships("/"),
		parts:    make(map[PackURI]*Part),
		partname: PackageURI,
	}
}

func (p *Package) AfterUnmarshal() {}

func (p *Package) CoreProperties() *Part {
	part, _ := p.PartRelatedBy(RTCoreProperties)
	return part
}

func (p *Package) IterRels() []*Relationship {
	visited := make(map[PackURI]bool)
	var rels []*Relationship
	p.walkRels(p, &visited, &rels)
	return rels
}

func (p *Package) walkRels(source interface{}, visited *map[PackURI]bool, rels *[]*Relationship) {
	var relsCollection *Relationships
	if pkg, ok := source.(*Package); ok {
		relsCollection = pkg.rels
	} else if part, ok := source.(*Part); ok {
		relsCollection = part.rels
	} else {
		return
	}

	for _, rel := range relsCollection.Values() {
		*rels = append(*rels, rel)
		if rel.isExternal {
			continue
		}
		targetPart, err := rel.TargetPart()
		if err != nil {
			continue
		}
		if (*visited)[targetPart.partname] {
			continue
		}
		(*visited)[targetPart.partname] = true
		p.walkRels(targetPart, visited, rels)
	}
}

func (p *Package) IterParts() []*Part {
	visited := make(map[PackURI]bool)
	var parts []*Part
	p.walkParts(p, &visited, &parts)
	return parts
}

func (p *Package) walkParts(source interface{}, visited *map[PackURI]bool, parts *[]*Part) {
	var relsCollection *Relationships
	if pkg, ok := source.(*Package); ok {
		relsCollection = pkg.rels
	} else if part, ok := source.(*Part); ok {
		relsCollection = part.rels
	} else {
		return
	}

	for _, rel := range relsCollection.Values() {
		if rel.isExternal {
			continue
		}
		targetPart, err := rel.TargetPart()
		if err != nil {
			continue
		}
		if (*visited)[targetPart.partname] {
			continue
		}
		(*visited)[targetPart.partname] = true
		*parts = append(*parts, targetPart)
		p.walkParts(targetPart, visited, parts)
	}
}

func (p *Package) LoadRel(reltype string, target interface{}, rID string, isExternal bool) *Relationship {
	return p.rels.AddRelationship(reltype, target, rID, isExternal)
}

func (p *Package) MainDocumentPart() (*Part, error) {
	return p.PartRelatedBy(RTOfficeDocument)
}

func (p *Package) NextPartname(template string) PackURI {
	partnames := make(map[string]bool)
	for _, part := range p.IterParts() {
		partnames[string(part.partname)] = true
	}
	for n := 1; n <= len(partnames)+1; n++ {
		candidate := fmt.Sprintf(template, n)
		if !partnames[candidate] {
			return PackURI(candidate)
		}
	}
	return PackURI(fmt.Sprintf(template, len(partnames)+1))
}

func Open(pkgFile interface{}) (*Package, error) {
	switch v := pkgFile.(type) {
	case string:
		zipReader, err := zip.OpenReader(v)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer zipReader.Close()
		return readZip(zipReader)
	case io.Reader:
		data, err := io.ReadAll(v)
		if err != nil {
			return nil, fmt.Errorf("failed to read: %w", err)
		}
		zipReader, err := zip.NewReader(strings.NewReader(string(data)), int64(len(data)))
		if err != nil {
			return nil, fmt.Errorf("failed to open zip: %w", err)
		}
		return readZipFromReader(zipReader)
	default:
		return nil, fmt.Errorf("unsupported package file type")
	}
}

func (p *Package) PartRelatedBy(reltype string) (*Part, error) {
	return p.rels.PartWithReltype(reltype)
}

func (p *Package) Parts() []*Part {
	return p.IterParts()
}

func (p *Package) RelateTo(part *Part, reltype string) string {
	rel := p.rels.GetOrAdd(reltype, part)
	return rel.rID
}

func (p *Package) Rels() *Relationships {
	return p.rels
}

func (p *Package) Save(pkgFile interface{}) error {
	for _, part := range p.Parts() {
		part.BeforeMarshal()
	}
	return nil
}

type Unmarshaller struct{}

func (u *Unmarshaller) Unmarshal(pkgReader *PackageReader, pkg *Package, partFactory *PartFactory) error {
	parts, err := u.unmarshalParts(pkgReader, pkg, partFactory)
	if err != nil {
		return err
	}
	if err := u.unmarshalRelationships(pkgReader, pkg, parts); err != nil {
		return err
	}
	for _, part := range parts {
		part.AfterUnmarshal()
	}
	pkg.AfterUnmarshal()
	return nil
}

func (u *Unmarshaller) unmarshalParts(pkgReader *PackageReader, pkg *Package, partFactory *PartFactory) (map[PackURI]*Part, error) {
	parts := make(map[PackURI]*Part)
	for spart := range pkgReader.IterSparts() {
		part := partFactory.CreatePart(spart.partname, spart.contentType, spart.reltype, spart.blob, pkg)
		parts[spart.partname] = part
	}
	return parts, nil
}

func (u *Unmarshaller) unmarshalRelationships(pkgReader *PackageReader, pkg *Package, parts map[PackURI]*Part) error {
	for srel := range pkgReader.IterSrels() {
		var source interface{}
		if srel.sourceURI == "/" {
			source = pkg
		} else {
			source = parts[PackURI(srel.sourceURI)]
		}
		var target interface{}
		if srel.isExternal {
			target = srel.targetRef
		} else {
			target = parts[srel.targetPartname]
		}
		if part, ok := source.(*Part); ok {
			part.Load(srel.reltype, target, srel.rID, srel.isExternal)
		} else if pkg, ok := source.(*Package); ok {
			pkg.LoadRel(srel.reltype, target, srel.rID, srel.isExternal)
		}
	}
	return nil
}

type SPart struct {
	partname    PackURI
	contentType string
	reltype     string
	blob        []byte
}

type SRel struct {
	sourceURI      string
	rID            string
	reltype        string
	targetRef      string
	targetPartname PackURI
	isExternal     bool
}

type PackageReader struct {
	sparts chan SPart
	srels  chan SRel
}

func NewPackageReader() *PackageReader {
	return &PackageReader{
		sparts: make(chan SPart, 100),
		srels:  make(chan SRel, 100),
	}
}

func (pr *PackageReader) IterSparts() <-chan SPart {
	return pr.sparts
}

func (pr *PackageReader) IterSrels() <-chan SRel {
	return pr.srels
}

func (pr *PackageReader) AddSPart(spart SPart) {
	pr.sparts <- spart
}

func (pr *PackageReader) AddSRel(srel SRel) {
	pr.srels <- srel
}

func (pr *PackageReader) Close() {
	close(pr.sparts)
	close(pr.srels)
}

type PackageWriter struct{}

func (pw *PackageWriter) Write(pkgFile interface{}, rels *Relationships, parts []*Part) error {
	var writer io.Writer
	switch v := pkgFile.(type) {
	case string:
		_ = v
		return nil
	case io.Writer:
		writer = v
	default:
		return fmt.Errorf("unsupported package file type")
	}
	_ = writer
	return nil
}

func parseContentTypesXML(data []byte) (map[string]string, error) {
	type Override struct {
		PartName    string `xml:"PartName,attr"`
		ContentType string `xml:"ContentType,attr"`
	}
	type Types struct {
		XMLName  xml.Name   `xml:"Types`
		Override []Override `xml:"Override"`
	}
	var types Types
	if err := xml.Unmarshal(data, &types); err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for _, o := range types.Override {
		result[o.PartName] = o.ContentType
	}
	return result, nil
}

func parseRelationshipsXML(data []byte) ([]SRel, error) {
	type Rel struct {
		ID         string `xml:"Id,attr"`
		Type       string `xml:"Type,attr"`
		Target     string `xml:"Target,attr"`
		TargetMode string `xml:"TargetMode,attr"`
	}
	type Relationships struct {
		XMLName      xml.Name `xml:"Relationships`
		Relationship []Rel    `xml:"Relationship"`
	}
	var rels Relationships
	if err := xml.Unmarshal(data, &rels); err != nil {
		return nil, err
	}
	var result []SRel
	for _, r := range rels.Relationship {
		srel := SRel{
			rID:        r.ID,
			reltype:    r.Type,
			isExternal: r.TargetMode == TargetModeExternal,
		}
		if srel.isExternal {
			srel.targetRef = r.Target
		} else {
			srel.targetPartname = PackURI(r.Target)
		}
		result = append(result, srel)
	}
	return result, nil
}

func readZip(zipReader *zip.ReadCloser) (*Package, error) {
	pkg := NewPackage()
	reader := NewPackageReader()

	contentTypes := make(map[string]string)

	for _, f := range zipReader.File {
		if f.Name == "[Content_Types].xml" {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			data, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, err
			}
			ct, err := parseContentTypesXML(data)
			if err != nil {
				return nil, err
			}
			contentTypes = ct
			continue
		}

		if strings.Contains(f.Name, "_rels/") && strings.HasSuffix(f.Name, ".rels") {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			data, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, err
			}
			rels, err := parseRelationshipsXML(data)
			if err != nil {
				return nil, err
			}
			for _, rel := range rels {
				if strings.HasSuffix(f.Name, ".rels") {
					basePath := strings.TrimSuffix(f.Name, "_rels/"+f.Name[strings.LastIndex(f.Name, "/")+1:])
					if basePath == "" {
						basePath = "/"
					}
					reader.AddSRel(SRel{
						sourceURI:      basePath,
						rID:            rel.rID,
						reltype:        rel.reltype,
						targetRef:      rel.targetRef,
						targetPartname: rel.targetPartname,
						isExternal:     rel.isExternal,
					})
				}
			}
			continue
		}

		contentType := contentTypes["/"+f.Name]
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		data, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, err
		}
		reader.AddSPart(SPart{
			partname:    PackURI("/" + f.Name),
			contentType: contentType,
			blob:        data,
		})
	}

	reader.Close()

	unmarshaller := &Unmarshaller{}
	if err := unmarshaller.Unmarshal(reader, pkg, GetPartFactory()); err != nil {
		return nil, err
	}

	return pkg, nil
}

func readZipFromReader(zipReader *zip.Reader) (*Package, error) {
	pkg := NewPackage()
	reader := NewPackageReader()

	contentTypes := make(map[string]string)

	for _, f := range zipReader.File {
		if f.Name == "[Content_Types].xml" {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			data, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, err
			}
			ct, err := parseContentTypesXML(data)
			if err != nil {
				return nil, err
			}
			contentTypes = ct
			continue
		}

		if strings.Contains(f.Name, "_rels/") && strings.HasSuffix(f.Name, ".rels") {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			data, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, err
			}
			rels, err := parseRelationshipsXML(data)
			if err != nil {
				return nil, err
			}
			for _, rel := range rels {
				if strings.HasSuffix(f.Name, ".rels") {
					basePath := strings.TrimSuffix(f.Name, "_rels/"+f.Name[strings.LastIndex(f.Name, "/")+1:])
					if basePath == "" {
						basePath = "/"
					}
					reader.AddSRel(SRel{
						sourceURI:      basePath,
						rID:            rel.rID,
						reltype:        rel.reltype,
						targetRef:      rel.targetRef,
						targetPartname: rel.targetPartname,
						isExternal:     rel.isExternal,
					})
				}
			}
			continue
		}

		contentType := contentTypes["/"+f.Name]
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		data, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, err
		}
		reader.AddSPart(SPart{
			partname:    PackURI("/" + f.Name),
			contentType: contentType,
			blob:        data,
		})
	}

	reader.Close()

	unmarshaller := &Unmarshaller{}
	if err := unmarshaller.Unmarshal(reader, pkg, GetPartFactory()); err != nil {
		return nil, err
	}

	return pkg, nil
}
