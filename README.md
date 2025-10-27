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

这是一个管理系统起始模板，采用现代化技术栈构建，遵循最佳实践和清晰的代码架构。适合作为中小型管理系统的开发基础，或用于学习全栈开发。

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

| 功能模块         | 描述                                        | 状态    |
| ---------------- | ------------------------------------------- | ------- |
| 🔐 用户认证      | 注册、登录、JWT Token 验证、刷新令牌        | ✅ 完成 |
| 🔑 密码重置      | 邮箱验证、重置令牌、密码找回功能            | ✅ 完成 |
| 🖼️ 图形验证码    | 防机器人注册/登录保护                       | ✅ 完成 |
| 🛡️ 登录限流      | 防暴力破解、登录频率限制                    | ✅ 完成 |
| 👥 用户管理      | 完整的用户 CRUD、用户状态管理、用户角色分配 | ✅ 完成 |
| 🔑 RBAC 权限系统 | 基于 Casbin 的角色权限管理                  | ✅ 完成 |
| 🎭 角色管理      | 角色 CRUD、角色权限分配                     | ✅ 完成 |
| 🔓 权限管理      | 权限 CRUD、权限树结构、资源权限控制         | ✅ 完成 |
| 📋 菜单管理      | 动态菜单、菜单树结构、用户菜单权限          | ✅ 完成 |
| 📚 数据字典      | 字典类型和字典项管理、支持 JSONB 扩展字段   | ✅ 完成 |
| 🎨 响应式界面    | 适配桌面和移动端、暗黑模式支持              | ✅ 完成 |
| ⚠️ 错误处理      | 统一的错误处理机制                          | ✅ 完成 |
| 📚 API 文档      | Swagger 自动生成文档                        | ✅ 完成 |

### 🚧 规划中的功能

- 📊 **数据统计面板** - 可视化数据展示和图表
- 🔔 **消息通知** - 系统消息推送和实时通知
- 📤 **文件上传** - 文件管理和对象存储集成
- 🔍 **高级搜索** - 全文搜索和复杂查询
- 📱 **移动端优化** - PWA 支持和移动端专属功能
- 🌐 **国际化** - 多语言支持

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

- `golang-jwt/jwt/v5` - JWT 认证和令牌管理
- `casbin/casbin/v2` - RBAC 权限控制
- `casbin/gorm-adapter/v3` - Casbin GORM 适配器
- `redis/go-redis/v9` - Redis 缓存和会话管理
- `spf13/viper` - 配置管理
- `uber.org/zap` - 结构化日志
- `natefinch/lumberjack` - 日志轮转
- `mojocn/base64Captcha` - 验证码生成
- `swaggo/swag` - API 文档生成
- `golang.org/x/crypto` - 密码加密
- `gopkg.in/gomail.v2` - 邮件发送服务
- `gorm.io/datatypes` - GORM 数据类型扩展

</td>
<td width="50%" valign="top">

### 前端 (manage-frontend)

**核心框架**

