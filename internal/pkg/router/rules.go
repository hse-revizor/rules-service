package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hse-revizor/rules-service/internal/pkg/models"
	"github.com/hse-revizor/rules-service/internal/pkg/router/dto"
)
// @Summary Create rule
// @Description In success case returns created rule model. Type must be equal (HasFile, HasStringInFile, HasExpectedValueInField, StrictEquality, HasSubstring, HasRegexMatch, NoSubstring, NotLongerThan, NotShorterThan, NotEmpty, DoesLLMSayThatRuleIsSatisfied)
// @Tags Rule
// @Accept json
// @Param data body RuleCreate true "Rule input"
// @Success 200 "" ""
// @Router /rules [post]
func (h *Handler) CreateRule(c *gin.Context) {
	var req dto.RuleCreate
	err := c.BindJSON(&req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	rule, err := h.service.CreateRule(c, &models.Rule{
		FilePath:    req.FilePath,
		Item:        req.Item,
		ShouldBe:    req.ShouldBe,
		Type:        models.RuleType(req.Type),
		WorkspaceId: req.WorkspaceId,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	responseOK(c, dto.Rule{
		Id:          rule.Id.String(),
		FilePath:    rule.FilePath,
		Item:        rule.Item,
		ShouldBe:    rule.ShouldBe,
		Type:        string(rule.Type),
		WorkspaceId: rule.WorkspaceId,
	})
}
// @Summary Get rule by id
// @Description In success case returns rule model with provided id
// @Tags Rule

// @Param ruleId query string true "Rule id input"
// @Success 200 "" ""
// @Router /rules [get]
func (h *Handler) GetRule(c *gin.Context) {
	ruleId := c.Query("ruleId")
	ruleUUID, err := uuid.Parse(ruleId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	rule, err := h.service.GetRuleById(c, ruleUUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	responseOK(c, dto.Rule{
		Id:          rule.Id.String(),
		FilePath:    rule.FilePath,
		Item:        rule.Item,
		ShouldBe:    rule.ShouldBe,
		Type:        string(rule.Type),
		WorkspaceId: rule.WorkspaceId,
	})
}
// @Summary Get rules with pagination
// @Description In success case returns rule models
// @Tags Rule

// @Param take query int true "Take"
// @Param skip query int true "Skip"
// @Success 200 "" ""
// @Router /rules/all [get]
func (h *Handler) GetRules(c *gin.Context) {
	limit := c.Query("take")
	skip := c.Query("skip")
	var limitInt int
	if limit != "" {
		var err error
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err)
			return
		}
	}
	var skipInt int
	if skip != "" {
		var err error
		skipInt, err = strconv.Atoi(skip)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err)
			return
		}
	}
	rule, err := h.service.GetAllRules(c, skipInt, limitInt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	var res dto.RuleList
	res.Count = int(rule.Count)
	res.Rules = make([]dto.Rule, len(rule.Rules))
	for i, r := range rule.Rules {
		res.Rules[i] = dto.Rule{
			Id:          r.Id.String(),
			FilePath:    r.FilePath,
			Item:        r.Item,
			ShouldBe:    r.ShouldBe,
			Type:        string(r.Type),
			WorkspaceId: r.WorkspaceId,
		}
	}
	responseOK(c, res)
}

// @Summary Update rule
// @Description In success case returns updated rule model. Type must be equal (HasFile, HasStringInFile, HasExpectedValueInField, StrictEquality, HasSubstring, HasRegexMatch, NoSubstring, NotLongerThan, NotShorterThan, NotEmpty, DoesLLMSayThatRuleIsSatisfied)
// @Tags Rule
// @Accept json
// @Param data body RuleUpdate true "Rule input"
// @Success 200 "" ""
// @Router /rules [put]
func (h *Handler) UpdateRule(c *gin.Context) {
	var req dto.RuleUpdate
	err := c.BindJSON(&req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	ruleUUID, err := uuid.Parse(req.Id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	rule, err := h.service.UpdateRule(c, &models.Rule{
		Id:       ruleUUID,
		FilePath: req.FilePath,
		Item:     req.Item,
		ShouldBe: req.ShouldBe,
		Type:     models.RuleType(req.Type),
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	responseOK(c, dto.Rule{
		Id:          rule.Id.String(),
		FilePath:    rule.FilePath,
		Item:        rule.Item,
		ShouldBe:    rule.ShouldBe,
		Type:        string(rule.Type),
		WorkspaceId: rule.WorkspaceId,
	})
}

// @Summary Delete rule
// @Description In success case returns deleted rule model
// @Tags Rule
// @Param ruleId query string true "Rule id input"
// @Success 200 "" ""
// @Router /rules [delete]
func (h *Handler) DeleteRule(c *gin.Context) {
	ruleId := c.Query("ruleId")
	ruleUUID, err := uuid.Parse(ruleId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	rule, err := h.service.DeleteRule(c, ruleUUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	responseOK(c, dto.Rule{
		Id:          rule.Id.String(),
		FilePath:    rule.FilePath,
		Item:        rule.Item,
		ShouldBe:    rule.ShouldBe,
		Type:        string(rule.Type),
		WorkspaceId: rule.WorkspaceId,
	})
}
