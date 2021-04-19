package stateCode

import (
	"encoding/json"
	"errors"
	"net/http"
	wejherrors "wejh-go/errors"
)

type StateCode int

const (
	OK                        StateCode = 1
	SystemError               StateCode = -1
	ParamError                StateCode = -2
	UserNotFind               StateCode = -301
	UserAlreadyExisted        StateCode = -302
	GetOpenIDFail             StateCode = -400
	NotLogin                  StateCode = -401
	NotAdmin                  StateCode = -403
	NoThatPasswordORWrong     StateCode = -413
	HttpTimeout               StateCode = -501
	UsernamePasswordUnmatched StateCode = -500
	Unknown                   StateCode = -1000
)

func ErrorToStateCode(err error) StateCode {
	var j *json.UnmarshalTypeError
	if errors.As(err, &j) || err.Error() == "EOF" {
		return ParamError
	}

	if errors.Is(err, wejherrors.PasswordWrong) {
		return NoThatPasswordORWrong
	}
	if errors.As(err, &http.ErrHandlerTimeout) {
		return HttpTimeout
	}
	return Unknown
}

func GetStateCodeMsg(code StateCode) string {
	switch code {
	case OK:
		return "OK"
	case SystemError:
		return "System Error"
	case ParamError:
		return "Params Error"
	case NotLogin:
		return "Not Login"
	case UserNotFind:
		return "User Not Find"
	case UserAlreadyExisted:
		return "User Already Existed"
	case UsernamePasswordUnmatched:
		return "Username Password Unmatched"
	case NoThatPasswordORWrong:
		return "No Password or Wrong"
	case GetOpenIDFail:
		return "Get WechatOpenID Fail"
	case HttpTimeout:
		return "Http Fetch timeout"
	case Unknown:
		return "Unknown"
	default:
		return "Unknown"

	}
}
