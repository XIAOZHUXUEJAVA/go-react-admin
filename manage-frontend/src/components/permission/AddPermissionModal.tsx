"use client";

import React from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
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
import { Button } from "@/components/ui/button";
import { Loader2 } from "lucide-react";
import { CreatePermissionRequest } from "@/types/permission";

// 表单验证 Schema
const permissionFormSchema = z.object({
  name: z
    .string()
    .min(2, { message: "权限名称至少需要2个字符" })
    .max(50, { message: "权限名称不能超过50个字符" }),
  code: z
    .string()
    .min(2, { message: "权限代码至少需要2个字符" })
    .max(100, { message: "权限代码不能超过100个字符" })
    .regex(/^[a-z0-9:_]+$/, {
      message: "权限代码只能包含小写字母、数字、冒号和下划线",
    }),
  resource: z
    .string()
    .min(2, { message: "资源名称至少需要2个字符" })
    .max(50, { message: "资源名称不能超过50个字符" }),
  action: z
    .string()
    .min(2, { message: "操作至少需要2个字符" })
    .max(50, { message: "操作不能超过50个字符" }),
  type: z.enum(["api", "menu", "button"]),
  path: z.string().max(200, { message: "路径不能超过200个字符" }).optional(),
  method: z.string().max(10, { message: "方法不能超过10个字符" }).optional(),
  description: z
    .string()
    .max(200, { message: "描述不能超过200个字符" })
    .optional(),
});

type PermissionFormValues = z.infer<typeof permissionFormSchema>;

interface AddPermissionModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSubmit: (data: CreatePermissionRequest) => Promise<void>;
  loading?: boolean;
}

/**
 * 添加权限对话框组件
 */
export const AddPermissionModal: React.FC<AddPermissionModalProps> = ({
  open,
  onOpenChange,
  onSubmit,
  loading = false,
}) => {
  const form = useForm<PermissionFormValues>({
    resolver: zodResolver(permissionFormSchema),
    defaultValues: {
      name: "",
      code: "",
      resource: "",
      action: "",
      type: "api",
      path: "",
      method: "",
      description: "",
    },
  });

  // 监听类型变化，自动填充建议值
  const watchType = form.watch("type");

  // 处理表单提交
  const handleSubmit = async (values: PermissionFormValues) => {
    try {
      await onSubmit(values);
      form.reset();
      onOpenChange(false);
    } catch (error) {
    }
  };

  // 处理对话框关闭
  const handleOpenChange = (newOpen: boolean) => {
    if (!newOpen && !loading) {
      form.reset();
    }
    onOpenChange(newOpen);
  };

  return (
    <Dialog open={open} onOpenChange={handleOpenChange}>
      <DialogContent className="sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>添加权限</DialogTitle>
          <DialogDescription>创建一个新的系统权限</DialogDescription>
        </DialogHeader>

        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="space-y-4"
          >
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

            {/* 权限代码 */}
            <FormField
              control={form.control}
              name="code"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    权限代码 <span className="text-red-500">*</span>
                  </FormLabel>
                  <FormControl>
                    <Input
                      placeholder="例如：user:list"
                      {...field}
                      disabled={loading}
                    />
                  </FormControl>
                  <FormDescription>
                    唯一标识符，格式：资源:操作（如 user:create）
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <div className="grid grid-cols-2 gap-4">
              {/* 资源名称 */}
              <FormField
                control={form.control}
                name="resource"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>
                      资源 <span className="text-red-500">*</span>
                    </FormLabel>
                    <FormControl>
                      <Input
                        placeholder="例如：user"
                        {...field}
                        disabled={loading}
                      />
                    </FormControl>
                    <FormDescription>资源类型</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              {/* 操作 */}
              <FormField
                control={form.control}
                name="action"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>
                      操作 <span className="text-red-500">*</span>
                    </FormLabel>
                    <FormControl>
                      <Input
                        placeholder="例如：list"
                        {...field}
                        disabled={loading}
                      />
                    </FormControl>
                    <FormDescription>操作类型</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>

            {/* 权限类型 */}
            <FormField
              control={form.control}
              name="type"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    权限类型 <span className="text-red-500">*</span>
                  </FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                    disabled={loading}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="选择权限类型" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value="api">API 权限</SelectItem>
                      <SelectItem value="menu">菜单权限</SelectItem>
                      <SelectItem value="button">按钮权限</SelectItem>
                    </SelectContent>
                  </Select>
                  <FormDescription>权限的应用场景</FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            {/* API 类型时显示路径和方法 */}
            {watchType === "api" && (
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
            {watchType === "menu" && (
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

            <DialogFooter>
              <Button
                type="button"
                variant="outline"
                onClick={() => handleOpenChange(false)}
                disabled={loading}
              >
                取消
              </Button>
              <Button type="submit" disabled={loading}>
                {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                {loading ? "创建中..." : "创建权限"}
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};
