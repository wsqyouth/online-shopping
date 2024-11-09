# online-shopping
基于容器化实现一个购物平台英语


### 实现步骤
1. v1 版本
确定 dockerfile, docker-compose, 选型 gin 框架，实现一个 简单的 curl 请求
```
curl -X POST http://localhost:3002/group/hello -H "Content-Type: application/json" -d '{"username": "xx"}'
curl -X POST http://localhost:3002/group/calc -H "Content-Type: application/json" -d '{"a": 1,"b":2}'
```