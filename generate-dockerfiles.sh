#!/bin/bash

# 微服务列表及其端口映射
services=(
  "auth"
  "cart"
  "checkout"
  "order"
  "payment"
  "product"
  "user"
  "eino"
)

ports=(
  "11801"
  "11802"
  "11803"
  "11804"
  "11805"
  "11806"
  "11807"
  "11808"
)

# 确保docker目录存在
mkdir -p docker

# 为每个微服务生成Dockerfile
for i in "${!services[@]}"; do
  service="${services[$i]}"
  port="${ports[$i]}"
  
  cat > "docker/Dockerfile.$service" << EOF
# 第一阶段：构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的构建工具
RUN apk add --no-cache git

ENV GO111MODULE=on
ENV GOWORK=/app/go.work

# 复制整个项目代码
COPY . .

# 编译${service}服务
RUN go build -o ${service}-service ./app/${service}

# 第二阶段：运行阶段
FROM alpine:3.17

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates tzdata && \\
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \\
    echo "Asia/Shanghai" > /etc/timezone

# 创建非root用户
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# 从构建阶段复制编译好的应用
COPY --from=builder /app/${service}-service .

# 使用非root用户运行
USER appuser

# 暴露服务端口
EXPOSE ${port}

# 运行应用
CMD ["./${service}-service"]
EOF

done

echo "已生成所有服务的 Dockerfile"