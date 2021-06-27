package routes

import (
	"log"
	"net/http"
	"strconv"

	pg_store "gin-api/storage"
	"gin-api/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProducts(c *gin.Context) {
	products := []pg_store.Product{}
	result := pg_store.DBClient.DB.Scopes(Paginate(c)).Find(&products)
	if result.Error != nil {
		log.Println("Failed to list products: ", result.Error)
		util.NotifyError(c, result.Error)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, products)
}

func Paginate(r *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(r.Query("per_page"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		r.Header("x-page", strconv.Itoa(page))
		r.Header("x-per-page", strconv.Itoa(pageSize))

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
