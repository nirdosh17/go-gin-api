package main

import (
	"net/http"
	"os"

	"gin-api/middlewares"
	"gin-api/routes"
	pg_store "gin-api/storage"
	"gin-api/util"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	// read these variables from environment later
	dbClient := pg_store.Storage{
		Host:     util.AppConfig.DBHost,
		Port:     util.AppConfig.DBPort,
		DBName:   util.AppConfig.DBName,
		User:     util.AppConfig.DBUser,
		Password: util.AppConfig.DBPassword,
	}

	dbClient.NewSession()
	// run migration when schema(model struct) is changed
	dbClient.DB.AutoMigrate(&pg_store.ApiKey{})
	dbClient.DB.AutoMigrate(&pg_store.Product{})

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	// ------------ Sentry ---------------
	app.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// attach tags if needed
	app.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			if os.Getenv("Environment") != "" {
				hub.Scope().SetTag("Environment", os.Getenv("Environment"))
			}
		}
		ctx.Next()
	})

	// ------------ Sentry ---------------

	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	app.Use(middlewares.Auth)
	app = setupRouter(app)
	app.Run(":8080")
}

func setupRouter(app *gin.Engine) *gin.Engine {
	app.GET("/products/:id", routes.GetProduct)
	app.GET("/products", routes.ListProducts)
	app.POST("/products", routes.CreateProduct)
	app.PUT("/products/:id", routes.UpdateProduct)
	app.DELETE("/products/:id", routes.DeleteProduct)
	return app
}