- ![Next.js](https://img.shields.io/badge/Next.js-15-black?style=flat&logo=next.js) **Next.js 15** - React 框架 (App Router)
- ![React](https://img.shields.io/badge/React-19-61DAFB?style=flat&logo=react) **React 19** - UI 库
- ![TypeScript](https://img.shields.io/badge/TypeScript-5-3178C6?style=flat&logo=typescript) **TypeScript** - 类型安全
- ![Tailwind](https://img.shields.io/badge/Tailwind-CSS-38B2AC?style=flat&logo=tailwind-css) **Tailwind CSS** - 样式框架

**核心依赖**

- `shadcn/ui` - UI 组件库（基于 Radix UI）
- `zustand` - 轻量级状态管理
- `axios` - HTTP 客户端
- `react-hook-form` - 表单管理和验证
- `zod` - TypeScript 优先的模式验证
- `@tanstack/react-query` - 数据获取和缓存
- `@tanstack/react-table` - 强大的表格组件
- `@tabler/icons-react` - 图标库
- `lucide-react` - 现代图标库
- `@dnd-kit` - 拖拽功能
- `dayjs` - 日期处理
- `framer-motion` - 动画库
- `recharts` - 图表库
- `sonner` - Toast 通知组件
- `vaul` - 抽屉组件
- `next-themes` - 主题切换
- `nprogress` - 页面加载进度条

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
│   │   ├── 📂 handler/                # HTTP 请求处理器
│   │   │   ├── user.go             # 用户管理
│   │   │   ├── role_handler.go     # 角色管理
│   │   │   ├── permission_handler.go # 权限管理
│   │   │   ├── menu_handler.go     # 菜单管理
│   │   │   ├── audit_log_handler.go # 审计日志
│   │   │   ├── dict_handler.go     # 数据字典
│   │   │   ├── captcha.go          # 验证码
│   │   │   ├── password_reset_handler.go # 密码重置
│   │   │   └── routes.go           # 路由配置
│   │   ├── middleware/             # 中间件
│   │   │   ├── auth.go             # JWT 认证
│   │   │   ├── permission.go       # 权限验证
│   │   │   ├── audit.go            # 审计日志
│   │   │   ├── cors.go             # CORS 配置
│   │   │   └── captcha.go          # 验证码验证
│   │   ├── 📂 model/                  # 数据模型
│   │   │   ├── user.go             # 用户模型
│   │   │   ├── role.go             # 角色模型
│   │   │   ├── permission.go       # 权限模型
│   │   │   ├── menu.go             # 菜单模型
│   │   │   ├── audit_log.go        # 审计日志模型
│   │   │   ├── dict.go             # 数据字典模型
│   │   │   └── password_reset.go   # 密码重置模型
│   │   ├── repository/             # 数据访问层
│   │   ├── 📂 service/                # 业务逻辑层
│   │   │   ├── user.go             # 用户服务
│   │   │   ├── role_service.go     # 角色服务
│   │   │   ├── permission_service.go # 权限服务
│   │   │   ├── menu_service.go     # 菜单服务
│   │   │   ├── casbin_service.go   # Casbin 服务
│   │   │   ├── session.go          # 会话管理
│   │   │   ├── captcha.go          # 验证码服务
│   │   │   ├── audit_log_service.go # 审计日志服务
│   │   │   ├── dict_service.go     # 数据字典服务
│   │   │   ├── email_service.go    # 邮件服务
│   │   │   ├── password_reset_service.go # 密码重置服务
│   │   │   ├── password_reset_redis.go # 密码重置 Redis 服务
│   │   │   └── login_rate_limit.go # 登录限流服务
│   │   └── utils/                  # 工具函数
│   ├── 📂 pkg/                     # 可复用的公共包
│   │   ├── auth/                   # JWT 认证
│   │   ├── cache/                  # Redis 缓存
│   │   ├── casbin/                 # Casbin 初始化
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
    │   │   ├── forgot-password/    # 忘记密码页面
    │   │   ├── reset-password/     # 重置密码页面
    │   │   ├── dashboard/          # 仪表盘
    │   │   │   └── welcome/        # 欢迎页面
    │   │   ├── logs/               # 审计日志页面
    │   │   ├── system/             # 系统管理
    │   │   │   ├── users/          # 用户管理
    │   │   │   ├── roles/          # 角色管理
    │   │   │   ├── permissions/    # 权限管理
    │   │   │   ├── menus/          # 菜单管理
    │   │   │   └── dict/           # 数据字典
    │   │   ├── unauthorized/       # 未授权页面
    │   │   └── forbidden/          # 禁止访问页面
    │   ├── 📂 components/          # React 组件
    │   │   ├── ui/                 # shadcn/ui 组件
    │   │   ├── layout/             # 布局组件
    │   │   ├── auth/               # 认证相关组件
    │   │   ├── dashboard/          # 仪表盘组件
    │   │   ├── features/           # 功能模块组件
    │   │   │   ├── audit/          # 审计日志组件
    │   │   │   └── system/         # 系统管理组件
    │   │   │       ├── user/       # 用户组件
    │   │   │       ├── role/       # 角色组件
    │   │   │       ├── permission/ # 权限组件
    │   │   │       ├── menu/       # 菜单组件
    │   │   │       └── dict/       # 数据字典组件
    │   │   ├── forms/              # 表单组件
    │   │   └── common/             # 通用组件
    │   ├── 📂 api/                 # API 请求封装
    │   │   ├── auth.ts             # 认证 API
    │   │   ├── user.ts             # 用户 API
    │   │   ├── role.ts             # 角色 API
    │   │   ├── permission.ts       # 权限 API
    │   │   ├── menu.ts             # 菜单 API
    │   │   ├── audit.ts            # 审计日志 API
    │   │   └── dict.ts             # 数据字典 API
    │   ├── 📂 hooks/               # 自定义 React Hooks
    │   ├── 📂 stores/              # Zustand 状态管理
    │   │   ├── authStore.ts        # 认证状态
    │   │   ├── themeStore.ts       # 主题状态
    │   │   └── sidebarStore.ts     # 侧边栏状态
    │   ├── 📂 types/               # TypeScript 类型定义
    │   ├── 📂 lib/                 # 工具库
    │   └── 📄 middleware.ts        # Next.js 中间件
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

# 4. 执行scripts目录数据库文件

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
2. 用户名：admin
3. 密码：admin123

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

| 方法   | 路径                        | 描述           | 认证 |
| ------ | --------------------------- | -------------- | ---- |
| `POST` | `/api/auth/register`        | 用户注册       | ❌   |
| `POST` | `/api/auth/login`           | 用户登录       | ❌   |
| `GET`  | `/api/auth/captcha`         | 获取图形验证码 | ❌   |
| `POST` | `/api/auth/refresh`         | 刷新访问令牌   | ❌   |
| `POST` | `/api/auth/logout`          | 用户登出       | ✅   |
| `POST` | `/api/auth/forgot-password` | 忘记密码       | ❌   |
| `POST` | `/api/auth/verify-reset-token` | 验证重置令牌   | ❌   |
| `POST` | `/api/auth/reset-password`  | 重置密码       | ❌   |

#### 用户管理

| 方法     | 路径                                  | 描述               | 认证 |
| -------- | ------------------------------------- | ------------------ | ---- |
| `GET`    | `/api/users`                          | 获取用户列表       | ✅   |
| `GET`    | `/api/users/:id`                      | 获取用户详情       | ✅   |
| `POST`   | `/api/users`                          | 创建用户           | ✅   |
| `PUT`    | `/api/users/:id`                      | 更新用户           | ✅   |
| `DELETE` | `/api/users/:id`                      | 删除用户           | ✅   |
| `GET`    | `/api/users/profile`                  | 获取当前用户信息   | ✅   |
| `PUT`    | `/api/users/profile`                  | 更新当前用户信息   | ✅   |
| `GET`    | `/api/users/:id/roles`                | 获取用户角色       | ✅   |
| `PUT`    | `/api/users/:id/roles`                | 分配用户角色       | ✅   |
| `GET`    | `/api/users/permissions`              | 获取用户权限       | ✅   |
| `GET`    | `/api/users/check-username/:username` | 检查用户名是否可用 | ❌   |
| `GET`    | `/api/users/check-email/:email`       | 检查邮箱是否可用   | ❌   |
| `POST`   | `/api/users/check-availability`       | 批量检查可用性     | ❌   |

#### 角色管理

| 方法     | 路径                         | 描述         | 认证 |
| -------- | ---------------------------- | ------------ | ---- |
| `GET`    | `/api/roles`                 | 获取角色列表 | ✅   |
| `GET`    | `/api/roles/all`             | 获取所有角色 | ✅   |
| `GET`    | `/api/roles/:id`             | 获取角色详情 | ✅   |
| `POST`   | `/api/roles`                 | 创建角色     | ✅   |
| `PUT`    | `/api/roles/:id`             | 更新角色     | ✅   |
| `DELETE` | `/api/roles/:id`             | 删除角色     | ✅   |
| `GET`    | `/api/roles/:id/permissions` | 获取角色权限 | ✅   |
| `PUT`    | `/api/roles/:id/permissions` | 分配角色权限 | ✅   |

#### 权限管理

| 方法     | 路径                                  | 描述           | 认证 |
| -------- | ------------------------------------- | -------------- | ---- |
| `GET`    | `/api/permissions`                    | 获取权限列表   | ✅   |
| `GET`    | `/api/permissions/all`                | 获取所有权限   | ✅   |
| `GET`    | `/api/permissions/tree`               | 获取权限树结构 | ✅   |
| `GET`    | `/api/permissions/:id`                | 获取权限详情   | ✅   |
| `POST`   | `/api/permissions`                    | 创建权限       | ✅   |
| `PUT`    | `/api/permissions/:id`                | 更新权限       | ✅   |
| `DELETE` | `/api/permissions/:id`                | 删除权限       | ✅   |
| `GET`    | `/api/permissions/resource/:resource` | 按资源获取权限 | ✅   |
| `GET`    | `/api/permissions/type/:type`         | 按类型获取权限 | ✅   |

#### 菜单管理

| 方法     | 路径                      | 描述           | 认证 |
| -------- | ------------------------- | -------------- | ---- |
| `GET`    | `/api/menus/tree`         | 获取菜单树结构 | ✅   |
| `GET`    | `/api/menus/tree/visible` | 获取可见菜单树 | ✅   |
| `GET`    | `/api/menus/user`         | 获取用户菜单树 | ✅   |
| `GET`    | `/api/menus/:id`          | 获取菜单详情   | ✅   |
| `POST`   | `/api/menus`              | 创建菜单       | ✅   |
| `PUT`    | `/api/menus/:id`          | 更新菜单       | ✅   |
| `PUT`    | `/api/menus/order`        | 更新菜单顺序   | ✅   |
| `DELETE` | `/api/menus/:id`          | 删除菜单       | ✅   |

#### 审计日志

| 方法   | 路径                    | 描述         | 认证 |
| ------ | ----------------------- | ------------ | ---- |
| `GET`  | `/api/audit-logs`       | 查询审计日志 | ✅   |
| `GET`  | `/api/audit-logs/:id`   | 获取日志详情 | ✅   |
| `POST` | `/api/audit-logs/clean` | 清理旧日志   | ✅   |

#### 数据字典

| 方法     | 路径                            | 描述             | 认证 |
| -------- | ------------------------------- | ---------------- | ---- |
| `GET`    | `/api/dict-types`               | 获取字典类型列表 | ✅   |
| `GET`    | `/api/dict-types/all`           | 获取所有字典类型 | ✅   |
| `GET`    | `/api/dict-types/:id`           | 获取字典类型详情 | ✅   |
| `POST`   | `/api/dict-types`               | 创建字典类型     | ✅   |
| `PUT`    | `/api/dict-types/:id`           | 更新字典类型     | ✅   |
| `DELETE` | `/api/dict-types/:id`           | 删除字典类型     | ✅   |
| `GET`    | `/api/dict-items`               | 获取字典项列表   | ✅   |
| `GET`    | `/api/dict-items/by-type/:code` | 按类型获取字典项 | ✅   |
| `GET`    | `/api/dict-items/:id`           | 获取字典项详情   | ✅   |
| `POST`   | `/api/dict-items`               | 创建字典项       | ✅   |
| `PUT`    | `/api/dict-items/:id`           | 更新字典项       | ✅   |
| `DELETE` | `/api/dict-items/:id`           | 删除字典项       | ✅   |

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
  level: info # 日志级别: debug, info, warn, error
  format: console # 日志格式: json, console
  output_path: logs/app.log # 输出路径: stdout 或文件路径
  max_size: 100 # 单个日志文件最大大小(MB)
  max_backups: 10 # 保留的旧日志文件最大数量
  max_age: 30 # 保留旧日志文件的最大天数
  compress: true # 是否压缩旧日志文件
```

### 📝 环境配置示例

**开发环境** (`config.development.yaml`)

```yaml
log:
  level: debug
  format: console # 彩色控制台输出，便于调试
  output_path: logs/dev.log # 同时输出到文件和控制台
```

**生产环境** (`config.production.yaml`)

```yaml
log:
  level: warn
  format: json # JSON 格式便于日志分析工具处理
  output_path: logs/app.log
  max_size: 100
  max_backups: 10
  max_age: 30
  compress: true # 压缩旧日志节省空间
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

### 🏗️ 架构特点

- **分层架构**: Handler → Service → Repository，职责清晰
- **依赖注入**: 通过构造函数注入依赖，便于测试
- **中间件模式**: 认证、权限、审计日志等功能模块化
- **RESTful API**: 遵循 REST 规范设计 API 接口
- **类型安全**: TypeScript + Zod 确保前端类型安全

### 🔒 权限系统说明

本项目实现了完整的 RBAC（基于角色的访问控制）权限系统：

- **Casbin 集成**: 使用 Casbin 作为权限引擎
- **动态权限**: 支持运行时动态修改权限规则
- **多级权限**: 支持资源级、操作级权限控制
- **权限继承**: 角色可以继承多个权限
- **菜单权限**: 前端菜单根据用户权限动态生成

### 🎯 适用场景

- ✅ 管理系统的快速开发
- ✅ 需要完整 RBAC 权限控制的项目
- ✅ 学习 Go + React 全栈开发的实践项目
- ✅ 作为其他项目的基础模板和脚手架

### 🔨 开发建议

1. **配置管理**: 使用不同的配置文件管理开发/测试/生产环境
2. **数据库迁移**: 通过 `migrations/` 目录管理数据库版本变更
3. **API 文档**: 使用 Swagger 注释自动生成和维护 API 文档
4. **代码规范**: 遵循 Go 和 TypeScript 的最佳实践和编码规范
5. **权限设计**: 合理规划资源和操作，避免权限粒度过细或过粗
6. **审计日志**: 关键操作务必记录审计日志，便于追溯
7. **数据字典**: 使用数据字典管理系统配置项，避免硬编码

### 🧪 测试建议

- 使用 `testify` 编写单元测试
- 为关键业务逻辑编写集成测试
- 使用 Swagger UI 进行 API 接口测试
- 前端使用 React Testing Library 测试组件

---

## 🎨 前端特性

### 组件库

- **shadcn/ui**: 基于 Radix UI 的高质量组件库
- **响应式设计**: 完美适配桌面和移动端
- **暗黑模式**: 支持亮色/暗色主题切换
- **表格组件**: 基于 TanStack Table 的强大表格功能
  - 排序、筛选、分页
  - 列显示/隐藏
  - 行选择
  - 自定义单元格渲染

### 状态管理

- **Zustand**: 轻量级状态管理
  - `authStore`: 用户认证状态
  - `themeStore`: 主题配置
  - `sidebarStore`: 侧边栏状态

### 表单处理

- **React Hook Form + Zod**: 类型安全的表单验证
- **实时验证**: 输入时即时反馈
- **错误提示**: 友好的错误信息展示

### 数据获取

- **TanStack Query**: 强大的数据获取和缓存
  - 自动缓存和重新验证
  - 乐观更新
  - 后台数据同步

---

## 🔐 安全特性

### 后端安全

- **密码加密**: 使用 bcrypt 加密存储密码
- **JWT 认证**: 访问令牌 + 刷新令牌双令牌机制
- **会话管理**: Redis 存储会话，支持强制登出
- **CORS 配置**: 可配置的跨域资源共享
- **SQL 注入防护**: GORM 参数化查询
- **XSS 防护**: 输入验证和输出转义

### 前端安全

- **Token 存储**: 安全的 Token 存储机制
- **路由守卫**: Next.js Middleware 实现路由保护
- **权限控制**: 基于角色的页面和组件显示控制

---

## 📊 数据库设计

### 核心表结构

- **users**: 用户基础信息
- **roles**: 角色定义
- **permissions**: 权限定义
- **user_roles**: 用户-角色关联
- **role_permissions**: 角色-权限关联
- **menus**: 菜单配置
- **audit_logs**: 审计日志
- **dict_types**: 字典类型
- **dict_items**: 字典项（支持 JSONB 扩展字段）
- **password_resets**: 密码重置记录
- **casbin_rule**: Casbin 权限规则

---

## 📄 许可证

本项目采用 [MIT License](./License) 开源协议

**Copyright © 2025 [xiaozhu](https://github.com/XIAOZHUXUEJAVA)**
