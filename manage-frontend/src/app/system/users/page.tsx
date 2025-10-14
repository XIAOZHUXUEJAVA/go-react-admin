"use client";

import React, { useState } from "react";
import { useUsers } from "@/hooks/useUsers";
import { User } from "@/types/api";
import { DashboardHeader } from "@/components/layout/dashboard-header";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Users, Plus, RefreshCw } from "lucide-react";
import {
  UserManagementTable,
  AddUserModal,
  UserStatsCards,
  UserSearchFilter,
} from "@/components/user";
import { userApi } from "@/api";
import { CreateUserRequest, UpdateUserRequest } from "@/types/api";
import { toast } from "sonner";
import { getErrorMessage } from "@/lib/error-handler";
import { PagePermissionGuard, PermissionButton } from "@/components/auth";

/**
 * Dashboard 用户管理页面
 */
export default function UsersManagePage() {
  const { users, pagination, loading, error, fetchUsers, refetch } = useUsers({
    page: 1,
    pageSize: 10,
  });

  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [searchTerm, setSearchTerm] = useState("");
  const [statusFilter, setStatusFilter] = useState<string>("all");
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [isCreating, setIsCreating] = useState(false);
  const [isUpdating, setIsUpdating] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

  // 处理用户详情查看
  const handleViewUser = (user: User) => {
    setSelectedUser(user);
    setIsDetailModalOpen(true);
  };

  // 处理分页
  const handlePageChange = (page: number) => {
    fetchUsers({
      page,
      pageSize: pagination?.page_size || 10,
    });
  };

  // 处理每页大小变更
  const handlePageSizeChange = (pageSize: number) => {
    fetchUsers({
      page: 1,
      pageSize,
    });
  };

  // 过滤用户
  const filteredUsers =
    users?.filter((user) => {
      const matchesSearch =
        user.username.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.email.toLowerCase().includes(searchTerm.toLowerCase());
      const matchesStatus =
        statusFilter === "all" || user.status === statusFilter;
      return matchesSearch && matchesStatus;
    }) || [];

  // 处理创建用户
  const handleCreateUser = async (userData: CreateUserRequest) => {
    setIsCreating(true);
    try {
      const response = await userApi.createUser(userData);
      if (response.code === 201) {
        toast.success("用户创建成功");
        refetch(); // 刷新用户列表
      } else {
        toast.error(response.message || "创建用户失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "创建用户失败，请稍后重试"));
    } finally {
      setIsCreating(false);
    }
  };

  // 处理编辑用户
  const handleEditUser = async (user: User) => {
    setIsUpdating(true);
    try {
      const updateData: UpdateUserRequest = {
        id: user.id,
        username: user.username,
        email: user.email,
        role: user.role,
        status: user.status,
      };

      const response = await userApi.updateUser(user.id, updateData);
      if (response.code === 200) {
        toast.success("用户更新成功");
        refetch(); // 刷新用户列表
      } else {
        toast.error(response.message || "更新用户失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "更新用户失败，请稍后重试"));
    } finally {
      setIsUpdating(false);
    }
  };

  // 处理删除用户
  const handleDeleteUser = async (user: User) => {
    setIsDeleting(true);
    try {
      const response = await userApi.deleteUser(user.id);
      if (response.code === 200) {
        toast.success("用户删除成功");
        refetch(); // 刷新用户列表
      } else {
        toast.error(response.message || "删除用户失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "删除用户失败，请稍后重试"));
    } finally {
      setIsDeleting(false);
    }
  };

  // 面包屑导航配置
  const breadcrumbs = [
    { label: "Dashboard", href: "/dashboard" },
    { label: "用户管理" },
  ];

  // 头部操作按钮
  const headerActions = (
    <div className="flex items-center gap-2">
      <Button variant="outline" size="sm" onClick={refetch} disabled={loading}>
        <RefreshCw className={`h-4 w-4 ${loading ? "animate-spin" : ""}`} />
        刷新
      </Button>
      <PermissionButton 
        permission="user:create"
        size="sm" 
        onClick={() => setIsAddModalOpen(true)}
        noPermissionTooltip="您没有创建用户的权限"
      >
        <Plus className="h-4 w-4" />
        添加用户
      </PermissionButton>
    </div>
  );

  return (
    <PagePermissionGuard permission="user:read">
      <DashboardHeader breadcrumbs={breadcrumbs} actions={headerActions} />

      {/* 主要内容区域 */}
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        {/* 页面标题 */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">用户管理</h1>
            <p className="text-muted-foreground">管理系统中的所有用户信息</p>
          </div>
        </div>

        {/* 统计卡片 */}
        <UserStatsCards pagination={pagination} users={filteredUsers} />

        {/* 搜索和过滤 */}
        <UserSearchFilter
          searchTerm={searchTerm}
          onSearchChange={setSearchTerm}
          statusFilter={statusFilter}
          onStatusFilterChange={setStatusFilter}
        />

        {/* 用户表格 */}
        <Card>
          <CardHeader>
            <CardTitle>用户列表</CardTitle>
            <CardDescription>
              显示 {filteredUsers.length} 个用户中的结果
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
            ) : filteredUsers.length === 0 && !loading ? (
              <div className="text-center py-8">
                <Users className="mx-auto h-12 w-12 text-gray-400" />
                <h3 className="mt-2 text-sm font-medium text-gray-900">
                  暂无用户数据
                </h3>
                <p className="mt-1 text-sm text-gray-500">
                  没有找到匹配的用户信息
                </p>
              </div>
            ) : (
              <UserManagementTable
                users={filteredUsers}
                loading={loading || isUpdating || isDeleting}
                onView={handleViewUser}
                onEdit={handleEditUser}
                onDelete={handleDeleteUser}
              />
            )}
          </CardContent>
        </Card>

        {/* 分页 */}
        {pagination && filteredUsers.length > 0 && (
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

        {/* 添加用户模态框 */}
        <AddUserModal
          open={isAddModalOpen}
          onOpenChange={setIsAddModalOpen}
          onSubmit={handleCreateUser}
          loading={isCreating}
        />
      </div>
    </PagePermissionGuard>
  );
}
