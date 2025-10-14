"use client";

import React, { useState, useEffect, useMemo } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Loader2, Search, CheckCircle2, Circle } from "lucide-react";
import { Role } from "@/types/role";
import { Permission } from "@/types/permission";
import { useAllPermissions } from "@/hooks/usePermissions";
import { roleApi } from "@/api";
import { toast } from "sonner";
import { getErrorMessage } from "@/lib/error-handler";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

interface AssignPermissionsModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  role: Role | null;
  onSuccess?: () => void;
}

/**
 * 权限分配对话框组件
 */
export const AssignPermissionsModal: React.FC<AssignPermissionsModalProps> = ({
  open,
  onOpenChange,
  role,
  onSuccess,
}) => {
  const { permissions, loading: permissionsLoading } = useAllPermissions();
  const [selectedPermissionIds, setSelectedPermissionIds] = useState<number[]>(
    []
  );
  const [searchTerm, setSearchTerm] = useState("");
  const [loading, setLoading] = useState(false);
  const [loadingRolePermissions, setLoadingRolePermissions] = useState(false);

  // 按类型分组权限
  const permissionsByType = useMemo(() => {
    const grouped: Record<string, Permission[]> = {};
    permissions.forEach((permission) => {
      const type = permission.type || "other";
      if (!grouped[type]) {
        grouped[type] = [];
      }
      grouped[type].push(permission);
    });
    return grouped;
  }, [permissions]);

  // 按资源分组权限
  const permissionsByResource = useMemo(() => {
    const grouped: Record<string, Permission[]> = {};
    permissions.forEach((permission) => {
      const resource = permission.resource || "other";
      if (!grouped[resource]) {
        grouped[resource] = [];
      }
      grouped[resource].push(permission);
    });
    return grouped;
  }, [permissions]);

  // 过滤权限
  const filteredPermissions = useMemo(() => {
    if (!searchTerm) return permissions;
    const term = searchTerm.toLowerCase();
    return permissions.filter(
      (p) =>
        p.name.toLowerCase().includes(term) ||
        p.code.toLowerCase().includes(term) ||
        p.resource.toLowerCase().includes(term)
    );
  }, [permissions, searchTerm]);

  // 加载角色的权限
  useEffect(() => {
    const loadRolePermissions = async () => {
      if (!role || !open) return;

      setLoadingRolePermissions(true);
      try {
        const response = await roleApi.getRolePermissions(role.id);
        if (response.code === 200 && response.data) {
          // 处理 permissions 可能为 null 的情况
          const permissions = response.data.permissions || [];
          const permIds = permissions.map((p) => p.id);
          setSelectedPermissionIds(permIds);
        }
      } catch (error) {
        toast.error(getErrorMessage(error, "加载角色权限失败"));
      } finally {
        setLoadingRolePermissions(false);
      }
    };

    loadRolePermissions();
  }, [role, open]);

  // 处理权限选择
  const handlePermissionToggle = (permissionId: number) => {
    setSelectedPermissionIds((prev) =>
      prev.includes(permissionId)
        ? prev.filter((id) => id !== permissionId)
        : [...prev, permissionId]
    );
  };

  // 处理全选/取消全选（按资源）
  const handleResourceToggle = (resource: string) => {
    const resourcePermissions = permissionsByResource[resource] || [];
    const resourcePermissionIds = resourcePermissions.map((p) => p.id);
    const allSelected = resourcePermissionIds.every((id) =>
      selectedPermissionIds.includes(id)
    );

    if (allSelected) {
      // 取消全选
      setSelectedPermissionIds((prev) =>
        prev.filter((id) => !resourcePermissionIds.includes(id))
      );
    } else {
      // 全选
      setSelectedPermissionIds((prev) => [
        ...new Set([...prev, ...resourcePermissionIds]),
      ]);
    }
  };

  // 处理提交
  const handleSubmit = async () => {
    if (!role) return;

    setLoading(true);
    try {
      const response = await roleApi.assignPermissions(role.id, {
        permission_ids: selectedPermissionIds,
      });

      if (response.code === 200) {
        toast.success("权限分配成功");
        onSuccess?.();
        onOpenChange(false);
      } else {
        toast.error(response.message || "权限分配失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "权限分配失败"));
    } finally {
      setLoading(false);
    }
  };

  // 处理对话框关闭
  const handleOpenChange = (newOpen: boolean) => {
    if (!newOpen && !loading) {
      setSearchTerm("");
    }
    onOpenChange(newOpen);
  };

  // 获取类型标签
  const getTypeLabel = (type: string) => {
    const labels: Record<string, string> = {
      api: "API",
      menu: "菜单",
      button: "按钮",
      other: "其他",
    };
    return labels[type] || type;
  };

  // 获取类型颜色
  const getTypeColor = (type: string) => {
    const colors: Record<string, string> = {
      api: "bg-blue-100 text-blue-800",
      menu: "bg-green-100 text-green-800",
      button: "bg-purple-100 text-purple-800",
      other: "bg-gray-100 text-gray-800",
    };
    return colors[type] || colors.other;
  };

  return (
    <Dialog open={open} onOpenChange={handleOpenChange}>
      <DialogContent className="sm:max-w-[700px] max-h-[80vh]">
        <DialogHeader>
          <DialogTitle>分配权限</DialogTitle>
          <DialogDescription>
            为角色 &quot;{role?.name}&quot; 分配权限
          </DialogDescription>
        </DialogHeader>

        {/* 搜索框 */}
        <div className="relative">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="搜索权限名称、代码或资源..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="pl-9"
          />
        </div>

        {/* 统计信息 */}
        <div className="flex items-center gap-4 text-sm text-muted-foreground">
          <span>
            已选择: <strong>{selectedPermissionIds.length}</strong> 个权限
          </span>
          <span>
            总计: <strong>{permissions.length}</strong> 个权限
          </span>
        </div>

        {permissionsLoading || loadingRolePermissions ? (
          <div className="flex items-center justify-center py-12">
            <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
          </div>
        ) : (
          <Tabs defaultValue="resource" className="w-full">
            <TabsList className="grid w-full grid-cols-2">
              <TabsTrigger value="resource">按资源分组</TabsTrigger>
              <TabsTrigger value="type">按类型分组</TabsTrigger>
            </TabsList>

            {/* 按资源分组 */}
            <TabsContent value="resource" className="mt-4">
              <ScrollArea className="h-[400px] pr-4">
                <div className="space-y-4">
                  {Object.entries(permissionsByResource).map(
                    ([resource, resourcePermissions]) => {
                      const allSelected = resourcePermissions.every((p) =>
                        selectedPermissionIds.includes(p.id)
                      );
                      const someSelected = resourcePermissions.some((p) =>
                        selectedPermissionIds.includes(p.id)
                      );

                      return (
                        <div
                          key={resource}
                          className="border rounded-lg p-4 space-y-3"
                        >
                          {/* 资源标题 */}
                          <div className="flex items-center justify-between">
                            <div className="flex items-center gap-2">
                              <button
                                onClick={() => handleResourceToggle(resource)}
                                className="flex items-center gap-2 hover:text-primary transition-colors"
                              >
                                {allSelected ? (
                                  <CheckCircle2 className="h-5 w-5 text-primary" />
                                ) : someSelected ? (
                                  <Circle className="h-5 w-5 text-primary fill-primary/20" />
                                ) : (
                                  <Circle className="h-5 w-5" />
                                )}
                                <h4 className="font-semibold capitalize">
                                  {resource}
                                </h4>
                              </button>
                              <Badge variant="outline">
                                {resourcePermissions.length}
                              </Badge>
                            </div>
                          </div>

                          {/* 权限列表 */}
                          <div className="grid grid-cols-1 gap-2 pl-7">
                            {resourcePermissions.map((permission) => (
                              <div
                                key={permission.id}
                                className="flex items-start space-x-3 p-2 rounded hover:bg-muted/50 transition-colors"
                              >
                                <Checkbox
                                  id={`permission-${permission.id}`}
                                  checked={selectedPermissionIds.includes(
                                    permission.id
                                  )}
                                  onCheckedChange={() =>
                                    handlePermissionToggle(permission.id)
                                  }
                                  className="mt-1"
                                />
                                <label
                                  htmlFor={`permission-${permission.id}`}
                                  className="flex-1 cursor-pointer"
                                >
                                  <div className="flex items-center gap-2">
                                    <span className="font-medium">
                                      {permission.name}
                                    </span>
                                    <Badge
                                      className={getTypeColor(permission.type)}
                                    >
                                      {getTypeLabel(permission.type)}
                                    </Badge>
                                  </div>
                                  <div className="text-xs text-muted-foreground mt-1">
                                    <code className="bg-muted px-1 py-0.5 rounded">
                                      {permission.code}
                                    </code>
                                    {permission.description && (
                                      <span className="ml-2">
                                        {permission.description}
                                      </span>
                                    )}
                                  </div>
                                </label>
                              </div>
                            ))}
                          </div>
                        </div>
                      );
                    }
                  )}
                </div>
              </ScrollArea>
            </TabsContent>

            {/* 按类型分组 */}
            <TabsContent value="type" className="mt-4">
              <ScrollArea className="h-[400px] pr-4">
                <div className="space-y-4">
                  {Object.entries(permissionsByType).map(
                    ([type, typePermissions]) => (
                      <div
                        key={type}
                        className="border rounded-lg p-4 space-y-3"
                      >
                        <div className="flex items-center gap-2">
                          <Badge className={getTypeColor(type)}>
                            {getTypeLabel(type)}
                          </Badge>
                          <Badge variant="outline">
                            {typePermissions.length}
                          </Badge>
                        </div>

                        <div className="grid grid-cols-1 gap-2">
                          {typePermissions.map((permission) => (
                            <div
                              key={permission.id}
                              className="flex items-start space-x-3 p-2 rounded hover:bg-muted/50 transition-colors"
                            >
                              <Checkbox
                                id={`permission-type-${permission.id}`}
                                checked={selectedPermissionIds.includes(
                                  permission.id
                                )}
                                onCheckedChange={() =>
                                  handlePermissionToggle(permission.id)
                                }
                                className="mt-1"
                              />
                              <label
                                htmlFor={`permission-type-${permission.id}`}
                                className="flex-1 cursor-pointer"
                              >
                                <div className="font-medium">
                                  {permission.name}
                                </div>
                                <div className="text-xs text-muted-foreground mt-1">
                                  <code className="bg-muted px-1 py-0.5 rounded">
                                    {permission.code}
                                  </code>
                                  <span className="mx-2">•</span>
                                  <span className="capitalize">
                                    {permission.resource}
                                  </span>
                                </div>
                              </label>
                            </div>
                          ))}
                        </div>
                      </div>
                    )
                  )}
                </div>
              </ScrollArea>
            </TabsContent>
          </Tabs>
        )}

        <DialogFooter>
          <Button
            type="button"
            variant="outline"
            onClick={() => handleOpenChange(false)}
            disabled={loading}
          >
            取消
          </Button>
          <Button onClick={handleSubmit} disabled={loading}>
            {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            {loading ? "保存中..." : "保存权限"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};
