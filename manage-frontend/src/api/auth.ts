/**
 * 认证相关 API
 */
import { ApiService } from "@/lib/api";
import { APIResponse } from "@/types/common";
import {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  AuthUser,
  RefreshTokenRequest,
  RefreshTokenResponse,
  LogoutRequest,
  CaptchaResponse,
} from "@/types/auth";

export const authApi = {
  /**
   * 用户登录
   */
  login: async (
    credentials: LoginRequest
  ): Promise<APIResponse<LoginResponse>> => {
    return ApiService.post<LoginResponse>("/auth/login", credentials);
  },

  /**
   * 用户注册
   */
  register: async (data: RegisterRequest): Promise<APIResponse<AuthUser>> => {
    return ApiService.post<AuthUser>("/auth/register", data);
  },

  /**
   * 刷新访问令牌
   */
  refreshToken: async (
    refreshToken: string
  ): Promise<APIResponse<RefreshTokenResponse>> => {
    const request: RefreshTokenRequest = { refresh_token: refreshToken };
    return ApiService.post<RefreshTokenResponse>("/auth/refresh", request);
  },

  /**
   * 用户登出
   */
  logout: async (refreshToken?: string): Promise<APIResponse<null>> => {
    const request: LogoutRequest = refreshToken
      ? { refresh_token: refreshToken }
      : {};
    return ApiService.post<null>("/auth/logout", request);
  },

  /**
   * 验证令牌有效性
   */
  validateToken: async (): Promise<APIResponse<{ valid: boolean }>> => {
    return ApiService.get<{ valid: boolean }>("/auth/validate");
  },

  /**
   * 发送密码重置邮件
   */
  forgotPassword: async (
    email: string
  ): Promise<APIResponse<{ message: string }>> => {
    return ApiService.post<{ message: string }>("/auth/forgot-password", {
      email,
    });
  },

  /**
   * 重置密码
   */
  resetPassword: async (
    token: string,
    newPassword: string
  ): Promise<APIResponse<{ message: string }>> => {
    return ApiService.post<{ message: string }>("/auth/reset-password", {
      token,
      new_password: newPassword,
    });
  },

  /**
   * 生成验证码
   */
  generateCaptcha: async (): Promise<APIResponse<CaptchaResponse>> => {
    return ApiService.get<CaptchaResponse>("/auth/captcha");
  },
} as const;
