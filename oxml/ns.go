package oxml

const (
	NSMAP_W   = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"
	NSMAP_R   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"
	NSMAP_A   = "http://schemas.openxmlformats.org/drawingml/2006/main"
	NSMAP_WP  = "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
	NSMAP_PIC = "http://schemas.openxmlformats.org/drawingml/2006/picture"
	NSMAP_C   = "http://schemas.openxmlformats.org/drawingml/2006/chart"
	NSMAP_V   = "urn:schemas-microsoft-com:vml"
	NSMAP_MC  = "http://schemas.openxmlformats.org/markup-compatibility/2006"
	NSMAP_MO  = "http://schemas.microsoft.com/office/mac/office/2008/main"
	NSMAP_MV  = "http://macVmlSchema.microsoft.com"
	NSMAP_O   = "urn:schemas-microsoft-com:office:office"
	NSMAP_RPT = "http://schemas.microsoft.com/office/word/2006/wordml"
	NSMAP_S   = "http://schemas.openxmlformats.org/officeDocument/2006/shared-types"
	NSMAP_WNE = "http://schemas.microsoft.com/office/word/2006/wordml"
	NSMAP_XSI = "http://www.w3.org/2001/XMLSchema-instance"
	NSMAP_W10 = "urn:schemas-microsoft-com:office:word"
	NSMAP_W14 = "http://schemas.microsoft.com/office/word/2010/wordml"
	NSMAP_W15 = "http://schemas.microsoft.com/office/word/2012/wordml"
	NSMAP_SL  = "http://schemas.openxmlformats.org/schemaLibrary/2006/main"
	NSMAP_DSP = "http://schemas.microsoft.com/office/drawing/2008/diagram"
	NSMAP_M   = "http://schemas.openxmlformats.org/officeDocument/2006/math"
	NSMAP_MSO = "urn:schemas-microsoft-com:office:office"
	NSMAP_MXL = "http://schemas.microsoft.com/office/excel/2006/main"
)

var nsmap = map[string]string{
	"w":   NSMAP_W,
	"r":   NSMAP_R,
	"a":   NSMAP_A,
	"wp":  NSMAP_WP,
	"pic": NSMAP_PIC,
	"c":   NSMAP_C,
	"v":   NSMAP_V,
	"mc":  NSMAP_MC,
	"mo":  NSMAP_MO,
	"mv":  NSMAP_MV,
	"o":   NSMAP_O,
	"rpt": NSMAP_RPT,
	"s":   NSMAP_S,
	"wne": NSMAP_WNE,
	"xsi": NSMAP_XSI,
	"w10": NSMAP_W10,
	"w14": NSMAP_W14,
	"w15": NSMAP_W15,
	"sl":  NSMAP_SL,
	"dsp": NSMAP_DSP,
	"m":   NSMAP_M,
	"mso": NSMAP_MSO,
	"mxl": NSMAP_MXL,
}

func NamespaceForPrefix(prefix string) string {
	if ns, ok := nsmap[prefix]; ok {
		return ns
	}
	return ""
}

func PrefixForNamespace(ns string) string {
	for prefix, namespace := range nsmap {
		if namespace == ns {
			return prefix
		}
	}
	return ""
}

func KnownNamespaces() map[string]string {
	result := make(map[string]string)
	for k, v := range nsmap {
		result[k] = v
	}
	return result
}
