package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dungvan2512/soccer-social-network-api/infrastructure"
	"github.com/dungvan2512/soccer-social-network-api/shared/auth"
	"github.com/stretchr/testify/assert"
)

const (
	invalidToken     = "Bearer ayJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI4MTc1MDE4MTQsImlhdCI6MTUxNjY4MjQzNCwiaXNzIjoiZnItY2lyY2xlLWFwaS5jb20iLCJuYmYiOjE1MTY2ODI0MzQsImNvbnRleHQiOnsidXNlciI6eyJpZCI6IjEifX19.Wqlc-9aPHc0QoVCRiGt4N5c5i065o55VbTxTIFzFqYo"
	wrongFormatToken = "Bearer eyJleHAiOjI4MTc1MDE4MTQsImlhdCI6MTUxNjY4MjQzNCwiaXNzIjoiZnItY2lyY2xlLWFwaS5jb20iLCJuYmYiOjE1MTY2ODI0MzQsImNvbnRleHQiOnsidXNlciI6eyJpZCI6IjEifX19.Wqlc-9aPHc0QoVCRiGt4N5c5i065o55VbTxTIFzFqYo"
)

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	return
})

var db = infrastructure.NewSQL().DB

type userMock struct {
	ID        uint
	GetClaims func() map[string]interface{}
}

func (uMock userMock) GetCustomClaims() map[string]interface{} {
	return uMock.GetClaims()
}

func getClaims() map[string]interface{} {
	claims := make(map[string]interface{})
	userClaim := struct {
		ID uint64 `json:"id"`
	}{
		ID: 123,
	}
	claims["user"] = userClaim
	return claims
}

func (uMock userMock) GetIdentifier() uint {
	return uMock.ID
}
func getTokenTest() string {
	uMock := userMock{
		ID:        123,
		GetClaims: getClaims,
	}
	token, _ := auth.GenerateToken(uMock)
	return token
}

func TestJwtAuthSuccess(t *testing.T) {
	token := getTokenTest()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rr := httptest.NewRecorder()
	JwtAuth(infrastructure.NewLogger(), db)(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestJwtAuthFailWithInvalidToken(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", invalidToken)
	rr := httptest.NewRecorder()
	JwtAuth(infrastructure.NewLogger(), db)(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestJwtAuthFailWithWrongFormatToken(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", wrongFormatToken)
	rr := httptest.NewRecorder()
	JwtAuth(infrastructure.NewLogger(), db)(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestJwtAuthFailWithTokenWithoutBearer(t *testing.T) {
	token := getTokenTest()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", token)
	rr := httptest.NewRecorder()
	JwtAuth(infrastructure.NewLogger(), db)(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestJwtAuthFailWithTokenOnlyBearer(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer")
	rr := httptest.NewRecorder()
	JwtAuth(infrastructure.NewLogger(), db)(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestJwtAuthFailWithEmptyToken(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "")
	rr := httptest.NewRecorder()
	JwtAuth(infrastructure.NewLogger(), db)(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestJwtAuthFailWithExpiredToken(t *testing.T) {
	infrastructure.SetConfig("jwt.claim.exp", 1)
	token := getTokenTest()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", token)
	rr := httptest.NewRecorder()
	JwtAuth(infrastructure.NewLogger(), db)(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	infrastructure.SetConfig("jwt.claim.exp", 604800)
}

func TestJwtAuthFailWithoutAuthorizationHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	JwtAuth(infrastructure.NewLogger(), db)(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
