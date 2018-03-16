package auth

import (
	rqContext "context"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/mitchellh/mapstructure"
)

// UserAuth user auth
type UserAuth struct {
	ID uint64 `json:"id"`
}
type context struct {
	User UserAuth `json:"user"`
}
type frClaims struct {
	jwt.StandardClaims
	Context context `json:"context"`
}

// ContextKey string type
type ContextKey string

var (
	// ContextKeyAuth contextKey for authenticate.
	ContextKeyAuth = ContextKey("user")
)

// ParseToken to userID
func ParseToken(tokenString string) (user UserAuth, err error) {
	// get private key.
	key := infrastructure.GetConfigByte("jwt.key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.ErrorsNew("Parse token error")
		}
		return key, nil
	})
	if err != nil {
		return user, utils.ErrorsNew("Parse token error")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var claim frClaims
		err = mapstructure.Decode(claims, &claim)
		if err != nil {
			err = utils.ErrorsNew("can't decode claims")
		}
		return claim.Context.User, nil
	}
	err = utils.ErrorsNew("No claims token")
	return UserAuth{}, err
}

// GetUserFromContext get user id from context
func GetUserFromContext(c rqContext.Context) UserAuth {
	return c.Value(ContextKeyAuth).(UserAuth)
}
