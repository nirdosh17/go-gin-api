package routes

import (
	"log"
	"net/http"

	pg_store "gin-api/storage"
	"gin-api/util"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var req pg_store.CreateRequest

	// does basic type validation based on the CreateRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	p := pg_store.Product{}
	tx := pg_store.DBClient.DB.Find(&p, id)
	if tx.Error != nil {
		log.Printf("Failed to read product id(%v): %v\n", p.ID, tx.Error)
		util.NotifyError(c, tx.Error)
		c.Status(http.StatusInternalServerError)
		return
	}

	if p.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	} else {
		// update single column
		tx := pg_store.DBClient.DB.Model(&p).Update("price", req.Price)
		if tx.Error != nil {
			log.Printf("Failed to update product id(%v): %v\n", p.ID, tx.Error)
			util.NotifyError(c, tx.Error)
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, p)
}
