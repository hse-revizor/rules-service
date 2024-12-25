package models

type RuleType string

const (
	RuleType_HasFile                       RuleType = "HasFile"
	RuleType_HasStringInFile               RuleType = "HasStringInFile"
	RuleType_HasExpectedValueInField       RuleType = "HasExpectedValueInField"
	RuleType_StrictEquality                RuleType = "StrictEquality"
	RuleType_HasSubstring                  RuleType = "HasSubstring"
	RuleType_HasRegexMatch                 RuleType = "HasRegexMatch"
	RuleType_NoSubstring                   RuleType = "NoSubstring"
	RuleType_NotLongerThan                 RuleType = "NotLongerThan"
	RuleType_NotShorterThan                RuleType = "NotShorterThan"
	RuleType_NotEmpty                      RuleType = "NotEmpty"
	RuleType_DoesLLMSayThatRuleIsSatisfied RuleType = "DoesLLMSayThatRuleIsSatisfied"
)
