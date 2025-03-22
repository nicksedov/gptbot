package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context) (interface{}, error)

func Wrap(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handler(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Status": "Error", 
				"ErrorMessage": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, data)
	}
}