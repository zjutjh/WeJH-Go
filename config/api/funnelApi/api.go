package funnelApi

type LoginType string

const (
	Oauth   LoginType = "OAUTH"
	ZF      LoginType = "ZF"
	Unknown LoginType = "Unknown"
)
