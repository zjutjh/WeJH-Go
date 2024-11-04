package zfController

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/funnelServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
	"wejh-go/config/redis"
)

type form struct {
	Year string `json:"year" binding:"required"`
	Term string `json:"term" binding:"required"`
}

func GetClassTable(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	loginType, err := genLoginType(user)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}

	result, err := funnelServices.GetClassTable(user, postForm.Year, postForm.Term, loginType)
	if err != nil {
		if err == apiException.NoThatPasswordOrWrong {
			userServices.DelPassword(user, loginType)
		}
		_ = c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

func GetScore(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	loginType, err := genLoginType(user)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}

	result, err := funnelServices.GetScore(user, postForm.Year, postForm.Term, loginType)
	if err != nil {
		if err == apiException.NoThatPasswordOrWrong {
			userServices.DelPassword(user, loginType)
		}
		_ = c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

func GetMidTermScore(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	loginType, err := genLoginType(user)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}

	result, err := funnelServices.GetMidTermScore(user, postForm.Year, postForm.Term, loginType)
	if err != nil {
		if err == apiException.NoThatPasswordOrWrong {
			userServices.DelPassword(user, loginType)
		}
		_ = c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, result)
}

func GetExam(c *gin.Context) {
	var postForm form
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	loginType, err := genLoginType(user)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}

	result, err := funnelServices.GetExam(user, postForm.Year, postForm.Term, loginType)
	if err != nil {
		if err == apiException.NoThatPasswordOrWrong {
			userServices.DelPassword(user, loginType)
		}
		_ = c.AbortWithError(200, err)
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
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	loginType, err := genLoginType(user)
	if err != nil {
		_ = c.AbortWithError(200, err)
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
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}

	result, err := funnelServices.GetRoom(user, postForm.Year, postForm.Term, postForm.Campus, postForm.Weekday, postForm.Week, postForm.Sections, loginType)
	if err != nil {
		if err == apiException.NoThatPasswordOrWrong {
			userServices.DelPassword(user, loginType)
			_ = c.AbortWithError(200, err)
			return
		}
		_ = c.AbortWithError(200, err)
		return
	}
	// 将结果缓存到 Redis 中
	if result != nil {
		resultJson, _ := json.Marshal(result)
		err = redis.RedisClient.Set(c, cacheKey, string(resultJson), 1*time.Hour).Err()
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}
	utils.JsonSuccessResponse(c, result)
}

func genLoginType(u *models.User) (string, error) {
	var loginType string
	rand.Seed(time.Now().UnixNano())
	oauthVal := rand.Intn(40)
	zfVal := rand.Intn(60)

	if u.OauthPassword != "" && u.ZFPassword != "" {
		if oauthVal > zfVal {
			loginType = "OAUTH"
		} else {
			loginType = "ZF"
		}
	} else if u.OauthPassword != "" {
		loginType = "OAUTH"
	} else if u.ZFPassword != "" {
		loginType = "ZF"
	} else {
		return "", apiException.NoThatPasswordOrWrong
	}

	return loginType, nil
}
