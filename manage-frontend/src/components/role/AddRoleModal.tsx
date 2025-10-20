"use client";

import React from "react";
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
import { CreateRoleRequest } from "@/types/role";
import { FormDialog } from "@/components/common";

// 表单验证 Schema
const roleFormSchema = z.object({
  name: z
    .string()
    .min(2, { message: "角色名称至少需要2个字符" })
    .max(50, { message: "角色名称不能超过50个字符" }),
  code: z
    .string()
    .min(2, { message: "角色代码至少需要2个字符" })
    .max(50, { message: "角色代码不能超过50个字符" })
    .regex(/^[a-z0-9_]+$/, {
      message: "角色代码只能包含小写字母、数字和下划线",
    }),
  description: z
    .string()
    .max(200, { message: "描述不能超过200个字符" })
    .optional(),
});

type RoleFormValues = z.infer<typeof roleFormSchema>;

interface AddRoleModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSubmit: (data: CreateRoleRequest) => Promise<void>;
  loading?: boolean;
}

/**
 * 添加角色对话框组件
 */
export const AddRoleModal: React.FC<AddRoleModalProps> = ({
  open,
  onOpenChange,
  onSubmit,
  loading = false,
}) => {
  const form = useForm<RoleFormValues>({
    resolver: zodResolver(roleFormSchema),
    defaultValues: {
      name: "",
      code: "",
      description: "",
    },
  });

  // 处理表单提交
  const handleSubmit = async (values: RoleFormValues) => {
    await onSubmit({
      ...values,
      description: values.description || "",
    });
  };

  return (
    <FormDialog
      open={open}
      onOpenChange={onOpenChange}
      title="添加角色"
      description="创建一个新的系统角色"
      form={form}
      onSubmit={handleSubmit}
      loading={loading}
      submitText="创建角色"
      maxWidth="sm:max-w-[500px]"
    >
      <Form {...form}>
        <div className="space-y-4">
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

          {/* 角色代码 */}
          <FormField
            control={form.control}
            name="code"
            render={({ field }) => (
              <FormItem>
                <FormLabel>
                  角色代码 <span className="text-red-500">*</span>
                </FormLabel>
                <FormControl>
                  <Input
                    placeholder="例如：admin"
                    {...field}
                    disabled={loading}
                  />
                </FormControl>
                <FormDescription>
                  唯一标识符，只能包含小写字母、数字和下划线
                </FormDescription>
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
        </div>
      </Form>
    </FormDialog>
  );
};
