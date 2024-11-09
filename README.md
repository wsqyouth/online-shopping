# online-shopping
基于容器化实现一个购物平台英语


### 实现步骤
1. v1 版本
确定 dockerfile, docker-compose, 选型 gin 框架，实现一个 简单的 curl 请求
```
docker build -t online-shopping .
docker run --name online-shopping -d online-shopping

docker-compose up --build -d
docker-compose logs online-shopping
curl -X POST http://localhost:3002/group/hello -H "Content-Type: application/json" -d '{"username": "xx"}'
curl -X POST http://localhost:3002/group/calc -H "Content-Type: application/json" -d '{"a": 1,"b":2}'
```
2. v2 版本
### 项目结构分析

项目结构符合领域驱动设计（DDD）的原则。以下是对各个部分的分析：

1. **Domains**：
   - `dto`：用于定义数据传输对象，用于在各层之间传递数据。
   - `entity`：定义领域实体，包含业务相关的数据和行为。
   - `repo`：定义仓储接口，负责数据的存储和检索操作。
   - `services`：包含业务逻辑，实现了领域服务。

2. **Infra**：
   - `memory`：实现了仓储接口的内存版本，适用于在开发和测试阶段不依赖于外部数据库。

3. **Server**：
   - `handlers`：定义 HTTP 请求处理逻辑，负责请求的接收、解析和响应。

4. **main.go**：
   - 负责应用程序的启动和配置。

这种结构清晰地分离了领域逻辑、基础设施实现和接口适配，便于项目的维护和扩展。

### 使用方法
```
curl -X POST http://localhost:3002/v1/accounts \
     -H "Content-Type: application/json" \
     -d '{
           "username": "nickname",
           "password": "yourpassword"
         }'

curl -X GET "http://localhost:3002/v1/accounts?id=d50af1aa-45de-417e-9520-3ea88ecaceb0"
curl -X GET "http://localhost:3002/v1/accounts/by-username?username=nickname"

```

### TODO:
* 容器云原生运行成功
* 更贴近实战

### 参考项目：

* [启发了我容器及业务] https://github.com/nurana88/online-shopping/blob/master/Makefile
* [启发了我容器及业务]https://github.com/furkanbaran5/E-Trade-Online-Shopping/blob/main/Controller/docker-compose.yml