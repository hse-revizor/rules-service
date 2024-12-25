package dto

type Rule struct {
	Id          string `json:"id"`
	FilePath    string `json:"filePath"`
	Item        string `json:"item"`
	ShouldBe    string `json:"shouldBe"`
	Type        string `json:"type"`
	WorkspaceId string `json:"workspaceId"`
} // @name Rule

type RuleList struct {
	Rules []Rule `json:"rules"`
	Count int    `json:"count"`
} // @name RuleList

type RuleCreate struct {
	FilePath    string `json:"filePath"`
	Item        string `json:"item"`
	ShouldBe    string `json:"shouldBe"`
	Type        string `json:"type"`
	WorkspaceId string `json:"workspaceId"`
} // @name RuleCreate

type RuleUpdate struct {
	Id          string `json:"id"`
	FilePath    string `json:"filePath"`
	Item        string `json:"item"`
	ShouldBe    string `json:"shouldBe"`
	Type        string `json:"type"`
	WorkspaceId string `json:"workspaceId"`
} // @name RuleUpdate
