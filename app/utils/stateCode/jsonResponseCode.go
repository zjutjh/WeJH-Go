package stateCode

var OK = 1
var SystemError = -1
var ParamError = -2
var UserNotFind = -3

func GetStateCodeMsg(code int) string {
	switch code {
	case OK:
		return "OK"
	case SystemError:
		return "System Error"
	case ParamError:
		return "Params Error"
	case UserNotFind:
		return "User Not Find"
	default:
		return "Unknown"

	}
}
