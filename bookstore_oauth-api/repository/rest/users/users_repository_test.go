package users

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"

	"github.com/bmizerany/assert"
	//"github.com/go-resty/resty/v2"
	//"github.com/jarcoal/httpmock"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("starting test cases")
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	//rest.FlushMockups()
	//rest.AddMockups(&rest.Mock{
	//	HTTPMethod:   http.MethodPost,
	//	URL:          "https://api.bookstore.com/users/login",
	//	ReqBody:      `{"email":"test@gmail.com","password":"password"}`,
	//	RespHTTPCode: -1,
	//	RespBody:     `{}`,
	//})
	//
	//repo := usersRepository{}
	//user, err := repo.LoginUser("test@gmail.com", "password")
	//
	//assert.Nil(t, user)
	//assert.NotNil(t, err)
	//assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	//assert.EqualValues(t, "Invalid rest-client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	// Create a Resty Client
	client := resty.New()

	// Get the underlying HTTP Client and set it to Mock
	httpmock.ActivateNonDefault(client.GetClient())

	resp, _ := client.R().SetBody(`{"test":"test"}`).Post("https://dummyurl.com/users/login")

	assert.Equal(t, resp.StatusCode(), 200)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginUserNoError(t *testing.T) {

}
