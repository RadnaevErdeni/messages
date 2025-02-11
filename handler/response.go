package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorz struct {
	Mes string `json:"message"`
}

func errResponse(c *gin.Context, statusCode int, mes string) {
	c.JSON(statusCode, errorz{Mes: mes})
}
func logError(err error, mes string) {
	if err != nil {
		logrus.WithError(err).Error(mes)
	} else {
		logrus.Error(mes)
	}
}
