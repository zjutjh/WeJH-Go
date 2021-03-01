package stateCode

const (
	OK                      = 1
	SystemError             = -1
	ParamError              = -2
	UserNotFind             = -3
	GetOpenIDFail           = -4
	UsernamePasswordUnMatch = -5
)

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
	case GetOpenIDFail:
		return "Get OpenID Fail"
	default:
		return "Unknown"

	}
}
