"use client";

import { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { useAuthStore } from "@/stores/authStore";
import { cn } from "@/lib/utils";
import { getErrorMessage } from "@/lib/error-handler";
import {
  Eye,
  EyeOff,
  Loader2,
  AlertCircle,
  CheckCircle,
  XCircle,
} from "lucide-react";
import { userSchemas, type RegisterFormData } from "@/lib/validations";
import { useAvailabilityCheck } from "@/hooks/useUserValidation";

// 使用统一的注册验证 schema
const registerSchema = userSchemas.register;

interface RegisterFormProps {
  className?: string;
}

export default function RegisterPage({ className }: RegisterFormProps) {
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [error, setError] = useState<string>("");
  const [usernameValidation, setUsernameValidation] = useState<{
    isValidating: boolean;
    isValid: boolean;
    message?: string;
  }>({ isValidating: false, isValid: true });

  const [emailValidation, setEmailValidation] = useState<{
    isValidating: boolean;
    isValid: boolean;
    message?: string;
  }>({ isValidating: false, isValid: true });

  const router = useRouter();
  const { register: registerUser, isLoading } = useAuthStore();
  const {
    checkUsernameAvailability,
    checkEmailAvailability,
    usernameAvailability: hookUsernameAvailability,
    emailAvailability: hookEmailAvailability,
  } = useAvailabilityCheck();

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
  });

  const watchedUsername = watch("username");
  const watchedEmail = watch("email");

  // 实时验证用户名
  useEffect(() => {
    if (watchedUsername && watchedUsername.length >= 3) {
      checkUsernameAvailability(watchedUsername);
    } else {
      // 当输入为空或长度不足时，重置验证状态
      setUsernameValidation({
        isValidating: false,
        isValid: true,
        message: undefined,
      });
    }
  }, [watchedUsername, checkUsernameAvailability]);

  // 实时验证邮箱
  useEffect(() => {
    if (watchedEmail && watchedEmail.includes("@")) {
      checkEmailAvailability(watchedEmail);
    } else {
      // 当输入为空或格式不正确时，重置验证状态
      setEmailValidation({
        isValidating: false,
        isValid: true,
        message: undefined,
      });
    }
  }, [watchedEmail, checkEmailAvailability]);

  // 同步用户名验证状态
  useEffect(() => {
    setUsernameValidation({
      isValidating: hookUsernameAvailability.isChecking,
      isValid: hookUsernameAvailability.isAvailable !== false,
      message: hookUsernameAvailability.message,
    });
  }, [hookUsernameAvailability]);

  // 同步邮箱验证状态
  useEffect(() => {
    setEmailValidation({
      isValidating: hookEmailAvailability.isChecking,
      isValid: hookEmailAvailability.isAvailable !== false,
      message: hookEmailAvailability.message,
    });
  }, [hookEmailAvailability]);

  const onSubmit = async (data: RegisterFormData) => {
    try {
      setError("");
      await registerUser({
        username: data.username,
        email: data.email,
        password: data.password,
      });
      // 注册成功后重定向到登录页面
      router.push("/login");
    } catch (error) {
      // 使用错误处理工具获取错误消息
      const errorMessage = getErrorMessage(error, "注册失败，请稍后重试");
      setError(errorMessage);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className={cn("w-full max-w-md space-y-8", className)}>
        <div className="text-center">
          <h2 className="mt-6 text-3xl font-bold tracking-tight text-gray-900">
            创建新账户
          </h2>
          <p className="mt-2 text-sm text-gray-600">
            或者{" "}
            <Link
              href="/login"
              className="font-medium text-blue-600 hover:text-blue-500"
            >
              登录已有账户
            </Link>
          </p>
        </div>

        <Card>
          <CardHeader>
            <CardTitle>注册信息</CardTitle>
            <CardDescription>请填写以下信息来创建您的账户</CardDescription>
          </CardHeader>
          <CardContent>
            {error && (
              <Alert variant="destructive" className="mb-6">
                <AlertCircle className="h-4 w-4" />
                <AlertDescription>{error}</AlertDescription>
              </Alert>
            )}

            <form onSubmit={handleSubmit(onSubmit)}>
              <div className="flex flex-col gap-6">
                <div className="grid gap-2">
                  <Label htmlFor="username">用户名</Label>
                  <div className="relative">
                    <Input
                      id="username"
                      placeholder="请输入用户名"
                      type="text"
                      autoCapitalize="none"
                      autoComplete="username"
                      autoCorrect="off"
                      {...register("username")}
                      className={cn(
                        "pr-10",
                        errors.username ? "border-red-500" : "",
                        !usernameValidation.isValid && !errors.username
                          ? "border-red-500"
                          : "",
                        usernameValidation.isValid &&
                          watchedUsername &&
                          watchedUsername.length >= 3 &&
                          !errors.username
                          ? "border-green-500"
                          : ""
                      )}
                    />
                    <div className="absolute right-3 top-1/2 -translate-y-1/2">
                      {usernameValidation.isValidating ? (
                        <Loader2 className="h-4 w-4 animate-spin text-gray-400" />
                      ) : watchedUsername && watchedUsername.length >= 3 ? (
                        usernameValidation.isValid ? (
                          <CheckCircle className="h-4 w-4 text-green-500" />
                        ) : (
                          <XCircle className="h-4 w-4 text-red-500" />
                        )
                      ) : null}
                    </div>
                  </div>
                  {errors.username && (
                    <p className="text-sm text-red-500 flex items-center gap-1">
                      <AlertCircle className="h-3 w-3" />
                      {errors.username.message}
                    </p>
                  )}
                  {!errors.username &&
                    !usernameValidation.isValidating &&
                    usernameValidation.message &&
                    !usernameValidation.isValid && (
                      <p className="text-sm text-red-500 flex items-center gap-1">
                        <AlertCircle className="h-3 w-3" />
                        {usernameValidation.message}
                      </p>
                    )}
                  {!errors.username &&
                    !usernameValidation.isValidating &&
                    usernameValidation.isValid &&
                    watchedUsername &&
                    watchedUsername.length >= 3 &&
                    usernameValidation.message && (
                      <p className="text-sm text-green-600 flex items-center gap-1">
                        <CheckCircle className="h-3 w-3" />
                        {usernameValidation.message}
                      </p>
                    )}
                </div>

                <div className="grid gap-2">
                  <Label htmlFor="email">邮箱</Label>
                  <div className="relative">
                    <Input
                      id="email"
                      placeholder="请输入邮箱地址"
                      type="email"
                      autoCapitalize="none"
                      autoComplete="email"
                      autoCorrect="off"
                      {...register("email")}
                      className={cn(
                        "pr-10",
                        errors.email ? "border-red-500" : "",
                        !emailValidation.isValid && !errors.email
                          ? "border-red-500"
                          : "",
                        emailValidation.isValid &&
                          watchedEmail &&
                          watchedEmail.includes("@") &&
                          !errors.email
                          ? "border-green-500"
                          : ""
                      )}
                    />
                    <div className="absolute right-3 top-1/2 -translate-y-1/2">
                      {emailValidation.isValidating ? (
                        <Loader2 className="h-4 w-4 animate-spin text-gray-400" />
                      ) : watchedEmail && watchedEmail.includes("@") ? (
                        emailValidation.isValid ? (
                          <CheckCircle className="h-4 w-4 text-green-500" />
                        ) : (
                          <XCircle className="h-4 w-4 text-red-500" />
                        )
                      ) : null}
                    </div>
                  </div>
                  {errors.email && (
                    <p className="text-sm text-red-500 flex items-center gap-1">
                      <AlertCircle className="h-3 w-3" />
                      {errors.email.message}
                    </p>
                  )}
                  {!errors.email &&
                    !emailValidation.isValidating &&
                    emailValidation.message &&
                    !emailValidation.isValid && (
                      <p className="text-sm text-red-500 flex items-center gap-1">
                        <AlertCircle className="h-3 w-3" />
                        {emailValidation.message}
                      </p>
                    )}
                  {!errors.email &&
                    !emailValidation.isValidating &&
                    emailValidation.isValid &&
                    watchedEmail &&
                    watchedEmail.includes("@") &&
                    emailValidation.message && (
                      <p className="text-sm text-green-600 flex items-center gap-1">
                        <CheckCircle className="h-3 w-3" />
                        {emailValidation.message}
                      </p>
                    )}
                </div>

                <div className="grid gap-2">
                  <Label htmlFor="password">密码</Label>
                  <div className="relative">
                    <Input
                      id="password"
                      type={showPassword ? "text" : "password"}
                      placeholder="请输入密码"
                      {...register("password")}
                      className={
                        errors.password ? "border-red-500 pr-10" : "pr-10"
                      }
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

                <div className="grid gap-2">
                  <Label htmlFor="confirmPassword">确认密码</Label>
                  <div className="relative">
                    <Input
                      id="confirmPassword"
                      type={showConfirmPassword ? "text" : "password"}
                      placeholder="请再次输入密码"
                      {...register("confirmPassword")}
                      className={
                        errors.confirmPassword
                          ? "border-red-500 pr-10"
                          : "pr-10"
                      }
                    />
                    <Button
                      type="button"
                      variant="ghost"
                      size="sm"
                      className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"
                      onClick={() =>
                        setShowConfirmPassword(!showConfirmPassword)
                      }
                    >
                      {showConfirmPassword ? (
                        <EyeOff className="h-4 w-4" />
                      ) : (
                        <Eye className="h-4 w-4" />
                      )}
                    </Button>
                  </div>
                  {errors.confirmPassword && (
                    <p className="text-sm text-red-500">
                      {errors.confirmPassword.message}
                    </p>
                  )}
                </div>

                <Button
                  type="submit"
                  className="w-full"
                  disabled={
                    isLoading ||
                    !usernameValidation.isValid ||
                    usernameValidation.isValidating ||
                    !emailValidation.isValid ||
                    emailValidation.isValidating
                  }
                >
                  {isLoading && (
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  )}
                  创建账户
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
