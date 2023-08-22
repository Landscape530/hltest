package val

import "awesomeProject/server/base/auth"

type InitConnectionReq struct {
	Remark string         `validate:"omitempty" json:"remark"`
	User   *auth.UserInfo `json:"user"`
}
