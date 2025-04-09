package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hse-revizor/rules-service/internal/pkg/models"
	"github.com/hse-revizor/rules-service/internal/pkg/router/dto"
)

// @Summary Create policy
// @Description Creates a new policy for applying rules to a project
// @Tags Policy
// @Accept json
// @Param data body dto.CreatePolicyDto true "Policy input"
// @Success 201 {object} dto.GetPolicyDto
// @Router /policy [post]
func (h *Handler) CreatePolicy(c *gin.Context) {
	var req dto.CreatePolicyDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rulesIDs := make([]models.StringStr, len(req.RulesIDs))
	for i, id := range req.RulesIDs {
		rulesIDs[i] = models.StringStr{Value: id}
	}

	policy := &models.Policy{
		ProjectID: req.ProjectID,
		RulesIDs:  rulesIDs,
	}

	createdPolicy, err := h.service.CreatePolicy(c, policy)
	if err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response := dto.GetPolicyDto{
		ID:        createdPolicy.ID.String(),
		ProjectID: createdPolicy.ProjectID,
		RulesIDs:  req.RulesIDs,
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary Get policy by id
// @Description Returns policy model with provided id
// @Tags Policy
// @Param id path string true "Policy id"
// @Success 200 {object} dto.GetPolicyDto
// @Router /policy/{id} [get]
func (h *Handler) GetPolicy(c *gin.Context) {
	id := c.Param("id")
	policyUUID, err := uuid.Parse(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	policy, err := h.service.GetPolicyById(c, policyUUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	rulesIDs := make([]string, len(policy.RulesIDs))
	for i, id := range policy.RulesIDs {
		rulesIDs[i] = id.Value
	}

	response := dto.GetPolicyDto{
		ID:        policy.ID.String(),
		ProjectID: policy.ProjectID,
		RulesIDs:  rulesIDs,
	}

	responseOK(c, response)
}

// @Summary Delete policy by id
// @Description Deletes policy with provided id
// @Tags Policy
// @Param id path string true "Policy id"
// @Success 200 {object} dto.GetPolicyDto
// @Router /policy/{id} [delete]
func (h *Handler) DeletePolicy(c *gin.Context) {
	id := c.Param("id")
	policyUUID, err := uuid.Parse(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	policy, err := h.service.DeletePolicy(c, policyUUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	rulesIDs := make([]string, len(policy.RulesIDs))
	for i, id := range policy.RulesIDs {
		rulesIDs[i] = id.Value
	}

	response := dto.GetPolicyDto{
		ID:        policy.ID.String(),
		ProjectID: policy.ProjectID,
		RulesIDs:  rulesIDs,
	}

	responseOK(c, response)
}
