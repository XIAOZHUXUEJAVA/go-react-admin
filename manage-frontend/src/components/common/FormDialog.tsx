import React, { useEffect } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Loader2 } from "lucide-react";
import { UseFormReturn, FieldValues } from "react-hook-form";

interface FormDialogProps<TFieldValues extends FieldValues> {
  /**
   * 对话框是否打开
   */
  open: boolean;
  /**
   * 对话框打开状态变化回调
   */
  onOpenChange: (open: boolean) => void;
  /**
   * 对话框标题
   */
  title: string;
  /**
   * 对话框描述
   */
  description?: string;
  /**
   * react-hook-form 实例
   */
  form: UseFormReturn<TFieldValues>;
  /**
   * 表单提交处理函数
   */
  onSubmit: (data: TFieldValues) => Promise<void> | void;
  /**
   * 表单内容（children）
   */
  children: React.ReactNode;
  /**
   * 提交按钮文本
   * @default "提交"
   */
  submitText?: string;
  /**
   * 取消按钮文本
   * @default "取消"
   */
  cancelText?: string;
  /**
   * 是否显示加载状态
   */
  loading?: boolean;
  /**
   * 是否禁用提交按钮（自定义逻辑）
   */
  disableSubmit?: boolean;
  /**
   * 对话框最大宽度
   * @default "max-w-md"
   */
  maxWidth?: string;
  /**
   * 关闭时是否重置表单
   * @default true
   */
  resetOnClose?: boolean;
  /**
   * 自定义 Footer（如果提供，会覆盖默认的提交/取消按钮）
   */
  customFooter?: React.ReactNode;
  /**
   * 关闭前的回调（可用于清理状态）
   */
  onClose?: () => void;
}

/**
 * 通用表单对话框组件
 * 
 * 封装了常见的表单对话框模式：
 * - Dialog 包装器
 * - 表单验证（react-hook-form）
 * - 加载状态处理
 * - 提交/取消按钮
 * - 表单重置逻辑
 * 
 * @example
 * ```tsx
 * const form = useForm<FormData>({
 *   resolver: zodResolver(schema),
 * });
 * 
 * <FormDialog
 *   open={isOpen}
 *   onOpenChange={setIsOpen}
 *   title="添加用户"
 *   description="创建一个新的用户账户"
 *   form={form}
 *   onSubmit={handleSubmit}
 *   loading={isSubmitting}
 * >
 *   <div className="space-y-4">
 *     <Input {...form.register("name")} />
 *     <Input {...form.register("email")} />
 *   </div>
 * </FormDialog>
 * ```
 */
export function FormDialog<TFieldValues extends FieldValues>({
  open,
  onOpenChange,
  title,
  description,
  form,
  onSubmit,
  children,
  submitText = "提交",
  cancelText = "取消",
  loading = false,
  disableSubmit = false,
  maxWidth = "max-w-md",
  resetOnClose = true,
  customFooter,
  onClose,
}: FormDialogProps<TFieldValues>) {
  const { handleSubmit, reset, formState } = form;

  // 当对话框关闭时重置表单
  useEffect(() => {
    if (!open && resetOnClose) {
      reset();
      onClose?.();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [open, resetOnClose]);

  const handleFormSubmit = async (data: TFieldValues) => {
    try {
      await onSubmit(data);
      if (resetOnClose) {
        reset();
      }
      onOpenChange(false);
    } catch (error) {
      // 错误由父组件处理
      console.error("Form submission error:", error);
    }
  };

  const handleClose = () => {
    if (!loading) {
      onClose?.();
      onOpenChange(false);
    }
  };

  const isSubmitDisabled =
    loading || disableSubmit || formState.isSubmitting || !formState.isValid;

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogContent className={maxWidth}>
        <DialogHeader>
          <DialogTitle>{title}</DialogTitle>
          {description && <DialogDescription>{description}</DialogDescription>}
        </DialogHeader>

        <form onSubmit={handleSubmit(handleFormSubmit)}>
          <div className="py-4">{children}</div>

          {customFooter || (
            <DialogFooter>
              <Button
                type="button"
                variant="outline"
                onClick={handleClose}
                disabled={loading}
              >
                {cancelText}
              </Button>
              <Button type="submit" disabled={isSubmitDisabled}>
                {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                {submitText}
              </Button>
            </DialogFooter>
          )}
        </form>
      </DialogContent>
    </Dialog>
  );
}
