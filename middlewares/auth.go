package middlewares

import (
	"log"
	"net/http"

	pg_store "gin-api/storage"
	"gin-api/util"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	apiKey := ctx.GetHeader("x-api-key")
	ak := pg_store.ApiKey{}
	tx := pg_store.DBClient.DB.Where(&pg_store.ApiKey{ApiKey: apiKey}).First(&ak)
	if tx.Error != nil {
		log.Println("Failed to read api key: ", tx.Error)
		util.NotifyError(ctx, tx.Error)
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
		return
	}

	if apiKey == "" || ak.ApiKey == "" {
		ctx.Status(http.StatusUnauthorized)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized. Invalid 'x-api-key' header."})
		ctx.Abort()
		return
	}

	ctx.Next()
}
