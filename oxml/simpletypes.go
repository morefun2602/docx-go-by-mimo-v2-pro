package oxml

type SimpleType string

const (
	STOnOffTrue  SimpleType = "true"
	STOnOffFalse SimpleType = "false"
	STOnOff1     SimpleType = "1"
	STOnOff0     SimpleType = "0"
)

func ParseOnOff(val string) bool {
	switch val {
	case "true", "1", "":
		return true
	case "false", "0":
		return false
	default:
		return false
	}
}

func FormatOnOff(on bool) string {
	if on {
		return "true"
	}
	return "false"
}

type STAlignH string

const (
	STAlignHLeft    STAlignH = "left"
	STAlignHCenter  STAlignH = "center"
	STAlignHRight   STAlignH = "right"
	STAlignHInside  STAlignH = "inside"
	STAlignHOutside STAlignH = "outside"
)

type STAlignV string

const (
	STAlignVTop     STAlignV = "top"
	STAlignVCenter  STAlignV = "center"
	STAlignVBottom  STAlignV = "bottom"
	STAlignVInside  STAlignV = "inside"
	STAlignVOutside STAlignV = "outside"
)

type STBrClear string

const (
	STBrClearNone  STBrClear = "none"
	STBrClearLeft  STBrClear = "left"
	STBrClearRight STBrClear = "right"
	STBrClearAll   STBrClear = "all"
)

type STBrType string

const (
	STBrTypePage         STBrType = "page"
	STBrTypeColumn       STBrType = "column"
	STBrTypeTextWrapping STBrType = "textWrapping"
)

type STCalendarType string

const (
	STCalendarTypeGregorian     STCalendarType = "gregorian"
	STCalendarTypeHebrew        STCalendarType = "hebrew"
	STCalendarTypeHijri         STCalendarType = "hijri"
	STCalendarTypeTaiwan        STCalendarType = "taiwan"
	STCalendarTypeJapan         STCalendarType = "japan"
	STCalendarTypeThai          STCalendarType = "thai"
	STCalendarTypeKorea         STCalendarType = "korea"
	STCalendarTypeSaka          STCalendarType = "saka"
	STCalendarTypeGregorianXlit STCalendarType = "gregorianXlit"
	STCalendarTypeGregorianUs   STCalendarType = "gregorianUs"
	STCalendarTypeJapanXlit     STCalendarType = "japanXlit"
	STCalendarTypeJapanXlit2    STCalendarType = "japanXlit2"
	STCalendarTypeNone          STCalendarType = "none"
)

type STChapterSep string

const (
	STChapterSepHyphen STChapterSep = "hyphen"
	STChapterSepPeriod STChapterSep = "period"
	STChapterSepColon  STChapterSep = "colon"
	STChapterSepEmDash STChapterSep = "emDash"
	STChapterSepEnDash STChapterSep = "endash"
)

type STCharacterSpacing string

const (
	STCharacterSpacingDoNotCompress                  STCharacterSpacing = "doNotCompress"
	STCharacterSpacingCompressPunctuation            STCharacterSpacing = "compressPunctuation"
	STCharacterSpacingCompressPunctuationAndJapanese STCharacterSpacing = "compressPunctuationAndJapanese"
)

type STColorSchemeIndex string

const (
	STColorSchemeIndexDk1      STColorSchemeIndex = "dk1"
	STColorSchemeIndexLt1      STColorSchemeIndex = "lt1"
	STColorSchemeIndexDk2      STColorSchemeIndex = "dk2"
	STColorSchemeIndexLt2      STColorSchemeIndex = "lt2"
	STColorSchemeIndexAccent1  STColorSchemeIndex = "accent1"
	STColorSchemeIndexAccent2  STColorSchemeIndex = "accent2"
	STColorSchemeIndexAccent3  STColorSchemeIndex = "accent3"
	STColorSchemeIndexAccent4  STColorSchemeIndex = "accent4"
	STColorSchemeIndexAccent5  STColorSchemeIndex = "accent5"
	STColorSchemeIndexAccent6  STColorSchemeIndex = "accent6"
	STColorSchemeIndexHlink    STColorSchemeIndex = "hlink"
	STColorSchemeIndexFolHlink STColorSchemeIndex = "folHlink"
)

type STCombineBrackets string

