package zfController

import (
	"encoding/json"
	"fmt"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/services/funnelServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/circuitBreaker"
	"wejh-go/config/redis"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type form struct {
	Year string `json:"year" binding:"required"`
	Term string `json:"term" binding:"required"`
}

func GetClassTable(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	api, loginType, err := circuitBreaker.CB.GetApi(user.ZFPassword != "", user.OauthPassword != "")
	if err != nil {
		apiException.AbortWithError(c, err)
		return
	}

	result, err := funnelServices.GetClassTable(user, postForm.Year, postForm.Term, api, loginType)
	if err != nil {
		zap.L().Error("获取课表失败", zap.Error(err),
			zap.String("studentID", user.StudentID),
			zap.String("loginType", string(loginType)),
			zap.String("api", api))
		userServices.DelPassword(err, user, string(loginType))
		apiException.AbortWithError(c, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

func GetScore(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	api, loginType, err := circuitBreaker.CB.GetApi(user.ZFPassword != "", user.OauthPassword != "")
	if err != nil {
		apiException.AbortWithError(c, err)
		return
	}

	result, err := funnelServices.GetScore(user, postForm.Year, postForm.Term, api, loginType)
	if err != nil {
		zap.L().Error("获取期末分数失败", zap.Error(err),
			zap.String("studentID", user.StudentID),
			zap.String("loginType", string(loginType)),
			zap.String("api", api))
		userServices.DelPassword(err, user, string(loginType))
		apiException.AbortWithError(c, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

func GetMidTermScore(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	api, loginType, err := circuitBreaker.CB.GetApi(user.ZFPassword != "", user.OauthPassword != "")
	if err != nil {
		apiException.AbortWithError(c, err)
		return
	}

	result, err := funnelServices.GetMidTermScore(user, postForm.Year, postForm.Term, api, loginType)
	if err != nil {
		zap.L().Error("获取期中分数失败", zap.Error(err),
			zap.String("studentID", user.StudentID),
			zap.String("loginType", string(loginType)),
			zap.String("api", api))
		userServices.DelPassword(err, user, string(loginType))
		apiException.AbortWithError(c, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

func GetExam(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	api, loginType, err := circuitBreaker.CB.GetApi(user.ZFPassword != "", user.OauthPassword != "")
	if err != nil {
		apiException.AbortWithError(c, err)
		return
	}

	result, err := funnelServices.GetExam(user, postForm.Year, postForm.Term, api, loginType)
	if err != nil {
		zap.L().Error("获取考试信息失败", zap.Error(err),
			zap.String("studentID", user.StudentID),
			zap.String("loginType", string(loginType)),
			zap.String("api", api))
		userServices.DelPassword(err, user, string(loginType))
		apiException.AbortWithError(c, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

type roomForm struct {
	Year     string `json:"year" binding:"required"`
	Term     string `json:"term" binding:"required"`
	Campus   string `json:"campus" binding:"required"`
	Weekday  string `json:"weekday" binding:"required"`
	Sections string `json:"sections" binding:"required"`
	Week     string `json:"week" binding:"required"`
}

func GetRoom(c *gin.Context) {
	var postForm roomForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	api, loginType, err := circuitBreaker.CB.GetApi(user.ZFPassword != "", user.OauthPassword != "")
	if err != nil {
		apiException.AbortWithError(c, err)
		return
	}

	// 使用 Redis 缓存键，包含查询参数
	cacheKey := fmt.Sprintf("room:%s:%s:%s:%s:%s:%s", postForm.Year, postForm.Term, postForm.Campus, postForm.Weekday, postForm.Week, postForm.Sections)

	// 从 Redis 中获取缓存结果
	cachedResult, cacheErr := redis.RedisClient.Get(c, cacheKey).Result()
	if cacheErr == nil {
		var result []map[string]interface{}
		if err := json.Unmarshal([]byte(cachedResult), &result); err == nil {
			utils.JsonSuccessResponse(c, result)
			return
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
			return
		}
	}

	result, err := funnelServices.GetRoom(user, postForm.Year, postForm.Term, postForm.Campus, postForm.Weekday, postForm.Week, postForm.Sections, api, loginType)
	if err != nil {
		zap.L().Error("获取空教室失败", zap.Error(err),
			zap.String("studentID", user.StudentID),
			zap.String("loginType", string(loginType)),
			zap.String("api", api))
		userServices.DelPassword(err, user, string(loginType))
		apiException.AbortWithError(c, err)
		return
	}
	// 将结果缓存到 Redis 中
	if result != nil {
		resultJson, _ := json.Marshal(result)
		err = redis.RedisClient.Set(c, cacheKey, string(resultJson), 1*time.Hour).Err()
		if err != nil {
			apiException.AbortWithException(c, apiException.ServerError, err)
			return
		}
	}
	utils.JsonSuccessResponse(c, result)
}
