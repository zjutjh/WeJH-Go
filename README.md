# WeJH-Go

[![Go](https://github.com/zjutjh/wejh-go/actions/workflows/go-build-test.yml/badge.svg)](https://github.com/zjutjh/wejh-go/actions/workflows/go-build-test.yml)

一个用 go 写成的微精弘后端服务项目

## 如何开始

首先先将动态库依赖准备好

```bash
# windows
cp ./lib/yxy.dll ./yxy.dll

# linux
cp ./lib/libyxy.so /lib/libyxy.so
```

在项目根目录下执行

```bash
go build -tags=jsoniter
```

然后直接运行编辑出来的程序就行了

> 这个 tag 的意思就是我们的代码用第三方库解析json，这样速度更快
> 程序能识别的变量类型更多（原生 json 解析库部分复杂的嵌套数据类型
> 没法解析成 json 格式)

## 风格规范说明

> 为了拯救有强迫症的开发者，请遵守下面的几个小小的约定吧（笑）

1. 所有的文件/文件夹名采用小驼峰命名(首单词全小写，后面的单词首字母大写）
2. 代码采用tab制表符缩进而不是空格（Go 的默认风格）

## 配置系统说明

采用了viper这个配置文件库，方便读取配置信息。

[viper 使用教程](https://www.liwenzhou.com/posts/Go/viper_tutorial/#autoid-1-4-3)

> 本程序采用的是自定义了 viper 对象的写法，和教程中的写法大同小异
> 有什么区别的话看看这个库的源码，和现在conf文件夹中的代码应该就懂了（吧

```go
// 就是这种写法，新建了一个 Viper 指针然后进行操作，而不是直接用 包内置的函数
x := viper.New()
y := viper.New()

x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")
//...
```

其中 config.yaml 是程序的配置文件，采用 yaml 格式。

> 为什么不把所有配置信息写在不同的go程序呢？
> 因为这样就能采用 yaml 之类的通用配置文件格式，
> 方便日后移植到不同的语言上~~（我这次搬运数据好痛苦啊）~~，
> 同时也方便日后程序写入配置

通过在项目根目录下执行以下命令来生成配置文件

```bash
cp config.example.yaml config.yaml
```

yaml 的格式就自己百度/Google吧，建议学会了以后在看下面的内容

### 如何添加自己的配置

如果要添加自己的配置，可以在 config 中添加一个对象（键值对），然后 在这个对象中添加你想要的配置 **（可以参考之前的设置）**。 之后所有的配置都会被读入 conf 包 中的 Config 变量

> PS: 程序会自动监视配置文件，如果配置文件发生了变化，程序会自动更新

## Contributing

If you are interested in reporting/fixing issues and contributing directly to the code base, please see [CONTRIBUTING.md](doc/CONTRIBUTING.md) for more information on what we're looking for and how to get started.

## License

[MIT](https://github.com/zjutjh/wejh-go/blob/master/LICENSE)
