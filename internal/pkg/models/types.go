package models

type RuleTemplate string

const (
	RuleTemplate_StrictEqualityTempl     RuleTemplate = "StrictEqualityTempl"
	RuleTemplate_StrictInequalityTempl   RuleTemplate = "StrictInequalityTempl"
	RuleTemplate_StrictGreaterTempl      RuleTemplate = "StrictGreaterTempl"
	RuleTemplate_StrictShorterThenTempl  RuleTemplate = "StrictShorterThenTempl"
	RuleTemplate_StrictMatchesRegexTempl RuleTemplate = "StrictMatchesRegexTempl"
	RuleTemplate_NonStrictMathTempl      RuleTemplate = "NonStrictMathTempl"
)
