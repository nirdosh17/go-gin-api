package routes

import (
	"log"
	"net/http"

	pg_store "gin-api/storage"
	"gin-api/util"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	p := pg_store.Product{}
	tx := pg_store.DBClient.DB.Find(&p, id)
	if tx.Error != nil {
		log.Printf("Failed to read product id: %v Err: %v\n", id, tx.Error)
		util.NotifyError(c, tx.Error)
		c.Status(http.StatusInternalServerError)
		return
	}
	if p.ID != 0 {
		c.JSON(http.StatusOK, p)
	} else {
		c.Status(http.StatusNotFound)
	}
}
