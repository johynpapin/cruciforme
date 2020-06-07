package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetFormsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"forms": []struct {
			id string
		}{
			{
				id: "AZk45AU",
			},
			{
				id: "0K58AT1",
			},
		},
	})
}

func (h *Handlers) CreateFormHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": "AZk45AU",
	})
}
