# 🐳 Docker 部署指南

本项目支持使用 Docker Compose 一键启动完整的开发环境，包括前端、后端、数据库和 Redis。

## 📋 前置要求

确保你的系统已安装：

- [Docker](https://www.docker.com/get-started) (20.10+)
- [Docker Compose](https://docs.docker.com/compose/install/) (2.0+)

验证安装：

```bash
docker --version
docker-compose --version
```

---

## 🚀 快速开始

### 方式 1: 使用启动脚本（推荐）

**Windows:**

```bash
docker-start.bat
```

**Linux/Mac:**

```bash
chmod +x docker-start.sh
./docker-start.sh
```

### 方式 2: 手动启动

1. **复制环境变量文件**

   ```bash
   cp .env.docker .env
   ```

2. **启动开发环境**

   ```bash
   docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
   ```

3. **查看日志**
   ```bash
   docker-compose logs -f
   ```

---

## 📍 访问地址

启动成功后，可以通过以下地址访问：

| 服务       | 地址                                     | 说明         |
| ---------- | ---------------------------------------- | ------------ |
| 前端       | http://localhost:3000                    | Next.js 应用 |
| 后端       | http://localhost:9000                    | Go API 服务  |
| Swagger    | http://localhost:9000/swagger/index.html | API 文档     |
| PostgreSQL | localhost:5432                           | 数据库       |
| Redis      | localhost:6379                           | 缓存         |

**默认登录信息：**

- 用户名: `admin`
- 密码: `admin123`

---

## 🛠️ 常用命令

### 服务管理

```bash
# 启动所有服务
docker-compose up -d

# 启动开发环境（支持热重载）
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

# 停止所有服务
docker-compose down

# 重启服务
docker-compose restart

# 重启单个服务
docker-compose restart backend
```

### 日志查看

```bash
# 查看所有服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f backend
docker-compose logs -f frontend
docker-compose logs -f postgres

# 查看最近 100 行日志
docker-compose logs --tail=100 backend
```

### 进入容器

```bash
# 进入后端容器
docker-compose exec backend sh

# 进入前端容器
docker-compose exec frontend sh

# 进入数据库容器
docker-compose exec postgres psql -U xiaozhu -d go_manage_starter
```

### 数据库操作

```bash
# 连接数据库
docker-compose exec postgres psql -U xiaozhu -d go_manage_starter

# 备份数据库
docker-compose exec postgres pg_dump -U xiaozhu go_manage_starter > backup.sql

# 恢复数据库
docker-compose exec -T postgres psql -U xiaozhu -d go_manage_starter < backup.sql
```

### 清理操作

```bash
# 停止并删除容器
docker-compose down

# 停止并删除容器、网络、卷（⚠️ 会删除所有数据）
docker-compose down -v

# 清理未使用的镜像
docker image prune -a

# 清理所有未使用的资源
docker system prune -a --volumes
```

---

## 🔧 配置说明

### 环境变量配置

编辑 `.env` 文件来修改配置：

```env
# 数据库配置
DB_USER=xiaozhu
DB_PASSWORD=12345679
DB_NAME=go_manage_starter

# Redis 配置
REDIS_PASSWORD=

# JWT 密钥
JWT_SECRET=dev-jwt-secret-key-change-this-in-production

# 服务端口
BACKEND_PORT=9000
FRONTEND_PORT=3000

# 前端 API 地址
NEXT_PUBLIC_API_URL=http://localhost:9000/api/v1
```

### 修改端口

如果默认端口被占用，可以在 `.env` 文件中修改：

```env
BACKEND_PORT=9001
FRONTEND_PORT=3001
DB_PORT=5433
REDIS_PORT=6380
```

---

## 📦 服务说明

### 1. PostgreSQL 数据库

- **镜像**: `postgres:16-alpine`
- **端口**: 5432
- **数据持久化**: `postgres_data` 卷
- **初始化脚本**: `manage-backend/scripts/` 目录下的 SQL 文件会在首次启动时自动执行

### 2. Redis 缓存

- **镜像**: `redis:7-alpine`
- **端口**: 6379
- **数据持久化**: `redis_data` 卷

### 3. Go 后端

- **开发模式**: 使用 Air 支持热重载
- **端口**: 9000
- **日志**: 挂载到 `backend_logs` 卷
- **代码挂载**: 本地代码实时同步到容器

### 4. Next.js 前端

- **开发模式**: 支持热重载
- **端口**: 3000
- **代码挂载**: 本地代码实时同步到容器

---

## 🐛 故障排查

### 1. 端口被占用

**错误信息**: `Bind for 0.0.0.0:9000 failed: port is already allocated`

**解决方案**:

```bash
# 查看端口占用
netstat -ano | findstr :9000  # Windows
lsof -i :9000                 # Linux/Mac

# 修改 .env 文件中的端口
BACKEND_PORT=9001
```

### 2. 数据库连接失败

**解决方案**:

```bash
# 检查数据库容器状态
docker-compose ps postgres

# 查看数据库日志
docker-compose logs postgres

# 重启数据库
docker-compose restart postgres
```

### 3. 前端无法连接后端

**解决方案**:

- 检查 `.env` 文件中的 `NEXT_PUBLIC_API_URL` 配置
- 确保后端服务已启动: `docker-compose ps backend`
- 查看后端日志: `docker-compose logs backend`

### 4. 热重载不工作

**解决方案**:

```bash
# 确保使用开发模式启动
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d

# 重启服务
docker-compose restart backend frontend
```

### 5. 数据库初始化失败

**解决方案**:

```bash
# 清理数据并重新初始化
docker-compose down -v
docker-compose up -d

# 手动执行初始化脚本
docker-compose exec postgres psql -U xiaozhu -d go_manage_starter -f /docker-entrypoint-initdb.d/manage_dev.sql
```

---

## 🔒 生产环境部署

### 1. 修改配置

编辑 `.env` 文件，设置生产环境配置：

```env
ENVIRONMENT=production
DB_PASSWORD=strong-production-password
REDIS_PASSWORD=strong-redis-password
JWT_SECRET=strong-jwt-secret-key
```

### 2. 构建生产镜像

```bash
docker-compose build --no-cache
```

### 3. 启动生产环境

```bash
docker-compose up -d
```

### 4. 配置反向代理

推荐使用 Nginx 作为反向代理：

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    location /api {
        proxy_pass http://localhost:9000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

---

## 📊 性能优化

### 1. 限制资源使用

在 `docker-compose.yml` 中添加资源限制：

```yaml
services:
  backend:
    deploy:
      resources:
        limits:
          cpus: "1.0"
          memory: 512M
```

### 2. 使用多阶段构建

生产环境已使用多阶段构建，镜像体积更小。

### 3. 启用 Redis 持久化

编辑 `docker-compose.yml`:

```yaml
redis:
  command: redis-server --appendonly yes
```

---

## 📝 开发建议

1. **使用开发模式**: 支持代码热重载，提高开发效率
2. **定期备份数据**: 使用 `pg_dump` 备份重要数据
3. **查看日志**: 遇到问题先查看日志 `docker-compose logs -f`
4. **清理资源**: 定期清理未使用的镜像和容器
5. **环境隔离**: 开发、测试、生产使用不同的配置文件

---

## 🆘 获取帮助

- 查看 [Docker 官方文档](https://docs.docker.com/)
- 查看 [Docker Compose 文档](https://docs.docker.com/compose/)

---

## 📄 许可证

本项目采用 MIT License 开源协议
