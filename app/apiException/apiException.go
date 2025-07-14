package apiException

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"wejh-go/config/logger"
)

type Error struct {
	StatusCode int          `json:"-"`
	Code       int          `json:"code"`
	Msg        string       `json:"msg"`
	Level      logger.Level `json:"-"`
}

var (
	ServerError               = NewError(http.StatusInternalServerError, 200500, logger.LevelError, "系统异常，请稍后重试!")
	OpenIDError               = NewError(http.StatusInternalServerError, 200500, logger.LevelError, "系统异常，请稍后重试!")
	ParamError                = NewError(http.StatusInternalServerError, 200501, logger.LevelInfo, "参数错误")
	NotAdmin                  = NewError(http.StatusInternalServerError, 200502, logger.LevelInfo, "该用户无管理员权限")
	UserNotFind               = NewError(http.StatusInternalServerError, 200503, logger.LevelInfo, "该用户不存在")
	NotLogin                  = NewError(http.StatusInternalServerError, 200503, logger.LevelInfo, "未登录")
	NoThatPasswordOrWrong     = NewError(http.StatusInternalServerError, 200504, logger.LevelInfo, "密码错误")
	HttpTimeout               = NewError(http.StatusInternalServerError, 200505, logger.LevelError, "系统异常，请稍后重试!")
	RequestError              = NewError(http.StatusInternalServerError, 200506, logger.LevelError, "系统异常，请稍后重试!")
	NotBindYxy                = NewError(http.StatusInternalServerError, 200507, logger.LevelInfo, "该手机号或用户未绑定或未注册易校园")
	UserAlreadyExisted        = NewError(http.StatusInternalServerError, 200508, logger.LevelInfo, "该用户已激活")
	WrongVerificationCode     = NewError(http.StatusInternalServerError, 200509, logger.LevelInfo, "手机验证码错误，错误3次将锁定15分钟")
	StudentNumAndIidError     = NewError(http.StatusInternalServerError, 200510, logger.LevelInfo, "该学号或身份证不存在或者不匹配，请重新输入")
	PwdError                  = NewError(http.StatusInternalServerError, 200511, logger.LevelInfo, "密码长度必须在6~20位之间")
	ReactiveError             = NewError(http.StatusInternalServerError, 200512, logger.LevelInfo, "该通行证已经存在，请重新输入")
	StudentIdError            = NewError(http.StatusInternalServerError, 200513, logger.LevelInfo, "学号格式不正确，请重新输入")
	YxySessionExpired         = NewError(http.StatusInternalServerError, 200514, logger.LevelInfo, "一卡通登陆过期，请重新登陆")
	WrongPhoneNum             = NewError(http.StatusInternalServerError, 200517, logger.LevelInfo, "手机号格式不正确")
	ImgTypeError              = NewError(http.StatusInternalServerError, 200518, logger.LevelInfo, "图片类型有误")
	PersonalInfoNotFill       = NewError(http.StatusInternalServerError, 200519, logger.LevelInfo, "请先填写个人基本信息")
	StockNotEnough            = NewError(http.StatusInternalServerError, 200520, logger.LevelInfo, "物资库存不足")
	RecordAlreadyExisted      = NewError(http.StatusInternalServerError, 200521, logger.LevelInfo, "该用户已经申请过该物资")
	RecordAlreadyRejected     = NewError(http.StatusInternalServerError, 200522, logger.LevelInfo, "含有已经被驳回的申请，请重新选择")
	NotBorrowingRecord        = NewError(http.StatusInternalServerError, 200523, logger.LevelInfo, "含有非借用中的记录，请重新选择")
	SendVerificationCodeLimit = NewError(http.StatusInternalServerError, 200524, logger.LevelInfo, "短信发送超限，请1分钟后再试")
	CampusMismatch            = NewError(http.StatusInternalServerError, 200525, logger.LevelInfo, "暂无该校区绑定信息")
	OAuthNotUpdate            = NewError(http.StatusInternalServerError, 200526, logger.LevelInfo, "统一身份认证密码未更新")
	NoApiAvailable            = NewError(http.StatusInternalServerError, 200527, logger.LevelInfo, "正方相关服务暂不可用")
	NotBindCard               = NewError(http.StatusInternalServerError, 200528, logger.LevelInfo, "请先在易校园APP绑定浙工大校园卡")
	NotInit                   = NewError(http.StatusNotFound, 200404, logger.LevelWarn, http.StatusText(http.StatusNotFound))
	NotFound                  = NewError(http.StatusNotFound, 200404, logger.LevelWarn, http.StatusText(http.StatusNotFound))
)

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, level logger.Level, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
		Level:      level,
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
	// 构建日志字段
	logFields := []zap.Field{
		zap.Int("error_code", apiErr.Code),
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.String("ip", c.ClientIP()),
		zap.Error(err), // 记录原始错误信息
	}
	logger.GetLogFunc(apiErr.Level)(apiErr.Msg, logFields...)
}
