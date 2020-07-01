package http

import (
	"github.com/gin-gonic/gin"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/domain/model"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/domain/services"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type accessTokenHandler struct {
	service services.AccessTokenService
}

func NewHandler(service services.AccessTokenService) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("id"))
	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at model.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors_utils.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}