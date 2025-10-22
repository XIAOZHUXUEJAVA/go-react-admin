"use client";

import React from "react";
import { User } from "@/types/api";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Users, Calendar, UserCheck, UserX } from "lucide-react";

interface Pagination {
  page: number;
  page_size: number;
  total: number;
  total_pages: number;
}

interface UserStatsCardsProps {
  pagination: Pagination | null;
  users: User[];
}

export function UserStatsCards({ pagination, users }: UserStatsCardsProps) {
  if (!pagination) {
    return null;
  }

  const activeUsersCount = users.filter((u) => u.status === "active").length;

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">总用户数</CardTitle>
          <Users className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{pagination.total}</div>
          <p className="text-xs text-muted-foreground">系统中的所有用户</p>
        </CardContent>
      </Card>
      
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">当前页</CardTitle>
          <Calendar className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{pagination.page}</div>
          <p className="text-xs text-muted-foreground">
            共 {pagination.total_pages} 页
          </p>
        </CardContent>
      </Card>
      
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">活跃用户</CardTitle>
          <UserCheck className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{activeUsersCount}</div>
          <p className="text-xs text-muted-foreground">状态为活跃的用户</p>
        </CardContent>
      </Card>
      
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium">每页显示</CardTitle>
          <UserX className="h-4 w-4 text-muted-foreground" />
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{pagination.page_size}</div>
          <p className="text-xs text-muted-foreground">当前页面大小</p>
        </CardContent>
      </Card>
    </div>
  );
}