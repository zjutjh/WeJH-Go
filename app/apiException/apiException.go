package apiException

import "net/http"

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	ServerError               = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	OpenIDError               = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	ParamError                = NewError(http.StatusInternalServerError, 200501, "参数错误")
	NotAdmin                  = NewError(http.StatusInternalServerError, 200502, "该用户无管理员权限")
	UserNotFind               = NewError(http.StatusInternalServerError, 200503, "该用户不存在")
	NotLogin                  = NewError(http.StatusInternalServerError, 200503, "未登录")
	NoThatPasswordOrWrong     = NewError(http.StatusInternalServerError, 200504, "密码错误")
	HttpTimeout               = NewError(http.StatusInternalServerError, 200505, "系统异常，请稍后重试!")
	RequestError              = NewError(http.StatusInternalServerError, 200506, "系统异常，请稍后重试!")
	NotBindYxy                = NewError(http.StatusInternalServerError, 200507, "该手机号或用户未绑定或未注册易校园")
	UserAlreadyExisted        = NewError(http.StatusInternalServerError, 200508, "该用户已激活")
	WrongVerificationCode     = NewError(http.StatusInternalServerError, 200509, "手机验证码错误，错误3次将锁定15分钟")
	StudentNumAndIidError     = NewError(http.StatusInternalServerError, 200510, "该学号或身份证不存在或者不匹配，请重新输入")
	PwdError                  = NewError(http.StatusInternalServerError, 200511, "密码长度必须在6~20位之间")
	ReactiveError             = NewError(http.StatusInternalServerError, 200512, "该通行证已经存在，请重新输入")
	StudentIdError            = NewError(http.StatusInternalServerError, 200513, "学号格式不正确，请重新输入")
	YxySessionExpired         = NewError(http.StatusInternalServerError, 200514, "一卡通登陆过期，请稍后再试")
	YxyNeedCaptcha            = NewError(http.StatusInternalServerError, 200515, "请输入验证码")
	WrongCaptcha              = NewError(http.StatusInternalServerError, 200516, "图形验证码错误")
	WrongPhoneNum             = NewError(http.StatusInternalServerError, 200517, "手机号格式不正确")
	ImgTypeError              = NewError(http.StatusInternalServerError, 200518, "图片类型有误")
	PersonalInfoNotFill       = NewError(http.StatusInternalServerError, 200519, "请先填写个人基本信息")
	StockNotEnough            = NewError(http.StatusInternalServerError, 200520, "物资库存不足")
	RecordAlreadyExisted      = NewError(http.StatusInternalServerError, 200521, "该用户已经申请过该物资")
	RecordAlreadyRejected     = NewError(http.StatusInternalServerError, 200522, "含有已经被驳回的申请，请重新选择")
	NotBorrowingRecord        = NewError(http.StatusInternalServerError, 200523, "含有非借用中的记录，请重新选择")
	SendVerificationCodeLimit = NewError(http.StatusInternalServerError, 200524, "短信发送超限，请1分钟后再试")
	CampusMismatch            = NewError(http.StatusInternalServerError, 200525, "暂无该校区绑定信息")
	OAuthNotUpdate            = NewError(http.StatusInternalServerError, 200526, "统一身份认证密码未更新")
	NotInit                   = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	NotFound                  = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	Unknown                   = NewError(http.StatusInternalServerError, 300500, "系统异常，请稍后重试!")
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}
