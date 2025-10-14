<div align="center">

# 🚀 Go 管理系统起始模板

**现代化的全栈管理系统开发脚手架**

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Next.js](https://img.shields.io/badge/Next.js-15-black?style=flat&logo=next.js)](https://nextjs.org/)
[![React](https://img.shields.io/badge/React-19-61DAFB?style=flat&logo=react)](https://react.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](./License)

_一个基于 Go + React + Next.js 构建的前后端分离管理系统模板，提供开箱即用的用户认证和基础架构_

[功能特性](#-功能特性) • [技术栈](#-技术栈) • [快速开始](#-快速开始) • [项目结构](#-项目结构) • [API 文档](#-api-接口) • [日志系统](#-日志系统)

</div>

---

## 📋 项目概述

这是一个**生产就绪**的管理系统起始模板，采用现代化技术栈构建，遵循最佳实践和清晰的代码架构。适合作为中小型管理系统的开发基础，或用于学习全栈开发。

### ✨ 核心亮点

- 🎯 **清晰的架构设计** - 前后端完全分离，代码结构清晰易维护
- 🔐 **完整的认证系统** - JWT Token + 图形验证码
- 🎨 **现代化 UI** - 基于 shadcn/ui 和 Tailwind CSS 的精美界面
- 📦 **开箱即用** - 包含完整的开发环境配置和数据库迁移
- 🚀 **高性能** - Go 后端 + Next.js 15 (Turbopack) 前端
- 📝 **API 文档** - 集成 Swagger 自动生成 API 文档

---

## 🎯 功能特性

### ✅ 已实现功能

| 功能模块      | 描述                       | 状态    |
| ------------- | -------------------------- | ------- |
| 🔐 用户认证   | 注册、登录、JWT Token 验证 | ✅ 完成 |
| 🖼️ 图形验证码 | 防机器人注册/登录保护      | ✅ 完成 |
| 👥 用户管理   | 基础的用户 CRUD 接口       | ✅ 完成 |
| 🎨 响应式界面 | 适配桌面和移动端           | ✅ 完成 |
| ⚠️ 错误处理   | 统一的错误处理机制         | ✅ 完成 |
| 📚 API 文档   | Swagger 自动生成文档       | ✅ 完成 |

### 🚧 规划中的功能

- 🔑 **RBAC 权限系统** - 角色和权限管理
- 📊 **数据统计面板** - 可视化数据展示
- 📝 **操作日志** - 完整的审计日志
- 🔔 **消息通知** - 系统消息推送
- 📤 **文件上传** - 文件管理功能

---

## 🛠 技术栈

<table>
<tr>
<td width="50%" valign="top">

### 后端 (manage-backend)

**核心框架**

- ![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go) **Go** - 高性能编程语言
- ![Gin](https://img.shields.io/badge/Gin-Web_Framework-00ADD8?style=flat) **Gin** - 轻量级 Web 框架
- ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791?style=flat&logo=postgresql) **PostgreSQL** - 关系型数据库
- ![GORM](https://img.shields.io/badge/GORM-ORM-00ADD8?style=flat) **GORM** - Go ORM 框架

**核心依赖**

- `golang-jwt/jwt` - JWT 认证
- `go-redis` - Redis 缓存
- `viper` - 配置管理
- `zap` - 结构化日志
- `lumberjack` - 日志轮转
- `base64Captcha` - 验证码生成
- `swaggo` - API 文档生成

</td>
<td width="50%" valign="top">

### 前端 (manage-frontend)

**核心框架**

- ![Next.js](https://img.shields.io/badge/Next.js-15-black?style=flat&logo=next.js) **Next.js 15** - React 框架 (App Router)
- ![React](https://img.shields.io/badge/React-19-61DAFB?style=flat&logo=react) **React 19** - UI 库
- ![TypeScript](https://img.shields.io/badge/TypeScript-5-3178C6?style=flat&logo=typescript) **TypeScript** - 类型安全
- ![Tailwind](https://img.shields.io/badge/Tailwind-CSS-38B2AC?style=flat&logo=tailwind-css) **Tailwind CSS** - 样式框架

**核心依赖**

- `shadcn/ui` - UI 组件库
- `zustand` - 状态管理
- `axios` - HTTP 客户端
- `react-hook-form` - 表单管理
- `@tanstack/react-query` - 数据获取
- `@tabler/icons-react` - 图标库

</td>
</tr>
</table>

---

## 📁 项目结构

```
go-manage-starter/
│
├── 📂 manage-backend/              # Go 后端服务
│   ├── 📂 cmd/                     # 应用程序入口
│   │   ├── server/                 # HTTP 服务器
│   │   └── migrate/                # 数据库迁移工具
│   ├── 📂 internal/                # 内部业务逻辑
│   │   ├── config/                 # 配置加载
│   │   ├── handler/                # HTTP 请求处理器
│   │   ├── middleware/             # 中间件 (认证、日志等)
│   │   ├── model/                  # 数据模型
│   │   ├── repository/             # 数据访问层
│   │   ├── service/                # 业务逻辑层
│   │   └── utils/                  # 工具函数
│   ├── 📂 pkg/                     # 可复用的公共包
│   │   ├── auth/                   # JWT 认证
│   │   ├── cache/                  # Redis 缓存
│   │   ├── database/               # 数据库连接
│   │   └── logger/                 # 日志工具 (zap + lumberjack)
│   ├── 📂 migrations/              # 数据库迁移文件
│   ├── 📂 config/                  # 配置文件
│   ├── 📂 docs/                    # Swagger API 文档
│   └── 📄 go.mod                   # Go 依赖管理
│
└── 📂 manage-frontend/             # Next.js 前端应用
    ├── 📂 src/
    │   ├── 📂 app/                 # Next.js App Router 页面
    │   │   ├── login/              # 登录页面
    │   │   ├── register/           # 注册页面
    │   │   ├── dashboard/          # 仪表盘
    │   │   └── unauthorized/       # 未授权页面
    │   ├── 📂 components/          # React 组件
    │   │   ├── ui/                 # shadcn/ui 组件
    │   │   ├── layout/             # 布局组件
    │   │   ├── auth/               # 认证相关组件
    │   │   └── dashboard/          # 仪表盘组件
    │   ├── 📂 api/                 # API 请求封装
    │   ├── 📂 hooks/               # 自定义 React Hooks
    │   ├── 📂 stores/              # Zustand 状态管理
    │   ├── 📂 types/               # TypeScript 类型定义
    │   └── 📂 lib/                 # 工具库
    └── 📄 package.json             # npm 依赖管理
```

---

## 🚀 快速开始

### 📋 环境要求

确保你的开发环境已安装以下工具：

| 工具       | 版本要求  | 下载链接                                     |
| ---------- | --------- | -------------------------------------------- |
| Go         | 1.21+     | [下载](https://go.dev/dl/)                   |
| Node.js    | 18+       | [下载](https://nodejs.org/)                  |
| PostgreSQL | 12+       | [下载](https://www.postgresql.org/download/) |
| Redis      | 6+ (可选) | [下载](https://redis.io/download)            |

### 🔧 后端设置

```bash
# 1. 进入后端目录
cd manage-backend

# 2. 安装 Go 依赖
go mod download

# 3. 配置环境文件
# 编辑 config/config.yaml config.development.yaml 配置数据库/Redis连接信息

# 4. 运行数据库迁移
go run cmd/migrate/main.go

# 5. 启动后端服务
go run cmd/server/main.go
```

**后端服务将运行在:** `http://localhost:8080`  
**Swagger API 文档:** `http://localhost:8080/swagger/index.html`

### 🎨 前端设置

```bash
# 1. 进入前端目录
cd manage-frontend

# 2. 安装 npm 依赖
npm install

# 3. 启动开发服务器 (使用 Turbopack)
npm run dev
```

**前端应用将运行在:** `http://localhost:3000`

### 🎉 开始使用

1. 访问 `http://localhost:3000`
2. 点击"注册"创建新账户
3. 使用注册的账户登录系统
4. 开始探索管理界面！

---

## 📡 API 接口

### 📚 Swagger API 文档

本项目集成了 Swagger 自动生成 API 文档，提供交互式的 API 测试界面。

#### 🚀 访问 Swagger 文档

1. **启动后端服务**

   ```bash
   cd manage-backend
   go run cmd/server/main.go
   ```

2. **访问 Swagger UI**
   ```
   http://localhost:8080/swagger/index.html
   ```

#### 🔄 更新 Swagger 文档

当你修改了 API 接口或添加了新的路由后，需要重新生成 Swagger 文档：

````bash
cd manage-backend

# 安装 swag 工具（首次使用）
go install github.com/swaggo/swag/cmd/swag@latest

# 生成/更新 Swagger 文档
swag init -g cmd/server/main.go


#### 📝 Swagger 注释示例

在代码中使用 Swagger 注释来描述 API：

```go
// @Summary 用户登录
// @Description 使用用户名、密码和验证码登录系统
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body model.LoginRequest true "登录凭证"
// @Success 200 {object} utils.APIResponse{data=model.LoginResponse}
// @Failure 401 {object} utils.APIResponse "认证失败"
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
    // 实现代码...
}
````

#### 🎯 Swagger 功能特性

- ✅ **交互式测试** - 直接在浏览器中测试 API
- ✅ **自动生成** - 从代码注释自动生成文档
- ✅ **类型定义** - 完整的请求/响应数据结构
- ✅ **认证支持** - 支持 JWT Token 认证测试

### 🔌 主要 API 端点

#### 认证接口

| 方法   | 路径                 | 描述           | 认证 |
| ------ | -------------------- | -------------- | ---- |
| `POST` | `/api/auth/register` | 用户注册       | ✅   |
| `POST` | `/api/auth/login`    | 用户登录       | ✅   |
| `GET`  | `/api/auth/captcha`  | 获取图形验证码 | ✅   |
| `POST` | `/api/auth/refresh`  | 刷新访问令牌   | ✅   |
| `POST` | `/api/auth/logout`   | 用户登出       | ✅   |

#### 用户管理

| 方法     | 路径             | 描述         | 认证 |
| -------- | ---------------- | ------------ | ---- |
| `GET`    | `/api/users`     | 获取用户列表 | ✅   |
| `GET`    | `/api/users/:id` | 获取用户详情 | ✅   |
| `POST`   | `/api/users`     | 创建用户     | ✅   |
| `PUT`    | `/api/users/:id` | 更新用户     | ✅   |
| `DELETE` | `/api/users/:id` | 删除用户     | ✅   |

> 💡 **提示:** 完整的 API 文档和详细说明请访问 Swagger UI

---

## 🛠️ Makefile 使用指南

后端项目提供了完整的 Makefile 来简化开发流程，包含构建、运行、测试、数据库管理等常用命令。

### 📊 完整命令列表

| 类别         | 命令                        | 描述             |
| ------------ | --------------------------- | ---------------- |
| **构建运行** | `make build`                | 构建应用程序     |
|              | `make run`                  | 运行应用程序     |
|              | `make dev`                  | 开发模式运行     |
| **环境启动** | `make dev-local`            | 本地开发环境     |
|              | `make run-test`             | 测试环境         |
|              | `make run-prod`             | 生产环境         |
| **测试**     | `make test`                 | 运行所有测试     |
|              | `make test-connection-dev`  | 测试连接（开发） |
|              | `make test-connection`      | 测试连接（测试） |
|              | `make test-connection-prod` | 测试连接（生产） |
| **数据库**   | `make migrate`              | 运行迁移（开发） |
|              | `make migrate-test`         | 运行迁移（测试） |
|              | `make migrate-prod`         | 运行迁移（生产） |
|              | `make seed`                 | 填充数据（开发） |
|              | `make seed-test`            | 填充数据（测试） |
|              | `make seed-prod`            | 填充数据（生产） |
|              | `make db-reset`             | 重置数据库       |
| **环境配置** | `make env-check`            | 检查环境变量     |
|              | `make env-setup-test`       | 设置测试环境     |
|              | `make env-setup-prod`       | 设置生产环境     |
| **工具**     | `make docs`                 | 生成 API 文档    |
|              | `make setup`                | 完整开发环境设置 |

### 💡 使用技巧

1. **首次设置**

   ```bash
   cd manage-backend
   make setup  # 自动完成：测试连接 → 迁移 → 填充数据
   ```

2. **日常开发**

   ```bash
   make dev    # 启动开发服务器
   make docs   # 修改 API 后更新文档
   ```

3. **数据库重置**

   ```bash
   make db-reset  # 快速重置数据库到初始状态
   ```

4. **多环境切换**

   ```bash
   # 使用环境变量文件
   source .env.test
   make run-test

   # 或直接指定环境
   ENVIRONMENT=test make run
   ```

---

## 📋 日志系统

本项目集成了生产级日志系统，基于 **zap + lumberjack** 实现。

### ✨ 功能特性

- ✅ **结构化日志** - 支持 JSON 和 Console 两种格式
- ✅ **文件输出** - 自动写入日志文件
- ✅ **日志轮转** - 按大小自动切割日志文件
- ✅ **日志压缩** - 自动压缩旧日志节省空间
- ✅ **环境配置** - 开发/生产环境独立配置
- ✅ **多输出** - 同时输出到控制台和文件

### 🔧 配置说明

在 `config/config.yaml` 或环境特定配置文件中：

```yaml
log:
  level: info              # 日志级别: debug, info, warn, error
  format: console          # 日志格式: json, console
  output_path: logs/app.log # 输出路径: stdout 或文件路径
  max_size: 100            # 单个日志文件最大大小(MB)
  max_backups: 10          # 保留的旧日志文件最大数量
  max_age: 30              # 保留旧日志文件的最大天数
  compress: true           # 是否压缩旧日志文件
```

### 📝 环境配置示例

**开发环境** (`config.development.yaml`)
```yaml
log:
  level: debug
  format: console          # 彩色控制台输出，便于调试
  output_path: logs/dev.log # 同时输出到文件和控制台
```

**生产环境** (`config.production.yaml`)
```yaml
log:
  level: warn
  format: json             # JSON 格式便于日志分析工具处理
  output_path: logs/app.log
  max_size: 100
  max_backups: 10
  max_age: 30
  compress: true           # 压缩旧日志节省空间
```

### 💡 使用示例

```go
import (
    "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
    "go.uber.org/zap"
)

// 普通日志
logger.Info("用户登录成功")

// 结构化日志
logger.Info("用户登录成功", 
    zap.String("username", "admin"),
    zap.String("ip", "192.168.1.1"),
    zap.Duration("duration", time.Since(start)))

// 错误日志
logger.Error("数据库连接失败", 
    zap.Error(err),
    zap.String("database", "postgres"))
```

### 📂 日志文件位置

- **开发环境**: `manage-backend/logs/dev.log`
- **生产环境**: `manage-backend/logs/app.log`
- **轮转文件**: `app-2025-10-09T11-05-14.123.log.gz`

> 💡 **提示**: 日志文件已添加到 `.gitignore`，不会提交到版本控制

---

## 📝 开发说明

### ⚠️ 当前限制

| 限制        | 说明                                        |
| ----------- | ------------------------------------------- |
| 🔐 权限系统 | 目前仅实现 JWT 认证，暂无 RBAC 角色权限控制 |
| 📄 功能范围 | 主要包含用户认证和用户管理功能              |
| 🎨 页面数量 | 前端页面相对简单，主要为登录/注册和基础布局 |

### 🎯 适用场景

- ✅ 中小型管理系统的快速原型开发
- ✅ 学习 Go + React 全栈开发的实践项目
- ✅ 作为其他项目的基础模板和脚手架

### 🔨 开发建议

1. **配置管理**: 使用不同的配置文件管理开发/生产环境
2. **数据库迁移**: 通过 `migrations/` 目录管理数据库版本
3. **API 文档**: 使用 Swagger 注释自动生成 API 文档
4. **代码规范**: 遵循 Go 和 TypeScript 的最佳实践

---

## 📄 许可证

本项目采用 [MIT License](./License) 开源协议

**Copyright © 2025 [xiaozhu](https://github.com/XIAOZHUXUEJAVA)**
