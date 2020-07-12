package http

import (
	"github.com/gin-gonic/gin"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-api/domain/access_token"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type accessTokenHandler struct {
	service access_token.AccessTokenService
}

func NewHandler(service access_token.AccessTokenService) AccessTokenHandler {
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
	var atr access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&atr); err != nil {
		restErr := errors_utils.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	at, err := h.service.Create(atr)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}
