package funnelApi

import "wejh-go/config/config"

var FunnelHost = config.Config.GetString("funnel.host")

type FunnelApi string

const (
	ZFClassTable   FunnelApi = "student/zf/table"
	ZFScore        FunnelApi = "student/zf/score"
	ZFMidTermScore FunnelApi = "student/zf/midtermscore"
	ZFExam         FunnelApi = "student/zf/exam"
	ZFRoom         FunnelApi = "student/zf/room"
	LibraryCurrent FunnelApi = "student/library/borrow/current"
	LibraryHistory FunnelApi = "student/library/borrow/history"
	CanteenFlow    FunnelApi = "canteen/flow"
)
