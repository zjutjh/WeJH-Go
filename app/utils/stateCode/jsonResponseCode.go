package stateCode

import (
	"encoding/json"
	"errors"
	"reflect"
)

type StateCode int

const (
	OK                      StateCode = 1
	SystemError             StateCode = -1
	ParamError              StateCode = -2
	UserNotFind             StateCode = -301
	UserAlreadyExisted      StateCode = -302
	GetOpenIDFail           StateCode = -400
	UsernamePasswordUnMatch StateCode = -500
	Unknown                 StateCode = -1000
)

func ErrorToStateCode(err error) StateCode {
	var j *json.UnmarshalTypeError
	if errors.As(err, &j) {
		return ParamError
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
	case UserNotFind:
		return "User Not Find"
	case GetOpenIDFail:
		return "Get OpenID Fail"
	default:
		return reflect.TypeOf(code).Name()

	}
}
