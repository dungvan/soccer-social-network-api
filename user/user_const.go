package user

import "errors"

const (
	duplidateUniquePreMessage = "duplicate key value violates unique constraint"
	userNameKey               = "user_user_name_key"
	emailKey                  = "user_email_key"
	userName                  = "user_name"
	email                     = "email"
)

var (
	uniqueKeys = []string{userNameKey, emailKey}
	fieldName  = map[string]string{userNameKey: userName, emailKey: email}

	errUserNameOrPassword = errors.New("user name or password dose not match")
)
