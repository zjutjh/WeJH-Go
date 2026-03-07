package apiException

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zjutjh/mygo/nlog"
)

type Error struct {
	Code  int
	Msg   string
	Level logrus.Level
}

var (
	ServerError               = NewError(200500, logrus.ErrorLevel, "系统异常，请稍后重试!")
	OpenIDError               = NewError(200500, logrus.ErrorLevel, "系统异常，请稍后重试!")
	ParamError                = NewError(200501, logrus.InfoLevel, "参数错误")
	NotAdmin                  = NewError(200502, logrus.InfoLevel, "该用户无管理员权限")
	UserNotFind               = NewError(200503, logrus.InfoLevel, "该用户不存在")
	NotLogin                  = NewError(200503, logrus.InfoLevel, "未登录")
	NoThatPasswordOrWrong     = NewError(200504, logrus.InfoLevel, "密码错误")
	HttpTimeout               = NewError(200505, logrus.ErrorLevel, "系统异常，请稍后重试!")
	RequestError              = NewError(200506, logrus.ErrorLevel, "系统异常，请稍后重试!")
	NotBindYxy                = NewError(200507, logrus.InfoLevel, "该手机号或用户未绑定或未注册易校园")
	UserAlreadyExisted        = NewError(200508, logrus.InfoLevel, "该用户已激活")
	WrongVerificationCode     = NewError(200509, logrus.InfoLevel, "手机验证码错误，错误3次将锁定15分钟")
	StudentNumAndIidError     = NewError(200510, logrus.InfoLevel, "该学号或身份证不存在或者不匹配，请重新输入")
	PwdError                  = NewError(200511, logrus.InfoLevel, "密码长度必须在6~20位之间")
	ReactiveError             = NewError(200512, logrus.InfoLevel, "该通行证已经存在，请重新输入")
	StudentIdError            = NewError(200513, logrus.InfoLevel, "学号格式不正确，请重新输入")
	YxySessionExpired         = NewError(200514, logrus.InfoLevel, "一卡通登陆过期，请重新登陆")
	WrongPhoneNum             = NewError(200517, logrus.InfoLevel, "手机号格式不正确")
	ImgTypeError              = NewError(200518, logrus.InfoLevel, "图片类型有误")
	PersonalInfoNotFill       = NewError(200519, logrus.InfoLevel, "请先填写个人基本信息")
	StockNotEnough            = NewError(200520, logrus.InfoLevel, "物资库存不足")
	RecordAlreadyExisted      = NewError(200521, logrus.InfoLevel, "该用户已经申请过该物资")
	RecordAlreadyRejected     = NewError(200522, logrus.InfoLevel, "含有已经被驳回的申请，请重新选择")
	NotBorrowingRecord        = NewError(200523, logrus.InfoLevel, "含有非借用中的记录，请重新选择")
	SendVerificationCodeLimit = NewError(200524, logrus.InfoLevel, "短信发送超限，请1分钟后再试")
	CampusMismatch            = NewError(200525, logrus.InfoLevel, "暂无该校区绑定信息")
	OAuthNotUpdate            = NewError(200526, logrus.InfoLevel, "统一身份认证密码未更新")
	NoApiAvailable            = NewError(200527, logrus.InfoLevel, "正方相关服务暂不可用")
	NotBindCard               = NewError(200528, logrus.InfoLevel, "请先在易校园APP绑定浙工大校园卡")
	NotInit                   = NewError(200404, logrus.WarnLevel, http.StatusText(http.StatusNotFound))
	NotFound                  = NewError(200404, logrus.WarnLevel, http.StatusText(http.StatusNotFound))
)

func (e *Error) Error() string {
	return e.Msg
}
func NewError(Code int, level logrus.Level, msg string) *Error {
	return &Error{
		Code:  Code,
		Msg:   msg,
		Level: level,
	}
}

// AbortWithException 用于返回自定义错误信息
func AbortWithException(c *gin.Context, apiError *Error, err error) {
	logError(c, apiError, err)
	_ = c.AbortWithError(200, apiError)
}

// AbortWithError 用于兼容原始错误返回
func AbortWithError(c *gin.Context, err error) {
	var apiError *Error
	if errors.As(err, &apiError) {
		logError(c, apiError, nil)
		_ = c.AbortWithError(200, apiError)
	} else {
		logError(c, ServerError, err)
		_ = c.AbortWithError(200, ServerError)
	}
}

// logError 记录错误日志
func logError(c *gin.Context, apiErr *Error, err error) {
	logEntry := nlog.Pick().WithContext(c).WithError(err).WithFields(logrus.Fields{
		"error_code": apiErr.Code,
		"path":       c.Request.URL.Path,
		"method":     c.Request.Method,
		"ip":         c.ClientIP(),
	})
	switch apiErr.Level {
	case logrus.DebugLevel:
		logEntry.Debug(apiErr.Msg)
	case logrus.InfoLevel:
		logEntry.Info(apiErr.Msg)
	case logrus.WarnLevel:
		logEntry.Warn(apiErr.Msg)
	case logrus.ErrorLevel:
		logEntry.Error(apiErr.Msg)
	case logrus.FatalLevel:
		logEntry.Fatal(apiErr.Msg)
	case logrus.PanicLevel:
		logEntry.Panic(apiErr.Msg)
	}
}
