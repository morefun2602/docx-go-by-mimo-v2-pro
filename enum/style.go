package enum

type BuiltinStyle int

const (
	BuiltinStyleBlockQuotation             BuiltinStyle = -85
	BuiltinStyleBodyText                   BuiltinStyle = -67
	BuiltinStyleBodyText2                  BuiltinStyle = -81
	BuiltinStyleBodyText3                  BuiltinStyle = -82
	BuiltinStyleBodyTextFirstIndent        BuiltinStyle = -78
	BuiltinStyleBodyTextFirstIndent2       BuiltinStyle = -79
	BuiltinStyleBodyTextIndent             BuiltinStyle = -68
	BuiltinStyleBodyTextIndent2            BuiltinStyle = -83
	BuiltinStyleBodyTextIndent3            BuiltinStyle = -84
	BuiltinStyleBookTitle                  BuiltinStyle = -265
	BuiltinStyleCaption                    BuiltinStyle = -35
	BuiltinStyleClosing                    BuiltinStyle = -64
	BuiltinStyleCommentReference           BuiltinStyle = -40
	BuiltinStyleCommentText                BuiltinStyle = -31
	BuiltinStyleDate                       BuiltinStyle = -77
	BuiltinStyleDefaultParagraphFont       BuiltinStyle = -66
	BuiltinStyleEmphasis                   BuiltinStyle = -89
	BuiltinStyleEndnoteReference           BuiltinStyle = -43
	BuiltinStyleEndnoteText                BuiltinStyle = -44
	BuiltinStyleEnvelopeAddress            BuiltinStyle = -37
	BuiltinStyleEnvelopeReturn             BuiltinStyle = -38
	BuiltinStyleFooter                     BuiltinStyle = -33
	BuiltinStyleFootnoteReference          BuiltinStyle = -39
	BuiltinStyleFootnoteText               BuiltinStyle = -30
	BuiltinStyleHeader                     BuiltinStyle = -32
	BuiltinStyleHeading1                   BuiltinStyle = -2
	BuiltinStyleHeading2                   BuiltinStyle = -3
	BuiltinStyleHeading3                   BuiltinStyle = -4
	BuiltinStyleHeading4                   BuiltinStyle = -5
	BuiltinStyleHeading5                   BuiltinStyle = -6
	BuiltinStyleHeading6                   BuiltinStyle = -7
	BuiltinStyleHeading7                   BuiltinStyle = -8
	BuiltinStyleHeading8                   BuiltinStyle = -9
	BuiltinStyleHeading9                   BuiltinStyle = -10
	BuiltinStyleHyperlink                  BuiltinStyle = -86
	BuiltinStyleHyperlinkFollowed          BuiltinStyle = -87
	BuiltinStyleIntenseEmphasis            BuiltinStyle = -262
	BuiltinStyleIntenseQuote               BuiltinStyle = -182
	BuiltinStyleIntenseReference           BuiltinStyle = -264
	BuiltinStyleList                       BuiltinStyle = -48
	BuiltinStyleListBullet                 BuiltinStyle = -49
	BuiltinStyleListBullet2                BuiltinStyle = -55
	BuiltinStyleListBullet3                BuiltinStyle = -56
	BuiltinStyleListBullet4                BuiltinStyle = -57
	BuiltinStyleListBullet5                BuiltinStyle = -58
	BuiltinStyleListContinue               BuiltinStyle = -69
	BuiltinStyleListContinue2              BuiltinStyle = -70
	BuiltinStyleListContinue3              BuiltinStyle = -71
	BuiltinStyleListContinue4              BuiltinStyle = -72
	BuiltinStyleListContinue5              BuiltinStyle = -73
	BuiltinStyleListNumber                 BuiltinStyle = -50
	BuiltinStyleListNumber2                BuiltinStyle = -59
	BuiltinStyleListNumber3                BuiltinStyle = -60
	BuiltinStyleListNumber4                BuiltinStyle = -61
	BuiltinStyleListNumber5                BuiltinStyle = -62
	BuiltinStyleListParagraph              BuiltinStyle = -180
	BuiltinStyleMacroText                  BuiltinStyle = -46
	BuiltinStyleMessageHeader              BuiltinStyle = -74
	BuiltinStyleNavPane                    BuiltinStyle = -90
	BuiltinStyleNormal                     BuiltinStyle = -1
	BuiltinStyleNormalIndent               BuiltinStyle = -29
	BuiltinStyleNormalObject               BuiltinStyle = -158
	BuiltinStyleNormalTable                BuiltinStyle = -106
	BuiltinStyleNoteHeading                BuiltinStyle = -80
	BuiltinStylePageNumber                 BuiltinStyle = -42
	BuiltinStylePlainText                  BuiltinStyle = -91
	BuiltinStyleQuote                      BuiltinStyle = -181
	BuiltinStyleSalutation                 BuiltinStyle = -76
	BuiltinStyleSignature                  BuiltinStyle = -65
	BuiltinStyleStrong                     BuiltinStyle = -88
	BuiltinStyleSubtitle                   BuiltinStyle = -75
	BuiltinStyleSubtleEmphasis             BuiltinStyle = -261
	BuiltinStyleSubtleReference            BuiltinStyle = -263
	BuiltinStyleTableColorfulGrid          BuiltinStyle = -172
	BuiltinStyleTableColorfulList          BuiltinStyle = -171
	BuiltinStyleTableColorfulShading       BuiltinStyle = -170
	BuiltinStyleTableDarkList              BuiltinStyle = -169
	BuiltinStyleTableLightGrid             BuiltinStyle = -161
	BuiltinStyleTableLightGridAccent1      BuiltinStyle = -175
	BuiltinStyleTableLightList             BuiltinStyle = -160
	BuiltinStyleTableLightListAccent1      BuiltinStyle = -174
	BuiltinStyleTableLightShading          BuiltinStyle = -159
	BuiltinStyleTableLightShadingAccent1   BuiltinStyle = -173
	BuiltinStyleTableMediumGrid1           BuiltinStyle = -166
	BuiltinStyleTableMediumGrid2           BuiltinStyle = -167
	BuiltinStyleTableMediumGrid3           BuiltinStyle = -168
	BuiltinStyleTableMediumList1           BuiltinStyle = -164
	BuiltinStyleTableMediumList1Accent1    BuiltinStyle = -178
	BuiltinStyleTableMediumList2           BuiltinStyle = -165
	BuiltinStyleTableMediumShading1        BuiltinStyle = -162
	BuiltinStyleTableMediumShading1Accent1 BuiltinStyle = -176
	BuiltinStyleTableMediumShading2        BuiltinStyle = -163
	BuiltinStyleTableMediumShading2Accent1 BuiltinStyle = -177
	BuiltinStyleTableOfAuthorities         BuiltinStyle = -45
	BuiltinStyleTableOfFigures             BuiltinStyle = -36
	BuiltinStyleTitle                      BuiltinStyle = -63
	BuiltinStyleToaHeading                 BuiltinStyle = -47
	BuiltinStyleToc1                       BuiltinStyle = -20
	BuiltinStyleToc2                       BuiltinStyle = -21
	BuiltinStyleToc3                       BuiltinStyle = -22
	BuiltinStyleToc4                       BuiltinStyle = -23
	BuiltinStyleToc5                       BuiltinStyle = -24
	BuiltinStyleToc6                       BuiltinStyle = -25
	BuiltinStyleToc7                       BuiltinStyle = -26
	BuiltinStyleToc8                       BuiltinStyle = -27
	BuiltinStyleToc9                       BuiltinStyle = -28
)

func (s BuiltinStyle) String() string {
	switch s {
	case BuiltinStyleBlockQuotation:
		return "Block Text"
	case BuiltinStyleBodyText:
		return "Body Text"
	case BuiltinStyleHeading1:
		return "Heading 1"
	case BuiltinStyleHeading2:
		return "Heading 2"
	case BuiltinStyleHeading3:
		return "Heading 3"
	case BuiltinStyleNormal:
		return "Normal"
	case BuiltinStyleTitle:
		return "Title"
	default:
		return ""
	}
}

type StyleType int

const (
	StyleTypeCharacter StyleType = 2
	StyleTypeList      StyleType = 4
	StyleTypeParagraph StyleType = 1
	StyleTypeTable     StyleType = 3
)

func (t StyleType) XmlValue() string {
	switch t {
	case StyleTypeCharacter:
		return "character"
	case StyleTypeList:
		return "numbering"
	case StyleTypeParagraph:
		return "paragraph"
	case StyleTypeTable:
		return "table"
	default:
		return ""
	}
}
