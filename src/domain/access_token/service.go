package access_token

import (
	"github.com/kirankothule/bookstore_oauth_api/utils/errors"
)

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}
