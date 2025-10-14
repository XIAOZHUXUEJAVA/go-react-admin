"use client";

import { useAuthStore } from "@/stores/authStore";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { LogOut } from "lucide-react";
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
  const breadcrumbs = [
    { label: "管理系统", href: "/dashboard" },
    { label: "Dashboard" },
  ];

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
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        <div className="grid auto-rows-min gap-4 md:grid-cols-3">
          <div className="bg-muted/50 aspect-video rounded-xl" />
          <div className="bg-muted/50 aspect-video rounded-xl" />
          <div className="bg-muted/50 aspect-video rounded-xl" />
        </div>
        <div className="bg-muted/50 min-h-[100vh] flex-1 rounded-xl md:min-h-min" />
      </div>
    </>
  );
}
