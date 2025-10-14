"use client";

import React, { useEffect, useCallback } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { RefreshCw, AlertCircle } from "lucide-react";
import { useCaptchaStore } from "@/stores/captchaStore";
import { cn } from "@/lib/utils";

interface CaptchaProps {
  value: string;
  onChange: (value: string) => void;
  error?: string;
  disabled?: boolean;
  className?: string;
  required?: boolean;
  placeholder?: string;
}

/**
 * 验证码组件
 * 包含验证码图片显示、输入框和刷新按钮
 */
export const Captcha: React.FC<CaptchaProps> = ({
  value,
  onChange,
  error,
  disabled = false,
  className,
  required = false,
  placeholder = "请输入验证码",
}) => {
  const {
    captchaId,
    captchaImage,
    isLoading,
    error: captchaError,
    generateCaptcha,
    refreshCaptcha,
  } = useCaptchaStore();

  // 组件挂载时生成验证码
  useEffect(() => {
    if (!captchaImage) {
      generateCaptcha();
    }
  }, [captchaImage, generateCaptcha]);

  // 刷新验证码的处理函数
  const handleRefresh = useCallback(async () => {
    try {
      await refreshCaptcha();
      // 清空输入框
      onChange("");
    } catch (error) {}
  }, [refreshCaptcha, onChange]);

  // 输入变化处理
  const handleInputChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      onChange(e.target.value);
    },
    [onChange]
  );

  // 显示的错误信息（优先显示表单验证错误）
  const displayError = error || captchaError;

  return (
    <div className={cn("space-y-2", className)}>
      <Label htmlFor="captcha" className="text-sm font-medium">
        验证码 {required && <span className="text-red-500">*</span>}
      </Label>

      <div className="flex items-center space-x-2">
        {/* 验证码图片 */}
        <div className="relative flex-shrink-0">
          {isLoading ? (
            <div className="flex items-center justify-center w-32 h-12 bg-gray-100 border border-gray-300 rounded">
              <RefreshCw className="w-4 h-4 animate-spin text-gray-500" />
            </div>
          ) : captchaImage ? (
            <img
              src={captchaImage}
              alt="验证码"
              className="w-32 h-12 border border-gray-300 rounded cursor-pointer hover:border-gray-400 transition-colors"
              onClick={handleRefresh}
              title="点击刷新验证码"
            />
          ) : (
            <div className="flex items-center justify-center w-32 h-12 bg-gray-100 border border-gray-300 rounded">
              <span className="text-xs text-gray-500">加载中...</span>
            </div>
          )}
        </div>

        {/* 验证码输入框 */}
        <div className="flex-1">
          <Input
            id="captcha"
            type="text"
            value={value}
            onChange={handleInputChange}
            placeholder={placeholder}
            disabled={disabled || isLoading}
            className={cn(
              "transition-colors",
              displayError && "border-red-500 focus:border-red-500"
            )}
            maxLength={6}
            autoComplete="off"
          />
        </div>

        {/* 刷新按钮 */}
        <Button
          type="button"
          variant="outline"
          size="sm"
          onClick={handleRefresh}
          disabled={disabled || isLoading}
          className="flex-shrink-0"
          title="刷新验证码"
        >
          <RefreshCw className={cn("w-4 h-4", isLoading && "animate-spin")} />
        </Button>
      </div>

      {/* 错误信息 */}
      {displayError && (
        <div className="flex items-center space-x-1 text-sm text-red-600">
          <AlertCircle className="w-4 h-4" />
          <span>{displayError}</span>
        </div>
      )}

      {/* 提示信息 */}
      {!displayError && (
        <p className="text-xs text-gray-500">
          看不清？点击图片或刷新按钮重新获取
        </p>
      )}
    </div>
  );
};

export default Captcha;
