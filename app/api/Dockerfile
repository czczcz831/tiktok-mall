# 第一阶段：构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api-server ./main.go

# 第二阶段：运行阶段
FROM alpine:3.17

RUN apk add --no-cache ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/api-server .
COPY --from=builder /app/casbin.conf .

USER appuser

# 暴露应用端口（请根据实际使用的端口进行修改）
EXPOSE 8888


# 运行应用
CMD ["./api-server"] 