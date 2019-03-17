# Material

|字段|类型|说明|
|:--|:--|:--|
|id|character varying [32]|唯一标识, 主键, GUID生成|
|name|character varying [32]|物料名称|
|ownerId|character varying[32]|初次上传用户id|
|location|character varying[100]|存放位置|
|type|character varying[10]|类型|
|count|integer|数量|
|provider|character varying[32]|提供者(商)|
|providerLink|character varying[100]|提供者(商)链接|
|images|character varying[36*5]|物料图片|
|createTime|bigint|物料加入时间|
|updateTime|bigint|物料更新时间|
|updateUserId|character varying[32]|最后更新的用户id|
