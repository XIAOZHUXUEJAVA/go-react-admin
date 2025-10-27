#!/bin/bash
set -e

# 数据库初始化脚本
# 此脚本会在 PostgreSQL 容器首次启动时自动执行

echo "🚀 开始初始化数据库..."
echo "📊 数据库: $POSTGRES_DB"
echo "👤 用户: $POSTGRES_USER"

# 创建 schema
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    -- 显示当前数据库
    SELECT current_database();
    
    -- 创建开发环境 schema
    CREATE SCHEMA IF NOT EXISTS manage_dev;
    
    -- 创建测试环境 schema
    CREATE SCHEMA IF NOT EXISTS manage_test;
    
    -- 创建生产环境 schema
    CREATE SCHEMA IF NOT EXISTS manage_prod;
    
    -- 授予权限
    GRANT ALL PRIVILEGES ON SCHEMA manage_dev TO $POSTGRES_USER;
    GRANT ALL PRIVILEGES ON SCHEMA manage_test TO $POSTGRES_USER;
    GRANT ALL PRIVILEGES ON SCHEMA manage_prod TO $POSTGRES_USER;
    
    -- 设置默认 schema
    ALTER DATABASE $POSTGRES_DB SET search_path TO manage_dev, public;
EOSQL

echo "✅ Schema 创建完成"

# 如果存在 manage_dev.sql，则执行它
if [ -f /docker-entrypoint-initdb.d/manage_dev.sql ]; then
    echo "📝 执行 manage_dev.sql..."
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -f /docker-entrypoint-initdb.d/manage_dev.sql
    echo "✅ manage_dev.sql 执行完成"
fi

echo "🎉 数据库初始化完成！"
