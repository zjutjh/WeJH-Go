package sessionServices

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
)

func SetWechatSession(c *gin.Context, session *auth.ResCode2Session) error {
	webSession := sessions.Default(c)
	webSession.Set("openid", session.OpenID)
	webSession.Set("sessionKey", session.SessionKey)
	webSession.Set("unionID", session.UnionID)
	return webSession.Save()
}

func GetWechatSession(c *gin.Context) (*auth.ResCode2Session, error) {
	webSession := sessions.Default(c)
	openid := webSession.Get("openid")
	sessionKey := webSession.Get("sessionKey")
	unionID := webSession.Get("unionID")
	if openid == nil || sessionKey == nil || unionID == nil {
		return nil, errors.New("")
	}
	session := auth.ResCode2Session{OpenID: openid.(string), SessionKey: sessionKey.(string), UnionID: unionID.(string)}

	return &session, nil
}
