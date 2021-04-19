package funnelServices

import (
	"net/url"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
	"wejh-go/errors"
)

func GetCurrentBorrow(u *models.User) (interface{}, error) {
	if u.LibPassword == "" {
		return nil, errors.PasswordWrong
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.LibPassword)

	return FetchHandleOfPost(form, funnelApi.LibraryCurrent)
}

func GetHistoryBorrow(u *models.User) (interface{}, error) {
	if u.LibPassword == "" {
		return nil, errors.PasswordWrong
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.LibPassword)

	return FetchHandleOfPost(form, funnelApi.LibraryHistory)
}
