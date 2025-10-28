"use client";

import { useAuthStore } from "@/stores/authStore";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { LogOut, Rocket, History, Bell, Database, Layout } from "lucide-react";
import { DashboardHeader } from "@/components/layout/dashboard-header";

/**
 * Dashboard 主页
 * 显示系统概览和主要功能入口
 */
export default function DashboardPage() {
  const { user, logout } = useAuthStore();

  const handleLogout = () => {
    logout();
  };

  // 面包屑导航配置
  const breadcrumbs = [{ label: "Dashboard" }];

  // 头部操作按钮
  const headerActions = (
    <div className="flex items-center space-x-4">
      <div className="flex items-center space-x-2">
        <Avatar className="h-8 w-8">
          <AvatarImage src={user?.avatar} />
          <AvatarFallback>
            {user?.username?.charAt(0).toUpperCase()}
          </AvatarFallback>
        </Avatar>
        <span className="text-sm font-medium text-gray-700">
          {user?.username}
        </span>
        <Badge variant="secondary">{user?.role}</Badge>
      </div>
      <Button
        variant="outline"
        size="sm"
        onClick={handleLogout}
        className="flex items-center space-x-1"
      >
        <LogOut className="h-4 w-4" />
        <span>退出</span>
      </Button>
    </div>
  );

  return (
    <>
      <DashboardHeader breadcrumbs={breadcrumbs} actions={headerActions} />
      <div className="flex flex-1 flex-col gap-5 p-4 pt-0">
        {/* 系统介绍 */}
        <div className="bg-white dark:bg-gray-900 rounded-lg p-5 border-l-4 border-blue-600 shadow-sm">
          <div className="flex items-center gap-3 mb-3">
            <Rocket className="h-5 w-5 text-blue-600" />
            <h2 className="text-lg font-semibold text-gray-800 dark:text-gray-200">
              Go 管理系统起始模板
            </h2>
          </div>
          <p className="text-sm text-gray-600 dark:text-gray-400 leading-relaxed mb-3">
            基于 Go + React + Next.js
            构建的前后端分离管理系统，提供完整的用户认证和权限管理功能。
          </p>
          <div className="flex flex-wrap gap-2">
            <span className="px-2 py-1 text-xs bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 rounded">
              架构清晰
            </span>
            <span className="px-2 py-1 text-xs bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 rounded">
              JWT认证
            </span>
            <span className="px-2 py-1 text-xs bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 rounded">
              RBAC权限
            </span>
            <span className="px-2 py-1 text-xs bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 rounded">
              开箱即用
            </span>
          </div>
        </div>

        {/* 技术栈 */}
        <div className="grid gap-4 md:grid-cols-2">
          <div className="bg-white dark:bg-gray-900 rounded-lg p-4 border border-gray-200 dark:border-gray-700">
            <h3 className="text-base font-medium text-gray-800 dark:text-gray-200 mb-3 flex items-center gap-2">
              <Database className="h-4 w-4" />
              后端技术
            </h3>
            <div className="space-y-2 text-sm">
              <div className="flex justify-between items-center">
                <span className="text-gray-600 dark:text-gray-400">
                  Go + Gin
                </span>
                <Badge variant="secondary" className="text-xs">
                  v1.25
                </Badge>
              </div>
              <div className="flex justify-between items-center">
                <span className="text-gray-600 dark:text-gray-400">
                  PostgreSQL + GORM
                </span>
                <Badge variant="secondary" className="text-xs">
                  ORM
                </Badge>
              </div>
              <div className="flex justify-between items-center">
                <span className="text-gray-600 dark:text-gray-400">
                  JWT + Casbin
                </span>
                <Badge variant="secondary" className="text-xs">
                  Auth
                </Badge>
              </div>
              <div className="flex justify-between items-center">
                <span className="text-gray-600 dark:text-gray-400">
                  Redis + Swagger
                </span>
                <Badge variant="secondary" className="text-xs">
                  Cache
                </Badge>
              </div>
            </div>
          </div>

          <div className="bg-white dark:bg-gray-900 rounded-lg p-4 border border-gray-200 dark:border-gray-700">
            <h3 className="text-base font-medium text-gray-800 dark:text-gray-200 mb-3 flex items-center gap-2">
              <Layout className="h-4 w-4" />
              前端技术
            </h3>
            <div className="space-y-2 text-sm">
              <div className="flex justify-between items-center">
                <span className="text-gray-600 dark:text-gray-400">
                  Next.js + React
                </span>
                <Badge variant="secondary" className="text-xs">
                  v15
                </Badge>
              </div>
              <div className="flex justify-between items-center">
                <span className="text-gray-600 dark:text-gray-400">
                  Tailwind + shadcn/ui
                </span>
                <Badge variant="secondary" className="text-xs">
                  UI
                </Badge>
              </div>
              <div className="flex justify-between items-center">
                <span className="text-gray-600 dark:text-gray-400">
                  Zustand + TanStack
                </span>
                <Badge variant="secondary" className="text-xs">
                  State
                </Badge>
              </div>
              <div className="flex justify-between items-center">
                <span className="text-gray-600 dark:text-gray-400">
                  TypeScript + Zod
                </span>
                <Badge variant="secondary" className="text-xs">
                  Type
                </Badge>
              </div>
            </div>
          </div>
        </div>

        <div className="grid gap-4 md:grid-cols-2">
          {/* 更新日志 */}
          <div className="bg-white dark:bg-gray-900 rounded-lg p-4 border border-gray-200 dark:border-gray-700">
            <h3 className="text-base font-medium text-gray-800 dark:text-gray-200 mb-3 flex items-center gap-2">
              <History className="h-4 w-4" />
              更新日志
            </h3>
            <div className="space-y-3">
              <div className="pb-3 border-b border-gray-100 dark:border-gray-800">
                <div className="flex items-center justify-between mb-1">
                  <span className="text-xs text-gray-500">2025-10-22</span>
                  <Badge className="bg-green-600 text-xs h-5">v1.0.0</Badge>
                </div>
                <p className="text-sm text-gray-700 dark:text-gray-300 mb-1">
                  系统正式发布
                </p>
                <p className="text-xs text-gray-500 dark:text-gray-500">
                  用户认证、RBAC权限、菜单管理
                </p>
              </div>
              <div className="pb-3 border-b border-gray-100 dark:border-gray-800">
                <div className="flex items-center justify-between mb-1">
                  <span className="text-xs text-gray-500">2025-10-15</span>
                  <Badge variant="secondary" className="text-xs h-5">
                    v0.9.0
                  </Badge>
                </div>
                <p className="text-sm text-gray-700 dark:text-gray-300 mb-1">
                  Beta版本
                </p>
                <p className="text-xs text-gray-500 dark:text-gray-500">
                  审计日志、响应式优化
                </p>
              </div>
              <div>
                <div className="flex items-center justify-between mb-1">
                  <span className="text-xs text-gray-500">2025-10-01</span>
                  <Badge variant="outline" className="text-xs h-5">
                    v0.5.0
                  </Badge>
                </div>
                <p className="text-sm text-gray-700 dark:text-gray-300 mb-1">
                  Alpha版本
                </p>
                <p className="text-xs text-gray-500 dark:text-gray-500">
                  基础框架搭建
                </p>
              </div>
            </div>
          </div>

          {/* 系统公告 */}
          <div className="bg-white dark:bg-gray-900 rounded-lg p-4 border border-gray-200 dark:border-gray-700">
            <h3 className="text-base font-medium text-gray-800 dark:text-gray-200 mb-3 flex items-center gap-2">
              <Bell className="h-4 w-4" />
              系统公告
            </h3>
            <div className="space-y-2">
              <div className="p-3 bg-blue-50 dark:bg-blue-950/30 rounded border-l-2 border-blue-500">
                <div className="flex items-start justify-between mb-1">
                  <span className="text-sm font-medium text-gray-800 dark:text-gray-200">
                    欢迎使用
                  </span>
                  <span className="text-xs text-gray-500">10-22</span>
                </div>
                <p className="text-xs text-gray-600 dark:text-gray-400">
                  功能完善的管理系统模板，包含用户管理、权限控制等核心功能
                </p>
              </div>

              <div className="p-3 bg-amber-50 dark:bg-amber-950/30 rounded border-l-2 border-amber-500">
                <div className="flex items-start justify-between mb-1">
                  <span className="text-sm font-medium text-gray-800 dark:text-gray-200">
                    功能更新
                  </span>
                  <span className="text-xs text-gray-500">10-20</span>
                </div>
                <p className="text-xs text-gray-600 dark:text-gray-400">
                  v1.0.0版本发布，新增数据字典、审计日志功能
                </p>
              </div>

              <div className="p-3 bg-green-50 dark:bg-green-950/30 rounded border-l-2 border-green-500">
                <div className="flex items-start justify-between mb-1">
                  <span className="text-sm font-medium text-gray-800 dark:text-gray-200">
                    维护完成
                  </span>
                  <span className="text-xs text-gray-500">10-18</span>
                </div>
                <p className="text-xs text-gray-600 dark:text-gray-400">
                  系统维护已完成，所有功能已恢复正常
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
