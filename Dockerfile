# 第一阶段：构建阶段
FROM golang:1.22-alpine as builder

# 设置工作目录
WORKDIR /app

# 将当前目录中的所有文件复制到容器内
COPY . .

# 下载项目依赖
RUN go mod tidy

# 构建 Go 程序，使用 go build 编译当前目录下的所有文件
RUN go build -o /app/online-shopping .

# 第二阶段：运行阶段，使用更小的基础镜像
FROM alpine:latest

# 安装运行所需的库
RUN apk add --no-cache ca-certificates

# 将构建阶段生成的可执行文件复制到当前镜像
COPY --from=builder /app/online-shopping /app/online-shopping

# 暴露应用使用的端口
EXPOSE 3002

# 设置容器启动时运行的命令
CMD ["/app/online-shopping"]
