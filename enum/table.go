package enum

type CellVerticalAlignment int

const (
	CellVerticalAlignmentTop    CellVerticalAlignment = 0
	CellVerticalAlignmentCenter CellVerticalAlignment = 1
	CellVerticalAlignmentBottom CellVerticalAlignment = 2
)

func (c CellVerticalAlignment) XmlValue() string {
	switch c {
	case CellVerticalAlignmentTop:
		return "top"
	case CellVerticalAlignmentCenter:
		return "center"
	case CellVerticalAlignmentBottom:
		return "bottom"
	default:
		return ""
	}
}

type RowHeightRule int

const (
	RowHeightRuleAuto    RowHeightRule = 0
	RowHeightRuleAtLeast RowHeightRule = 1
	RowHeightRuleExact   RowHeightRule = 2
)

func (r RowHeightRule) XmlValue() string {
	switch r {
	case RowHeightRuleAuto:
		return "auto"
	case RowHeightRuleAtLeast:
		return "atLeast"
	case RowHeightRuleExact:
		return "exact"
	default:
		return ""
	}
}

type TableAlignment int

const (
	TableAlignmentLeft   TableAlignment = 0
	TableAlignmentCenter TableAlignment = 1
	TableAlignmentRight  TableAlignment = 2
)

func (t TableAlignment) XmlValue() string {
	switch t {
	case TableAlignmentLeft:
		return "left"
	case TableAlignmentCenter:
		return "center"
	case TableAlignmentRight:
		return "right"
	default:
		return ""
	}
}

type TableDirection int

const (
	TableDirectionLTR TableDirection = 0
	TableDirectionRTL TableDirection = 1
)

func (t TableDirection) XmlValue() string {
	switch t {
	case TableDirectionLTR:
		return "ltr"
	case TableDirectionRTL:
		return "rtl"
	default:
		return ""
	}
}
