package routes

import (
	"log"
	"net/http"

	pg_store "gin-api/storage"
	"gin-api/util"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var req pg_store.CreateRequest

	// does basic type validation based on the CreateRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	p := pg_store.Product{Category: req.Category, Price: req.Price, Name: req.Name}
	tx := pg_store.DBClient.DB.Create(&p)
	if tx.Error != nil {
		log.Printf("Failed to save product %+v in database: %v\n", req, tx.Error)
		util.NotifyError(c, tx.Error)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, p)
}
