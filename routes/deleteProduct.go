package routes

import (
	"log"
	"net/http"

	pg_store "gin-api/storage"
	"gin-api/util"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	tx := pg_store.DBClient.DB.Delete(&pg_store.Product{}, id)
	if tx.Error != nil {
		log.Printf("Failed to delete product id: %v Err: %v\n", id, tx.Error)
		util.NotifyError(c, tx.Error)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