const (
	STCombineBracketsNone   STCombineBrackets = "none"
	STCombineBracketsParens STCombineBrackets = "parens"
	STCombineBracketsSquare STCombineBrackets = "square"
	STCombineBracketsAngle  STCombineBrackets = "angle"
	STCombineBracketsCurly  STCombineBrackets = "curly"
)

type STCryptProv string

const (
	STCryptProvRSAAES  STCryptProv = "rsaAES"
	STCryptProvRSAFull STCryptProv = "rsaFull"
	STCryptProvCustom  STCryptProv = "custom"
)

type STDirection string

const (
	STDirectionLTR STDirection = "ltr"
	STDirectionRTL STDirection = "rtl"
)

type STDisplacedByCustomXml string

const (
	STDisplacedByCustomXmlNext STDisplacedByCustomXml = "next"
	STDisplacedByCustomXmlPrev STDisplacedByCustomXml = "prev"
)

type STDocProtect string

const (
	STDocProtectNone           STDocProtect = "none"
	STDocProtectReadOnly       STDocProtect = "readOnly"
	STDocProtectComments       STDocProtect = "comments"
	STDocProtectTrackedChanges STDocProtect = "trackedChanges"
	STDocProtectForms          STDocProtect = "forms"
)

type STDocType string

const (
	STDocTypeNotSpecified STDocType = "notSpecified"
	STDocTypeLetter       STDocType = "letter"
	STDocTypeEmail        STDocType = "e-mail"
)

type STDraftCap string

const (
	STDraftCapOn  STDraftCap = "on"
	STDraftCapOff STDraftCap = "off"
)

type STDropCap string

const (
	STDropCapNone   STDropCap = "none"
	STDropCapDrop   STDropCap = "drop"
	STDropCapMargin STDropCap = "margin"
)

type STEdGrp string

const (
	STEdGrpNone         STEdGrp = "none"
	STEdGrpEveryone     STEdGrp = "everyone"
	STEdGrpAdmins       STEdGrp = "administrators"
	STEdGrpContributors STEdGrp = "contributors"
	STEdGrpEditors      STEdGrp = "editors"
	STEdGrpOwners       STEdGrp = "owners"
	STEdGrpCurrent      STEdGrp = "current"
)

type STEm string

const (
	STEmNone     STEm = "none"
	STEmDot      STEm = "dot"
	STEmComma    STEm = "comma"
	STEmCircle   STEm = "circle"
	STEmUnderDot STEm = "underDot"
)

type STFFTextType string

const (
	STFFTextTypeRegular     STFFTextType = "regular"
	STFFTextTypeNumber      STFFTextType = "number"
	STFFTextTypeDate        STFFTextType = "date"
	STFFTextTypeCurrentDate STFFTextType = "currentDate"
	STFFTextTypeCurrentTime STFFTextType = "currentTime"
	STFFTextTypeCalc        STFFTextType = "calc"
)

type STFldCharType string

const (
	STFldCharTypeBegin    STFldCharType = "begin"
	STFldCharTypeSeparate STFldCharType = "separate"
	STFldCharTypeEnd      STFldCharType = "end"
)

type STFontFamily string

const (
	STFontFamilyDecorative STFontFamily = "decorative"
	STFontFamilyModern     STFontFamily = "modern"
	STFontFamilyRoman      STFontFamily = "roman"
	STFontFamilyScript     STFontFamily = "script"
	STFontFamilySwiss      STFontFamily = "swiss"
	STFontFamilyAuto       STFontFamily = "auto"
)

type STFrameLayout string

const (
	STFrameLayoutDefault STFrameLayout = "default"
	STFrameLayoutRows    STFrameLayout = "rows"
	STFrameLayoutCols    STFrameLayout = "cols"
	STFrameLayoutNone    STFrameLayout = "none"
)

type STFrameScrollbar string

const (
	STFrameScrollbarAuto STFrameScrollbar = "auto"
	STFrameScrollbarYes  STFrameScrollbar = "yes"
	STFrameScrollbarNo   STFrameScrollbar = "no"
)

type STFtnEdn string

const (
	STFtnEdnNormal                STFtnEdn = "normal"
	STFtnEdnSep                   STFtnEdn = "separator"
	STFtnEdnContinuationSeparator STFtnEdn = "continuationSeparator"
	STFtnEdnContinuationNotice    STFtnEdn = "continuationNotice"
)

