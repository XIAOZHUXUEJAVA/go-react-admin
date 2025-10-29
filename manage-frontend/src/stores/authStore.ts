import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";
import { AuthStore, LoginRequest, RegisterRequest } from "@/types/auth";
import { authApi, userApi } from "@/api";
import { toast } from "sonner";
import { ErrorHandler, parseError } from "@/lib/errorHandler";
import {
  getAccessToken,
  getRefreshToken,
  getTokenExpiresAt,
  setTokens,
  removeTokens,
  isTokenExpiringSoon,
} from "@/lib/tokenUtils";
import { usePermissionStore } from "./permissionStore";

/**
 * 认证状态管理 Store
 * 使用 Zustand 进行状态管理，支持持久化存储
 */
export const useAuthStore = create<AuthStore>()(
  persist(
    (set, get) => ({
      // 初始状态
      user: null,
      accessToken: null,
      refreshToken: null,
      tokenExpiresAt: null,
      isAuthenticated: false,
      isLoading: false,

      // 设置加载状态
      setLoading: (loading: boolean) => {
        set({ isLoading: loading });
      },

      // 用户登录
      login: async (credentials: LoginRequest) => {
        try {
          set({ isLoading: true });

          const response = await authApi.login(credentials);

          if (response.data) {
            const { access_token, refresh_token, expires_in, user } =
              response.data;

            // 计算token过期时间
            const tokenExpiresAt = Date.now() + expires_in * 1000;

            // 设置认证状态
            set({
              user,
              accessToken: access_token,
              refreshToken: refresh_token,
              tokenExpiresAt,
              isAuthenticated: true,
              isLoading: false,
            });

            // 保存tokens到localStorage
            setTokens(access_token, refresh_token, expires_in);

            // 加载用户权限和菜单
            try {
              const permissionStore = usePermissionStore.getState();
              await Promise.all([
                permissionStore.loadPermissions(),
                permissionStore.loadUserMenus(),
              ]);
            } catch (permError) {}

            toast.success("登录成功！正在跳转...");
            // 保持 loading 状态，直到 AuthGuard 完成重定向
          } else {
            set({ isLoading: false });
          }
        } catch (error) {
          set({ isLoading: false });

          // 使用错误处理工具解析错误
          const standardError = parseError(error, {
            defaultMessage: "登录失败，请稍后重试",
            logError: false,
          });

          // 根据错误类型和错误信息提供更精确的提示
          let errorMessage = standardError.message;
          let errorDescription = "如果问题持续存在，请联系技术支持";

          // 处理 API 错误
          if (ErrorHandler.isAPIError(error)) {
            // 验证码错误
            if (
              error.message?.includes("captcha") ||
              error.message?.includes("验证码") ||
              error.error === "invalid captcha"
            ) {
              errorMessage = "验证码错误，请重新输入";
              errorDescription = "验证码已刷新，请查看新的验证码";
            }
            // 用户名或密码错误
            else if (
              error.code === 401 &&
              (error.error === "invalid credentials" ||
                error.message?.includes("credentials"))
            ) {
              errorMessage = "用户名或密码错误";
              errorDescription = "请检查您的用户名和密码后重试";
            }
            // 其他 401 错误
            else if (error.code === 401) {
              errorMessage = "认证失败";
              errorDescription = "请检查您的登录信息";
            }
            // 请求参数错误
            else if (error.code === 400) {
              errorMessage = "请求参数错误";
              errorDescription = "请检查输入信息是否完整";
            }
            // 请求过于频繁
            else if (error.code === 429) {
              errorMessage = "登录尝试过于频繁";
              errorDescription = "请稍后再试，或联系管理员";
            }
            // 服务器错误
            else if (error.code && error.code >= 500) {
              errorMessage = "服务器错误";
              errorDescription = "请稍后重试或联系技术支持";
            }
          }

          toast.error(errorMessage, {
            description: errorDescription,
            duration: 4000,
          });
          throw error;
        }
      },

      // 用户注册
      register: async (data: RegisterRequest) => {
        try {
          set({ isLoading: true });

          const response = await authApi.register(data);

          if (response.data) {
            toast.success("注册成功！请登录");
            set({ isLoading: false });
          }
        } catch (error) {
          set({ isLoading: false });

          // 使用错误处理工具
          const standardError = parseError(error, {
            defaultMessage: "注册失败，请稍后重试",
          });

          let errorMessage = standardError.message;

          // 处理常见注册错误
          if (ErrorHandler.isAPIError(error)) {
            if (
              error.message?.includes("username already exists") ||
              error.message?.includes("用户名已存在")
            ) {
              errorMessage = "用户名已被使用，请选择其他用户名";
            } else if (
              error.message?.includes("email already exists") ||
              error.message?.includes("邮箱已存在")
            ) {
              errorMessage = "邮箱已被注册，请使用其他邮箱";
            }
          }

          toast.error(errorMessage);
          throw error;
        }
      },

      // 检查认证状态
      checkAuth: async () => {
        try {
          // 从localStorage获取最新的token信息
          const accessToken = getAccessToken();
          const refreshToken = getRefreshToken();
          const tokenExpiresAt = getTokenExpiresAt();
          // 如果没有 access token，直接设置为未认证状态
          if (!accessToken) {
            set({
              user: null,
              accessToken: null,
              refreshToken: null,
              tokenExpiresAt: null,
              isAuthenticated: false,
              isLoading: false,
            });
            return;
          }

          // 检查token是否即将过期（提前5分钟刷新）
          if (isTokenExpiringSoon() && refreshToken) {
            try {
              const refreshResponse = await authApi.refreshToken(refreshToken);
              if (refreshResponse.data) {
                const { access_token, expires_in } = refreshResponse.data;
                const newTokenExpiresAt = Date.now() + expires_in * 1000;

                // 更新store状态
                set({
                  accessToken: access_token,
                  tokenExpiresAt: newTokenExpiresAt,
                });

                // 更新localStorage
                setTokens(access_token, refreshToken, expires_in);
              }
            } catch (refreshError) {
              // 刷新失败，清除认证状态
              removeTokens();
              set({
                user: null,
                accessToken: null,
                refreshToken: null,
                tokenExpiresAt: null,
                isAuthenticated: false,
                isLoading: false,
              });
              return;
            }
          }

          // 同步store状态与localStorage
          set({
            accessToken,
            refreshToken,
            tokenExpiresAt,
          });

          set({ isLoading: true });

          const response = await userApi.getCurrentUser();

          if (response.data) {
            set({
              user: response.data,
              isAuthenticated: true,
              isLoading: false,
            });

            // 同步权限数据
            try {
              const permissionStore = usePermissionStore.getState();
              // 如果权限未加载，则加载权限
              if (!permissionStore.isLoaded) {
                await Promise.all([
                  permissionStore.loadPermissions(),
                  permissionStore.loadUserMenus(),
                ]);
                console.log("✅ CheckAuth - 权限和菜单同步成功");
              }
            } catch (permError) {
              // 权限同步失败不影响认证流程
            }
          } else {
            set({
              user: null,
              accessToken: null,
              refreshToken: null,
              tokenExpiresAt: null,
              isAuthenticated: false,
              isLoading: false,
            });
          }
        } catch (error) {
          // Token 无效，清除认证状态
          set({
            user: null,
            accessToken: null,
            refreshToken: null,
            tokenExpiresAt: null,
            isAuthenticated: false,
            isLoading: false,
          });

          removeTokens();
        }
      },

      // 用户登出
      logout: async () => {
        const { refreshToken } = get();

        try {
          // 等待后端登出接口完成
          await authApi.logout(refreshToken ?? undefined);
        } catch (error) {
          // 忽略登出接口错误，继续清除本地状态
        } finally {
          // 清除本地状态
          set({
            user: null,
            accessToken: null,
            refreshToken: null,
            tokenExpiresAt: null,
            isAuthenticated: false,
            isLoading: false,
          });

          // 清除权限数据
          const permissionStore = usePermissionStore.getState();
          permissionStore.clearPermissions();

          // 清除本地存储并立即跳转
          removeTokens();
          if (typeof window !== "undefined") {
            window.location.href = "/login";
          }

          toast.success("已退出登录");
        }
      },
    }),
    {
      name: "auth-storage",
      storage: createJSONStorage(() => localStorage),
      // 只持久化必要的字段
      partialize: (state) => ({
        user: state.user,
        accessToken: state.accessToken,
        refreshToken: state.refreshToken,
        tokenExpiresAt: state.tokenExpiresAt,
        isAuthenticated: state.isAuthenticated,
      }),
    }
  )
);
