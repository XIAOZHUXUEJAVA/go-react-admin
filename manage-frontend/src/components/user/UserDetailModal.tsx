"use client";

import React from "react";
import { User } from "@/types/api";
import { Badge } from "@/components/ui/badge";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Label } from "@/components/ui/label";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Calendar, Mail, Shield, User as UserIcon } from "lucide-react";
import { formatDateDetail } from "@/lib/date";

interface UserDetailModalProps {
  user: User | null;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function UserDetailModal({
  user,
  open,
  onOpenChange,
}: UserDetailModalProps) {
  // 获取状态颜色
  const getStatusColor = (status: string) => {
    switch (status.toLowerCase()) {
      case "active":
        return "bg-green-100 text-green-800 hover:bg-green-200";
      case "inactive":
        return "bg-gray-100 text-gray-800 hover:bg-gray-200";
      case "pending":
        return "bg-yellow-100 text-yellow-800 hover:bg-yellow-200";
      default:
        return "bg-gray-100 text-gray-800 hover:bg-gray-200";
    }
  };

  // 获取角色颜色
  const getRoleColor = (role: string) => {
    switch (role.toLowerCase()) {
      case "admin":
        return "bg-red-100 text-red-800 hover:bg-red-200";
      case "moderator":
        return "bg-blue-100 text-blue-800 hover:bg-blue-200";
      case "user":
        return "bg-green-100 text-green-800 hover:bg-green-200";
      default:
        return "bg-gray-100 text-gray-800 hover:bg-gray-200";
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>用户详情</DialogTitle>
          <DialogDescription>查看用户的详细信息</DialogDescription>
        </DialogHeader>
        {user && (
          <div className="grid gap-4">
            <div className="flex items-center gap-4">
              <Avatar className="h-16 w-16">
                <AvatarFallback className="text-lg">
                  {user.username.charAt(0).toUpperCase()}
                </AvatarFallback>
              </Avatar>
              <div>
                <h3 className="text-lg font-semibold">{user.username}</h3>
                <p className="text-muted-foreground">{user.email}</p>
                <div className="flex items-center gap-2 mt-2">
                  <Badge className={getRoleColor(user.role)}>
                    <Shield className="mr-1 h-3 w-3" />
                    {user.role}
                  </Badge>
                  <Badge className={getStatusColor(user.status)}>
                    {user.status}
                  </Badge>
                </div>
              </div>
            </div>
            <div className="grid gap-4 pt-4 border-t">
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label className="text-sm font-medium flex items-center gap-2">
                    <UserIcon className="h-4 w-4" />
                    用户ID
                  </Label>
                  <p className="text-sm text-muted-foreground">{user.id}</p>
                </div>
                <div className="space-y-2">
                  <Label className="text-sm font-medium flex items-center gap-2">
                    <Mail className="h-4 w-4" />
                    邮箱地址
                  </Label>
                  <p className="text-sm text-muted-foreground">{user.email}</p>
                </div>
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label className="text-sm font-medium flex items-center gap-2">
                    <Calendar className="h-4 w-4" />
                    创建时间
                  </Label>
                  <p className="text-sm text-muted-foreground">
                    {formatDateDetail(user.created_at)}
                  </p>
                </div>
                <div className="space-y-2">
                  <Label className="text-sm font-medium flex items-center gap-2">
                    <Calendar className="h-4 w-4" />
                    更新时间
                  </Label>
                  <p className="text-sm text-muted-foreground">
                    {formatDateDetail(user.updated_at)}
                  </p>
                </div>
              </div>
            </div>
          </div>
        )}
      </DialogContent>
    </Dialog>
  );
}
