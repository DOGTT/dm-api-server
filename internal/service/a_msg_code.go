package service

import (
	"fmt"
	"net/http"

	"github.com/DOGTT/dm-api-server/internal/data/rds"
)

// Error Msg
var (
	EM_OK = &ErrMsg{HttpStatus: 200, Code: "OK", Desc: "Success"}

	EM_CommonFail_ParamsInvalid = &ErrMsg{HttpStatus: http.StatusBadRequest, Code: "CommonFail.ParamsInvalid"}
	EM_CommonFail_AuthFail      = &ErrMsg{HttpStatus: http.StatusUnauthorized, Code: "CommonFail.AuthFail"}
	EM_CommonFail_BadRequest    = &ErrMsg{HttpStatus: http.StatusBadRequest, Code: "CommonFail.BadRequest"}
	EM_CommonFail_Forbidden     = &ErrMsg{HttpStatus: http.StatusForbidden, Code: "CommonFail.Forbidden"}
	EM_CommonFail_Internal      = &ErrMsg{HttpStatus: http.StatusInternalServerError, Code: "CommonFail.Internal"}
	EM_CommonFail_DBError       = &ErrMsg{HttpStatus: http.StatusInternalServerError, Code: "CommonFail.DBError"}

	EM_AuthFail_WX       = &ErrMsg{HttpStatus: http.StatusBadRequest, Code: "AuthFail.WeChat"}
	EM_AuthFail_NotFound = &ErrMsg{HttpStatus: http.StatusNotFound, Code: "AuthFail.NotFound"}

	EM_UserFail_AlreadyExist = &ErrMsg{HttpStatus: http.StatusBadRequest, Code: "UserFail.AlreadyExist"}
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

func putDescByDBErr(err error) *ErrMsg {
	if err == nil {
		return nil
	}
	orErr, ok := err.(*ErrMsg)
	if ok {
		return orErr
	}
	n := *EM_CommonFail_DBError
	if rds.IsNotFound(err) {
		n.Desc = "db not found"
	} else if rds.IsDuplicateErr(err) {
		n.Desc = "db already exist"
	}
	return &n
}

func (e *ErrMsg) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Desc)
}
