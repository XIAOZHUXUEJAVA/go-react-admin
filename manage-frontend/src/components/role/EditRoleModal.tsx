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
import { Role, UpdateRoleRequest } from "@/types/role";
import { FormDialog } from "@/components/common";

// 表单验证 Schema
const editRoleFormSchema = z.object({
  name: z
    .string()
    .min(2, { message: "角色名称至少需要2个字符" })
    .max(50, { message: "角色名称不能超过50个字符" }),
  description: z
    .string()
    .max(200, { message: "描述不能超过200个字符" })
    .optional(),
  status: z.enum(["active", "inactive"]),
});

type EditRoleFormValues = z.infer<typeof editRoleFormSchema>;

interface EditRoleModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  role: Role | null;
  onSubmit: (id: number, data: UpdateRoleRequest) => Promise<void>;
  loading?: boolean;
}

/**
 * 编辑角色对话框组件
 */
export const EditRoleModal: React.FC<EditRoleModalProps> = ({
  open,
  onOpenChange,
  role,
  onSubmit,
  loading = false,
}) => {
  const form = useForm<EditRoleFormValues>({
    resolver: zodResolver(editRoleFormSchema),
    defaultValues: {
      name: "",
      description: "",
      status: "active",
    },
  });

  // 当角色数据变化时更新表单
  useEffect(() => {
    if (role) {
      form.reset({
        name: role.name,
        description: role.description || "",
        status: role.status as "active" | "inactive",
      });
    }
  }, [role, form]);

  // 处理表单提交
  const handleSubmit = async (values: EditRoleFormValues) => {
    if (!role) return;
    await onSubmit(role.id, {
      ...values,
      description: values.description || "",
    });
  };

  return (
    <FormDialog
      open={open}
      onOpenChange={onOpenChange}
      title="编辑角色"
      description={`修改角色 "${role?.name}" 的信息`}
      form={form}
      onSubmit={handleSubmit}
      loading={loading}
      submitText="保存更改"
      maxWidth="sm:max-w-[500px]"
      resetOnClose={false}
    >
      <Form {...form}>
        <div className="space-y-4">
          {/* 角色代码（只读） */}
          <div className="space-y-2">
            <label className="text-sm font-medium">角色代码</label>
            <div className="flex items-center gap-2">
              <code className="text-sm bg-muted px-3 py-2 rounded-md flex-1">
                {role?.code}
              </code>
              <span className="text-xs text-muted-foreground">
                (不可修改)
              </span>
            </div>
          </div>

          {/* 角色名称 */}
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>
                  角色名称 <span className="text-red-500">*</span>
                </FormLabel>
                <FormControl>
                  <Input
                    placeholder="例如：管理员"
                    {...field}
                    disabled={loading}
                  />
                </FormControl>
                <FormDescription>角色的显示名称</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* 描述 */}
          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel>描述</FormLabel>
                <FormControl>
                  <Textarea
                    placeholder="角色的详细描述..."
                    className="resize-none"
                    rows={3}
                    {...field}
                    disabled={loading}
                  />
                </FormControl>
                <FormDescription>角色的功能说明（可选）</FormDescription>
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
                  disabled={loading || role?.is_system}
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
                <FormDescription>
                  {role?.is_system
                    ? "系统角色状态不可修改"
                    : "角色的启用状态"}
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
      </Form>
    </FormDialog>
  );
};
