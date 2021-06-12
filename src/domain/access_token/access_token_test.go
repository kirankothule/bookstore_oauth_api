package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "The expiration time should be 24 hr")
}
func TestGetAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "New access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "New access token should not be defined")
	assert.True(t, at.UserID == 0, "New access token should not have associated user user id")

}
func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "The access token should be valid")
}