type STFtnPos string

const (
	STFtnPosPageBottom  STFtnPos = "pageBottom"
	STFtnPosBeneathText STFtnPos = "beneathText"
	STFtnPosDocEnd      STFtnPos = "docEnd"
	STFtnPosSectEnd     STFtnPos = "sectEnd"
)

type STHAnchor string

const (
	STHAnchorMargin STHAnchor = "margin"
	STHAnchorPage   STHAnchor = "page"
	STHAnchorText   STHAnchor = "text"
)

type STHdrFtr string

const (
	STHdrFtrDefault STHdrFtr = "default"
	STHdrFtrEven    STHdrFtr = "even"
	STHdrFtrFirst   STHdrFtr = "first"
)

type STHeightRule string

const (
	STHeightRuleAuto    STHeightRule = "auto"
	STHeightRuleExact   STHeightRule = "exact"
	STHeightRuleAtLeast STHeightRule = "atLeast"
)

type STHexColor string

const (
	STHexColorAuto STHexColor = "auto"
)

type STHint string

const (
	STHintDefault  STHint = "default"
	STHintEastAsia STHint = "eastAsia"
)

type STHpsMeasure string

type STInfoDocType string

const (
	STInfoDocTypeNotSpecified STInfoDocType = "notSpecified"
	STInfoDocTypeLetter       STInfoDocType = "letter"
	STInfoDocTypeEmail        STInfoDocType = "e-mail"
)

type STJc string

const (
	STJcStart          STJc = "start"
	STJcCenter         STJc = "center"
	STJcEnd            STJc = "end"
	STJcBoth           STJc = "both"
	STJcMediumKashida  STJc = "mediumKashida"
	STJcDistribute     STJc = "distribute"
	STJcNumTab         STJc = "numTab"
	STJcHighKashida    STJc = "highKashida"
	STJcLowKashida     STJc = "lowKashida"
	STJcThaiDistribute STJc = "thaiDistribute"
)

type STJcTable string

const (
	STJcTableStart  STJcTable = "start"
	STJcTableCenter STJcTable = "center"
	STJcTableEnd    STJcTable = "end"
)

type STLevelSuffix string

const (
	STLevelSuffixTab     STLevelSuffix = "tab"
	STLevelSuffixSpace   STLevelSuffix = "space"
	STLevelSuffixNothing STLevelSuffix = "nothing"
)

type STLineNumberRestart string

const (
	STLineNumberRestartNewPage    STLineNumberRestart = "newPage"
	STLineNumberRestartNewSection STLineNumberRestart = "newSection"
	STLineNumberRestartContinuous STLineNumberRestart = "continuous"
)

type STLineSpacingRule string

const (
	STLineSpacingRuleAuto    STLineSpacingRule = "auto"
	STLineSpacingRuleExact   STLineSpacingRule = "exact"
	STLineSpacingRuleAtLeast STLineSpacingRule = "atLeast"
)

type STLock string

const (
	STLockUnlocked         STLock = "unlocked"
	STLockContentLocked    STLock = "contentLocked"
	STLockSdtLocked        STLock = "sdtLocked"
	STLockSdtContentLocked STLock = "sdtContentLocked"
)

type STMacroName string

type STMerge string

const (
	STMergeContinue STMerge = "continue"
	STMergeRestart  STMerge = "restart"
)

type STMultiLevelType string

const (
	STMultiLevelTypeHybridMultilevel STMultiLevelType = "hybridMultilevel"
	STMultiLevelTypeMultilevel       STMultiLevelType = "multilevel"
	STMultiLevelTypeSingleLevel      STMultiLevelType = "singleLevel"
)

type STNumberFormat string

