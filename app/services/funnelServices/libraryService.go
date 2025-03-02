package funnelServices

import (
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
)

func GetCurrentBorrow(u *models.User) (interface{}, error) {
	if u.OauthPassword == "" {
		return nil, apiException.NoThatPasswordOrWrong
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.OauthPassword)
	return FetchHandleOfPost(form, funnelApi.FunnelHost, funnelApi.LibraryCurrent)
}

func GetHistoryBorrow(u *models.User) (interface{}, error) {
	if u.OauthPassword == "" {
		return nil, apiException.NoThatPasswordOrWrong
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.OauthPassword)

	return FetchHandleOfPost(form, funnelApi.FunnelHost, funnelApi.LibraryHistory)
}
