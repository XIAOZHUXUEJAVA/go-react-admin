"use client";

import { useAuthStore } from "@/stores/authStore";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { DashboardHeader } from "@/components/layout/dashboard-header";
import {
  Sparkles,
  Users,
  Shield,
  Settings,
  BarChart3,
  FileText,
  Clock,
  TrendingUp,
  ArrowRight,
} from "lucide-react";
import Link from "next/link";

/**
 * 欢迎页面
 * 展示系统概览和快速入口
 */
export default function WelcomePage() {
  const { user } = useAuthStore();

  // 面包屑导航配置
  const breadcrumbs = [
    { label: "Dashboard", href: "/dashboard" },
    { label: "欢迎" },
  ];

  // 获取当前时间问候语
  const getGreeting = () => {
    const hour = new Date().getHours();
    if (hour < 6) return "夜深了";
    if (hour < 9) return "早上好";
    if (hour < 12) return "上午好";
    if (hour < 14) return "中午好";
    if (hour < 18) return "下午好";
    if (hour < 22) return "晚上好";
    return "夜深了";
  };

  // 快速入口
  const quickLinks = [
    {
      title: "用户管理",
      description: "管理系统用户和权限",
      icon: Users,
      href: "/system/users",
      color: "text-blue-600",
      bgColor: "bg-blue-50",
    },
    {
      title: "角色管理",
      description: "配置角色和权限",
      icon: Shield,
      href: "/system/roles",
      color: "text-purple-600",
      bgColor: "bg-purple-50",
    },
    {
      title: "菜单管理",
      description: "配置系统菜单结构",
      icon: FileText,
      href: "/system/menus",
      color: "text-green-600",
      bgColor: "bg-green-50",
    },
    {
      title: "系统设置",
      description: "系统参数配置",
      icon: Settings,
      href: "/system/settings",
      color: "text-orange-600",
      bgColor: "bg-orange-50",
    },
  ];

  // 统计数据（示例）
  const stats = [
    {
      title: "总用户数",
      value: "1,234",
      change: "+12.5%",
      trend: "up",
      icon: Users,
    },
    {
      title: "活跃用户",
      value: "856",
      change: "+8.2%",
      trend: "up",
      icon: TrendingUp,
    },
    {
      title: "今日访问",
      value: "3,456",
      change: "+23.1%",
      trend: "up",
      icon: BarChart3,
    },
    {
      title: "系统运行",
      value: "99.9%",
      change: "稳定",
      trend: "stable",
      icon: Clock,
    },
  ];

  return (
    <div className="flex flex-col gap-6">
      <DashboardHeader breadcrumbs={breadcrumbs} />

      <div className="grid gap-6 px-4">
        {/* 欢迎卡片 */}
        <Card className="border-none shadow-lg bg-gradient-to-br from-blue-50 via-white to-purple-50">
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div className="space-y-2">
                <div className="flex items-center gap-3">
                  <Avatar className="h-16 w-16 border-2 border-white shadow-md">
                    <AvatarImage src={user?.avatar} />
                    <AvatarFallback className="text-xl bg-gradient-to-br from-blue-500 to-purple-500 text-white">
                      {user?.username?.charAt(0).toUpperCase()}
                    </AvatarFallback>
                  </Avatar>
                  <div>
                    <h1 className="text-3xl font-bold tracking-tight">
                      {getGreeting()}，{user?.username}！
                    </h1>
                    <p className="text-muted-foreground mt-1">
                      欢迎回到管理系统
                    </p>
                  </div>
                </div>
                <div className="flex items-center gap-2 mt-4">
                  <Badge variant="secondary" className="text-sm">
                    {user?.role}
                  </Badge>
                  <Badge variant="outline" className="text-sm">
                    {user?.email}
                  </Badge>
                </div>
              </div>
              <Sparkles className="h-24 w-24 text-yellow-400 opacity-50" />
            </div>
          </CardContent>
        </Card>

        {/* 统计卡片 */}
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
          {stats.map((stat, index) => {
            const Icon = stat.icon;
            return (
              <Card key={index}>
                <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                  <CardTitle className="text-sm font-medium">
                    {stat.title}
                  </CardTitle>
                  <Icon className="h-4 w-4 text-muted-foreground" />
                </CardHeader>
                <CardContent>
                  <div className="text-2xl font-bold">{stat.value}</div>
                  <p className="text-xs text-muted-foreground mt-1">
                    {stat.trend === "up" && (
                      <span className="text-green-600">{stat.change}</span>
                    )}
                    {stat.trend === "down" && (
                      <span className="text-red-600">{stat.change}</span>
                    )}
                    {stat.trend === "stable" && (
                      <span className="text-blue-600">{stat.change}</span>
                    )}
                    {stat.trend === "up" && " 较上周"}
                  </p>
                </CardContent>
              </Card>
            );
          })}
        </div>

        {/* 快速入口 */}
        <Card>
          <CardHeader>
            <CardTitle>快速入口</CardTitle>
            <CardDescription>
              快速访问常用功能模块
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
              {quickLinks.map((link, index) => {
                const Icon = link.icon;
                return (
                  <Link key={index} href={link.href}>
                    <Card className="hover:shadow-md transition-shadow cursor-pointer border-2 hover:border-primary">
                      <CardContent className="pt-6">
                        <div className="flex flex-col items-center text-center space-y-3">
                          <div
                            className={`${link.bgColor} p-3 rounded-lg`}
                          >
                            <Icon className={`h-6 w-6 ${link.color}`} />
                          </div>
                          <div>
                            <h3 className="font-semibold">{link.title}</h3>
                            <p className="text-xs text-muted-foreground mt-1">
                              {link.description}
                            </p>
                          </div>
                          <Button
                            variant="ghost"
                            size="sm"
                            className="w-full"
                          >
                            进入
                            <ArrowRight className="ml-2 h-4 w-4" />
                          </Button>
                        </div>
                      </CardContent>
                    </Card>
                  </Link>
                );
              })}
            </div>
          </CardContent>
        </Card>

        {/* 最近活动 */}
        <div className="grid gap-4 md:grid-cols-2">
          <Card>
            <CardHeader>
              <CardTitle>最近活动</CardTitle>
              <CardDescription>系统最新动态</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="flex items-start gap-3">
                  <div className="h-2 w-2 rounded-full bg-blue-500 mt-2" />
                  <div className="flex-1">
                    <p className="text-sm font-medium">新用户注册</p>
                    <p className="text-xs text-muted-foreground">
                      张三 刚刚注册了账号
                    </p>
                    <p className="text-xs text-muted-foreground">5分钟前</p>
                  </div>
                </div>
                <div className="flex items-start gap-3">
                  <div className="h-2 w-2 rounded-full bg-green-500 mt-2" />
                  <div className="flex-1">
                    <p className="text-sm font-medium">权限更新</p>
                    <p className="text-xs text-muted-foreground">
                      管理员角色权限已更新
                    </p>
                    <p className="text-xs text-muted-foreground">1小时前</p>
                  </div>
                </div>
                <div className="flex items-start gap-3">
                  <div className="h-2 w-2 rounded-full bg-purple-500 mt-2" />
                  <div className="flex-1">
                    <p className="text-sm font-medium">系统维护</p>
                    <p className="text-xs text-muted-foreground">
                      系统将于今晚进行例行维护
                    </p>
                    <p className="text-xs text-muted-foreground">2小时前</p>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>系统公告</CardTitle>
              <CardDescription>重要通知</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="p-3 bg-blue-50 rounded-lg border border-blue-200">
                  <p className="text-sm font-medium text-blue-900">
                    系统升级通知
                  </p>
                  <p className="text-xs text-blue-700 mt-1">
                    系统将于本周末进行版本升级，届时将暂停服务2小时
                  </p>
                </div>
                <div className="p-3 bg-green-50 rounded-lg border border-green-200">
                  <p className="text-sm font-medium text-green-900">
                    新功能上线
                  </p>
                  <p className="text-xs text-green-700 mt-1">
                    菜单管理功能已上线，支持拖拽排序和权限关联
                  </p>
                </div>
                <div className="p-3 bg-orange-50 rounded-lg border border-orange-200">
                  <p className="text-sm font-medium text-orange-900">
                    安全提醒
                  </p>
                  <p className="text-xs text-orange-700 mt-1">
                    请定期修改密码，确保账号安全
                  </p>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
