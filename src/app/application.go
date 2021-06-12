package app

import (
	"github.com/gin-gonic/gin"

	"github.com/kirankothule/bookstore_oauth_api/src/http"
	"github.com/kirankothule/bookstore_oauth_api/src/repository/db"
	"github.com/kirankothule/bookstore_oauth_api/src/repository/rest"
	"github.com/kirankothule/bookstore_oauth_api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(
		access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
