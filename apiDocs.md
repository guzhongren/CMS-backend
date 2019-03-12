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

**/users**

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
**GET**

**/users/:id**

获取用户信息

**Result**

```json
{
    "success": true,
    "result":
        {
        "id": 1,
        "name": "张三",
        "roleid": "1"
        }
    ],
    "message": ""
}
```

**POST**

**/users/:id**

新增用户

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

**PUT**

**/users/:id**

更新用户

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

**DELETE**

**/users/:id**

删除用户

**Result**

```json
{
    "success": true,
    "result": 5,
    "message": ""
}
```