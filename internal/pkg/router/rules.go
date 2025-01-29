package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hse-revizor/rules-service/internal/pkg/router/dto"
	"github.com/hse-revizor/rules-service/internal/pkg/service/rule"
)

// @Summary Create rule
// @Description In success case returns created rule model. Type must be equal (HasFile, HasStringInFile, HasExpectedValueInField, StrictEquality, HasSubstring, HasRegexMatch, NoSubstring, NotLongerThan, NotShorterThan, NotEmpty, DoesLLMSayThatRuleIsSatisfied)
// @Tags Rule
// @Accept json
// @Param data body dto.CreateRuleDto true "Rule input"
// @Success 200 "" ""
// @Router /rules [post]
func (h *Handler) CreateRule(c *gin.Context) {
	var req dto.CreateRuleDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rule, err := h.service.CreateRule(c, &rule.CreateRule{
		FilePath: req.ApplyToURI,
		Value:    req.Value,
		Template: req.TemplateID,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	responseOK(c, rule)
}

// @Summary Get rule by id
// @Description In success case returns rule model with provided id
// @Tags Rule
// @Param id path string true "Rule id input"
// @Success 200 "" ""
// @Router /rule/{id} [get]
func (h *Handler) GetRule(c *gin.Context) {
	id := c.Param("id")
	ruleUUID, err := uuid.Parse(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	rule, err := h.service.GetRuleById(c, ruleUUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	responseOK(c, rule)
}
// @Summary Delete rule by id
// @Description In success case delete rule model with provided id
// @Tags Rule
// @Param id path string true "Rule id input"
// @Success 200 "" ""
// @Router /rule/{id} [delete]
func (h *Handler) DeleteRule(c *gin.Context) {
	id := c.Param("id")
	ruleUUID, err := uuid.Parse(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	rule, err := h.service.DeleteRule(c, ruleUUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	responseOK(c, rule)
}
