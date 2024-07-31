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
	message := ms.NewMessage{
		Id:      id,
		Key:     "",
		Payload: input.Payload,
	}

	err = h.services.SendToKafka(c.Request.Context(), message)
	if err != nil {
		err = h.services.UpdateStatusErr(c.Request.Context(), id, "error")
		logError(err, "Failed to send message to Kafka")
		errResponse(c, http.StatusInternalServerError, "Failed to send message to Kafka")
		return
	} else {
		err = h.services.Message.UpdateStatus(c.Request.Context(), id, "processed")
		if err != nil {
			err = h.services.UpdateStatusErr(c.Request.Context(), id, "error")
			logError(err, "Failed to mark message as processed")
			errResponse(c, http.StatusInternalServerError, "Failed to mark message as processed")
			return
		}
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

/*
func (h *Handler) readMessageFromKafka(c *gin.Context) {
	ctx := context.Background()
	message, err := h.services.ReadMessageFromKafka(ctx)
	if err != nil {
		logrus.Errorf("Failed to read message from Kafka: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
*/
