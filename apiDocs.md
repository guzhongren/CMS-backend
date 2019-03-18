# API 详细说明

## 接口前置路径


> baseUrl = http://localhost:1234/api/v1


## 参数位置说明

> **默认以Form放在body中**

## 静态资源位置

> baseUrl/static + 资源名称.ext

## user

用户操作

### login

**POST**

**Form Parameters**

```json
{
    username: jon       // 用户名
    password: password  // 密码
}
```

**Result**

```json
{
    success: true,
    result:{
        token: "TOKEN",
    },
    message: ""
}

```

### logout

**GET**

Result

```json
{
    success: true,
    result: true,
    message: ""
}
```

# HEADERS

> **以下接口请求需要设置Header**

```
Authorization: "Bearer TOKEN" # TOKEN是登录的返回值
```
### users

**GET**

**/users**

**Result**

```json
{
    success: true,
    result:[
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
    message: ""
}
```
**GET**

**/users/:id**

获取用户信息

**Result**

```json
{
    success: true,
    result:
        {
        id: 1,
        name: "张三",
        roleid: "1"
        }
    ],
    message: ""
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
    success: true,
    result:
        {
            name: 'admin',
            roleid: 1,
            id: 5
        },
    message: ""
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
    success: true,
    result:
        {
        name: 'admin',
        roleid: 1,
        id: 5
        },
    message: ""
}
```

**DELETE**

**/users/:id**

删除用户

**Result**

```json
{
    success: true,
    result: 5,
    message: ""
}
```

**GET**

/users/:id/resetPassword

重置密码

**Form Paramaters**

```json
{
    password: string
}
```

**Result**

```json
{
    success: true,
    result: true,
    message: ""
}
```


## Material

/materials

**GET**

获取所有物料

**Result**

```json
{
    success: true,
    result: [
        {
            id": "84970af63baa41923539ad19a468d645",
            "name": "测试11",
            "location":{"String": "西安市", "Valid": true},
            "typeName": "书籍",
            "count":{"Int64": 10, "Valid": true},
            "provider":{"String": "guzhongren", "Valid": true},
            "providerLink":{"String": "https://guzhongren.github.io", "Valid": true},
            "images":{"String": "aab608e25cc70959b5ff0f5045f09a19.jpg,c696c19afdcc03bb26e5a3d19ec02759.jpg,7237fad134b6c2e59ba99452a8dc94c3.jpg,1936b4649666ec0c25ae2447ed35a29e.jpg,4179869dbd9c95d917737a4b93f48816.jpg", "Valid": true},
            "createTime": 1552821860,
            "updateTime":{"Int64": 1552825504, "Valid": true},
            "price":{"Float64": 200, "Valid": true},
            "owner":{"id": "6a67b113839fdcf8527be29b723dc859", "name": "admin"},
            "updateUser":{"id": "", "name": ""}
        }
    ],
    message: ""
}
```

**POST**

新增物料

**Form Parameters**

```json
{
    name: "测试11",
    location:"西安市",
    type: 1,
    count:10,
    provider:"guzhongren",
    providerLink:"https://guzhongren.github.io",
    images:"aab608e25cc70959b5ff0f5045f09a19.jpc696c19afdcc03bb26e5a3d19ec02759.jpg,7237fad134b6c2e59ba99452a8dc94c3.jp1936b4649666ec0c25ae2447ed35a29e.jpg,4179869dbd9c95d917737a4b93f48816.jpg", //form
    price: 200,
}

```
**Result**

