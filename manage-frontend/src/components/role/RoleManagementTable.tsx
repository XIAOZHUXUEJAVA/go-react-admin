"use client";

import React, { useState } from "react";
import { Role } from "@/types/role";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { StatusBadge } from "@/components/ui";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { DeleteConfirmDialog, TableEmptyState, LoadingState } from "@/components/common";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { MoreHorizontal, Edit, Trash2, Shield, Key } from "lucide-react";
import { formatDateTable } from "@/lib/date";

interface RoleManagementTableProps {
  roles: Role[];
  loading?: boolean;
  onEdit?: (role: Role) => void;
  onDelete?: (role: Role) => void;
  onAssignPermissions?: (role: Role) => void;
}

/**
 * 角色管理表格组件
 */
export const RoleManagementTable: React.FC<RoleManagementTableProps> = ({
  roles,
  loading = false,
  onEdit,
  onDelete,
  onAssignPermissions,
}) => {
  const [selectedRole, setSelectedRole] = useState<Role | null>(null);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);

  // 处理编辑
  const handleEdit = (role: Role) => {
    setSelectedRole(role);
    onEdit?.(role);
  };

  // 处理删除确认
  const handleDeleteClick = (role: Role) => {
    setSelectedRole(role);
    setIsDeleteDialogOpen(true);
  };

  // 确认删除
  const handleDeleteConfirm = () => {
    if (selectedRole) {
      onDelete?.(selectedRole);
      setIsDeleteDialogOpen(false);
      setSelectedRole(null);
    }
  };

  // 处理权限分配
  const handleAssignPermissions = (role: Role) => {
    onAssignPermissions?.(role);
  };

  if (loading) {
    return <LoadingState mode="spinner" />;
  }

  return (
    <>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>角色名称</TableHead>
              <TableHead>角色代码</TableHead>
              <TableHead>描述</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>类型</TableHead>
              <TableHead>创建时间</TableHead>
              <TableHead className="text-right">操作</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {roles.length === 0 ? (
              <TableEmptyState
                colSpan={7}
                icon={Shield}
                title="暂无角色数据"
                description="还没有创建任何角色，点击上方按钮添加新角色"
              />
            ) : (
              roles.map((role) => (
                <TableRow key={role.id}>
                  <TableCell className="font-medium">
                    <div className="flex items-center gap-2">
                      <Shield className="h-4 w-4 text-muted-foreground" />
                      {role.name}
                    </div>
                  </TableCell>
                  <TableCell>
                    <code className="text-xs bg-muted px-2 py-1 rounded">
                      {role.code}
                    </code>
                  </TableCell>
                  <TableCell className="max-w-xs truncate">
                    {role.description || "-"}
                  </TableCell>
                  <TableCell>
                    <StatusBadge status={role.status as "active" | "inactive"}>
                      {role.status === "active" ? "启用" : "禁用"}
                    </StatusBadge>
                  </TableCell>
                  <TableCell>
                    {role.is_system ? (
                      <Badge variant="outline" className="text-purple-600">
                        系统角色
                      </Badge>
                    ) : (
                      <Badge variant="outline">自定义</Badge>
                    )}
                  </TableCell>
                  <TableCell className="text-muted-foreground">
                    {formatDateTable(role.created_at)}
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
                        <DropdownMenuSeparator />
                        <DropdownMenuItem
                          onClick={() => handleAssignPermissions(role)}
                        >
                          <Key className="mr-2 h-4 w-4" />
                          分配权限
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={() => handleEdit(role)}>
                          <Edit className="mr-2 h-4 w-4" />
                          编辑
                        </DropdownMenuItem>
                        {!role.is_system && (
                          <>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem
                              onClick={() => handleDeleteClick(role)}
                              className="text-red-600"
                            >
                              <Trash2 className="mr-2 h-4 w-4" />
                              删除
                            </DropdownMenuItem>
                          </>
                        )}
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </div>

      {/* 删除确认对话框 */}
      <DeleteConfirmDialog
        open={isDeleteDialogOpen}
        onOpenChange={setIsDeleteDialogOpen}
        onConfirm={handleDeleteConfirm}
        resourceName={selectedRole?.name}
        resourceType="角色"
        title="确认删除角色"
        description={`您确定要删除角色 "${selectedRole?.name}" 吗？此操作无法撤销，该角色下的所有权限关联也将被移除。`}
        confirmText="确认删除"
      />
    </>
  );
};
