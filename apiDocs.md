# API

## Login

**Url**

`http://localhost:1234/api/v1/login`

**Queyr Parameters**

username: jon
password: password

Result

```json
success": true,
"result":{
    "token": TOKEN,
},
"message": ""

```

## getUsers

**Url**

http://localhost:1234/api/v1/users

**HEADERS**

Authorization: "Bearer TOKEN"

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