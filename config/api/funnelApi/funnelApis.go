package funnelApi

import "wejh-go/config/config"

var FunnelHost = config.Config.GetString("funnel.host")

type FunnelApi string

const (
	ZFClassTable   FunnelApi = "student/zf/table"
	ZFScore        FunnelApi = "student/zf/score"
	ZFExam         FunnelApi = "student/zf/exam"
	CardHistory    FunnelApi = "student/card/history"
	CardToday      FunnelApi = "student/card/today"
	CardBalance    FunnelApi = "student/card/balance"
	LibraryCurrent FunnelApi = "student/library/borrow/current"
	LibraryHistory FunnelApi = "student/library/borrow/history"
)
