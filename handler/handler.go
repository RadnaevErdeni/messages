package handler

import (
	"messageService/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/api")
	{
		user.POST("/message", h.createMessage)
		user.GET("/message", h.statusMessage)
	}

	return router
}
