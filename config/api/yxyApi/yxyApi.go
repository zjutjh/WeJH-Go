package yxyApi

import "wejh-go/config/config"

var YxyHost = config.Config.GetString("yxy.host")

type YxyApi string

const (
	SecurityToken        YxyApi = "v1/campus/login/security_token"
	CaptchaImage         YxyApi = "v1/campus/login/captcha_image"
	SendVerificationCode YxyApi = "v1/campus/login/send_verification_code"
	LoginByCode          YxyApi = "v1/campus/login/by_code"
	LoginByPassword      YxyApi = "v1/campus/login/by_password"
	PublicKey            YxyApi = "v1/campus/login/public_key"
	LoginSilent          YxyApi = "v1/campus/login/silent"
	CardBalance          YxyApi = "v1/campus/user/card_balance"
	ConsumptionRecords   YxyApi = "v1/campus/user/consumption_records"
	Consumption          YxyApi = "v1/app/electricity/consumption/"
)
