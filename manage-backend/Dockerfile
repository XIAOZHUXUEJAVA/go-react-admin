# ==================== 开发阶段 ====================
FROM golang:1.25-alpine AS development

# 设置工作目录
WORKDIR /app

# 安装必要的工具
# RUN apk add --no-cache git make curl
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache git make curl
# 设置Go代理和环境变量
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

# 安装 Air 用于热重载
RUN go install github.com/air-verse/air@latest

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 创建日志目录
RUN mkdir -p logs

# 暴露端口
EXPOSE 9000

# 使用 Air 启动（支持热重载）
CMD ["air", "-c", ".air.toml"]


# ==================== 构建阶段 ====================
FROM golang:1.25-alpine AS builder

WORKDIR /app

# 安装构建依赖
# RUN apk add --no-cache git make
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache git make curl

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/main.go


# ==================== 生产阶段 ====================
FROM alpine:latest AS production

# 安装 ca-certificates 用于 HTTPS 请求
# RUN apk --no-cache add ca-certificates tzdata
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .

# 复制配置文件
COPY --from=builder /app/config ./config

# 创建日志目录
RUN mkdir -p logs

# 暴露端口
EXPOSE 9000

# 运行应用
CMD ["./main"]