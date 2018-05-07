package user

import "errors"

const (
	duplidateUniquePreMessage = "duplicate key value violates unique constraint"
	userNameKey               = "users_user_name_key"
	emailKey                  = "users_email_key"
	userName                  = "user_name"
	email                     = "email"
	pagingLimit               = 10
)

var (
	uniqueKeys = []string{userNameKey, emailKey}
	fieldName  = map[string]string{userNameKey: userName, emailKey: email}

	errUserNameOrPassword = errors.New("user name or password dose not match")
)
