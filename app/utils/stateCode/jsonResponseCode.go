package stateCode

import (
	"encoding/json"
	"errors"
	"net/http"
	wejherrors "wejh-go/errors"
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
	case NotInit:
		return "Not Init"
	case Unknown:
		return "Unknown"
	default:
		return "Unknown"

	}
}
