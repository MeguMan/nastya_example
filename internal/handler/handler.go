package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router  *gin.Engine
	storage storage
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) initRoutes() {
	h.router.POST("/test", h.AddTest)
	h.router.GET("/test/:id", h.GetTest)
	h.router.GET("/test", h.GetTests)
}

func New(storage storage) *Handler {
	h := &Handler{
		router:  gin.New(),
		storage: storage,
	}

	h.initRoutes()

	return h
}
