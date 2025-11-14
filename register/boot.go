package register

import (
	"wejh-go/register/generate"

	"github.com/zjutjh/mygo/feishu"
	"github.com/zjutjh/mygo/foundation/kernel"
	"github.com/zjutjh/mygo/nlog"
	"github.com/zjutjh/mygo/ndb"
	"github.com/zjutjh/mygo/nedis"
	"github.com/zjutjh/mygo/nesty"
)

func Boot() kernel.BootList {
	return kernel.BootList{
		// 基础引导器 
		feishu.Boot(),    
		nlog.Boot(),
		generate.Boot(), // 导入生成代码

		// Client引导器
		ndb.Boot(),   // DB
		nedis.Boot(), // Redis
		nesty.Boot(), // HTTP Client
	}
}

