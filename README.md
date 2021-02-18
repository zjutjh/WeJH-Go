# wejh-go
一个用 go 写成的微精弘后端服务项目

## 配置系统说明
采用了viper这个配置文件库，方便读取配置信息。

[viper 使用教程](https://www.liwenzhou.com/posts/Go/viper_tutorial/#autoid-1-4-3)

> 本程序采用的是自定义了 viper 对象的写法，和教程中的写法大同小异
> 有什么区别的话看看这个库的源码，和现在conf文件夹中的代码应该就懂了（吧

```go
// 就是这种写法，新建了一个 Viper 指针然后进行操作，而不是直接用
// 包内置的函数
x := viper.New()
y := viper.New()

x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")
//...
```

其中 config.yaml 是程序的配置文件，采用 yaml 格式。

> 为什么不把所有配置信息写在不同的go程序呢？
> 因为这样就能采用 yaml 之类的通用配置文件格式，
> 方便日后移植到不同的语言上，
> 同时也方便日后程序写入配置

通过在conf文件夹下执行以下命令来生成配置文件

```bash
cp config.yaml.example config.yaml
```

yaml 的格式就自己百度/Google吧，建议学会了以后在看下面的内容

### 如何添加自己的配置

如果要添加自己的配置，可以在 config 中添加一个对象（键值对），然后
在这个对象中添加你想要的配置。**（可以参考之前的设置）**

之后所有的配置都会被读入 conf 包 中的 Config 变量