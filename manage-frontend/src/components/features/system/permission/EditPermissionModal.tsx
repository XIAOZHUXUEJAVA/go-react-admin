"use client";

import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Permission, UpdatePermissionRequest } from "@/types/permission";
import { FormDialog } from "@/components/common";

// 表单验证 Schema
const editPermissionFormSchema = z.object({
  name: z
    .string()
    .min(2, { message: "权限名称至少需要2个字符" })
    .max(50, { message: "权限名称不能超过50个字符" }),
  path: z.string().max(200, { message: "路径不能超过200个字符" }).optional(),
  method: z.string().max(10, { message: "方法不能超过10个字符" }).optional(),
  description: z
    .string()
    .max(200, { message: "描述不能超过200个字符" })
    .optional(),
  status: z.enum(["active", "inactive"]),
});

type EditPermissionFormValues = z.infer<typeof editPermissionFormSchema>;

interface EditPermissionModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  permission: Permission | null;
  onSubmit: (id: number, data: UpdatePermissionRequest) => Promise<void>;
  loading?: boolean;
}

/**
 * 编辑权限对话框组件
 */
export const EditPermissionModal: React.FC<EditPermissionModalProps> = ({
  open,
  onOpenChange,
  permission,
  onSubmit,
  loading = false,
}) => {
  const form = useForm<EditPermissionFormValues>({
    resolver: zodResolver(editPermissionFormSchema),
    defaultValues: {
      name: "",
      path: "",
      method: "",
      description: "",
      status: "active",
    },
  });

  // 当权限数据变化时更新表单
  useEffect(() => {
    if (permission) {
      form.reset({
        name: permission.name,
        path: permission.path || "",
        method: permission.method || "",
        description: permission.description || "",
        status: permission.status as "active" | "inactive",
      });
    }
  }, [permission, form]);

  // 处理表单提交
  const handleSubmit = async (values: EditPermissionFormValues) => {
    if (!permission) return;
    await onSubmit(permission.id, {
      ...values,
      path: values.path || "",
      method: values.method || "",
      description: values.description || "",
    });
  };

  // 获取类型标签
  const getTypeLabel = (type: string) => {
    const labels: Record<string, string> = {
      api: "API 权限",
      menu: "菜单权限",
      button: "按钮权限",
    };
    return labels[type] || type;
  };

  return (
    <FormDialog
      open={open}
      onOpenChange={onOpenChange}
      title="编辑权限"
      description={`修改权限 "${permission?.name}" 的信息`}
      form={form}
      onSubmit={handleSubmit}
      loading={loading}
      submitText="保存更改"
      maxWidth="sm:max-w-[600px]"
      resetOnClose={false}
    >
      <Form {...form}>
        <div className="space-y-4 max-h-[60vh] overflow-y-auto pr-2">
            {/* 权限代码（只读） */}
            <div className="space-y-2">
              <label className="text-sm font-medium">权限代码</label>
              <div className="flex items-center gap-2">
                <code className="text-sm bg-muted px-3 py-2 rounded-md flex-1">
                  {permission?.code}
                </code>
                <span className="text-xs text-muted-foreground">
                  (不可修改)
                </span>
              </div>
            </div>

            {/* 资源和操作（只读） */}
            <div className="grid grid-cols-3 gap-4">
              <div className="space-y-2">
                <label className="text-sm font-medium">资源</label>
                <div className="text-sm bg-muted px-3 py-2 rounded-md capitalize">
                  {permission?.resource}
                </div>
              </div>
              <div className="space-y-2">
                <label className="text-sm font-medium">操作</label>
                <div className="text-sm bg-muted px-3 py-2 rounded-md capitalize">
                  {permission?.action}
                </div>
              </div>
              <div className="space-y-2">
                <label className="text-sm font-medium">类型</label>
                <div className="text-sm bg-muted px-3 py-2 rounded-md">
                  {permission && getTypeLabel(permission.type)}
                </div>
              </div>
            </div>

            {/* 权限名称 */}
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    权限名称 <span className="text-red-500">*</span>
                  </FormLabel>
                  <FormControl>
                    <Input
                      placeholder="例如：查看用户列表"
                      {...field}
                      disabled={loading}
                    />
                  </FormControl>
                  <FormDescription>权限的显示名称</FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            {/* API 类型时显示路径和方法 */}
            {permission?.type === "api" && (
              <div className="grid grid-cols-2 gap-4">
                {/* 路径 */}
                <FormField
                  control={form.control}
                  name="path"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>API 路径</FormLabel>
                      <FormControl>
                        <Input
                          placeholder="/api/v1/users"
                          {...field}
                          disabled={loading}
                        />
                      </FormControl>
                      <FormDescription>API 端点路径</FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                {/* HTTP 方法 */}
                <FormField
                  control={form.control}
                  name="method"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>HTTP 方法</FormLabel>
                      <Select
                        onValueChange={field.onChange}
                        defaultValue={field.value}
                        disabled={loading}
                      >
                        <FormControl>
                          <SelectTrigger>
                            <SelectValue placeholder="选择方法" />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          <SelectItem value="GET">GET</SelectItem>
                          <SelectItem value="POST">POST</SelectItem>
                          <SelectItem value="PUT">PUT</SelectItem>
                          <SelectItem value="DELETE">DELETE</SelectItem>
                          <SelectItem value="PATCH">PATCH</SelectItem>
                        </SelectContent>
                      </Select>
                      <FormDescription>HTTP 请求方法</FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
            )}

            {/* 菜单类型时显示路径 */}
            {permission?.type === "menu" && (
              <FormField
                control={form.control}
                name="path"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>菜单路径</FormLabel>
                    <FormControl>
                      <Input
                        placeholder="/dashboard/users"
                        {...field}
                        disabled={loading}
                      />
                    </FormControl>
                    <FormDescription>前端路由路径</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            )}

            {/* 描述 */}
            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>描述</FormLabel>
                  <FormControl>
                    <Textarea
                      placeholder="权限的详细描述..."
                      className="resize-none"
                      rows={3}
                      {...field}
                      disabled={loading}
                    />
                  </FormControl>
                  <FormDescription>权限的功能说明（可选）</FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            {/* 状态 */}
            <FormField
              control={form.control}
              name="status"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    状态 <span className="text-red-500">*</span>
                  </FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                    disabled={loading}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="选择状态" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="active">启用</SelectItem>
                      <SelectItem value="inactive">禁用</SelectItem>
                    </SelectContent>
                  </Select>
                  <FormDescription>权限的启用状态</FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

        </div>
      </Form>
    </FormDialog>
  );
};
