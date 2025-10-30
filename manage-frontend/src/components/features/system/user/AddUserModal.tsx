"use client";

import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Eye, EyeOff, Loader2 } from "lucide-react";
import { CreateUserRequest } from "@/types/api";
import { userSchemas, type CreateUserFormData } from "@/lib/validations";
import { useAvailabilityCheck } from "@/hooks/useUserValidation";
import { useAllRoles } from "@/hooks/useRoles";
import { FormDialog } from "@/components/common";

interface AddUserModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSubmit: (userData: CreateUserRequest) => Promise<void>;
  loading?: boolean;
}

/**
 * 添加用户模态框组件（使用 FormDialog 重构版本）
 */
export const AddUserModal: React.FC<AddUserModalProps> = ({
  open,
  onOpenChange,
  onSubmit,
  loading = false,
}) => {
  const [showPassword, setShowPassword] = useState(false);
  const { roles, loading: rolesLoading } = useAllRoles();

  // 表单实例
  const form = useForm<CreateUserFormData>({
    // @ts-expect-error - zod resolver 类型兼容性问题
    resolver: zodResolver(userSchemas.create),
    defaultValues: {
      role: "user" as const,
    },
    mode: "onChange",
  });

  const {
    register,
    setValue,
    watch,
    formState: { errors },
  } = form;

  // 用户验证 hook
  const {
    usernameAvailability,
    emailAvailability,
    checkUsernameAvailability,
    checkEmailAvailability,
    resetAvailabilityCheck,
  } = useAvailabilityCheck();

  const roleValue = watch("role");
  const usernameValue = watch("username");
  const emailValue = watch("email");

  // 监听用户名变化并进行实时验证
  useEffect(() => {
    if (usernameValue && usernameValue.length >= 3) {
      checkUsernameAvailability(usernameValue);
    }
  }, [usernameValue, checkUsernameAvailability]);

  // 监听邮箱变化并进行实时验证
  useEffect(() => {
    if (emailValue && emailValue.includes("@")) {
      checkEmailAvailability(emailValue);
    }
  }, [emailValue, checkEmailAvailability]);

  const handleFormSubmit = async (data: CreateUserFormData) => {
    // 检查用户名和邮箱是否可用
    if (usernameAvailability.isAvailable === false) {
      return;
    }
    if (emailAvailability.isAvailable === false) {
      return;
    }

    await onSubmit(data);
  };

  // 检查是否可以提交表单
  const canSubmit =
    !loading &&
    usernameAvailability.isAvailable !== false &&
    emailAvailability.isAvailable !== false &&
    !usernameAvailability.isChecking &&
    !emailAvailability.isChecking;

  const handleClose = () => {
    resetAvailabilityCheck();
    setShowPassword(false);
  };

  return (
    <FormDialog
      open={open}
      onOpenChange={onOpenChange}
      title="添加新用户"
      description="创建一个新的用户账户，请填写所有必需的信息。"
      form={form}
      onSubmit={handleFormSubmit}
      loading={loading}
      disableSubmit={!canSubmit}
      submitText="创建用户"
      onClose={handleClose}
    >
      <div className="grid gap-4">
        {/* 用户名 */}
        <div className="space-y-2">
          <Label htmlFor="username">用户名 *</Label>
          <div className="relative">
            <Input
              id="username"
              placeholder="请输入用户名"
              {...register("username")}
              className={
                errors.username || usernameAvailability.isAvailable === false
                  ? "border-red-500"
                  : usernameAvailability.isAvailable === true
                  ? "border-green-500"
                  : ""
              }
            />
            {usernameAvailability.isChecking && (
              <div className="absolute right-3 top-1/2 -translate-y-1/2">
                <Loader2 className="h-4 w-4 animate-spin text-gray-400" />
              </div>
            )}
          </div>
          {errors.username && (
            <p className="text-sm text-red-500">{errors.username.message}</p>
          )}
          {!errors.username && usernameAvailability.message && (
            <p
              className={`text-sm ${
                usernameAvailability.isAvailable
                  ? "text-green-600"
                  : "text-red-500"
              }`}
            >
              {usernameAvailability.message}
            </p>
          )}
        </div>

        {/* 邮箱 */}
        <div className="space-y-2">
          <Label htmlFor="email">邮箱地址 *</Label>
          <div className="relative">
            <Input
              id="email"
              type="email"
              placeholder="请输入邮箱地址"
              {...register("email")}
              className={
                errors.email || emailAvailability.isAvailable === false
                  ? "border-red-500"
                  : emailAvailability.isAvailable === true
                  ? "border-green-500"
                  : ""
              }
            />
            {emailAvailability.isChecking && (
              <div className="absolute right-3 top-1/2 -translate-y-1/2">
                <Loader2 className="h-4 w-4 animate-spin text-gray-400" />
              </div>
            )}
          </div>
          {errors.email && (
            <p className="text-sm text-red-500">{errors.email.message}</p>
          )}
          {!errors.email && emailAvailability.message && (
            <p
              className={`text-sm ${
                emailAvailability.isAvailable
                  ? "text-green-600"
                  : "text-red-500"
              }`}
            >
              {emailAvailability.message}
            </p>
          )}
        </div>

        {/* 密码 */}
        <div className="space-y-2">
          <Label htmlFor="password">密码 *</Label>
          <div className="relative">
            <Input
              id="password"
              type={showPassword ? "text" : "password"}
              placeholder="请输入密码"
              {...register("password")}
              className={errors.password ? "border-red-500" : ""}
            />
            <button
              type="button"
              onClick={() => setShowPassword(!showPassword)}
              className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
            >
              {showPassword ? (
                <EyeOff className="h-4 w-4" />
              ) : (
                <Eye className="h-4 w-4" />
              )}
            </button>
          </div>
          {errors.password && (
            <p className="text-sm text-red-500">{errors.password.message}</p>
          )}
          <p className="text-xs text-muted-foreground">
            密码至少8个字符，包含大小写字母和数字
          </p>
        </div>

        {/* 角色选择 */}
        <div className="space-y-2">
          <Label htmlFor="role">角色 *</Label>
          <Select
            value={roleValue}
            onValueChange={(value) =>
              setValue("role", value, {
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
};
