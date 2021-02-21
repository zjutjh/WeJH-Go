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
> 没法解析成 json 格式)

## 风格规范说明

> 为了拯救有强迫症的开发者，请遵守下面的几个小小的约定吧（笑）

1. 所有的文件/文件夹名采用蛇形命名(所有字符均为小写，并且两个单词之间用空格分开）
2. 控制器函数采用驼峰命名，并且所有以Controller结尾（控制器也就是你的接口会调用的那个函数）
3. 一个代码文件只能有中一个控制器函数
4. 控制器函数所在的那个文件的文件名为 控制器函数名去掉 controller（并且转换成小写)
5. 代码采用tab制表符缩进而不是空格（Go 的默认风格）

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
cp config.example.yaml config.yaml
```

yaml 的格式就自己百度/Google吧，建议学会了以后在看下面的内容

### 如何添加自己的配置

如果要添加自己的配置，可以在 config 中添加一个对象（键值对），然后 在这个对象中添加你想要的配置 **（可以参考之前的设置）**。 之后所有的配置都会被读入 conf 包 中的 Config 变量

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

**获取小程序首页列表**

接口信息

| | 内容 |
|--- | --- |
| 接口 | /api/app-list |
| 参数 | 无 |

返回值样例

```json
{
  "data": {
    "app-list": [
      {
        "bg": "blue",
        "icon": "http:// *** ",
        "route": "/pages/timetable/timetable",
        "title": "课表查询"
      },
      {
        "bg": "blue",
        "icon": "http:// *** ",
        "route": "/pages/borrow/borrow",
        "title": "借阅信息"
      }
    ],
    "icons": {
      "borrow": {
        "bg": "blue",
        "card": "http:// *** ",
        "icon": "http:// *** "
      },
      "tri": {
        "bg": "blue",
        "card": "http:// *** ",
        "icon": ""
      }
    }
  },
  "errcode": 1,
  "errmsg": "ok",
  "redirect": null
}
```

### user 接口

> 这里的接口主要都是和用户信息相关的接口

#### 登陆体系介绍

在整个小程序中，最关键的就是 openid，这个 openid 是微信小程序 用来标识用户的（可以理解成一个贼NB的用户名）

openid 由我们后端接口这里生成，小程序通过接口 `/api/code/weapp`
来获得该用户 openid，然后将 openid 作为小程序里的登陆认证凭据

只有将 openid 和微精弘账号绑定后的用户才能继续使用小程序

#### 接口说明

**接口路由**：`/api/code/weapp`

**接口方法**：POST

**body 格式**：json

| 参数 | 类型 | 解释 |
| --- | --- | --- |
| code | string | 小程序端生成的 code |

**返回数据格式：** json

**返回数据样例:**

```json
{
  "errcode": 200,
  "errmsg": "获取 openID 成功",
  "data": {
    "openid": "A string here"
  }
}
```

**接口路由**：`/bind/jh`

**接口方法**：POST

**body 格式**：json

| 参数 | 类型 | 解释 |
| --- | --- | --- |
| openid | string | 当前用户的 openid |
| password | string | 当前用户的精弘账号密码 |
| username | string | 当前用户的精弘账号用户名 |
| type | string | 当前用户的登陆类型（暂时用不到）|

**返回数据格式：** json

**返回数据样例：**

```json
{
  "errcode":  200,
  "errmsg":   "绑定成功",
  "redirect": null,
  "token": "The open ID",
  "user": {
    "uno": "学号"
  }
}
```

> 为了兼容以前的小程序前端，现在还会再返回一次 open ID



**接口路由**：`/api/autologin`

**接口方法**：POST

**body 格式** ：json

| 参数 | 类型 | 解释 |
| --- | --- | --- |
| openid | string | 当前用户的 openid |
| type | string | 当前用户的登陆类型（暂时用不到）|

**返回数据格式：** json

**返回数据样例：**

```json
{
  "errcode": 200,
  "errmsg": "绑定成功",
  "token": "The open ID",
  "user": {
    "uno": "学号"
  }
}
```
