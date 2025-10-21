"use client";

import React, { useState } from "react";
import { User } from "@/types/api";
import { Button } from "@/components/ui/button";
import { StatusBadge, RoleBadge } from "@/components/ui";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { DeleteConfirmDialog, TableLoadingSkeleton } from "@/components/common";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { MoreHorizontal, Eye, Edit, Trash2, Mail } from "lucide-react";
import { UserDetailModal } from "./UserDetailModal";
import { formatDateTable } from "@/lib/date";
import { EditUserModal } from "./EditUserModal";
import { PermissionDropdownMenuItem } from "@/components/auth";

interface UserManagementTableProps {
  users: User[];
  loading?: boolean;
  onEdit?: (user: User) => void;
  onDelete?: (user: User) => void;
  onView?: (user: User) => void;
}

/**
 * 用户管理表格组件
 */
export const UserManagementTable: React.FC<UserManagementTableProps> = ({
  users,
  loading = false,
  onEdit,
  onDelete,
  onView,
}) => {
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);

  // 处理查看用户详情
  const handleViewUser = (user: User) => {
    setSelectedUser(user);
    setIsDetailModalOpen(true);
    onView?.(user);
  };

  // 处理编辑用户
  const handleEditUser = (user: User) => {
    setSelectedUser(user);
    setIsEditModalOpen(true);
  };

  // 处理删除用户
  const handleDeleteUser = (user: User) => {
    setSelectedUser(user);
    setIsDeleteDialogOpen(true);
  };

  // 保存编辑
  const handleSaveEdit = (updatedUser: User) => {
    onEdit?.(updatedUser);
    setSelectedUser(null);
  };

  // 确认删除用户
  const confirmDeleteUser = () => {
    if (selectedUser) {
      onDelete?.(selectedUser);
      setSelectedUser(null);
      setIsDeleteDialogOpen(false);
    }
  };

  // 加载状态
  if (loading) {
    return (
      <TableLoadingSkeleton
        rows={5}
        columns={[
          {
            header: "用户",
            skeleton: (
              <div className="flex items-center gap-3">
                <div className="h-8 w-8 bg-gray-200 rounded-full animate-pulse" />
                <div className="space-y-2">
                  <div className="h-4 w-24 bg-gray-200 rounded animate-pulse" />
                  <div className="h-3 w-16 bg-gray-200 rounded animate-pulse" />
                </div>
              </div>
            ),
          },
          { header: "邮箱", skeleton: "w-32" },
          { header: "角色", skeleton: "w-16" },
          { header: "状态", skeleton: "w-16" },
          { header: "创建时间", skeleton: "w-24" },
          {
            header: "操作",
            skeleton: (
              <div className="h-8 w-8 bg-gray-200 rounded animate-pulse ml-auto" />
            ),
          },
        ]}
      />
    );
  }

  return (
    <>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>用户</TableHead>
              <TableHead>邮箱</TableHead>
              <TableHead>角色</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>创建时间</TableHead>
              <TableHead className="text-right">操作</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {users.map((user) => (
              <TableRow key={user.id}>
                <TableCell className="font-medium">
                  <div className="flex items-center gap-3">
                    <Avatar className="h-8 w-8">
                      <AvatarFallback>
                        {user.username.charAt(0).toUpperCase()}
                      </AvatarFallback>
                    </Avatar>
                    <div>
                      <div className="font-medium">{user.username}</div>
                      <div className="text-sm text-muted-foreground">
                        ID: {user.id}
                      </div>
                    </div>
                  </div>
                </TableCell>
                <TableCell>
                  <div className="flex items-center gap-2">
                    <Mail className="h-4 w-4 text-muted-foreground" />
                    {user.email}
                  </div>
                </TableCell>
                <TableCell>
                  <RoleBadge role={user.role as "admin" | "moderator" | "user"}>
                    {user.role}
                  </RoleBadge>
                </TableCell>
                <TableCell>
                  <StatusBadge
                    status={user.status as "active" | "inactive" | "pending"}
                  >
                    {user.status}
                  </StatusBadge>
                </TableCell>
                <TableCell>
                  <div className="text-sm">
                    {formatDateTable(user.created_at)}
                  </div>
                </TableCell>
                <TableCell className="text-right">
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" className="h-8 w-8 p-0">
                        <span className="sr-only">打开菜单</span>
                        <MoreHorizontal className="h-4 w-4" />
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                      <DropdownMenuLabel>操作</DropdownMenuLabel>
                      <PermissionDropdownMenuItem
                        permission="user:read"
                        onClick={() => handleViewUser(user)}
                      >
                        <Eye className="mr-2 h-4 w-4" />
                        查看详情
                      </PermissionDropdownMenuItem>
                      <PermissionDropdownMenuItem
                        permission="user:update"
                        onClick={() => handleEditUser(user)}
                      >
                        <Edit className="mr-2 h-4 w-4" />
                        编辑用户
                      </PermissionDropdownMenuItem>
                      <DropdownMenuSeparator />
                      <PermissionDropdownMenuItem
                        permission="user:delete"
                        className="text-red-600"
                        onClick={() => handleDeleteUser(user)}
                        hideWhenNoPermission
                      >
                        <Trash2 className="mr-2 h-4 w-4" />
                        删除用户
                      </PermissionDropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      {/* 使用新的独立组件 */}
      <UserDetailModal
        user={selectedUser}
        open={isDetailModalOpen}
        onOpenChange={setIsDetailModalOpen}
      />

      <EditUserModal
        user={selectedUser}
        open={isEditModalOpen}
        onOpenChange={setIsEditModalOpen}
        onSave={handleSaveEdit}
      />

      {/* 删除确认对话框 */}
      <DeleteConfirmDialog
        open={isDeleteDialogOpen}
        onOpenChange={setIsDeleteDialogOpen}
        onConfirm={confirmDeleteUser}
        resourceName={selectedUser?.username}
        resourceType="用户"
        title="确认删除用户"
      />
    </>
  );
};
