# 程序接口说明
## user 接口

> 这里的接口主要都是和用户信息相关的接口

### 登陆体系介绍

在整个小程序中，最关键的就是 openid，这个 openid 是微信小程序 用来标识用户的（可以理解成一个贼NB的用户名）

openid 由我们后端接口这里生成，小程序通过接口 `/api/code/weapp`
来获得该用户 openid，然后将 openid 作为小程序里的登陆认证凭据

只有将 openid 和微精弘账号绑定后的用户才能继续使用小程序

### 接口说明

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