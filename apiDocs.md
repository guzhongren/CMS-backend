# API 详细说明

## 接口前置路径

```string
http://localhost:1234/api/v1
```

## login

**POST**

**Queyr Parameters**

```json
{
    username: jon       用户名
    password: password  密码
}
```

**Result**

```json
{
    success": true,
    "result":{
        "token": "TOKEN",
    },
    "message": ""
}

```

## logout

**GET**

# HEADERS

> **以下接口请求需要设置Header**

```
Authorization: "Bearer TOKEN" # TOKEN是登录的返回值
```
## users

**GET**



**Result**

```json
{
    "success": true,
    "result":[
        {
        "id": 1,
        "name": "张三",
        "roleid": "1"
        },
        {
        "id": 2,
        "name": "李四",
        "roleid": "2"
        }
    ],
    "message": ""
}
```

**POST**

**Parameters**

```json
{
    name: 'admin',
    roleid: 1,
    id: 5
}
```

**Result**

```json
{
    "success": true,
    "result":
        {
        name: 'admin',
        roleid: 1,
        id: 5
        },
    "message": ""
}
```