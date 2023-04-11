package yxyApi

import "wejh-go/config/config"

var YxyHost = config.Config.GetString("yxy.host")

type YxyApi string

const (
	SecurityToken          YxyApi = "v1/campus/login/security_token"
	CaptchaImage           YxyApi = "v1/campus/login/captcha_image"
	SendVerificationCode   YxyApi = "v1/campus/login/send_verification_code"
	LoginByCode            YxyApi = "v1/campus/login/by_code"
	CardBalance            YxyApi = "v1/campus/user/card_balance"
	ConsumptionRecords     YxyApi = "v1/campus/user/consumption_records"
	Auth                   YxyApi = "v1/app/auth"
	Bind                   YxyApi = "v1/app/electricity/bind"
	ElectricityBalance     YxyApi = "v1/app/electricity/subsidy/by_user"
	RechargeRecords        YxyApi = "v1/app/electricity/recharge/by_room"
	ElectricityConsumption YxyApi = "v1/app/electricity/consumption"
	SlientLogin            YxyApi = "v1/campus/login/silent"
)
