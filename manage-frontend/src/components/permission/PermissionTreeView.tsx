"use client";

import React, { useState } from "react";
import { Permission, PermissionTree } from "@/types/permission";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible";
import { ChevronDown, ChevronRight, Folder, Key } from "lucide-react";
import { formatDateTable } from "@/lib/date";
import { DeleteConfirmDialog, EmptyState } from "@/components/common";

interface PermissionTreeViewProps {
  permissionTree: PermissionTree[];
  onEdit?: (permission: Permission) => void;
  onDelete?: (permission: Permission) => void;
}

/**
 * 权限树形视图组件
 */
export const PermissionTreeView: React.FC<PermissionTreeViewProps> = ({
  permissionTree,
  onEdit,
  onDelete,
}) => {
  const [expandedResources, setExpandedResources] = useState<Set<string>>(
    new Set(permissionTree.map((tree) => tree.resource))
  );
  const [selectedPermission, setSelectedPermission] = useState<Permission | null>(null);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);

  // 处理删除点击
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

  // 切换资源展开/折叠
  const toggleResource = (resource: string) => {
    setExpandedResources((prev) => {
      const newSet = new Set(prev);
      if (newSet.has(resource)) {
        newSet.delete(resource);
      } else {
        newSet.add(resource);
      }
      return newSet;
    });
  };

  // 获取类型颜色
  const getTypeColor = (type: string) => {
    switch (type.toLowerCase()) {
      case "api":
        return "bg-purple-100 text-purple-800";
      case "menu":
        return "bg-green-100 text-green-800";
      case "button":
        return "bg-orange-100 text-orange-800";
      default:
        return "bg-gray-100 text-gray-800";
    }
  };

  // 获取类型标签
  const getTypeLabel = (type: string) => {
    const labels: Record<string, string> = {
      api: "API",
      menu: "菜单",
      button: "按钮",
    };
    return labels[type] || type;
  };

  // 获取状态颜色
  const getStatusColor = (status: string) => {
    return status === "active"
      ? "bg-green-100 text-green-800"
      : "bg-gray-100 text-gray-800";
  };

  if (permissionTree.length === 0) {
    return (
      <EmptyState
        icon={Key}
        title="暂无权限数据"
        description="还没有创建任何权限，点击上方按钮添加新权限"
        className="py-12"
      />
    );
  }

  return (
    <div className="space-y-2">
      {permissionTree.map((tree) => {
        const isExpanded = expandedResources.has(tree.resource);

        return (
          <Collapsible
            key={tree.resource}
            open={isExpanded}
            onOpenChange={() => toggleResource(tree.resource)}
          >
            <div className="border rounded-lg">
              {/* 资源标题 */}
              <CollapsibleTrigger className="w-full">
                <div className="flex items-center justify-between p-4 hover:bg-muted/50 transition-colors">
                  <div className="flex items-center gap-3">
                    {isExpanded ? (
                      <ChevronDown className="h-5 w-5 text-muted-foreground" />
                    ) : (
                      <ChevronRight className="h-5 w-5 text-muted-foreground" />
                    )}
                    <Folder className="h-5 w-5 text-blue-600" />
                    <div className="text-left">
                      <h3 className="font-semibold capitalize">
                        {tree.resource}
                      </h3>
                      <p className="text-sm text-muted-foreground">
                        {tree.permissions.length} 个权限
                      </p>
                    </div>
                  </div>
                  <Badge variant="outline">{tree.permissions.length}</Badge>
                </div>
              </CollapsibleTrigger>

              {/* 权限列表 */}
              <CollapsibleContent>
                <div className="border-t">
                  {tree.permissions.map((permission, index) => (
                    <div
                      key={permission.id}
                      className={`p-4 hover:bg-muted/30 transition-colors ${
                        index !== tree.permissions.length - 1
                          ? "border-b"
                          : ""
                      }`}
                    >
                      <div className="flex items-start justify-between gap-4">
                        <div className="flex-1 space-y-2">
                          {/* 权限名称和代码 */}
                          <div className="flex items-center gap-2">
                            <Key className="h-4 w-4 text-muted-foreground" />
                            <span className="font-medium">
                              {permission.name}
                            </span>
                            <code className="text-xs bg-muted px-2 py-1 rounded">
                              {permission.code}
                            </code>
                          </div>

                          {/* 权限详情 */}
                          <div className="flex flex-wrap items-center gap-2 text-sm">
                            <Badge className={getTypeColor(permission.type)}>
                              {getTypeLabel(permission.type)}
                            </Badge>
                            <Badge className={getStatusColor(permission.status)}>
                              {permission.status === "active" ? "启用" : "禁用"}
                            </Badge>
                            <span className="text-muted-foreground">
                              操作: <span className="capitalize">{permission.action}</span>
                            </span>
                          </div>

                          {/* 路径和方法 */}
                          {(permission.path || permission.method) && (
                            <div className="text-xs space-y-1">
                              {permission.path && (
                                <div className="flex items-center gap-2">
                                  <span className="text-muted-foreground">
                                    路径:
                                  </span>
                                  <code className="bg-muted px-2 py-1 rounded">
                                    {permission.path}
                                  </code>
                                </div>
                              )}
                              {permission.method && (
                                <div className="flex items-center gap-2">
                                  <span className="text-muted-foreground">
                                    方法:
                                  </span>
                                  <Badge variant="outline" className="text-xs">
                                    {permission.method}
                                  </Badge>
                                </div>
                              )}
                            </div>
                          )}

                          {/* 描述 */}
                          {permission.description && (
                            <p className="text-sm text-muted-foreground">
                              {permission.description}
                            </p>
                          )}

                          {/* 创建时间 */}
                          <p className="text-xs text-muted-foreground">
                            创建于 {formatDateTable(permission.created_at)}
                          </p>
                        </div>

                        {/* 操作按钮 */}
                        <div className="flex items-center gap-2">
                          <Button
                            variant="outline"
                            size="sm"
                            onClick={() => onEdit?.(permission)}
                          >
                            编辑
                          </Button>
                          <Button
                            variant="outline"
                            size="sm"
                            onClick={() => handleDeleteClick(permission)}
                            className="text-red-600 hover:text-red-700"
                          >
                            删除
                          </Button>
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
              </CollapsibleContent>
            </div>
          </Collapsible>
        );
      })}

      {/* 删除确认对话框 */}
      <DeleteConfirmDialog
        open={isDeleteDialogOpen}
        onOpenChange={setIsDeleteDialogOpen}
        onConfirm={handleDeleteConfirm}
        resourceName={selectedPermission?.name}
        resourceType="权限"
        title="确认删除权限"
        description={`您确定要删除权限 "${selectedPermission?.name}" 吗？此操作无法撤销，该权限的所有关联也将被移除。`}
        confirmText="确认删除"
      />
    </div>
  );
};
