package users

import (
	"github.com/gin-gonic/gin"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_oauth-go/oauth"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/domain/users"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_users-api/services"
	"github.com/olegnalivajev/learning_go_microservices/bookstore_utils-go/errors_utils"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors_utils.RestErr)  {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors_utils.NewBadRequestErr("user id should be a number")
	}
	return userId, nil
}

func Get(c *gin.Context)  {
	if err := oauth.AuthenticateRequest(c.Request); err != nil{
		c.JSON(err.Status, err)
		return
	}

	if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
		err := errors_utils.NewNotAuthorizedError("Resource not available.")
		c.JSON(err.Status, err)
		return
	}

	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors_utils.NewBadRequestErr("Id is invalid or not present.")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

func Create(c *gin.Context)  {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors_utils.NewBadRequestErr("Invalid JSON body.")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors_utils.NewBadRequestErr("Invalid JSON body.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, updateErr := services.UserService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if deleteErr := services.UserService.DeleteUser(userId); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status" : "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	userz, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, userz.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors_utils.NewBadRequestErr("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
