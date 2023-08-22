package auth

var CtxLoginUserKey = "ctxLoginUserKey"

type UserInfo struct {
	Username string `json:"username"`

	Authority string `json:"authority"`
}
