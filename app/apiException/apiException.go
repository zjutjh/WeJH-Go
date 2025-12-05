package apiException

import (
	"errors"
	"net/http"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"github.com/zjutjh/mygo/nlog"
)

type Error struct {
	Code  int
	Msg   string
	Level logger.Level
}

var (
	ServerError               = NewError(200500, logger.LevelError, "系统异常，请稍后重试!")
	OpenIDError               = NewError(200500, logger.LevelError, "系统异常，请稍后重试!")
	ParamError                = NewError(200501, logger.LevelInfo, "参数错误")
	NotAdmin                  = NewError(200502, logger.LevelInfo, "该用户无管理员权限")
	UserNotFind               = NewError(200503, logger.LevelInfo, "该用户不存在")
	NotLogin                  = NewError(200503, logger.LevelInfo, "未登录")
	NoThatPasswordOrWrong     = NewError(200504, logger.LevelInfo, "密码错误")
	HttpTimeout               = NewError(200505, logger.LevelError, "系统异常，请稍后重试!")
	RequestError              = NewError(200506, logger.LevelError, "系统异常，请稍后重试!")
	NotBindYxy                = NewError(200507, logger.LevelInfo, "该手机号或用户未绑定或未注册易校园")
	UserAlreadyExisted        = NewError(200508, logger.LevelInfo, "该用户已激活")
	WrongVerificationCode     = NewError(200509, logger.LevelInfo, "手机验证码错误，错误3次将锁定15分钟")
	StudentNumAndIidError     = NewError(200510, logger.LevelInfo, "该学号或身份证不存在或者不匹配，请重新输入")
	PwdError                  = NewError(200511, logger.LevelInfo, "密码长度必须在6~20位之间")
	ReactiveError             = NewError(200512, logger.LevelInfo, "该通行证已经存在，请重新输入")
	StudentIdError            = NewError(200513, logger.LevelInfo, "学号格式不正确，请重新输入")
	YxySessionExpired         = NewError(200514, logger.LevelInfo, "一卡通登陆过期，请重新登陆")
	WrongPhoneNum             = NewError(200517, logger.LevelInfo, "手机号格式不正确")
	ImgTypeError              = NewError(200518, logger.LevelInfo, "图片类型有误")
	PersonalInfoNotFill       = NewError(200519, logger.LevelInfo, "请先填写个人基本信息")
	StockNotEnough            = NewError(200520, logger.LevelInfo, "物资库存不足")
	RecordAlreadyExisted      = NewError(200521, logger.LevelInfo, "该用户已经申请过该物资")
	RecordAlreadyRejected     = NewError(200522, logger.LevelInfo, "含有已经被驳回的申请，请重新选择")
	NotBorrowingRecord        = NewError(200523, logger.LevelInfo, "含有非借用中的记录，请重新选择")
	SendVerificationCodeLimit = NewError(200524, logger.LevelInfo, "短信发送超限，请1分钟后再试")
	CampusMismatch            = NewError(200525, logger.LevelInfo, "暂无该校区绑定信息")
	OAuthNotUpdate            = NewError(200526, logger.LevelInfo, "统一身份认证密码未更新")
	NoApiAvailable            = NewError(200527, logger.LevelInfo, "正方相关服务暂不可用")
	NotBindCard               = NewError(200528, logger.LevelInfo, "请先在易校园APP绑定浙工大校园卡")
	NotInit                   = NewError(200404, logger.LevelWarn, http.StatusText(http.StatusNotFound))
	NotFound                  = NewError(200404, logger.LevelWarn, http.StatusText(http.StatusNotFound))
)

func (e *Error) Error() string {
	return e.Msg
}

func NewError(Code int, level logger.Level, msg string) *Error {
	return &Error{
		Code:  Code,
		Msg:   msg,
		Level: level,
	}
}

// AbortWithException 用于返回自定义错误信息
func AbortWithException(c *gin.Context, apiError *Error, err error) {
	if apiError.Level == 4 {
		nlog.Pick().WithContext(c).WithError(err).Info(apiError.Msg)
	} else if apiError.Level == 3 {
		nlog.Pick().WithContext(c).WithError(err).Warn(apiError.Msg)
	} else if apiError.Level == 2 {
		nlog.Pick().WithContext(c).WithError(err).Error(apiError.Msg)
	}
	_ = c.AbortWithError(200, apiError)
}

// AbortWithError 用于兼容原始错误返回
func AbortWithError(c *gin.Context, err error) {
	var apiError *Error
	if errors.As(err, &apiError) {
		AbortWithException(c, apiError, err)
	} else {
		nlog.Pick().WithContext(c).WithError(err).Error(ServerError.Msg)
		_ = c.AbortWithError(200, ServerError)
	}
}
