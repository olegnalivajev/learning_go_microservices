package users

import (
	"github.com/gin-gonic/gin"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/domain/users"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/services"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context)  {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestErr("Id is invalid or not present.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context)  {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestErr("Invalid JSON body.")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
