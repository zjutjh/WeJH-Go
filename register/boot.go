package register

import (
	"github.com/zjutjh/mygo/feishu"
	"github.com/zjutjh/mygo/foundation/kernel"
	"github.com/zjutjh/mygo/ndb"
	"github.com/zjutjh/mygo/nedis"
	"github.com/zjutjh/mygo/nesty"
	"github.com/zjutjh/mygo/nlog"
	"github.com/zjutjh/mygo/wechat/miniprogram"
)

func Boot() kernel.BootList {
	return kernel.BootList{
		// 基础引导器
		miniprogram.Boot(),
		feishu.Boot(),
		nlog.Boot(),

		// Client引导器

		ndb.Boot(),   // DB
		nedis.Boot(), // Redis
		nesty.Boot(), // HTTP Client
	}
}
