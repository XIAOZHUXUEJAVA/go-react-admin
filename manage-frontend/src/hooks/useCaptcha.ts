"use client";

import { useState, useEffect } from "react";
import { useCaptchaStore } from "@/stores/captchaStore";

interface UseCaptchaOptions {
  autoGenerate?: boolean; // 是否自动生成验证码
  clearOnSuccess?: boolean; // 成功后是否清除验证码
}

interface UseCaptchaReturn {
  // 验证码状态
  captchaId: string | null;
  captchaImage: string | null;
  isLoading: boolean;
  error: string | null;

  // 验证码输入
  captchaCode: string;
  setCaptchaCode: (code: string) => void;

  // 验证码操作
  generateCaptcha: () => Promise<void>;
  refreshCaptcha: () => Promise<void>;
  clearCaptcha: () => void;

  // 验证码验证
  isValid: boolean;
  hasRequiredData: boolean;
}

/**
 * 验证码 Hook
 * 提供验证码相关的状态和操作方法
 */
export const useCaptcha = (
  options: UseCaptchaOptions = {}
): UseCaptchaReturn => {
  const { autoGenerate = true, clearOnSuccess = true } = options;

  const {
    captchaId,
    captchaImage,
    isLoading,
    error,
    generateCaptcha,
    refreshCaptcha,
    clearCaptcha,
  } = useCaptchaStore();

  const [captchaCode, setCaptchaCode] = useState("");

  // 自动生成验证码
  useEffect(() => {
    if (autoGenerate && !captchaImage && !isLoading) {
      generateCaptcha();
    }
  }, [autoGenerate, captchaImage, isLoading, generateCaptcha]);

  // 验证码是否有效（有ID和用户输入了验证码）
  const isValid = Boolean(captchaId && captchaCode.trim());

  // 是否有必需的数据（验证码ID和图片）
  const hasRequiredData = Boolean(captchaId && captchaImage);

  // 包装刷新方法，清空输入
  const handleRefreshCaptcha = async () => {
    setCaptchaCode("");
    await refreshCaptcha();
  };

  // 包装清除方法
  const handleClearCaptcha = () => {
    setCaptchaCode("");
    clearCaptcha();
  };

  return {
    // 验证码状态
    captchaId,
    captchaImage,
    isLoading,
    error,

    // 验证码输入
    captchaCode,
    setCaptchaCode,

    // 验证码操作
    generateCaptcha,
    refreshCaptcha: handleRefreshCaptcha,
    clearCaptcha: handleClearCaptcha,

    // 验证码验证
    isValid,
    hasRequiredData,
  };
};

/**
 * 检查是否需要验证码的 Hook
 * 可以根据后端配置或其他条件来决定是否显示验证码
 */
export const useCaptchaRequired = (): boolean => {
  // 这里可以根据实际需求来实现
  // 例如：从配置文件读取、根据登录失败次数判断等

  // 目前默认总是需要验证码
  // 在实际项目中，可以通过以下方式来控制：
  // 1. 从环境变量读取
  // 2. 从后端API获取配置
  // 3. 根据用户登录失败次数判断

  return true;
};

export default useCaptcha;
