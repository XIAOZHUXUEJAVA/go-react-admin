"use client";

import React, { useState, useEffect } from "react";
import { toast } from "sonner";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Switch } from "@/components/ui/switch";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { menuApi } from "@/api/menu";
import {
  type Menu,
  type CreateMenuRequest,
  type UpdateMenuRequest,
} from "@/types/menu";
import { permissionApi } from "@/api/permission";
import type { Permission } from "@/types/permission";
import { IconPicker } from "./IconPicker";

interface MenuEditDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  menu: Menu | null;
  parentMenu: Menu | null;
  onSave: () => void;
}

export function MenuEditDialog({
  open,
  onOpenChange,
  menu,
  parentMenu,
  onSave,
}: MenuEditDialogProps) {
  const [loading, setLoading] = useState(false);
  const [permissions, setPermissions] = useState<Permission[]>([]);
  const [formData, setFormData] = useState({
    parent_id: null as number | null,
    name: "",
    title: "",
    path: "",
    component: "",
    icon: "FileText",
    order_num: 0,
    type: "menu",
    permission_code: "",
    visible: true,
    status: "active",
  });

  // 加载权限列表
  useEffect(() => {
    const loadPermissions = async () => {
      try {
        const response = await permissionApi.getAllPermissions();
        if (response.code === 200 && response.data) {
          setPermissions(response.data || []);
        }
      } catch (error) {}
    };

    if (open) {
      loadPermissions();
    }
  }, [open]);

  // 初始化表单数据
  useEffect(() => {
    if (menu) {
      // 编辑模式
      setFormData({
        parent_id: menu.parent_id,
        name: menu.name,
        title: menu.title,
        path: menu.path,
        component: menu.component || "",
        icon: menu.icon || "FileText",
        order_num: menu.order_num,
        type: menu.type,
        permission_code: menu.permission_code || "",
        visible: menu.visible,
        status: menu.status,
      });
    } else if (parentMenu) {
      // 创建子菜单
      setFormData({
        parent_id: parentMenu.id,
        name: "",
        title: "",
        path: "",
        component: "",
        icon: "FileText",
        order_num: 0,
        type: "menu",
        permission_code: "",
        visible: true,
        status: "active",
      });
    } else {
      // 创建根菜单
      setFormData({
        parent_id: null,
        name: "",
        title: "",
        path: "",
        component: "",
        icon: "FileText",
        order_num: 0,
        type: "menu",
        permission_code: "",
        visible: true,
        status: "active",
      });
    }
  }, [menu, parentMenu, open]);

  // 处理提交
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // 验证
    if (!formData.title.trim()) {
      toast.error("请输入菜单标题");
      return;
    }
    if (!formData.path.trim()) {
      toast.error("请输入菜单路径");
      return;
    }

    setLoading(true);
    try {
      if (menu) {
        // 更新
        const updateData: UpdateMenuRequest = {
          ...formData,
          parent_id: formData.parent_id,
        };
        const response = await menuApi.updateMenu(menu.id, updateData);
        if (response.code === 200) {
          toast.success("更新成功");
          onSave();
        }
      } else {
        // 创建
        const createData: CreateMenuRequest = {
          ...formData,
          parent_id: formData.parent_id,
        };
        const response = await menuApi.createMenu(createData);
        if (response.code === 201 || response.code === 200) {
          toast.success("创建成功");
          onSave();
        }
      }
    } catch (error) {
      toast.error(error instanceof Error ? error.message : "保存失败");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-2xl max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>
            {menu ? "编辑菜单" : parentMenu ? "添加子菜单" : "新建菜单"}
          </DialogTitle>
          <DialogDescription>
            {parentMenu && `父菜单: ${parentMenu.title}`}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit}>
          <Tabs defaultValue="basic" className="w-full">
            <TabsList className="grid w-full grid-cols-2">
              <TabsTrigger value="basic">基本信息</TabsTrigger>
              <TabsTrigger value="advanced">高级设置</TabsTrigger>
            </TabsList>

            <TabsContent value="basic" className="space-y-4 mt-4">
              {/* 菜单名称 */}
              <div className="space-y-2">
                <Label htmlFor="name">
                  菜单名称
                </Label>
                <Input
                  id="name"
                  value={formData.name}
                  onChange={(e) =>
                    setFormData({ ...formData, name: e.target.value })
                  }
                  placeholder="例如: user-management"
                />
              </div>

              {/* 菜单标题 */}
              <div className="space-y-2">
                <Label htmlFor="title">
                  菜单标题 <span className="text-red-500">*</span>
                </Label>
                <Input
                  id="title"
                  value={formData.title}
                  onChange={(e) =>
                    setFormData({ ...formData, title: e.target.value })
                  }
                  placeholder="例如: 用户管理"
                />
              </div>

              {/* 菜单路径 */}
              <div className="space-y-2">
                <Label htmlFor="path">
                  菜单路径 <span className="text-red-500">*</span>
                </Label>
                <Input
                  id="path"
                  value={formData.path}
                  onChange={(e) =>
                    setFormData({ ...formData, path: e.target.value })
                  }
                  placeholder="例如: /system/users"
                />
              </div>

              {/* 图标选择 */}
              <div className="space-y-2">
                <Label htmlFor="icon">菜单图标</Label>
                <IconPicker
                  value={formData.icon}
                  onChange={(icon: string) =>
                    setFormData({ ...formData, icon })
                  }
                />
              </div>

              {/* 菜单类型 */}
              <div className="space-y-2">
                <Label htmlFor="type">菜单类型</Label>
                <Select
                  value={formData.type}
                  onValueChange={(value) =>
                    setFormData({ ...formData, type: value })
                  }
                >
                  <SelectTrigger>
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="menu">菜单</SelectItem>
                    <SelectItem value="button">按钮</SelectItem>
                  </SelectContent>
                </Select>
              </div>

              {/* 排序 */}
              <div className="space-y-2">
                <Label htmlFor="order_num">排序</Label>
                <Input
                  id="order_num"
                  type="number"
                  value={formData.order_num}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      order_num: parseInt(e.target.value) || 0,
                    })
                  }
                />
                <p className="text-xs text-muted-foreground">数字越小越靠前</p>
              </div>
            </TabsContent>

            <TabsContent value="advanced" className="space-y-4 mt-4">
              {/* 组件路径 */}
              <div className="space-y-2">
                <Label htmlFor="component">组件路径</Label>
                <Input
                  id="component"
                  value={formData.component}
                  onChange={(e) =>
                    setFormData({ ...formData, component: e.target.value })
                  }
                  placeholder="例如: @/app/dashboard/welcome/page"
                />
                <p className="text-xs text-muted-foreground">
                  可选，格式: @/app/路径/page
                </p>
              </div>

              {/* 权限代码 */}
              <div className="space-y-2">
                <Label htmlFor="permission_code">关联权限</Label>
                <Select
                  value={formData.permission_code || "none"}
                  onValueChange={(value) =>
                    setFormData({
                      ...formData,
                      permission_code: value === "none" ? "" : value,
                    })
                  }
                >
                  <SelectTrigger>
                    <SelectValue placeholder="选择权限（可选）" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="none">无</SelectItem>
                    {permissions.map((perm) => (
                      <SelectItem key={perm.id} value={perm.code}>
                        {perm.name} ({perm.code})
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                <p className="text-xs text-muted-foreground">
                  关联权限后，只有拥有该权限的用户才能看到此菜单
                </p>
              </div>

              {/* 可见性 */}
              <div className="flex items-center justify-between">
                <div className="space-y-0.5">
                  <Label>是否可见</Label>
                  <p className="text-xs text-muted-foreground">
                    隐藏后不会在菜单中显示
                  </p>
                </div>
                <Switch
                  checked={formData.visible}
                  onCheckedChange={(checked) =>
                    setFormData({ ...formData, visible: checked })
                  }
                />
              </div>

              {/* 状态 */}
              <div className="space-y-2">
                <Label htmlFor="status">状态</Label>
                <Select
                  value={formData.status}
                  onValueChange={(value) =>
                    setFormData({ ...formData, status: value })
                  }
                >
                  <SelectTrigger>
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="active">启用</SelectItem>
                    <SelectItem value="inactive">禁用</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </TabsContent>
          </Tabs>

          <DialogFooter className="mt-6">
            <Button
              type="button"
              variant="outline"
              onClick={() => onOpenChange(false)}
            >
              取消
            </Button>
            <Button type="submit" disabled={loading}>
              {loading ? "保存中..." : "保存"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
