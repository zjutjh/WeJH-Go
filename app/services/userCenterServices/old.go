package userCenterServices

import (
	"net/url"
	"wejh-go/app/utils/fetch"
	"wejh-go/config/api/userCenterApi"
)

func OldAuthStudent(username, password string) error {
	f := fetch.Fetch{}
	f.Init()
	form := url.Values{}
	form.Add("app", "passport")
	form.Add("action", "login")
	form.Add("username", username)
	form.Add("password", password)
	_, err := f.Get(userCenterApi.UserCenterHost + "?" + form.Encode())
	if err != nil {
		return err
	}
	return nil
}

func OldActiveStudent(username, password, iid, email string) error {
	f := fetch.Fetch{}
	f.Init()
	form := url.Values{}
	form.Add("app", "passport")
	form.Add("action", "active")
	form.Add("username", username)
	form.Add("password", password)
	form.Add("iid", iid)
	form.Add("email", email)
	_, err := f.Get(userCenterApi.UserCenterHost + "?" + form.Encode())
	if err != nil {
		return err
	}
	return nil // To-do impl it
}

func OldResetStudent(username, password, iid string) error {
	f := fetch.Fetch{}
	f.Init()
	form := url.Values{}
	form.Add("app", "passport")
	form.Add("action", "reset")
	form.Add("username", username)
	form.Add("password", password)
	form.Add("iid", iid)
	_, err := f.Get(userCenterApi.UserCenterHost + "?" + form.Encode())
	if err != nil {
		return err
	}
	return nil
}