const (
	STNumberFormatDecimal                    STNumberFormat = "decimal"
	STNumberFormatUpperRoman                 STNumberFormat = "upperRoman"
	STNumberFormatLowerRoman                 STNumberFormat = "lowerRoman"
	STNumberFormatUpperLetter                STNumberFormat = "upperLetter"
	STNumberFormatLowerLetter                STNumberFormat = "lowerLetter"
	STNumberFormatOrdinal                    STNumberFormat = "ordinal"
	STNumberFormatCardinalText               STNumberFormat = "cardinalText"
	STNumberFormatOrdinalText                STNumberFormat = "ordinalText"
	STNumberFormatHex                        STNumberFormat = "hex"
	STNumberFormatChicago                    STNumberFormat = "chicago"
	STNumberFormatIdeographDigital           STNumberFormat = "ideographDigital"
	STNumberFormatJapaneseCounting           STNumberFormat = "japaneseCounting"
	STNumberFormatAiueo                      STNumberFormat = "aiueo"
	STNumberFormatIroha                      STNumberFormat = "iroha"
	STNumberFormatDecimalHalfWidth           STNumberFormat = "decimalHalfWidth"
	STNumberFormatJapaneseLegal              STNumberFormat = "japaneseLegal"
	STNumberFormatJapaneseDigitalTenThousand STNumberFormat = "japaneseDigitalTenThousand"
	STNumberFormatKoreanCounting             STNumberFormat = "koreanCounting"
	STNumberFormatKoreanDigital              STNumberFormat = "koreanDigital"
	STNumberFormatKoreanLegal                STNumberFormat = "koreanLegal"
	STNumberFormatKoreanDigital2             STNumberFormat = "koreanDigital2"
	STNumberFormatHebrew1                    STNumberFormat = "hebrew1"
	STNumberFormatArabicAlpha                STNumberFormat = "arabicAlpha"
	STNumberFormatHebrew2                    STNumberFormat = "hebrew2"
	STNumberFormatArabicAbjad                STNumberFormat = "arabicAbjad"
	STNumberFormatHindiVowels                STNumberFormat = "hindiVowels"
	STNumberFormatHindiConsonants            STNumberFormat = "hindiConsonants"
	STNumberFormatHindiNumbers               STNumberFormat = "hindiNumbers"
	STNumberFormatHindiCounting              STNumberFormat = "hindiCounting"
	STNumberFormatThaiLetters                STNumberFormat = "thaiLetters"
	STNumberFormatThaiNumbers                STNumberFormat = "thaiNumbers"
	STNumberFormatThaiCounting               STNumberFormat = "thaiCounting"
	STNumberFormatVietnameseCounting         STNumberFormat = "vietnameseCounting"
	STNumberFormatNumberInDash               STNumberFormat = "numberInDash"
	STNumberFormatEnclosedCircle             STNumberFormat = "enclosedCircle"
	STNumberFormatEnclosedFullstop           STNumberFormat = "enclosedFullstop"
	STNumberFormatEnclosedParen              STNumberFormat = "enclosedParen"
	STNumberFormatEnclosedCircleChinese      STNumberFormat = "enclosedCircleChinese"
	STNumberFormatIdeographEnclosedCircle    STNumberFormat = "ideographEnclosedCircle"
	STNumberFormatIdeographTraditional       STNumberFormat = "ideographTraditional"
	STNumberFormatIdeographZodiac            STNumberFormat = "ideographZodiac"
	STNumberFormatIdeographZodiacTraditional STNumberFormat = "ideographZodiacTraditional"
	STNumberFormatTaiwaneseCounting          STNumberFormat = "taiwaneseCounting"
	STNumberFormatIdeographLegalTraditional  STNumberFormat = "ideographLegalTraditional"
	STNumberFormatTaiwaneseCountingThousand  STNumberFormat = "taiwaneseCountingThousand"
	STNumberFormatTaiwaneseDigital           STNumberFormat = "taiwaneseDigital"
	STNumberFormatChineseCounting            STNumberFormat = "chineseCounting"
	STNumberFormatChineseLegalSimplified     STNumberFormat = "chineseLegalSimplified"
	STNumberFormatChineseCountingThousand    STNumberFormat = "chineseCountingThousand"
	STNumberFormatVietnameseCounting2        STNumberFormat = "vietnameseCounting2"
	STNumberFormatBullet                     STNumberFormat = "bullet"
	STNumberFormatGanada                     STNumberFormat = "ganada"
	STNumberFormatChosung                    STNumberFormat = "chosung"
	STNumberFormatGB1                        STNumberFormat = "gb1"
	STNumberFormatGB2                        STNumberFormat = "gb2"
	STNumberFormatGB3                        STNumberFormat = "gb3"
	STNumberFormatGB4                        STNumberFormat = "gb4"
	STNumberFormatZodiac1                    STNumberFormat = "zodiac1"
	STNumberFormatZodiac2                    STNumberFormat = "zodiac2"
	STNumberFormatZodiac3                    STNumberFormat = "zodiac3"
	STNumberFormatTradChinNum1               STNumberFormat = "tradChinNum1"
	STNumberFormatTradChinNum2               STNumberFormat = "tradChinNum2"
	STNumberFormatTradChinNum3               STNumberFormat = "tradChinNum3"
	STNumberFormatTradChinNum4               STNumberFormat = "tradChinNum4"
	STNumberFormatSimpChinNum1               STNumberFormat = "simpChinNum1"
	STNumberFormatSimpChinNum2               STNumberFormat = "simpChinNum2"
	STNumberFormatSimpChinNum3               STNumberFormat = "simpChinNum3"
	STNumberFormatSimpChinNum4               STNumberFormat = "simpChinNum4"
	STNumberFormatHanzi1                     STNumberFormat = "hanja1"
	STNumberFormatHanzi2                     STNumberFormat = "hanja2"
	STNumberFormatHanzi3                     STNumberFormat = "hanja3"
	STNumberFormatHanzi4                     STNumberFormat = "hanja4"
	STNumberFormatNone                       STNumberFormat = "none"
	STNumberFormatCustom                     STNumberFormat = "custom"
)

