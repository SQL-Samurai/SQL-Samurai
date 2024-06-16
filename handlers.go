package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your_username/your_project/database"
	"github.com/your_username/your_project/models"
)

func HandleQuery(databaseURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.QueryRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		result, err := database.RunQuery(databaseURL, request.Query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.QueryResult{Result: result})
	}
}
