import React from "react";
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

interface DeleteConfirmDialogProps {
  /**
   * 对话框是否打开
   */
  open: boolean;
  /**
   * 对话框打开状态变化回调
   */
  onOpenChange: (open: boolean) => void;
  /**
   * 确认删除回调
   */
  onConfirm: () => void;
  /**
   * 对话框标题
   * @default "确认删除"
   */
  title?: string;
  /**
   * 要删除的资源名称，会显示在描述中
   */
  resourceName?: string;
  /**
   * 资源类型（用于描述文本）
   * @example "用户", "角色", "权限"
   */
  resourceType?: string;
  /**
   * 自定义描述文本（如果提供，会覆盖默认描述）
   */
  description?: string;
  /**
   * 确认按钮文本
   * @default "删除"
   */
  confirmText?: string;
  /**
   * 取消按钮文本
   * @default "取消"
   */
  cancelText?: string;
  /**
   * 是否显示加载状态
   */
  loading?: boolean;
}

/**
 * 通用删除确认对话框组件
 * 
 * @example
 * ```tsx
 * <DeleteConfirmDialog
 *   open={isOpen}
 *   onOpenChange={setIsOpen}
 *   onConfirm={handleDelete}
 *   resourceName={user.username}
 *   resourceType="用户"
 * />
 * ```
 */
export const DeleteConfirmDialog: React.FC<DeleteConfirmDialogProps> = ({
  open,
  onOpenChange,
  onConfirm,
  title = "确认删除",
  resourceName,
  resourceType,
  description,
  confirmText = "删除",
  cancelText = "取消",
  loading = false,
}) => {
  // 生成默认描述文本
  const getDefaultDescription = () => {
    if (resourceName && resourceType) {
      return `您确定要删除${resourceType} "${resourceName}" 吗？此操作无法撤销。`;
    }
    if (resourceType) {
      return `您确定要删除此${resourceType}吗？此操作无法撤销。`;
    }
    return "您确定要删除吗？此操作无法撤销。";
  };

  const finalDescription = description || getDefaultDescription();

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>{title}</AlertDialogTitle>
          <AlertDialogDescription>{finalDescription}</AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel disabled={loading}>{cancelText}</AlertDialogCancel>
          <AlertDialogAction
            onClick={onConfirm}
            disabled={loading}
            className="bg-red-600 hover:bg-red-700"
          >
            {loading ? "删除中..." : confirmText}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
};
