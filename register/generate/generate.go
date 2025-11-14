package generate

// 用于生成和注册生成模块的启动函数，请不要在此目录下防止自定义的业务代码

// Boot 用于自动注册 generate 模块下的所有组件 (本身无任何操作)
func Boot() func() error {
	return func() error {
		return nil
	}
}
