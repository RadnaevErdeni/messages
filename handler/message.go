package handler

import (
	ms "messageService"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createMessage(c *gin.Context) {
	var input ms.NewMessage
	if err := c.BindJSON(&input); err != nil {
		logError(err, "Invalid message")
		errResponse(c, http.StatusBadRequest, "Invalid message")
		return
	}
	id, err := h.services.Message.CreateMessage(input)
	if err != nil {
		logError(err, "failed to create message")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) statusMessage(c *gin.Context) {
	mes, err := h.services.Message.StatusMessage()
	if err != nil {
		logError(err, "failed to fetch message")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	logrus.WithFields(logrus.Fields{
		"userCount": len(mes),
	}).Debug("Fetched users successfully")

	c.JSON(http.StatusOK, mes)
}