type STObjectDrawAspect string

const (
	STObjectDrawAspectContent   STObjectDrawAspect = "content"
	STObjectDrawAspectIcon      STObjectDrawAspect = "icon"
	STObjectDrawAspectThumbnail STObjectDrawAspect = "thumbnail"
	STObjectDrawAspectPrint     STObjectDrawAspect = "print"
)

type STObjectUpdateMode string

const (
	STObjectUpdateModeAlways STObjectUpdateMode = "always"
	STObjectUpdateModeOnCall STObjectUpdateMode = "onCall"
)

type STOleUpdateMode string

const (
	STOleUpdateModeAlways STOleUpdateMode = "always"
	STOleUpdateModeOnCall STOleUpdateMode = "onCall"
)

type STPageBorderDisplay string

const (
	STPageBorderDisplayAllPages     STPageBorderDisplay = "allPages"
	STPageBorderDisplayFirstPage    STPageBorderDisplay = "firstPage"
	STPageBorderDisplayNotFirstPage STPageBorderDisplay = "notFirstPage"
)

type STPageBorderOffset string

const (
	STPageBorderOffsetPage STPageBorderOffset = "page"
	STPageBorderOffsetText STPageBorderOffset = "text"
)

type STPageBorderZOrder string

const (
	STPageBorderZOrderFront STPageBorderZOrder = "front"
	STPageBorderZOrderBack  STPageBorderZOrder = "back"
)

type STPageOrientation string

const (
	STPageOrientationPortrait  STPageOrientation = "portrait"
	STPageOrientationLandscape STPageOrientation = "landscape"
)

type STPitch string

const (
	STPitchDefault  STPitch = "default"
	STPitchFixed    STPitch = "fixed"
	STPitchVariable STPitch = "variable"
)

type STProofErr string

const (
	STProofErrSpellStart STProofErr = "spellStart"
	STProofErrSpellEnd   STProofErr = "spellEnd"
	STProofErrGramStart  STProofErr = "gramStart"
	STProofErrGramEnd    STProofErr = "gramEnd"
)

type STProofState string

const (
	STProofStateClean STProofState = "clean"
	STProofStateDirty STProofState = "dirty"
)

type STPTabAlignment string

const (
	STPTabAlignmentLeft   STPTabAlignment = "left"
	STPTabAlignmentCenter STPTabAlignment = "center"
	STPTabAlignmentRight  STPTabAlignment = "right"
)

type STPTabRelativeTo string

const (
	STPTabRelativeToMargin STPTabRelativeTo = "margin"
	STPTabRelativeToIndent STPTabRelativeTo = "indent"
)

type STPTabLeader string

const (
	STPTabLeaderNone       STPTabLeader = "none"
	STPTabLeaderDot        STPTabLeader = "dot"
	STPTabLeaderHyphen     STPTabLeader = "hyphen"
	STPTabLeaderUnderscore STPTabLeader = "underscore"
	STPTabLeaderMiddleDot  STPTabLeader = "middleDot"
)

