package userCenterServices

import (
	"encoding/json"
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
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
	res, err := f.Get(userCenterApi.UserCenterHost + "?" + form.Encode())
	if err != nil {
		return err
	}

	var resp models.OldInfo

	err = json.Unmarshal(res, &resp)
	if err != nil {
		return err
	}

	if resp.State == "error" {
		if resp.Info == "该学号和身份证不存在或者不匹配，请重新输入" {
			return apiException.StudentNumAndIidError
		} else if resp.Info == "密码长度必须在6~20位之间" {
			return apiException.PwdError
		} else if resp.Info == "该通行证已经存在，请重新输入" {
			return apiException.ReactiveError
		}
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
	res, err := f.Get(userCenterApi.UserCenterHost + "api.php?" + form.Encode())
	if err != nil {
		return err
	}

	var resp models.OldInfo

	err = json.Unmarshal(res, &resp)
	if err != nil {
		return err
	}
	if resp.State == "error" {
		if resp.Info == "该学号和身份证不存在或者不匹配，请重新输入" {
			return apiException.StudentNumAndIidError
		} else if resp.Info == "密码长度必须在6~20位之间" {
			return apiException.PwdError
		} else if resp.Info == "该通行证已经存在，请重新输入" {
			return apiException.ReactiveError
		} else if resp.Info == "学号格式不正确，请重新输入" {
			return apiException.StudentIdError
		}
	}
	return nil
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
	res, err := f.Get(userCenterApi.UserCenterHost + "?" + form.Encode())
	if err != nil {
		return err
	}

	var resp models.OldInfo

	err = json.Unmarshal(res, &resp)
	if err != nil {
		return err
	}

	if resp.State == "error" {
		if resp.Info == "该学号和身份证不存在或者不匹配，请重新输入" {
			return apiException.StudentNumAndIidError
		} else if resp.Info == "密码长度必须在6~20位之间" {
			return apiException.PwdError
		} else if resp.Info == "该通行证已经存在，请重新输入" {
			return apiException.ReactiveError
		}
	}
	return nil
}