```json
{
    success: true,
    result: [
        {
            id": "84970af63baa41923539ad19a468d645",
            "name": "测试11",
            "location":{"String": "西安市", "Valid": true},
            "typeName": "书籍",
            "count":{"Int64": 10, "Valid": true},
            "provider":{"String": "guzhongren", "Valid": true},
            "providerLink":{"String": "https://guzhongren.github.io", "Valid": true},
            "images":{"String": "aab608e25cc70959b5ff0f5045f09a19.jpg,c696c19afdcc03bb26e5a3d19ec02759.jpg,7237fad134b6c2e59ba99452a8dc94c3.jpg,1936b4649666ec0c25ae2447ed35a29e.jpg,4179869dbd9c95d917737a4b93f48816.jpg", "Valid": true},
            "createTime": 1552821860,
            "updateTime":{"Int64": 1552825504, "Valid": true},
            "price":{"Float64": 200, "Valid": true},
            "owner":{"id": "6a67b113839fdcf8527be29b723dc859", "name": "admin"},
            "updateUser":{"id": "", "name": ""}
        }
    ],
    message: ""
}
```



### /materials/:id

**GET**

查询物料详情

**Result**

```json
{
    success: true,
    result: {
            id": "84970af63baa41923539ad19a468d645",
            "name": "测试11",
            "location":{"String": "西安市", "Valid": true},
            "typeName": "书籍",
            "count":{"Int64": 10, "Valid": true},
            "provider":{"String": "guzhongren", "Valid": true},
            "providerLink":{"String": "https://guzhongren.github.io", "Valid": true},
            "images":{"String": "aab608e25cc70959b5ff0f5045f09a19.jpg,c696c19afdcc03bb26e5a3d19ec02759.jpg,7237fad134b6c2e59ba99452a8dc94c3.jpg,1936b4649666ec0c25ae2447ed35a29e.jpg,4179869dbd9c95d917737a4b93f48816.jpg", "Valid": true},
            "createTime": 1552821860,
            "updateTime":{"Int64": 1552825504, "Valid": true},
            "price":{"Float64": 200, "Valid": true},
            "owner":{"id": "6a67b113839fdcf8527be29b723dc859", "name": "admin"},
            "updateUser":{"id": "", "name": ""}
        },
    message: ""
}

```

**PUT**

更新物料

**Form Parameters**

```json
{
    name: "测试11",
    location:"西安市",
    type: 1,
    count:10,
    provider:"guzhongren",
    providerLink:"https://guzhongren.github.io",
    images:"aab608e25cc70959b5ff0f5045f09a19.jpc696c19afdcc03bb26e5a3d19ec02759.jpg,7237fad134b6c2e59ba99452a8dc94c3.jp1936b4649666ec0c25ae2447ed35a29e.jpg,4179869dbd9c95d917737a4b93f48816.jpg", //form
    price: 200,
}

```
**Result**

```json
{
    success: true,
    result: {
            id": "84970af63baa41923539ad19a468d645",
            "name": "测试11",
            "location":{"String": "西安市", "Valid": true},
            "typeName": "书籍",
            "count":{"Int64": 10, "Valid": true},
            "provider":{"String": "guzhongren", "Valid": true},
            "providerLink":{"String": "https://guzhongren.github.io", "Valid": true},
            "images":{"String": "aab608e25cc70959b5ff0f5045f09a19.jpg,c696c19afdcc03bb26e5a3d19ec02759.jpg,7237fad134b6c2e59ba99452a8dc94c3.jpg,1936b4649666ec0c25ae2447ed35a29e.jpg,4179869dbd9c95d917737a4b93f48816.jpg", "Valid": true},
            "createTime": 1552821860,
            "updateTime":{"Int64": 1552825504, "Valid": true},
            "price":{"Float64": 200, "Valid": true},
            "owner":{"id": "6a67b113839fdcf8527be29b723dc859", "name": "admin"},
            "updateUser":{"id": "", "name": ""}
        },
    message: ""
}
```

**DELETE**

删除物料

**Result**

```json
{
    success: true,
    result: true,
    message: ""
}
```

### /material/type

**GET**

查询物料类型

**Result**

```json
{
    success: true,
    result: [
        {
            id: 1,
            name: "书籍"
        }
    ],
    message: ""
}
```

### /material/type/:id

**GET**

查询具体物料类型

**Result**

```json
{
    success: true,
    result: {
        id: 1,
        name: "书籍"
    },
    message: ""
}
```

