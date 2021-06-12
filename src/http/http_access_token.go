package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	atDomian "github.com/kirankothule/bookstore_oauth_api/src/domain/access_token"
	"github.com/kirankothule/bookstore_oauth_api/src/services/access_token"
	"github.com/kirankothule/bookstore_oauth_api/src/utils/errors"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := handler.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at atDomian.AccessTokenRequest
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	accessToken, err := handler.service.Create(at)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
