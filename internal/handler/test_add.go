package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddTest(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req AddTestRequest
	)

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.storage.CreateTest(ctx, req.Text)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
}
