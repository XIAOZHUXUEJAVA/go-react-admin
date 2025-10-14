"use client";

import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
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

interface AddUserModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSubmit: (userData: CreateUserRequest) => Promise<void>;
  loading?: boolean;
}

/**
 * 添加用户模态框组件
 */
export const AddUserModal: React.FC<AddUserModalProps> = ({
  open,
  onOpenChange,
  onSubmit,
  loading = false,
}) => {
  const [showPassword, setShowPassword] = useState(false);
  const { roles, loading: rolesLoading } = useAllRoles();

  // 用户验证 hook
  const {
    usernameAvailability,
    emailAvailability,
    checkUsernameAvailability,
    checkEmailAvailability,
    resetAvailabilityCheck,
  } = useAvailabilityCheck();

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
    watch,
    reset,
  } = useForm<CreateUserFormData>({
    resolver: zodResolver(userSchemas.create),
    defaultValues: {
      role: "user",
    },
  });

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

  // 重置验证状态当模态框关闭时
  useEffect(() => {
    if (!open) {
      resetAvailabilityCheck();
    }
  }, [open, resetAvailabilityCheck]);

  const handleFormSubmit = async (data: CreateUserFormData) => {
    // 检查用户名和邮箱是否可用
    if (usernameAvailability.isAvailable === false) {
      return;
    }
    if (emailAvailability.isAvailable === false) {
      return;
    }

    try {
      await onSubmit(data);
      reset();
      resetAvailabilityCheck();
      onOpenChange(false);
    } catch (error) {}
  };

  // 检查是否可以提交表单
  const canSubmit =
    !loading &&
    usernameAvailability.isAvailable !== false &&
    emailAvailability.isAvailable !== false &&
    !usernameAvailability.isChecking &&
    !emailAvailability.isChecking;

  const handleClose = () => {
    reset();
    onOpenChange(false);
  };

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogContent className="max-w-md">
        <DialogHeader>
          <DialogTitle>添加新用户</DialogTitle>
          <DialogDescription>
            创建一个新的用户账户，请填写所有必需的信息。
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit(handleFormSubmit)}>
          <div className="grid gap-4 py-4">
            {/* 用户名 */}
            <div className="space-y-2">
              <Label htmlFor="username">用户名 *</Label>
              <div className="relative">
                <Input
                  id="username"
                  placeholder="请输入用户名"
                  {...register("username")}
                  className={
                    errors.username ||
                    usernameAvailability.isAvailable === false
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
                <p className="text-sm text-red-500">
                  {errors.username.message}
                </p>
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
                  className={errors.password ? "border-red-500 pr-10" : "pr-10"}
                />
                <Button
                  type="button"
                  variant="ghost"
                  size="sm"
                  className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? (
                    <EyeOff className="h-4 w-4" />
                  ) : (
                    <Eye className="h-4 w-4" />
                  )}
                </Button>
              </div>
              {errors.password && (
                <p className="text-sm text-red-500">
                  {errors.password.message}
                </p>
              )}
            </div>

            {/* 角色 */}
            <div className="space-y-2">
              <Label htmlFor="role">用户角色 *</Label>
              <Select
                value={roleValue}
                onValueChange={(value) => setValue("role", value)}
                disabled={rolesLoading}
              >
                <SelectTrigger className={errors.role ? "border-red-500" : ""}>
                  <SelectValue placeholder={rolesLoading ? "加载中..." : "选择用户角色"} />
                </SelectTrigger>
                <SelectContent>
                  {rolesLoading ? (
                    <div className="flex items-center justify-center p-2">
                      <Loader2 className="h-4 w-4 animate-spin" />
                    </div>
                  ) : roles.length > 0 ? (
                    roles
                      .filter((role) => role.status === "active")
                      .map((role) => (
                        <SelectItem key={role.id} value={role.code}>
                          {role.name}
                        </SelectItem>
                      ))
                  ) : (
                    <div className="p-2 text-sm text-muted-foreground">
                      暂无可用角色
                    </div>
                  )}
                </SelectContent>
              </Select>
              {errors.role && (
                <p className="text-sm text-red-500">{errors.role.message}</p>
              )}
            </div>
          </div>

          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={handleClose}
              disabled={loading}
            >
              取消
            </Button>
            <Button type="submit" disabled={!canSubmit}>
              {loading ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  创建中...
                </>
              ) : usernameAvailability.isChecking ||
                emailAvailability.isChecking ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  验证中...
                </>
              ) : (
                "创建用户"
              )}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
};
