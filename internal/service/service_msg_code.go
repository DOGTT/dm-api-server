package service

import (
	"fmt"
	"net/http"
)

// Error Msg
var (
	EM_OK = &ErrMsg{HttpStatus: 200, Code: "OK", Desc: "Success"}

	EM_CommonFail_ParamsInvalid = &ErrMsg{HttpStatus: http.StatusBadRequest, Code: "CommonFail.ParamsInvalid"}
	EM_CommonFail_AuthFail      = &ErrMsg{HttpStatus: http.StatusUnauthorized, Code: "CommonFail.AuthFail"}
	EM_CommonFail_BadRequest    = &ErrMsg{HttpStatus: http.StatusBadRequest, Code: "CommonFail.BadRequest"}
	EM_CommonFail_Internal      = &ErrMsg{HttpStatus: http.StatusInternalServerError, Code: "CommonFail.Internal"}

	EM_AuthFail_WX       = &ErrMsg{HttpStatus: http.StatusBadRequest, Code: "AuthFail.WeChat"}
	EM_AuthFail_NotFound = &ErrMsg{HttpStatus: http.StatusBadRequest, Code: "AuthFail.NotFound"}
)

type ErrMsg struct {
	HttpStatus int    `json:"-"`
	Code       string `json:"code"`
	Desc       string `json:"desc"`
}

func (e *ErrMsg) PutDesc(desc string) *ErrMsg {
	n := *e
	n.Desc = desc
	return &n
}

func (e *ErrMsg) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Desc)
}
