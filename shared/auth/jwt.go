package auth

import (
	rqContext "context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/dungvan2512/socker-social-network/model"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/mitchellh/mapstructure"
)

// UserAuth user auth
type UserAuth struct {
	ID       uint64 `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
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

// JWTObject interface
type JWTObject interface {
	GetCustomClaims() map[string]interface{}
	GetIdentifier() uint64
}

// GenerateToken Generate token
func GenerateToken(object JWTObject) (accessToken string, err error) {
	if object == nil {
		err = utils.ErrorsNew("Object is nil")
		return
	}
	emptyID := uint64(0)
	if object.GetIdentifier() == emptyID {
		err = utils.ErrorsNew("Object is empty")
		return
	}
	exp := infrastructure.GetConfigInt64("jwt.claim.exp")
	issuer := infrastructure.GetConfigString("jwt.claim.issuer")
	customClaims := object.GetCustomClaims()
	standardClaims := jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Duration(exp) * time.Second).Unix(),
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
	}
	pay := payload{
		StandardClaims: standardClaims,
		Context:        customClaims,
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = pay
	key := infrastructure.GetConfigByte("jwt.key")
	accessToken, err = token.SignedString(key)
	return
}

type payload struct {
	jwt.StandardClaims
	Context map[string]interface{} `json:"context"`
}

// GetUserFromContext get user id from context
func GetUserFromContext(c rqContext.Context) model.User {
	return c.Value(ContextKeyAuth).(model.User)
}
