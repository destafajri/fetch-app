package routes

import (
	"github.com/destafajri/auth-app/applications/middlewares"
	"github.com/destafajri/fetch-app/applications/repository"
	"github.com/destafajri/fetch-app/applications/ucase/fetch"
	"github.com/destafajri/fetch-app/db"
	"github.com/gin-gonic/gin"
)

func Router() {
	db := db.DBcon()

	//Repository
	fetchRepository := repository.NewFetch(db)

	//ucase
	getFetchData := fetch.NewGetFetchData(fetchRepository)

	//router default setting
	router := gin.Default()
	//versioning api
	api := router.Group("/api")
	
	//middleware api
	api.Use(middlewares.AuthMiddleware())

	//end point
	api.GET("/fetch", getFetchData.FetchDataHandler)
	
	router.Run("0.0.0.0:9000")

}