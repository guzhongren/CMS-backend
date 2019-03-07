# Material

|字段|类型|说明|
|:--|:--|:--|
|ID|character varying [32]|唯一标识, 主键, GUID生成|
|name|character varying [32]|物料名称|
|userid|character varying[32]|上传用户id|
|location|character varying[100]|存放位置|
|type|character varying[10]|类型|
|count|integer|数量|
|provider|character varying[32]|提供者(商)|
|providerlink|character varying[100]|提供者(商)链接|
|images|character varying[36*5]|物料图片|
