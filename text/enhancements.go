package text

import (
	"github.com/docx-go/oxml"
)

type ParagraphFormat struct {
	element *oxml.Element
}

func NewParagraphFormat(element *oxml.Element) *ParagraphFormat {
	return &ParagraphFormat{element: element}
}

func (pf *ParagraphFormat) Element() *oxml.Element {
	return pf.element
}

func (pf *ParagraphFormat) getOrAddPPr() *oxml.Element {
	pPr := pf.element.Find("w:pPr")
	if pPr == nil {
		pPr = oxml.NewElement("w:pPr")
		children := pf.element.Children()
		if len(children) > 0 {
			pf.element.InsertBefore(pPr, children[0])
		} else {
			pf.element.AddChild(pPr)
		}
	}
	return pPr
}

func formatInt(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(rune('0'+n%10)) + result
		n /= 10
	}
	return result
}

func (pf *ParagraphFormat) getOrAddSpacing() *oxml.Element {
	pPr := pf.getOrAddPPr()
	spacing := pPr.Find("w:spacing")
	if spacing == nil {
		spacing = oxml.NewElement("w:spacing")
		pPr.AddChild(spacing)
	}
	return spacing
}

func (pf *ParagraphFormat) getOrAddInd() *oxml.Element {
	pPr := pf.getOrAddPPr()
	ind := pPr.Find("w:ind")
	if ind == nil {
		ind = oxml.NewElement("w:ind")
		pPr.AddChild(ind)
	}
	return ind
}

