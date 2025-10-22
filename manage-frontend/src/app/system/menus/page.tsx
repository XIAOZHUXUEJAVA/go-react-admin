"use client";

import React, { useState, useEffect } from "react";
import { Plus, RefreshCw, Menu as MenuIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { toast } from "sonner";
import { DashboardHeader } from "@/components/layout/dashboard-header";
import {
  MenuTreeTable,
  MenuEditDialog,
  MenuStatsCards,
} from "@/components/features/system/menu";
import { menuApi } from "@/api/menu";
import { type Menu } from "@/types/menu";
import { getErrorMessage } from "@/lib/errorHandler";
import { PagePermissionGuard, PermissionButton } from "@/components/auth";

export default function MenusManagePage() {
  const [menus, setMenus] = useState<Menu[]>([]);
  const [loading, setLoading] = useState(true);
  const [editDialogOpen, setEditDialogOpen] = useState(false);
  const [selectedMenu, setSelectedMenu] = useState<Menu | null>(null);
  const [parentMenu, setParentMenu] = useState<Menu | null>(null);

  // 面包屑导航
  const breadcrumbs = [
    { label: "Dashboard", href: "/dashboard" },
    { label: "系统管理" },
    { label: "菜单管理" },
  ];

  // 加载菜单数据
  const loadMenus = async () => {
    setLoading(true);
    try {
      const response = await menuApi.getMenuTree();
      if (response.code === 200) {
        setMenus(response.data || []);
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "加载菜单失败"));
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadMenus();
  }, []);

  // 处理创建菜单
  const handleCreate = (parent?: Menu) => {
    setSelectedMenu(null);
    setParentMenu(parent || null);
    setEditDialogOpen(true);
  };

  // 处理编辑菜单
  const handleEdit = (menu: Menu) => {
    setSelectedMenu(menu);
    setParentMenu(null);
    setEditDialogOpen(true);
  };

  // 处理删除菜单
  const handleDelete = async (menu: Menu) => {
    if (menu.children && menu.children.length > 0) {
      toast.error("该菜单存在子菜单，无法删除");
      return;
    }

    try {
      const response = await menuApi.deleteMenu(menu.id);
      if (response.code === 200) {
        toast.success("删除成功");
        loadMenus();
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "删除菜单失败"));
    }
  };

  // 处理保存
  const handleSave = () => {
    setEditDialogOpen(false);
    loadMenus();
  };

  // 处理顺序更新
  const handleOrderUpdate = async (
    updatedMenus: { id: number; order_num: number; parent_id: number | null }[]
  ) => {
    try {
      const response = await menuApi.updateMenuOrder(updatedMenus);
      if (response.code === 200) {
        toast.success("顺序更新成功");
        loadMenus();
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "更新顺序失败"));
    }
  };

  // 处理添加菜单
  const handleAddMenu = () => {
    handleCreate();
  };

  // 头部操作按钮
  const headerActions = (
    <div className="flex items-center gap-2">
      <Button
        variant="outline"
        size="sm"
        onClick={loadMenus}
        disabled={loading}
      >
        <RefreshCw className={`h-4 w-4 ${loading ? "animate-spin" : ""}`} />
        刷新
      </Button>
      <PermissionButton
        permission="menu:create"
        size="sm"
        onClick={handleAddMenu}
        noPermissionTooltip="您没有创建菜单的权限"
      >
        <Plus className="h-4 w-4" />
        添加菜单
      </PermissionButton>
    </div>
  );

  return (
    <PagePermissionGuard permission="menu:read">
      <DashboardHeader breadcrumbs={breadcrumbs} actions={headerActions} />

      {/* 主要内容区域 */}
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        {/* 页面标题 */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">菜单管理</h1>
            <p className="text-muted-foreground">管理系统中的所有菜单和导航</p>
          </div>
        </div>

        {/* 统计卡片 */}
        <MenuStatsCards menus={menus} />

        {/* 菜单表格 */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <MenuIcon className="h-5 w-5" />
              菜单列表
            </CardTitle>
            <CardDescription>
              拖拽菜单可调整顺序，点击操作按钮可编辑或删除菜单
            </CardDescription>
          </CardHeader>
          <CardContent>
            <MenuTreeTable
              menus={menus}
              loading={loading}
              onEdit={handleEdit}
              onDelete={handleDelete}
              onCreate={handleCreate}
              onOrderUpdate={handleOrderUpdate}
            />
          </CardContent>
        </Card>

        {/* 编辑对话框 */}
        <MenuEditDialog
          open={editDialogOpen}
          onOpenChange={setEditDialogOpen}
          menu={selectedMenu}
          parentMenu={parentMenu}
          onSave={handleSave}
        />
      </div>
    </PagePermissionGuard>
  );
}
