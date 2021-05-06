package stateCode

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
	NotInit                   StateCode = -502
	Unknown                   StateCode = -1000
)
