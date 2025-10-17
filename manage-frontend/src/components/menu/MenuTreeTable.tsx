"use client";

import React, { useState } from "react";
import {
  ChevronRight,
  ChevronDown,
  Edit,
  Trash2,
  Plus,
  Eye,
  EyeOff,
  GripVertical,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
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
import { type Menu } from "@/types/menu";
import * as LucideIcons from "lucide-react";

interface MenuTreeTableProps {
  menus: Menu[];
  loading: boolean;
  onEdit: (menu: Menu) => void;
  onDelete: (menu: Menu) => void;
  onCreate: (parent?: Menu) => void;
  onOrderUpdate: (
    menus: { id: number; order_num: number; parent_id: number | null }[]
  ) => void;
}

export function MenuTreeTable({
  menus,
  loading,
  onEdit,
  onDelete,
  onCreate,
  onOrderUpdate,
}: MenuTreeTableProps) {
  const [expandedIds, setExpandedIds] = useState<Set<number>>(new Set());
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [menuToDelete, setMenuToDelete] = useState<Menu | null>(null);
  const [draggedItem, setDraggedItem] = useState<Menu | null>(null);

  // 切换展开/折叠
  const toggleExpand = (id: number) => {
    const newExpanded = new Set(expandedIds);
    if (newExpanded.has(id)) {
      newExpanded.delete(id);
    } else {
      newExpanded.add(id);
    }
    setExpandedIds(newExpanded);
  };

  // 获取图标组件
  const getIconComponent = (iconName: string) => {
    const Icon = LucideIcons[
      iconName as keyof typeof LucideIcons
    ] as React.ComponentType<{ className?: string }>;
    return Icon ? <Icon className="h-4 w-4" /> : null;
  };

  // 处理删除确认
  const handleDeleteClick = (menu: Menu) => {
    setMenuToDelete(menu);
    setDeleteDialogOpen(true);
  };

  const handleDeleteConfirm = () => {
    if (menuToDelete) {
      onDelete(menuToDelete);
    }
    setDeleteDialogOpen(false);
    setMenuToDelete(null);
  };

  // 拖拽处理
  const handleDragStart = (e: React.DragEvent, menu: Menu) => {
    setDraggedItem(menu);
    e.dataTransfer.effectAllowed = "move";
  };

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
    e.dataTransfer.dropEffect = "move";
  };

  const handleDrop = (e: React.DragEvent, targetMenu: Menu) => {
    e.preventDefault();
    if (!draggedItem || draggedItem.id === targetMenu.id) return;

    // 不允许拖到自己的子菜单
    if (
      draggedItem.parent_id === null &&
      targetMenu.parent_id === draggedItem.id
    ) {
      return;
    }

    // 交换顺序
    const updatedMenus = [
      {
        id: draggedItem.id,
        order_num: targetMenu.order_num,
        parent_id: draggedItem.parent_id,
      },
      {
        id: targetMenu.id,
        order_num: draggedItem.order_num,
        parent_id: targetMenu.parent_id,
      },
    ];

    onOrderUpdate(updatedMenus);
    setDraggedItem(null);
  };

  // 渲染菜单行
  const renderMenuRow = (menu: Menu, level: number = 0) => {
    const hasChildren = menu.children && menu.children.length > 0;
    const isExpanded = expandedIds.has(menu.id);

    return (
      <React.Fragment key={menu.id}>
        <TableRow
          draggable
          onDragStart={(e) => handleDragStart(e, menu)}
          onDragOver={handleDragOver}
          onDrop={(e) => handleDrop(e, menu)}
          className="cursor-move hover:bg-muted/50"
        >
          {/* 菜单名称 */}
          <TableCell>
            <div
              className="flex items-center gap-2"
              style={{ paddingLeft: `${level * 24}px` }}
            >
              <GripVertical className="h-4 w-4 text-muted-foreground" />
              {hasChildren && (
                <Button
                  variant="ghost"
                  size="sm"
                  className="h-6 w-6 p-0"
                  onClick={() => toggleExpand(menu.id)}
                >
                  {isExpanded ? (
                    <ChevronDown className="h-4 w-4" />
                  ) : (
                    <ChevronRight className="h-4 w-4" />
                  )}
                </Button>
              )}
              {!hasChildren && <div className="w-6" />}
              <div className="flex items-center gap-2">
                {getIconComponent(menu.icon)}
                <span className="font-medium">{menu.title}</span>
              </div>
            </div>
          </TableCell>

          {/* 路径 */}
          <TableCell>
            <code className="text-xs bg-muted px-2 py-1 rounded">
              {menu.path}
            </code>
          </TableCell>

          {/* 类型 */}
          <TableCell>
            <Badge variant={menu.type === "menu" ? "default" : "secondary"}>
              {menu.type === "menu" ? "菜单" : "按钮"}
            </Badge>
          </TableCell>

          {/* 权限代码 */}
          <TableCell>
            {menu.permission_code ? (
              <code className="text-xs text-muted-foreground">
                {menu.permission_code}
              </code>
            ) : (
              <span className="text-muted-foreground">-</span>
            )}
          </TableCell>

          {/* 排序 */}
          <TableCell className="text-center">{menu.order_num}</TableCell>

          {/* 可见性 */}
          <TableCell>
            {menu.visible ? (
              <Eye className="h-4 w-4 text-green-600" />
            ) : (
              <EyeOff className="h-4 w-4 text-muted-foreground" />
            )}
          </TableCell>

          {/* 状态 */}
          <TableCell>
            <Badge variant={menu.status === "active" ? "default" : "secondary"}>
              {menu.status === "active" ? "启用" : "禁用"}
            </Badge>
          </TableCell>

          {/* 操作 */}
          <TableCell>
            <div className="flex items-center gap-1">
              <Button
                variant="ghost"
                size="sm"
                onClick={() => onCreate(menu)}
                title="添加子菜单"
              >
                <Plus className="h-4 w-4" />
              </Button>
              <Button
                variant="ghost"
                size="sm"
                onClick={() => onEdit(menu)}
                title="编辑"
              >
                <Edit className="h-4 w-4" />
              </Button>
              <Button
                variant="ghost"
                size="sm"
                onClick={() => handleDeleteClick(menu)}
                title="删除"
                disabled={hasChildren}
              >
                <Trash2 className="h-4 w-4" />
              </Button>
            </div>
          </TableCell>
        </TableRow>

        {/* 递归渲染子菜单 */}
        {hasChildren &&
          isExpanded &&
          menu.children!.map((child) => renderMenuRow(child, level + 1))}
      </React.Fragment>
    );
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="text-muted-foreground">加载中...</div>
      </div>
    );
  }

  return (
    <>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>菜单名称</TableHead>
              <TableHead>路径</TableHead>
              <TableHead>类型</TableHead>
              <TableHead>权限代码</TableHead>
              <TableHead className="text-center">排序</TableHead>
              <TableHead>可见</TableHead>
              <TableHead>状态</TableHead>
              <TableHead className="text-right">操作</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {menus.length === 0 ? (
              <TableRow>
                <TableCell
                  colSpan={8}
                  className="text-center text-muted-foreground"
                >
                  暂无菜单数据
                </TableCell>
              </TableRow>
            ) : (
              menus.map((menu) => renderMenuRow(menu))
            )}
          </TableBody>
        </Table>
      </div>

      {/* 删除确认对话框 */}
      <AlertDialog open={deleteDialogOpen} onOpenChange={setDeleteDialogOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>确认删除</AlertDialogTitle>
            <AlertDialogDescription>
              确定要删除菜单 &quot;{menuToDelete?.title}&quot;
              吗？此操作无法撤销。
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>取消</AlertDialogCancel>
            <AlertDialogAction onClick={handleDeleteConfirm}>
              确认删除
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  );
}
