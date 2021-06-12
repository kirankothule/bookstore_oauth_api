package rest

import (
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"

	"github.com/kirankothule/bookstore_oauth_api/src/domain/users"
	"github.com/kirankothule/bookstore_oauth_api/src/utils/errors"
)

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Microsecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email, password string) (*users.User, *errors.RestErr) {

	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := userRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface whrn trying to login user")
		}
		return nil, &restErr
	}
	var user users.User

	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error while tring to unmarshell user response")
	}
	return &user, nil
}
