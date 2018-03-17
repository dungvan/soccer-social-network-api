package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type jwtObjectmock struct {
	ID        uint64
	GetClaims func() map[string]interface{}
}

func (mock jwtObjectmock) GetCustomClaims() map[string]interface{} {
	return mock.GetClaims()
}
func (mock jwtObjectmock) GetIdentifier() uint64 {
	return mock.ID
}

func setUpClaims() map[string]interface{} {
	userclaim := struct {
		ID uint64 `json:"id"`
	}{
		ID: 123,
	}
	mapData := make(map[string]interface{})
	mapData["user"] = userclaim
	return mapData
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestGenerateTokenSuccess(t *testing.T) {
	mock := jwtObjectmock{
		ID:        123,
		GetClaims: setUpClaims,
	}
	token, err := GenerateToken(mock)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGenerateTokenFailWithNilObject(t *testing.T) {
	token, err := GenerateToken(nil)
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestGenerateTokenFailWithEmptyObject(t *testing.T) {
	mock := jwtObjectmock{}
	token, err := GenerateToken(mock)
	assert.Error(t, err)
	assert.Empty(t, token)
}