type STRestartNumber string

const (
	STRestartNumberContinuous STRestartNumber = "continuous"
	STRestartNumberEachPage   STRestartNumber = "eachPage"
	STRestartNumberEachSect   STRestartNumber = "eachSect"
)

type STRubyAlign string

const (
	STRubyAlignCenter           STRubyAlign = "center"
	STRubyAlignDistributeLetter STRubyAlign = "distributeLetter"
	STRubyAlignDistributeSpace  STRubyAlign = "distributeSpace"
	STRubyAlignLeft             STRubyAlign = "left"
	STRubyAlignRight            STRubyAlign = "right"
	STRubyAlignRightVertical    STRubyAlign = "rightVertical"
)

type STSectionMark string

const (
	STSectionMarkContinuous STSectionMark = "continuous"
	STSectionMarkNextColumn STSectionMark = "nextColumn"
	STSectionMarkNextPage   STSectionMark = "nextPage"
	STSectionMarkEvenPage   STSectionMark = "evenPage"
	STSectionMarkOddPage    STSectionMark = "oddPage"
)

type STShd string

const (
	STShdClear                 STShd = "clear"
	STShdDiagCross             STShd = "diagCross"
	STShdDiagStripe            STShd = "diagStripe"
	STShdHorzCross             STShd = "horzCross"
	STShdHorzStripe            STShd = "horzStripe"
	STShdNil                   STShd = "nil"
	STShdPct10                 STShd = "pct10"
	STShdPct12                 STShd = "pct12"
	STShdPct15                 STShd = "pct15"
	STShdPct20                 STShd = "pct20"
	STShdPct25                 STShd = "pct25"
	STShdPct30                 STShd = "pct30"
	STShdPct35                 STShd = "pct35"
	STShdPct37                 STShd = "pct37"
	STShdPct40                 STShd = "pct40"
	STShdPct45                 STShd = "pct45"
	STShdPct5                  STShd = "pct5"
	STShdPct50                 STShd = "pct50"
	STShdPct55                 STShd = "pct55"
	STShdPct60                 STShd = "pct60"
	STShdPct62                 STShd = "pct62"
	STShdPct65                 STShd = "pct65"
	STShdPct70                 STShd = "pct70"
	STShdPct75                 STShd = "pct75"
	STShdPct80                 STShd = "pct80"
	STShdPct85                 STShd = "pct85"
	STShdPct87                 STShd = "pct87"
	STShdPct90                 STShd = "pct90"
	STShdPct95                 STShd = "pct95"
	STShdReverseDiagStripe     STShd = "reverseDiagStripe"
	STShdSolid                 STShd = "solid"
	STShdThinDiagCross         STShd = "thinDiagCross"
	STShdThinDiagStripe        STShd = "thinDiagStripe"
	STShdThinHorzCross         STShd = "thinHorzCross"
	STShdThinHorzStripe        STShd = "thinHorzStripe"
	STShdThinReverseDiagStripe STShd = "thinReverseDiagStripe"
	STShdThinVertStripe        STShd = "thinVertStripe"
	STShdVertStripe            STShd = "vertStripe"
)

type STSignedHpsMeasure int
type STSignedTwipsMeasure int

type STStyleLink string

type STStyleType string

const (
	STStyleTypeParagraph STStyleType = "paragraph"
	STStyleTypeCharacter STStyleType = "character"
	STStyleTypeTable     STStyleType = "table"
	STStyleTypeNumbering STStyleType = "numbering"
)

type STTabJc string

const (
	STTabJcClear   STTabJc = "clear"
	STTabJcStart   STTabJc = "start"
	STTabJcCenter  STTabJc = "center"
	STTabJcEnd     STTabJc = "end"
	STTabJcDecimal STTabJc = "decimal"
	STTabJcBar     STTabJc = "bar"
	STTabJcList    STTabJc = "list"
)

type STTabTlc string

const (
	STTabTlcNone       STTabTlc = "none"
	STTabTlcDot        STTabTlc = "dot"
	STTabTlcHyphen     STTabTlc = "hyphen"
	STTabTlcUnderscore STTabTlc = "underscore"
	STTabTlcHeavy      STTabTlc = "heavy"
	STTabTlcMiddleDot  STTabTlc = "middleDot"
)

