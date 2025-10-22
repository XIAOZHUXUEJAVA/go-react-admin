"use client";

import React, { useState } from "react";
import { useRoles } from "@/hooks/useRoles";
import { Role, CreateRoleRequest, UpdateRoleRequest } from "@/types/role";
import { DashboardHeader } from "@/components/layout/dashboard-header";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Shield, Plus, RefreshCw } from "lucide-react";
import {
  RoleManagementTable,
  AddRoleModal,
  EditRoleModal,
  RoleStatsCards,
  AssignPermissionsModal,
} from "@/components/features/system/role";
import { roleApi } from "@/api";
import { toast } from "sonner";
import { getErrorMessage } from "@/lib/errorHandler";
import { PagePermissionGuard, PermissionButton } from "@/components/auth";

/**
 * Dashboard 角色管理页面
 */
export default function RolesManagePage() {
  const { roles, pagination, loading, error, fetchRoles, refetch } = useRoles({
    page: 1,
    page_size: 10,
  });

  const [selectedRole, setSelectedRole] = useState<Role | null>(null);
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isAssignPermissionsModalOpen, setIsAssignPermissionsModalOpen] =
    useState(false);
  const [isCreating, setIsCreating] = useState(false);
  const [isUpdating, setIsUpdating] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

  // 处理分页
  const handlePageChange = (page: number) => {
    fetchRoles({
      page,
      page_size: pagination?.page_size || 10,
    });
  };

  // 处理创建角色
  const handleCreateRole = async (roleData: CreateRoleRequest) => {
    setIsCreating(true);
    try {
      const response = await roleApi.createRole(roleData);
      if (response.code === 201) {
        toast.success("角色创建成功");
        refetch();
        setIsAddModalOpen(false);
      } else {
        toast.error(response.message || "创建角色失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "创建角色失败，请稍后重试"));
    } finally {
      setIsCreating(false);
    }
  };

  // 处理编辑角色
  const handleEditRole = (role: Role) => {
    setSelectedRole(role);
    setIsEditModalOpen(true);
  };

  // 处理更新角色
  const handleUpdateRole = async (id: number, data: UpdateRoleRequest) => {
    setIsUpdating(true);
    try {
      const response = await roleApi.updateRole(id, data);
      if (response.code === 200) {
        toast.success("角色更新成功");
        refetch();
        setIsEditModalOpen(false);
      } else {
        toast.error(response.message || "更新角色失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "更新角色失败，请稍后重试"));
    } finally {
      setIsUpdating(false);
    }
  };

  // 处理删除角色
  const handleDeleteRole = async (role: Role) => {
    setIsDeleting(true);
    try {
      const response = await roleApi.deleteRole(role.id);
      if (response.code === 200) {
        toast.success("角色删除成功");
        refetch();
      } else {
        toast.error(response.message || "删除角色失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "删除角色失败，请稍后重试"));
    } finally {
      setIsDeleting(false);
    }
  };

  // 处理分配权限
  const handleAssignPermissions = (role: Role) => {
    setSelectedRole(role);
    setIsAssignPermissionsModalOpen(true);
  };

  // 权限分配成功回调
  const handlePermissionsAssigned = () => {
    refetch();
  };

  // 面包屑导航配置
  const breadcrumbs = [
    { label: "Dashboard", href: "/dashboard" },
    { label: "系统管理" },
    { label: "角色管理" },
  ];

  // 头部操作按钮
  const headerActions = (
    <div className="flex items-center gap-2">
      <Button variant="outline" size="sm" onClick={refetch} disabled={loading}>
        <RefreshCw className={`h-4 w-4 ${loading ? "animate-spin" : ""}`} />
        刷新
      </Button>
      <PermissionButton
        permission="role:create"
        size="sm"
        onClick={() => setIsAddModalOpen(true)}
        noPermissionTooltip="您没有创建角色的权限"
      >
        <Plus className="h-4 w-4" />
        添加角色
      </PermissionButton>
    </div>
  );

  return (
    <PagePermissionGuard permission="role:read">
      <DashboardHeader breadcrumbs={breadcrumbs} actions={headerActions} />

      {/* 主要内容区域 */}
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        {/* 页面标题 */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">角色管理</h1>
            <p className="text-muted-foreground">管理系统中的所有角色和权限</p>
          </div>
        </div>

        {/* 统计卡片 */}
        <RoleStatsCards pagination={pagination} roles={roles} />

        {/* 角色表格 */}
        <Card>
          <CardHeader>
            <CardTitle>角色列表</CardTitle>
            <CardDescription>
              显示 {roles.length} 个角色
              {pagination && ` (共 ${pagination.total} 个)`}
            </CardDescription>
          </CardHeader>
          <CardContent>
            {error ? (
              <div className="text-center py-8">
                <p className="text-red-500">加载失败: {error.message}</p>
                <Button onClick={refetch} className="mt-2">
                  重试
                </Button>
              </div>
            ) : roles.length === 0 && !loading ? (
              <div className="text-center py-8">
                <Shield className="mx-auto h-12 w-12 text-gray-400" />
                <h3 className="mt-2 text-sm font-medium text-gray-900">
                  暂无角色数据
                </h3>
                <p className="mt-1 text-sm text-gray-500">开始创建第一个角色</p>
                <Button
                  onClick={() => setIsAddModalOpen(true)}
                  className="mt-4"
                >
                  <Plus className="h-4 w-4 mr-2" />
                  添加角色
                </Button>
              </div>
            ) : (
              <RoleManagementTable
                roles={roles}
                loading={loading || isUpdating || isDeleting}
                onEdit={handleEditRole}
                onDelete={handleDeleteRole}
                onAssignPermissions={handleAssignPermissions}
              />
            )}
          </CardContent>
        </Card>

        {/* 分页 */}
        {pagination && roles.length > 0 && (
          <Card>
            <CardContent className="pt-6">
              <div className="flex items-center justify-between">
                <div className="text-sm text-muted-foreground">
                  显示第 {(pagination.page - 1) * pagination.page_size + 1} -{" "}
                  {Math.min(
                    pagination.page * pagination.page_size,
                    pagination.total
                  )}{" "}
                  条，共 {pagination.total} 条记录
                </div>
                <div className="flex items-center gap-2">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => handlePageChange(pagination.page - 1)}
                    disabled={pagination.page <= 1}
                  >
                    上一页
                  </Button>
                  <div className="flex items-center gap-1">
                    {Array.from(
                      { length: Math.min(5, pagination.total_pages) },
                      (_, i) => {
                        const pageNum = i + 1;
                        return (
                          <Button
                            key={pageNum}
                            variant={
                              pagination.page === pageNum
                                ? "default"
                                : "outline"
                            }
                            size="sm"
                            onClick={() => handlePageChange(pageNum)}
                          >
                            {pageNum}
                          </Button>
                        );
                      }
                    )}
                  </div>
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => handlePageChange(pagination.page + 1)}
                    disabled={pagination.page >= pagination.total_pages}
                  >
                    下一页
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>
        )}

        {/* 添加角色模态框 */}
        <AddRoleModal
          open={isAddModalOpen}
          onOpenChange={setIsAddModalOpen}
          onSubmit={handleCreateRole}
          loading={isCreating}
        />

        {/* 编辑角色模态框 */}
        <EditRoleModal
          open={isEditModalOpen}
          onOpenChange={setIsEditModalOpen}
          role={selectedRole}
          onSubmit={handleUpdateRole}
          loading={isUpdating}
        />

        {/* 分配权限模态框 */}
        <AssignPermissionsModal
          open={isAssignPermissionsModalOpen}
          onOpenChange={setIsAssignPermissionsModalOpen}
          role={selectedRole}
          onSuccess={handlePermissionsAssigned}
        />
      </div>
    </PagePermissionGuard>
  );
}
