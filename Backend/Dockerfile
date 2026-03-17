# 构建阶段
FROM golang:1.22-alpine AS builder

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git

# 复制 go mod 文件
COPY go.mod go.sum ./
COPY third_party ./third_party

# 下载依赖
RUN go mod download

# 复制源码
COPY . .

# 构建
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# 运行阶段
FROM alpine:3.19

WORKDIR /app

# 安装 ca-certificates（HTTPS 等）
RUN apk add --no-cache ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 从构建阶段复制二进制
COPY --from=builder /app/main .

# 复制 manifest（含 SQL、配置模板等）
COPY manifest ./manifest

EXPOSE 8888

CMD ["./main"]
