#url和Action对应关系

HTTP方法|URL|Action|意义
---|---|---|---
GET|/items/|Index()| 获取所有items单元
GET|/items/123|Query(key string)|查询key的item信息
GET|/items/create|CreateFromGet()|通过GET方法创建
|||
PUT|/items/create|CreateFromPut()|通过PUT方法创建
POST|/items/create|CreateFromPost()|通过Post方法创建
GET|/itmes/123/update|UpdateFromGet(key string)|通过GET方法更新
PUT|/items/123/update|UpdateFromPut(key string)|通过PUT方法更新
POST|/items/123/update|UpdateFromPost(key string)| 通过POST方法更新
DELETE|/items/123|Delete(key string)|删除item
|||
HEAD|/itmes/123|Brief(key string)|获取Head信息
OPTIONS|/items|Options()|获取支持的方法