type STTargetScreenSz string

const (
	STTargetScreenSz544x376   STTargetScreenSz = "544x376"
	STTargetScreenSz640x480   STTargetScreenSz = "640x480"
	STTargetScreenSz720x512   STTargetScreenSz = "720x512"
	STTargetScreenSz800x600   STTargetScreenSz = "800x600"
	STTargetScreenSz1024x768  STTargetScreenSz = "1024x768"
	STTargetScreenSz1152x882  STTargetScreenSz = "1152x882"
	STTargetScreenSz1152x900  STTargetScreenSz = "1152x900"
	STTargetScreenSz1280x1024 STTargetScreenSz = "1280x1024"
	STTargetScreenSz1600x1200 STTargetScreenSz = "1600x1200"
	STTargetScreenSz1800x1440 STTargetScreenSz = "1800x1440"
	STTargetScreenSz1920x1200 STTargetScreenSz = "1920x1200"
)

type STTblLayoutType string

const (
	STTblLayoutTypeFixed   STTblLayoutType = "fixed"
	STTblLayoutTypeAutofit STTblLayoutType = "autofit"
)

type STTblOverlap string

const (
	STTblOverlapNever   STTblOverlap = "never"
	STTblOverlapOverlap STTblOverlap = "overlap"
)

type STTblStyleOverrideType string

const (
	STTblStyleOverrideTypeWholeTable STTblStyleOverrideType = "wholeTable"
	STTblStyleOverrideTypeFirstRow   STTblStyleOverrideType = "firstRow"
	STTblStyleOverrideTypeLastRow    STTblStyleOverrideType = "lastRow"
	STTblStyleOverrideTypeFirstCol   STTblStyleOverrideType = "firstCol"
	STTblStyleOverrideTypeLastCol    STTblStyleOverrideType = "lastCol"
	STTblStyleOverrideTypeBand1Vert  STTblStyleOverrideType = "band1Vert"
	STTblStyleOverrideTypeBand2Vert  STTblStyleOverrideType = "band2Vert"
	STTblStyleOverrideTypeBand1Horz  STTblStyleOverrideType = "band1Horz"
	STTblStyleOverrideTypeBand2Horz  STTblStyleOverrideType = "band2Horz"
	STTblStyleOverrideTypeNeCell     STTblStyleOverrideType = "neCell"
	STTblStyleOverrideTypeNwCell     STTblStyleOverrideType = "nwCell"
	STTblStyleOverrideTypeSeCell     STTblStyleOverrideType = "seCell"
	STTblStyleOverrideTypeSwCell     STTblStyleOverrideType = "swCell"
)

type STTblWidth string

const (
	STTblWidthNil  STTblWidth = "nil"
	STTblWidthAuto STTblWidth = "auto"
	STTblWidthDxa  STTblWidth = "dxa"
	STTblWidthPct  STTblWidth = "pct"
)

type STTextAlignment string

const (
	STTextAlignmentTop      STTextAlignment = "top"
	STTextAlignmentCenter   STTextAlignment = "center"
	STTextAlignmentBaseline STTextAlignment = "baseline"
	STTextAlignmentBottom   STTextAlignment = "bottom"
	STTextAlignmentAuto     STTextAlignment = "auto"
)

type STTextboxTightWrap string

const (
	STTextboxTightWrapNone             STTextboxTightWrap = "none"
	STTextboxTightWrapAllLines         STTextboxTightWrap = "allLines"
	STTextboxTightWrapFirstAndLastLine STTextboxTightWrap = "firstAndLastLine"
	STTextboxTightWrapFirstLineOnly    STTextboxTightWrap = "firstLineOnly"
	STTextboxTightWrapLastLineOnly     STTextboxTightWrap = "lastLineOnly"
)

type STTheme string

const (
	STThemeMajorEastAsia STTheme = "majorEastAsia"
	STThemeMajorBidi     STTheme = "majorBidi"
	STThemeMajorAscii    STTheme = "majorAscii"
	STThemeMajorHAnsi    STTheme = "majorHAnsi"
	STThemeMinorEastAsia STTheme = "minorEastAsia"
	STThemeMinorBidi     STTheme = "minorBidi"
	STThemeMinorAscii    STTheme = "minorAscii"
	STThemeMinorHAnsi    STTheme = "minorHAnsi"
)