func (pf *ParagraphFormat) LineSpacing() *int {
	spacing := pf.element.Find("w:pPr/w:spacing")
	if spacing == nil {
		return nil
	}
	val, exists := spacing.GetAttr("line")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (pf *ParagraphFormat) SetLineSpacing(spacing *int) {
	sp := pf.getOrAddSpacing()
	if spacing == nil {
		sp.RemoveAttr("line")
	} else {
		sp.SetAttr("line", formatInt(*spacing))
		sp.SetAttr("lineRule", "auto")
	}
}

func (pf *ParagraphFormat) SpaceBefore() *int {
	spacing := pf.element.Find("w:pPr/w:spacing")
	if spacing == nil {
		return nil
	}
	val, exists := spacing.GetAttr("before")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (pf *ParagraphFormat) SetSpaceBefore(space *int) {
	sp := pf.getOrAddSpacing()
	if space == nil {
		sp.RemoveAttr("before")
	} else {
		sp.SetAttr("before", formatInt(*space))
	}
}

func (pf *ParagraphFormat) SpaceAfter() *int {
	spacing := pf.element.Find("w:pPr/w:spacing")
	if spacing == nil {
		return nil
	}
	val, exists := spacing.GetAttr("after")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (pf *ParagraphFormat) SetSpaceAfter(space *int) {
	sp := pf.getOrAddSpacing()
	if space == nil {
		sp.RemoveAttr("after")
	} else {
		sp.SetAttr("after", formatInt(*space))
	}
}

func (pf *ParagraphFormat) LeftIndent() *int {
	ind := pf.element.Find("w:pPr/w:ind")
	if ind == nil {
		return nil
	}
	val, exists := ind.GetAttr("left")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (pf *ParagraphFormat) SetLeftIndent(indent *int) {
	i := pf.getOrAddInd()
	if indent == nil {
		i.RemoveAttr("left")
	} else {
		i.SetAttr("left", formatInt(*indent))
	}
}

func (pf *ParagraphFormat) RightIndent() *int {
	ind := pf.element.Find("w:pPr/w:ind")
	if ind == nil {
		return nil
	}
	val, exists := ind.GetAttr("right")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (pf *ParagraphFormat) SetRightIndent(indent *int) {
	i := pf.getOrAddInd()
	if indent == nil {
		i.RemoveAttr("right")
	} else {
		i.SetAttr("right", formatInt(*indent))
	}
}

func (pf *ParagraphFormat) FirstLineIndent() *int {
	ind := pf.element.Find("w:pPr/w:ind")
	if ind == nil {
		return nil
	}
	val, exists := ind.GetAttr("firstLine")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (pf *ParagraphFormat) SetFirstLineIndent(indent *int) {
	i := pf.getOrAddInd()
	if indent == nil {
		i.RemoveAttr("firstLine")
	} else {
		i.SetAttr("firstLine", formatInt(*indent))
	}
}

func (pf *ParagraphFormat) HangingIndent() *int {
	ind := pf.element.Find("w:pPr/w:ind")
	if ind == nil {
		return nil
	}
	val, exists := ind.GetAttr("hanging")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (pf *ParagraphFormat) SetHangingIndent(indent *int) {
	i := pf.getOrAddInd()
	if indent == nil {
		i.RemoveAttr("hanging")
	} else {
		i.SetAttr("hanging", formatInt(*indent))
	}
}

func (pf *ParagraphFormat) KeepNext() bool {
	pPr := pf.element.Find("w:pPr")
	if pPr == nil {
		return false
	}
	return pPr.Find("w:keepNext") != nil
}

func (pf *ParagraphFormat) SetKeepNext(keep bool) {
	pPr := pf.getOrAddPPr()
	elem := pPr.Find("w:keepNext")
	if keep {
		if elem == nil {
			elem = oxml.NewElement("w:keepNext")
			pPr.AddChild(elem)
		}
	} else {
		if elem != nil {
			pPr.RemoveChild(elem)
		}
	}
}

func (pf *ParagraphFormat) KeepLines() bool {
	pPr := pf.element.Find("w:pPr")
	if pPr == nil {
		return false
	}
	return pPr.Find("w:keepLines") != nil
}

func (pf *ParagraphFormat) SetKeepLines(keep bool) {
	pPr := pf.getOrAddPPr()
	elem := pPr.Find("w:keepLines")
	if keep {
		if elem == nil {
			elem = oxml.NewElement("w:keepLines")
			pPr.AddChild(elem)
		}
	} else {
		if elem != nil {
			pPr.RemoveChild(elem)
		}
	}
}

func (pf *ParagraphFormat) PageBreakBefore() bool {
	pPr := pf.element.Find("w:pPr")
	if pPr == nil {
		return false
	}
	return pPr.Find("w:pageBreakBefore") != nil
}

func (pf *ParagraphFormat) SetPageBreakBefore(pageBreak bool) {
	pPr := pf.getOrAddPPr()
	elem := pPr.Find("w:pageBreakBefore")
	if pageBreak {
		if elem == nil {
			elem = oxml.NewElement("w:pageBreakBefore")
			pPr.AddChild(elem)
		}
	} else {
		if elem != nil {
			pPr.RemoveChild(elem)
		}
	}
}

func (pf *ParagraphFormat) WidowControl() bool {
	pPr := pf.element.Find("w:pPr")
	if pPr == nil {
		return true
	}
	elem := pPr.Find("w:widowControl")
	if elem == nil {
		return true
	}
	val, exists := elem.GetAttr("val")
	if !exists {
		return true
	}
	return val == "true" || val == "1"
}

func (pf *ParagraphFormat) SetWidowControl(control bool) {
	pPr := pf.getOrAddPPr()
	elem := pPr.Find("w:widowControl")
	if elem == nil {
		elem = oxml.NewElement("w:widowControl")
		pPr.AddChild(elem)
	}
	if control {
		elem.SetAttr("val", "true")
	} else {
		elem.SetAttr("val", "false")
	}
}

func (pf *ParagraphFormat) OutlineLevel() *int {
	pPr := pf.element.Find("w:pPr")
	if pPr == nil {
		return nil
	}
	elem := pPr.Find("w:outlineLvl")
	if elem == nil {
		return nil
	}
	val, exists := elem.GetAttr("val")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (pf *ParagraphFormat) SetOutlineLevel(level *int) {
	pPr := pf.getOrAddPPr()
	elem := pPr.Find("w:outlineLvl")
	if level == nil {
		if elem != nil {
			pPr.RemoveChild(elem)
		}
		return
	}
	if elem == nil {
		elem = oxml.NewElement("w:outlineLvl")
		pPr.AddChild(elem)
	}
	elem.SetAttr("val", formatInt(*level))
}

type Hyperlink struct {
	element *oxml.Element
	parent  interface{}
}

func NewHyperlink(element *oxml.Element, parent interface{}) *Hyperlink {
	return &Hyperlink{
		element: element,
		parent:  parent,
	}
}

func (h *Hyperlink) Element() *oxml.Element {
	return h.element
}

func (h *Hyperlink) RID() string {
	val, _ := h.element.GetAttr("r:id")
	return val
}

func (h *Hyperlink) SetRID(rID string) {
	h.element.SetAttr("r:id", rID)
}

func (h *Hyperlink) Anchor() string {
	val, _ := h.element.GetAttr("w:anchor")
	return val
}

func (h *Hyperlink) SetAnchor(anchor string) {
	h.element.SetAttr("w:anchor", anchor)
}

func (h *Hyperlink) Runs() []*Run {
	var runs []*Run
	for _, child := range h.element.FindAll("w:r") {
		runs = append(runs, NewRun(child, nil))
	}
	return runs
}

func (h *Hyperlink) Text() string {
	var text string
	for _, run := range h.Runs() {
		text += run.Text()
	}
	return text
}

type TabStop struct {
	element *oxml.Element
}

func NewTabStop(element *oxml.Element) *TabStop {
	return &TabStop{element: element}
}

func (ts *TabStop) Element() *oxml.Element {
	return ts.element
}

func (ts *TabStop) Val() string {
	val, _ := ts.element.GetAttr("val")
	return val
}

func (ts *TabStop) SetVal(val string) {
	ts.element.SetAttr("val", val)
}

func (ts *TabStop) Pos() *int {
	val, exists := ts.element.GetAttr("pos")
	if !exists {
		return nil
	}
	result := 0
	for _, c := range val {
		result = result*10 + int(c-'0')
	}
	return &result
}

func (ts *TabStop) SetPos(pos *int) {
	if pos == nil {
		ts.element.RemoveAttr("pos")
	} else {
		ts.element.SetAttr("pos", formatInt(*pos))
	}
}

func (ts *TabStop) Leader() string {
	val, _ := ts.element.GetAttr("leader")
	return val
}

func (ts *TabStop) SetLeader(leader string) {
	if leader == "" {
		ts.element.RemoveAttr("leader")
	} else {
		ts.element.SetAttr("leader", leader)
	}
}

type TabStops struct {
	element *oxml.Element
}

func NewTabStops(element *oxml.Element) *TabStops {
	return &TabStops{element: element}
}

func (ts *TabStops) Element() *oxml.Element {
	return ts.element
}

func (ts *TabStops) AddTab(val string, pos int, leader string) *TabStop {
	elem := oxml.NewElement("w:tab")
	elem.SetAttr("val", val)
	elem.SetAttr("pos", formatInt(pos))
	if leader != "" {
		elem.SetAttr("leader", leader)
	}
	ts.element.AddChild(elem)
	return NewTabStop(elem)
}

func (ts *TabStops) ClearTab(pos int) {
	for _, child := range ts.element.FindAll("w:tab") {
		tabPos := NewTabStop(child).Pos()
		if tabPos != nil && *tabPos == pos {
			ts.element.RemoveChild(child)
			return
		}
	}
}

func (ts *TabStops) All() []*TabStop {
	var tabs []*TabStop
	for _, child := range ts.element.FindAll("w:tab") {
		tabs = append(tabs, NewTabStop(child))
	}
	return tabs
}

type RenderedPageBreak struct {
	element *oxml.Element
	parent  interface{}
}

func NewRenderedPageBreak(element *oxml.Element, parent interface{}) *RenderedPageBreak {
	return &RenderedPageBreak{
		element: element,
		parent:  parent,
	}
}

func (rpb *RenderedPageBreak) Element() *oxml.Element {
	return rpb.element
}
