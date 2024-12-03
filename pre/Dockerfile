# 使用轻量级基础镜像
FROM golang:1.22.3-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将 Go 源码复制到容器中
COPY . .

# 下载依赖
RUN go mod tidy

# 构建 Go 应用程序
RUN go build -o main .

# 使用更轻量级的镜像运行二进制文件
FROM alpine:latest

# 创建一个非root用户
RUN adduser -D appuser

# 设置工作目录
WORKDIR /root/

# 复制二进制文件到最终镜像中
COPY --from=builder /app/main .

# 切换为非root用户
USER appuser

# 启动应用程序
CMD ["./main"]
