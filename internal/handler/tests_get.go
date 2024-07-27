package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetTests(c *gin.Context) {
	ctx := c.Request.Context()

	test, err := h.storage.GetTests(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, test)
}
