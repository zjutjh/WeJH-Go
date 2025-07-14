package stateCode

const (
	OK                        = 1
	SystemError               = -1
	ParamError                = -2
	UserNotFind               = -301
	UserAlreadyExisted        = -302
	GetOpenIDFail             = -400
	NotLogin                  = -401
	NotAdmin                  = -403
	NoThatPasswordORWrong     = -413
	HttpTimeout               = -501
	UsernamePasswordUnmatched = -500
	NotInit                   = -502
	Unknown                   = -1000
)
