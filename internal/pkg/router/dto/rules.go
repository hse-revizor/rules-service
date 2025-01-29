package dto

type GetRuleDto struct {
    ID         string `json:"id"`
    TemplateID string `json:"templateId"`
    ApplyToURI string `json:"applyToURI"`
    Value      string `json:"value"`
}

type CreateRuleDto struct {
    TemplateID string `json:"templateId" binding:"required"`
    ApplyToURI string `json:"applyToURI" binding:"required"`
    Value      string `json:"value" binding:"required"`
}

type GetPolicyDto struct {
    ID        string   `json:"id"`
    ProjectID string   `json:"projectId"`
    RulesIDs  []string `json:"rulesIds"`
}

type CreatePolicyDto struct {
    ProjectID string   `json:"projectId" binding:"required"`
    RulesIDs  []string `json:"rulesIds" binding:"required"`
}