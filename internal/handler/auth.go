package handler

import (
	"net/http"
	"standard/internal/model"

	"github.com/gin-gonic/gin"
)

// Login
// @Description Login User
// @Summary Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body model.UserLoginRequest true "Login"
// @Success 200 {object} baseResponse
// @Failure 400 {object} baseResponse
// @Failure 404 {object} baseResponse
// @Failure 500 {object} baseResponse
// @Router /api/auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input model.UserLoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newAbortResponse(c, err.Error())
		return
	}

	newSuccessResponse(c, http.StatusOK, "login success", token)
}
