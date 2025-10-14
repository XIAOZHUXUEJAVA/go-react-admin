"use client";

import React, { useState } from "react";
import { useAllPermissions, usePermissionTree } from "@/hooks/usePermissions";
import {
  Permission,
  CreatePermissionRequest,
  UpdatePermissionRequest,
} from "@/types/permission";
import { DashboardHeader } from "@/components/layout/dashboard-header";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Key, Plus, RefreshCw, Table, TreePine } from "lucide-react";
import {
  PermissionManagementTable,
  PermissionTreeView,
  AddPermissionModal,
  EditPermissionModal,
  PermissionStatsCards,
} from "@/components/permission";
import { permissionApi } from "@/api";
import { toast } from "sonner";
import { getErrorMessage } from "@/lib/error-handler";
import { PagePermissionGuard, PermissionButton } from "@/components/auth";

/**
 * Dashboard 权限管理页面
 */
export default function PermissionsManagePage() {
  const {
    permissions,
    loading: permissionsLoading,
    refetch: refetchPermissions,
  } = useAllPermissions();
  const {
    permissionTree,
    loading: treeLoading,
    refetch: refetchTree,
  } = usePermissionTree();

  const [selectedPermission, setSelectedPermission] =
    useState<Permission | null>(null);
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isCreating, setIsCreating] = useState(false);
  const [isUpdating, setIsUpdating] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const [activeTab, setActiveTab] = useState<"table" | "tree">("tree");

  const loading = permissionsLoading || treeLoading;

  // 刷新数据
  const handleRefresh = () => {
    refetchPermissions();
    refetchTree();
  };

  // 处理创建权限
  const handleCreatePermission = async (data: CreatePermissionRequest) => {
    setIsCreating(true);
    try {
      const response = await permissionApi.createPermission(data);
      if (response.code === 201) {
        toast.success("权限创建成功");
        handleRefresh();
        setIsAddModalOpen(false);
      } else {
        toast.error(response.message || "创建权限失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "创建权限失败，请稍后重试"));
    } finally {
      setIsCreating(false);
    }
  };

  // 处理编辑权限
  const handleEditPermission = (permission: Permission) => {
    setSelectedPermission(permission);
    setIsEditModalOpen(true);
  };

  // 处理更新权限
  const handleUpdatePermission = async (
    id: number,
    data: UpdatePermissionRequest
  ) => {
    setIsUpdating(true);
    try {
      const response = await permissionApi.updatePermission(id, data);
      if (response.code === 200) {
        toast.success("权限更新成功");
        handleRefresh();
        setIsEditModalOpen(false);
      } else {
        toast.error(response.message || "更新权限失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "更新权限失败，请稍后重试"));
    } finally {
      setIsUpdating(false);
    }
  };

  // 处理删除权限
  const handleDeletePermission = async (permission: Permission) => {
    setIsDeleting(true);
    try {
      const response = await permissionApi.deletePermission(permission.id);
      if (response.code === 200) {
        toast.success("权限删除成功");
        handleRefresh();
      } else {
        toast.error(response.message || "删除权限失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "删除权限失败，请稍后重试"));
    } finally {
      setIsDeleting(false);
    }
  };

  // 面包屑导航配置
  const breadcrumbs = [
    { label: "Dashboard", href: "/dashboard" },
    { label: "权限管理" },
  ];

  // 头部操作按钮
  const headerActions = (
    <div className="flex items-center gap-2">
      <Button
        variant="outline"
        size="sm"
        onClick={handleRefresh}
        disabled={permissionsLoading || treeLoading}
      >
        <RefreshCw
          className={`h-4 w-4 ${
            permissionsLoading || treeLoading ? "animate-spin" : ""
          }`}
        />
        刷新
      </Button>
      <PermissionButton
        permission="permission:create"
        size="sm"
        onClick={() => setIsAddModalOpen(true)}
        noPermissionTooltip="您没有创建权限的权限"
      >
        <Plus className="h-4 w-4" />
        添加权限
      </PermissionButton>
    </div>
  );

  return (
    <PagePermissionGuard permission="permission:read">
      <DashboardHeader breadcrumbs={breadcrumbs} actions={headerActions} />

      {/* 主要内容区域 */}
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        {/* 页面标题 */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">权限管理</h1>
            <p className="text-muted-foreground">管理系统中的所有权限配置</p>
          </div>
        </div>

        {/* 统计卡片 */}
        <PermissionStatsCards permissions={permissions} />

        {/* 权限列表/树形视图 */}
        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <div>
                <CardTitle>权限列表</CardTitle>
                <CardDescription>
                  显示 {permissions.length} 个权限
                </CardDescription>
              </div>
              <Tabs
                value={activeTab}
                onValueChange={(value) =>
                  setActiveTab(value as "table" | "tree")
                }
              >
                <TabsList>
                  <TabsTrigger value="tree" className="gap-2">
                    <TreePine className="h-4 w-4" />
                    树形视图
                  </TabsTrigger>
                  <TabsTrigger value="table" className="gap-2">
                    <Table className="h-4 w-4" />
                    表格视图
                  </TabsTrigger>
                </TabsList>
              </Tabs>
            </div>
          </CardHeader>
          <CardContent>
            {loading && permissions.length === 0 ? (
              <div className="flex items-center justify-center py-12">
                <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
              </div>
            ) : permissions.length === 0 ? (
              <div className="text-center py-12">
                <Key className="mx-auto h-12 w-12 text-gray-400" />
                <h3 className="mt-2 text-sm font-medium text-gray-900">
                  暂无权限数据
                </h3>
                <p className="mt-1 text-sm text-gray-500">开始创建第一个权限</p>
                <Button
                  onClick={() => setIsAddModalOpen(true)}
                  className="mt-4"
                >
                  <Plus className="h-4 w-4 mr-2" />
                  添加权限
                </Button>
              </div>
            ) : (
              <Tabs value={activeTab} className="w-full">
                <TabsContent value="tree" className="mt-0">
                  <PermissionTreeView
                    permissionTree={permissionTree}
                    onEdit={handleEditPermission}
                    onDelete={handleDeletePermission}
                  />
                </TabsContent>
                <TabsContent value="table" className="mt-0">
                  <PermissionManagementTable
                    permissions={permissions}
                    loading={isUpdating || isDeleting}
                    onEdit={handleEditPermission}
                    onDelete={handleDeletePermission}
                  />
                </TabsContent>
              </Tabs>
            )}
          </CardContent>
        </Card>

        {/* 添加权限模态框 */}
        <AddPermissionModal
          open={isAddModalOpen}
          onOpenChange={setIsAddModalOpen}
          onSubmit={handleCreatePermission}
          loading={isCreating}
        />

        {/* 编辑权限模态框 */}
        <EditPermissionModal
          open={isEditModalOpen}
          onOpenChange={setIsEditModalOpen}
          permission={selectedPermission}
          onSubmit={handleUpdatePermission}
          loading={isUpdating}
        />
      </div>
    </PagePermissionGuard>
  );
}
