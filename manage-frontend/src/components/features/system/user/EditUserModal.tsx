"use client";

import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { User } from "@/types/api";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useAllRoles } from "@/hooks/useRoles";
import { Loader2 } from "lucide-react";
import { FormDialog } from "@/components/common";

// 编辑用户表单验证 schema
const editUserSchema = z.object({
  username: z.string().min(3, "用户名至少3个字符"),
  email: z.string().email("请输入有效的邮箱地址"),
  role: z.enum(["admin", "user", "moderator"]),
});

type EditUserFormData = z.infer<typeof editUserSchema>;

interface EditUserModalProps {
  user: User | null;
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSave: (user: User) => void;
}

/**
 * 编辑用户模态框组件（使用 FormDialog 重构版本）
 */
export function EditUserModal({
  user,
  open,
  onOpenChange,
  onSave,
}: EditUserModalProps) {
  const { roles, loading: rolesLoading } = useAllRoles();

  const form = useForm<EditUserFormData>({
    resolver: zodResolver(editUserSchema),
    mode: "onChange",
  });

  const { register, setValue, watch, reset, formState: { errors } } = form;

  const roleValue = watch("role");

  // 当用户或模态框状态改变时，重置表单数据
  useEffect(() => {
    if (user && open) {
      reset({
        username: user.username,
        email: user.email,
        role: user.role as "admin" | "user" | "moderator",
      });
    }
  }, [user, open, reset]);

  const handleFormSubmit = async (data: EditUserFormData) => {
    if (!user) return;

    const updatedUser: User = {
      ...user,
      username: data.username,
      email: data.email,
      role: data.role,
    };

    onSave(updatedUser);
  };

  return (
    <FormDialog
      open={open}
      onOpenChange={onOpenChange}
      title="编辑用户"
      description="修改用户的基本信息"
      form={form}
      onSubmit={handleFormSubmit}
      submitText="保存更改"
      resetOnClose={false}
    >
      <div className="grid gap-4">
        {/* 用户名 */}
        <div className="space-y-2">
          <Label htmlFor="edit-username">用户名</Label>
          <Input
            id="edit-username"
            placeholder="请输入用户名"
            {...register("username")}
            className={errors.username ? "border-red-500" : ""}
          />
          {errors.username && (
            <p className="text-sm text-red-500">{errors.username.message}</p>
          )}
        </div>

        {/* 邮箱 */}
        <div className="space-y-2">
          <Label htmlFor="edit-email">邮箱地址</Label>
          <Input
            id="edit-email"
            type="email"
            placeholder="请输入邮箱地址"
            {...register("email")}
            className={errors.email ? "border-red-500" : ""}
          />
          {errors.email && (
            <p className="text-sm text-red-500">{errors.email.message}</p>
          )}
        </div>

        {/* 角色选择 */}
        <div className="space-y-2">
          <Label htmlFor="edit-role">角色</Label>
          <Select
            value={roleValue}
            onValueChange={(value) =>
              setValue("role", value as "admin" | "user" | "moderator", {
                shouldValidate: true,
              })
            }
          >
            <SelectTrigger>
              <SelectValue placeholder="选择角色" />
            </SelectTrigger>
            <SelectContent>
              {rolesLoading ? (
                <div className="flex items-center justify-center py-2">
                  <Loader2 className="h-4 w-4 animate-spin" />
                  <span className="ml-2 text-sm">加载中...</span>
                </div>
              ) : (
                roles.map((role) => (
                  <SelectItem key={role.id} value={role.code}>
                    {role.name}
                  </SelectItem>
                ))
              )}
            </SelectContent>
          </Select>
          {errors.role && (
            <p className="text-sm text-red-500">{errors.role.message}</p>
          )}
        </div>
      </div>
    </FormDialog>
  );
}
