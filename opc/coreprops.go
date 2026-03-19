package opc

import (
	"encoding/xml"
	"time"
)

type CoreProperties struct {
	part    *Part
	element *CorePropertiesXML
}

type CorePropertiesXML struct {
	XMLName        xml.Name `xml:"coreProperties"`
	Title          string   `xml:"dc:title,omitempty"`
	Subject        string   `xml:"dc:subject,omitempty"`
	Author         string   `xml:"dc:creator,omitempty"`
	Keywords       string   `xml:"cp:keywords,omitempty"`
	Description    string   `xml:"dc:description,omitempty"`
	LastModifiedBy string   `xml:"cp:lastModifiedBy,omitempty"`
	Revision       int      `xml:"cp:revision,omitempty"`
	Created        string   `xml:"dcterms:created,omitempty"`
	Modified       string   `xml:"dcterms:modified,omitempty"`
	Category       string   `xml:"cp:category,omitempty"`
	Comments       string   `xml:"dc:description,omitempty"`
	ContentType    string   `xml:"cp:contentType,omitempty"`
	Language       string   `xml:"dc:language,omitempty"`
	Version        string   `xml:"cp:version,omitempty"`
}

func NewCoreProperties(part *Part) *CoreProperties {
	cp := &CoreProperties{part: part}
	cp.parse()
	return cp
}

func (cp *CoreProperties) parse() {
	cp.element = &CorePropertiesXML{}
	if cp.part != nil && len(cp.part.blob) > 0 {
		xml.Unmarshal(cp.part.blob, cp.element)
	}
}

func (cp *CoreProperties) Title() string {
	if cp.element == nil {
		return ""
	}
	return cp.element.Title
}

func (cp *CoreProperties) SetTitle(title string) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Title = title
}

func (cp *CoreProperties) Subject() string {
	if cp.element == nil {
		return ""
	}
	return cp.element.Subject
}

func (cp *CoreProperties) SetSubject(subject string) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Subject = subject
}

func (cp *CoreProperties) Author() string {
	if cp.element == nil {
		return ""
	}
	return cp.element.Author
}

func (cp *CoreProperties) SetAuthor(author string) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Author = author
}

func (cp *CoreProperties) Keywords() string {
	if cp.element == nil {
		return ""
	}
	return cp.element.Keywords
}

func (cp *CoreProperties) SetKeywords(keywords string) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Keywords = keywords
}

func (cp *CoreProperties) Comments() string {
	if cp.element == nil {
		return ""
	}
	return cp.element.Comments
}

func (cp *CoreProperties) SetComments(comments string) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Comments = comments
}

func (cp *CoreProperties) Category() string {
	if cp.element == nil {
		return ""
	}
	return cp.element.Category
}

func (cp *CoreProperties) SetCategory(category string) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Category = category
}

func (cp *CoreProperties) Revision() int {
	if cp.element == nil {
		return 0
	}
	return cp.element.Revision
}

func (cp *CoreProperties) SetRevision(revision int) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Revision = revision
}

func (cp *CoreProperties) LastModifiedBy() string {
	if cp.element == nil {
		return ""
	}
	return cp.element.LastModifiedBy
}

func (cp *CoreProperties) SetLastModifiedBy(author string) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.LastModifiedBy = author
}

func (cp *CoreProperties) Language() string {
	if cp.element == nil {
		return ""
	}
	return cp.element.Language
}

func (cp *CoreProperties) SetLanguage(language string) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Language = language
}

func (cp *CoreProperties) Version() string {
	if cp.element == nil {
		return ""
	}
	return cp.element.Version
}

func (cp *CoreProperties) SetVersion(version string) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Version = version
}

func (cp *CoreProperties) Created() time.Time {
	if cp.element == nil || cp.element.Created == "" {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, cp.element.Created)
	return t
}

func (cp *CoreProperties) SetCreated(t time.Time) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Created = t.Format(time.RFC3339)
}

func (cp *CoreProperties) Modified() time.Time {
	if cp.element == nil || cp.element.Modified == "" {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, cp.element.Modified)
	return t
}

func (cp *CoreProperties) SetModified(t time.Time) {
	if cp.element == nil {
		cp.element = &CorePropertiesXML{}
	}
	cp.element.Modified = t.Format(time.RFC3339)
}

func (cp *CoreProperties) Blob() []byte {
	if cp.element == nil {
		return nil
	}
	data, _ := xml.Marshal(cp.element)
	return data
}
