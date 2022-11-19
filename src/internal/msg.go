package internal

type Messages struct {
	UserNotFound     string
	UserAlreadyExist string
	WrongPassword    string
	RegisterSuccess  string
	LoginSuccess     string
	LogoutSuccess    string
	AddRoleSuccess   string
	UnAuthorized     string
	UnAuthenticated  string
}

type HandlerKeys struct {
	Token string
}

var messages = Messages{
	UserNotFound:     "user_not_found",
	UserAlreadyExist: "user_already_exist",
	WrongPassword:    "wrong_password",
	RegisterSuccess:  "success_register",
	LoginSuccess:     "success_login",
	LogoutSuccess:    "success_logout",
	AddRoleSuccess:   "success_add_role",
	UnAuthorized:     "auth_unauthorized",
	UnAuthenticated:  "auth_unauthenticated",
}

var handlerKeys = HandlerKeys{
	Token: "token",
}