type STThemeColor string

const (
	STThemeColorDark1             STThemeColor = "dk1"
	STThemeColorLight1            STThemeColor = "lt1"
	STThemeColorDark2             STThemeColor = "dk2"
	STThemeColorLight2            STThemeColor = "lt2"
	STThemeColorAccent1           STThemeColor = "accent1"
	STThemeColorAccent2           STThemeColor = "accent2"
	STThemeColorAccent3           STThemeColor = "accent3"
	STThemeColorAccent4           STThemeColor = "accent4"
	STThemeColorAccent5           STThemeColor = "accent5"
	STThemeColorAccent6           STThemeColor = "accent6"
	STThemeColorHyperlink         STThemeColor = "hlink"
	STThemeColorFollowedHyperlink STThemeColor = "folHlink"
	STThemeColorNone              STThemeColor = "none"
	STThemeColorBackground1       STThemeColor = "bg1"
	STThemeColorBackground2       STThemeColor = "bg2"
)

type STTwipsMeasure int

type STUnderline string

const (
	STUnderlineSingle          STUnderline = "single"
	STUnderlineWords           STUnderline = "words"
	STUnderlineDouble          STUnderline = "double"
	STUnderlineDotted          STUnderline = "dotted"
	STUnderlineThick           STUnderline = "thick"
	STUnderlineDash            STUnderline = "dash"
	STUnderlineDotDash         STUnderline = "dotDash"
	STUnderlineDotDotDash      STUnderline = "dotDotDash"
	STUnderlineWave            STUnderline = "wave"
	STUnderlineWavyHeavy       STUnderline = "wavyHeavy"
	STUnderlineWavyDouble      STUnderline = "wavyDouble"
	STUnderlineDottedHeavy     STUnderline = "dottedHeavy"
	STUnderlineDashedHeavy     STUnderline = "dashedHeavy"
	STUnderlineDashDotHeavy    STUnderline = "dashDotHeavy"
	STUnderlineDashDotDotHeavy STUnderline = "dashDotDotHeavy"
	STUnderlineDashLong        STUnderline = "dashLong"
	STUnderlineDashLongHeavy   STUnderline = "dashLongHeavy"
	STUnderlineNone            STUnderline = "none"
)

type STVerticalAlignRun string

const (
	STVerticalAlignRunBaseline    STVerticalAlignRun = "baseline"
	STVerticalAlignRunSuperscript STVerticalAlignRun = "superscript"
	STVerticalAlignRunSubscript   STVerticalAlignRun = "subscript"
)

type STVerticalJc string

const (
	STVerticalJcTop    STVerticalJc = "top"
	STVerticalJcCenter STVerticalJc = "center"
	STVerticalJcBoth   STVerticalJc = "both"
	STVerticalJcBottom STVerticalJc = "bottom"
)

type STView string

const (
	STViewNone        STView = "none"
	STViewPrint       STView = "print"
	STViewOutline     STView = "outline"
	STViewMasterPages STView = "masterPages"
	STViewNormal      STView = "normal"
	STViewWeb         STView = "web"
)

type STWrap string

const (
	STWrapNone         STWrap = "none"
	STWrapSquare       STWrap = "square"
	STWrapTight        STWrap = "tight"
	STWrapThrough      STWrap = "through"
	STWrapTopAndBottom STWrap = "topAndBottom"
)

type STXAlign string

const (
	STXAlignLeft    STXAlign = "left"
	STXAlignCenter  STXAlign = "center"
	STXAlignRight   STXAlign = "right"
	STXAlignInside  STXAlign = "inside"
	STXAlignOutside STXAlign = "outside"
)

type STYAlign string

const (
	STYAlignTop     STYAlign = "top"
	STYAlignCenter  STYAlign = "center"
	STYAlignBottom  STYAlign = "bottom"
	STYAlignInside  STYAlign = "inside"
	STYAlignOutside STYAlign = "outside"
)

type STZoom int

const (
	STZoomNone     STZoom = 0
	STZoomFullPage STZoom = 1
	STZoomTextFit  STZoom = 2
)
