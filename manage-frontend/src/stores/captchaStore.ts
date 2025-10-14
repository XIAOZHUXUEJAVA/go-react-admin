import { create } from "zustand";
import { CaptchaStore } from "@/types/auth";
import { authApi } from "@/api";
import { toast } from "sonner";
import { APIError } from "@/types/common";

/**
 * 验证码状态管理 Store
 * 使用 Zustand 进行状态管理
 */
export const useCaptchaStore = create<CaptchaStore>((set, get) => ({
  // 初始状态
  captchaId: null,
  captchaImage: null,
  isLoading: false,
  error: null,

  // 设置加载状态
  setLoading: (loading: boolean) => {
    set({ isLoading: loading });
  },

  // 设置错误信息
  setError: (error: string | null) => {
    set({ error });
  },

  // 生成验证码
  generateCaptcha: async () => {
    try {
      set({ isLoading: true, error: null });

      const response = await authApi.generateCaptcha();

      if (response.data) {
        const { captcha_id, captcha_data } = response.data;

        set({
          captchaId: captcha_id,
          captchaImage: captcha_data,
          isLoading: false,
          error: null,
        });
      } else {
        set({
          isLoading: false,
          error: "验证码生成失败",
        });
      }
    } catch (error) {
      let errorMessage = "验证码生成失败，请稍后重试";
      const apiError = error as APIError;

      if (apiError.message) {
        errorMessage = apiError.message;
      } else if (apiError.code === 500) {
        errorMessage = "服务器错误，请稍后重试";
      }

      set({
        isLoading: false,
        error: errorMessage,
        captchaId: null,
        captchaImage: null,
      });

      toast.error(errorMessage);
      throw error;
    }
  },

  // 刷新验证码
  refreshCaptcha: async () => {
    const { generateCaptcha } = get();
    await generateCaptcha();
  },

  // 清除验证码
  clearCaptcha: () => {
    set({
      captchaId: null,
      captchaImage: null,
      error: null,
    });
  },
}));
