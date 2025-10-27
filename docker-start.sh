#!/bin/bash

# Docker Compose 启动脚本
# 用于快速启动开发环境

set -e

echo "🐳 Go 管理系统 - Docker 启动脚本"
echo "=================================="

# 检查 .env 文件是否存在
if [ ! -f .env ]; then
    echo "📝 未找到 .env 文件，从 .env.docker 复制..."
    cp .env.docker .env
    echo "✅ .env 文件创建完成"
    echo "⚠️  请检查并修改 .env 文件中的配置"
fi

# 选择启动模式
echo ""
echo "请选择启动模式："
echo "1) 开发模式 (支持热重载)"
echo "2) 生产模式"
echo "3) 仅启动数据库和 Redis"
echo "4) 停止所有服务"
echo "5) 清理所有数据（危险操作）"
echo ""
read -p "请输入选项 (1-5): " choice

case $choice in
    1)
        echo "🚀 启动开发环境..."
        docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
        echo "✅ 开发环境启动完成！"
        echo ""
        echo "📍 访问地址："
        echo "   前端: http://localhost:3000"
        echo "   后端: http://localhost:9000"
        echo "   Swagger: http://localhost:9000/swagger/index.html"
        echo ""
        echo "📊 查看日志: docker-compose logs -f"
        ;;
    2)
        echo "🚀 启动生产环境..."
        docker-compose up -d
        echo "✅ 生产环境启动完成！"
        ;;
    3)
        echo "🚀 仅启动数据库和 Redis..."
        docker-compose up -d postgres redis
        echo "✅ 数据库和 Redis 启动完成！"
        echo ""
        echo "📍 连接信息："
        echo "   PostgreSQL: localhost:5432"
        echo "   Redis: localhost:6379"
        ;;
    4)
        echo "🛑 停止所有服务..."
        docker-compose down
        echo "✅ 所有服务已停止"
        ;;
    5)
        read -p "⚠️  确定要清理所有数据吗？这将删除数据库中的所有数据！(yes/no): " confirm
        if [ "$confirm" = "yes" ]; then
            echo "🗑️  清理所有数据..."
            docker-compose down -v
            echo "✅ 数据清理完成"
        else
            echo "❌ 取消清理操作"
        fi
        ;;
    *)
        echo "❌ 无效选项"
        exit 1
        ;;
esac

echo ""
echo "=================================="
echo "常用命令："
echo "  查看日志: docker-compose logs -f [service]"
echo "  进入容器: docker-compose exec [service] sh"
echo "  重启服务: docker-compose restart [service]"
echo "  停止服务: docker-compose stop"
echo "=================================="
