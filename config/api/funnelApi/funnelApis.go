package funnelApi

import "wejh-go/config/config"

var FunnelHost = config.Config.GetString("funnel.host")

var CardHistory = "student/card/history"
