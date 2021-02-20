# wejh-go

一个用 go 写成的微精弘后端服务项目

## 如何开始

在项目根目录下执行
```bash
go build -tags=jsoniter
```

然后直接运行编辑出来的程序就行了

> 这个 tag 的意思就是我们的代码用第三方库解析json，这样速度更快
> 程序能识别的变量类型更多（原生 json 解析库部分复杂的嵌套数据类型
> 没发解析成 json 格式)

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
> 方便日后移植到不同的语言上~~（我这次搬运数据好痛苦啊）~~，
> 同时也方便日后程序写入配置

通过在conf文件夹下执行以下命令来生成配置文件

```bash
cp config.yaml.example config.yaml
```

yaml 的格式就自己百度/Google吧，建议学会了以后在看下面的内容

### 如何添加自己的配置

如果要添加自己的配置，可以在 config 中添加一个对象（键值对），然后 在这个对象中添加你想要的配置 **（可以参考之前的设置）**。
之后所有的配置都会被读入 conf 包 中的 Config 变量

> PS: 程序会自动监视配置文件，如果配置文件发生了变化，程序会自动更新

## 程序接口说明

由于历史遗留原因，现在api设计的时候都会在变量前加上 "/api" 前缀 日后记得删掉。

### misc 接口

> misc 意思就是这些是一些杂项接口，不太好分类

**获取学期相关信息**

接口信息

|  | 内容 | 
| --- | --- |
| 接口 | /api/time |
| 参数 | 无 |

返回值样例

```json
{
  "data": {
    "day": 5,
    "is_begin": true,
    "month": "2",
    "term": "2020/2021(1)",
    "week": 21
  },
  "errcode": 1,
  "errmsg": "获取时间成功"
}
```

**获取后台通知** 

接口信息

|  | 内容 | 
| --- | --- |
| 接口 | /api/announcement |
| 参数 | 无 |

返回值样例
```json
{
    "data": {
        "clipboard": "",
        "clipboardtip": "",
        "content": "由于学校寒假期间正在进行停电检修及网络维护，目前微精弘无法查询到最新的课表、成绩等信息",
        "footer": "有任何微精弘问题，请加QQ群:462530805（一群）282402782（二群）",
        "id": 36,
        "show": true,
        "title": "公告"
    },
    "errcode": 1,
    "errmsg": "ok",
    "redirect": null
}
```

