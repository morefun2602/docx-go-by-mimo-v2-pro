package opc

import (
	"fmt"
	"strings"
)

type Relationship struct {
	rID        string
	reltype    string
	target     interface{}
	baseURI    string
	isExternal bool
}

func newRelationship(rID, reltype string, target interface{}, baseURI string, external bool) *Relationship {
	return &Relationship{
		rID:        rID,
		reltype:    reltype,
		target:     target,
		baseURI:    baseURI,
		isExternal: external,
	}
}

func (r *Relationship) IsExternal() bool {
	return r.isExternal
}

func (r *Relationship) RelType() string {
	return r.reltype
}

func (r *Relationship) RID() string {
	return r.rID
}

func (r *Relationship) TargetPart() (*Part, error) {
	if r.isExternal {
		return nil, fmt.Errorf("TargetPart is undefined when target mode is External")
	}
	part, ok := r.target.(*Part)
	if !ok {
		return nil, fmt.Errorf("target is not a Part")
	}
	return part, nil
}

func (r *Relationship) TargetRef() (string, error) {
	if r.isExternal {
		ref, ok := r.target.(string)
		if !ok {
			return "", fmt.Errorf("external target is not a string")
		}
		return ref, nil
	}
	part, ok := r.target.(*Part)
	if !ok {
		return "", fmt.Errorf("target is not a Part")
	}
	return part.Partname().RelativeRef(r.baseURI), nil
}

type Relationships struct {
	rels             map[string]*Relationship
	baseURI          string
	targetPartsByRID map[string]interface{}
}

func NewRelationships(baseURI string) *Relationships {
	return &Relationships{
		rels:             make(map[string]*Relationship),
		baseURI:          baseURI,
		targetPartsByRID: make(map[string]interface{}),
	}
}

func (r *Relationships) AddRelationship(reltype string, target interface{}, rID string, isExternal bool) *Relationship {
	rel := newRelationship(rID, reltype, target, r.baseURI, isExternal)
	r.rels[rID] = rel
	if !isExternal {
		r.targetPartsByRID[rID] = target
	}
	return rel
}

func (r *Relationships) GetOrAdd(reltype string, targetPart *Part) *Relationship {
	rel := r.getMatching(reltype, targetPart, false)
	if rel == nil {
		rID := r.nextRID()
		rel = r.AddRelationship(reltype, targetPart, rID, false)
	}
	return rel
}

func (r *Relationships) GetOrAddExtRel(reltype, targetRef string) string {
	rel := r.getMatching(reltype, targetRef, true)
	if rel == nil {
		rID := r.nextRID()
		rel = r.AddRelationship(reltype, targetRef, rID, true)
	}
	return rel.rID
}

func (r *Relationships) PartWithReltype(reltype string) (*Part, error) {
	rel, err := r.getRelOfType(reltype)
	if err != nil {
		return nil, err
	}
	return rel.TargetPart()
}

func (r *Relationships) RelatedParts() map[string]interface{} {
	return r.targetPartsByRID
}

func (r *Relationships) Values() []*Relationship {
	var rels []*Relationship
	for _, rel := range r.rels {
		rels = append(rels, rel)
	}
	return rels
}

func (r *Relationships) Get(rID string) (*Relationship, bool) {
	rel, ok := r.rels[rID]
	return rel, ok
}

func (r *Relationships) Delete(rID string) {
	delete(r.rels, rID)
	delete(r.targetPartsByRID, rID)
}

func (r *Relationships) XML() (string, error) {
	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`)
	sb.WriteString(`<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">`)
	for _, rel := range r.rels {
		sb.WriteString(`<Relationship`)
		sb.WriteString(fmt.Sprintf(` Id="%s"`, rel.rID))
		sb.WriteString(fmt.Sprintf(` Type="%s"`, rel.reltype))
		ref, _ := rel.TargetRef()
		sb.WriteString(fmt.Sprintf(` Target="%s"`, ref))
		if rel.isExternal {
			sb.WriteString(` TargetMode="External"`)
		}
		sb.WriteString(`/>`)
	}
	sb.WriteString(`</Relationships>`)
	return sb.String(), nil
}

func (r *Relationships) getMatching(reltype string, target interface{}, isExternal bool) *Relationship {
	for _, rel := range r.rels {
		if rel.reltype != reltype {
			continue
		}
		if rel.isExternal != isExternal {
			continue
		}
		var relTarget interface{}
		if rel.isExternal {
			relTarget = rel.target
		} else {
			relTarget = rel.target
		}
		if relTarget == target {
			return rel
		}
	}
	return nil
}

func (r *Relationships) getRelOfType(reltype string) (*Relationship, error) {
	var matching []*Relationship
	for _, rel := range r.rels {
		if rel.reltype == reltype {
			matching = append(matching, rel)
		}
	}
	if len(matching) == 0 {
		return nil, fmt.Errorf("no relationship of type '%s' in collection", reltype)
	}
	if len(matching) > 1 {
		return nil, fmt.Errorf("multiple relationships of type '%s' in collection", reltype)
	}
	return matching[0], nil
}

func (r *Relationships) nextRID() string {
	for n := 1; n <= len(r.rels)+1; n++ {
		rIDCandidate := fmt.Sprintf("rId%d", n)
		if _, exists := r.rels[rIDCandidate]; !exists {
			return rIDCandidate
		}
	}
	return fmt.Sprintf("rId%d", len(r.rels)+1)
}
