package register

import (
	"wejh-go/app/utils/circuitBreaker"

	"github.com/zjutjh/mygo/feishu"
	"github.com/zjutjh/mygo/foundation/kernel"
	"github.com/zjutjh/mygo/ndb"
	"github.com/zjutjh/mygo/nedis"
	"github.com/zjutjh/mygo/nesty"
	"github.com/zjutjh/mygo/nlog"
	"github.com/zjutjh/mygo/wechat/miniProgram"
)

func Boot() kernel.BootList {
	return kernel.BootList{
		// 基础引导器
		feishu.Boot(),
		nlog.Boot(),

		// Client引导器
		ndb.Boot(),   // DB
		nedis.Boot(), // Redis
		nesty.Boot(), // HTTP Client
		miniProgram.Boot(),

		AppBoot(),
	}
}



// AppBoot 应用定制引导器
func AppBoot() func() error {
	return func() error {
		//do.ProvideValue(nil,config.Pick())
		// 可以在这里编写业务初始引导逻辑
		circuitBreaker.Init() 
		return nil
	}
}
