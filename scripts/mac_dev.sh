#!/bin/bash

echo "开始启动本地开发环境..."

# 启动 db 和 redis 服务，不启动 app 服务
docker-compose up -d db redis

# 等待服务健康检查通过
echo "等待 db & redis  容器启动成功..."
while ! (docker-compose ps | grep -q "db.*healthy") || ! (docker-compose ps | grep -q "redis.*healthy"); do
    sleep 1
done

echo "db & redis 启动成功 !"

# 检查是否安装了 air
if ! command -v air &>/dev/null; then
    echo "air 未安装, 开始安装..."
    go install github.com/air-verse/air@latest
    # 将 GOPATH/bin 添加到 PATH 以确保可以执行 air
    export PATH=$PATH:$(go env GOPATH)/bin
else
    echo "air 已安装 !"
fi

# 后台启动 air 热重载
echo "启动 air 热重载..."
air &

# 等待应用在 8080 端口启动
echo "应用正在全力启动中..."
while ! curl -s http://localhost:8080/ping >/dev/null; do
    sleep 1
done

echo "本地开发环境启动成功 !"

# 保持脚本运行，防止后台进程被终止
wait