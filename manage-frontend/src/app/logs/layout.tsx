"use client";

import { AppSidebar } from "@/components/layout/app-sidebar";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { useEffect, useState } from "react";

interface LogsLayoutProps {
  children: React.ReactNode;
}

/**
 * Logs 布局组件
 * 为日志管理页面提供统一的侧边栏和布局结构
 */
export default function LogsLayout({ children }: LogsLayoutProps) {
  const [defaultOpen, setDefaultOpen] = useState(true);

  useEffect(() => {
    // 从 cookie 中读取侧边栏状态
    const cookies = document.cookie.split("; ");
    const sidebarCookie = cookies.find((c) => c.startsWith("sidebar_state="));
    if (sidebarCookie) {
      const value = sidebarCookie.split("=")[1];
      setDefaultOpen(value === "true");
    }
  }, []);

  return (
    <SidebarProvider defaultOpen={defaultOpen}>
      <AppSidebar />
      <SidebarInset>{children}</SidebarInset>
    </SidebarProvider>
  );
}
