// Package opc implements Open Packaging Conventions (OPC) support for .docx files.
package opc

import (
	"regexp"
	"strings"
)

// PackURI represents a pack URI in the OPC package.
type PackURI string

var filenameRe = regexp.MustCompile(`^([a-zA-Z]+)([1-9][0-9]*)?$`)

// NewPackURI creates a new PackURI from a string.
func NewPackURI(s string) (PackURI, error) {
	if !strings.HasPrefix(s, "/") {
		return "", ErrInvalidPackURI(s)
	}
	return PackURI(s), nil
}

// BaseURI returns the base URI (directory portion) of the pack URI.
func (p PackURI) BaseURI() string {
	idx := strings.LastIndex(string(p), "/")
	if idx == 0 {
		return "/"
	}
	return string(p)[:idx]
}

// Ext returns the file extension without the leading dot.
func (p PackURI) Ext() string {
	filename := p.Filename()
	idx := strings.LastIndex(filename, ".")
	if idx == -1 {
		return ""
	}
	return filename[idx+1:]
}

// Filename returns the filename portion of the pack URI.
func (p PackURI) Filename() string {
	idx := strings.LastIndex(string(p), "/")
	return string(p)[idx+1:]
}

// Idx returns the numeric index from a tuple partname, or 0 for singleton partnames.
func (p PackURI) Idx() int {
	name := strings.TrimSuffix(p.Filename(), "."+p.Ext())
	matches := filenameRe.FindStringSubmatch(name)
	if matches == nil || matches[2] == "" {
		return 0
	}
	idx := 0
	for _, c := range matches[2] {
		idx = idx*10 + int(c-'0')
	}
	return idx
}

// Membername returns the pack URI without the leading slash.
func (p PackURI) Membername() string {
	if len(p) <= 1 {
		return ""
	}
	return string(p)[1:]
}

// RelativeRef returns a relative reference from the given base URI.
func (p PackURI) RelativeRef(baseURI string) string {
	if baseURI == "/" {
		return string(p)[1:]
	}
	// Simple relative path calculation
	baseParts := strings.Split(baseURI[1:], "/")
	targetParts := strings.Split(string(p)[1:], "/")

	// Find common prefix
	common := 0
	for i := 0; i < len(baseParts) && i < len(targetParts); i++ {
		if baseParts[i] == targetParts[i] {
			common++
		} else {
			break
		}
	}

	// Build relative path
	var result []string
	for i := common; i < len(baseParts); i++ {
		result = append(result, "..")
	}
	result = append(result, targetParts[common:]...)

	return strings.Join(result, "/")
}

// RelsURI returns the pack URI of the .rels part corresponding to this pack URI.
func (p PackURI) RelsURI() PackURI {
	relsFilename := p.Filename() + ".rels"
	relsURI := p.BaseURI() + "/_rels/" + relsFilename
	return PackURI(relsURI)
}

// PackageURI is the pack URI for the package root.
const PackageURI PackURI = "/"

// ContentTypesURI is the pack URI for [Content_Types].xml.
const ContentTypesURI PackURI = "/[Content_Types].xml"
