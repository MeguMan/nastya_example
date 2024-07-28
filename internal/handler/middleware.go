package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRole(c *gin.Context) {
	ctx := c.Request.Context()

	token := c.Request.Header["Token"][0]

	role, err := h.storage.GetRole(ctx, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	if role != "admin" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
