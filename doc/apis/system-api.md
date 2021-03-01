# 程序接口说明
## system 接口

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
