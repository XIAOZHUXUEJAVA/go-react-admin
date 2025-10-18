"use client";

import React, { useState } from "react";
import { Permission } from "@/types/permission";
import { Button } from "@/components/ui/button";
import { StatusBadge, PermissionTypeBadge } from "@/components/ui";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { MoreHorizontal, Edit, Trash2, Key } from "lucide-react";
import { formatDateTable } from "@/lib/date";

interface PermissionManagementTableProps {
  permissions: Permission[];
  loading?: boolean;
  onEdit?: (permission: Permission) => void;
  onDelete?: (permission: Permission) => void;
}

/**
 * 权限管理表格组件
 */
export const PermissionManagementTable: React.FC<
  PermissionManagementTableProps
> = ({ permissions, loading = false, onEdit, onDelete }) => {
  const [selectedPermission, setSelectedPermission] =
    useState<Permission | null>(null);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);

  // 获取类型标签
  const getTypeLabel = (type: string) => {
    const labels: Record<string, string> = {
      api: "API",
      menu: "菜单",
      button: "按钮",
    };
    return labels[type] || type;
  };

  // 处理编辑
  const handleEdit = (permission: Permission) => {
    setSelectedPermission(permission);
    onEdit?.(permission);
  };

  // 处理删除确认
  const handleDeleteClick = (permission: Permission) => {
    setSelectedPermission(permission);
    setIsDeleteDialogOpen(true);
  };

  // 确认删除
  const handleDeleteConfirm = () => {
    if (selectedPermission) {
      onDelete?.(selectedPermission);
      setIsDeleteDialogOpen(false);
      setSelectedPermission(null);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-8">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  return (
    <>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>权限名称</TableHead>
              <TableHead>权限代码</TableHead>
              <TableHead>资源</TableHead>
              <TableHead>操作</TableHead>
              <TableHead>类型</TableHead>
              <TableHead>路径/方法</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>创建时间</TableHead>
              <TableHead>操作</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {permissions.length === 0 ? (
              <TableRow>
                <TableCell colSpan={9} className="text-center py-8">
                  <div className="flex flex-col items-center justify-center text-muted-foreground">
                    <Key className="h-12 w-12 mb-2 opacity-50" />
                    <p>暂无权限数据</p>
                  </div>
                </TableCell>
              </TableRow>
            ) : (
              permissions.map((permission) => (
                <TableRow key={permission.id}>
                  <TableCell className="font-medium">
                    <div className="flex items-center gap-2">
                      <Key className="h-4 w-4 text-muted-foreground" />
                      {permission.name}
                    </div>
                  </TableCell>
                  <TableCell>
                    <code className="text-xs bg-muted px-2 py-1 rounded">
                      {permission.code}
                    </code>
                  </TableCell>
                  <TableCell>
                    <span className="capitalize">{permission.resource}</span>
                  </TableCell>
                  <TableCell>
                    <span className="capitalize">{permission.action}</span>
                  </TableCell>
                  <TableCell>
                    <PermissionTypeBadge type={permission.type as "api" | "menu" | "button"}>
                      {getTypeLabel(permission.type)}
                    </PermissionTypeBadge>
                  </TableCell>
                  <TableCell className="max-w-xs">
                    <div className="text-xs space-y-1">
                      {permission.path && (
                        <div className="truncate">
                          <span className="text-muted-foreground">路径: </span>
                          <code className="bg-muted px-1 rounded">
                            {permission.path}
                          </code>
                        </div>
                      )}
                      {permission.method && (
                        <div>
                          <span className="text-muted-foreground">方法: </span>
                          <span className="inline-flex items-center justify-center rounded-md border px-2 py-0.5 text-xs font-medium">
                            {permission.method}
                          </span>
                        </div>
                      )}
                    </div>
                  </TableCell>
                  <TableCell>
                    <StatusBadge status={permission.status as "active" | "inactive"}>
                      {permission.status === "active" ? "启用" : "禁用"}
                    </StatusBadge>
                  </TableCell>
                  <TableCell className="text-muted-foreground">
                    {formatDateTable(permission.created_at)}
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
                          onClick={() => handleEdit(permission)}
                        >
                          <Edit className="mr-2 h-4 w-4" />
                          编辑
                        </DropdownMenuItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem
                          onClick={() => handleDeleteClick(permission)}
                          className="text-red-600"
                        >
                          <Trash2 className="mr-2 h-4 w-4" />
                          删除
                        </DropdownMenuItem>
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
      <AlertDialog
        open={isDeleteDialogOpen}
        onOpenChange={setIsDeleteDialogOpen}
      >
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>确认删除权限</AlertDialogTitle>
            <AlertDialogDescription>
              您确定要删除权限 &quot;{selectedPermission?.name}&quot; 吗？
              此操作无法撤销，该权限的所有关联也将被移除。
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>取消</AlertDialogCancel>
            <AlertDialogAction
              onClick={handleDeleteConfirm}
              className="bg-red-600 hover:bg-red-700"
            >
              确认删除
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  );
};
