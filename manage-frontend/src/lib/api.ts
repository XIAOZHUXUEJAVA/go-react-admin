import axios, {
  AxiosInstance,
  AxiosResponse,
  AxiosError,
  InternalAxiosRequestConfig,
} from "axios";
import { APIResponse, APIError } from "@/types/api";
import { RefreshTokenResponse } from "@/types/auth";

// 创建 axios 实例
const apiClient: AxiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || "http://localhost:9000/api/v1",
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});

// Token刷新状态管理
let isRefreshing = false;
let failedQueue: Array<{
  resolve: (value?: string | null) => void;
  reject: (reason?: unknown) => void;
}> = [];

const processQueue = (error: unknown, token: string | null = null) => {
  failedQueue.forEach(({ resolve, reject }) => {
    if (error) {
      reject(error);
    } else {
      resolve(token);
    }
  });

  failedQueue = [];
};

// 请求拦截器 - 添加认证 token
apiClient.interceptors.request.use(
  (config) => {
    // 从 localStorage 获取 access token
    if (typeof window !== "undefined") {
      const accessToken = localStorage.getItem("access-token");
      if (accessToken) {
        config.headers.Authorization = `Bearer ${accessToken}`;
      }
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器 - 统一处理响应和token刷新
apiClient.interceptors.response.use(
  (response: AxiosResponse<APIResponse>) => {
    // 成功响应直接返回
    return response;
  },
  async (error: AxiosError<APIResponse>) => {
    const originalRequest = error.config as InternalAxiosRequestConfig & {
      _retry?: boolean;
    };

    // 401 未授权处理 - 但排除登录和注册接口
    if (error.response?.status === 401 && !originalRequest._retry) {
      // 如果是登录或注册接口的401错误，直接抛出，让业务逻辑处理
      const isAuthEndpoint =
        originalRequest.url?.includes("/auth/login") ||
        originalRequest.url?.includes("/auth/register");

      if (isAuthEndpoint) {
        // 统一错误处理
        const apiError: APIError = {
          code: error.response?.data?.code || error.response?.status || 500,
          message: error.response?.data?.message || error.message || "请求失败",
          error: error.response?.data?.error,
        };
        return Promise.reject(apiError);
      }
      if (isRefreshing) {
        // 如果正在刷新token，将请求加入队列
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        })
          .then((token) => {
            originalRequest.headers.Authorization = `Bearer ${token}`;
            return apiClient(originalRequest);
          })
          .catch((err) => {
            return Promise.reject(err);
          });
      }

      originalRequest._retry = true;
      isRefreshing = true;

      if (typeof window !== "undefined") {
        const refreshToken = localStorage.getItem("refresh-token");

        if (refreshToken) {
          try {
            // 尝试刷新token
            const response = await axios.post<
              APIResponse<RefreshTokenResponse>
            >(
              `${
                process.env.NEXT_PUBLIC_API_URL ||
                "http://localhost:9000/api/v1"
              }/auth/refresh`,
              { refresh_token: refreshToken }
            );

            if (response.data.data) {
              const { access_token, expires_in } = response.data.data;

              // 保存新的access token
              localStorage.setItem("access-token", access_token);

              // 计算过期时间
              const expiresAt = Date.now() + expires_in * 1000;
              localStorage.setItem("token-expires-at", expiresAt.toString());

              // 更新默认请求头
              apiClient.defaults.headers.common.Authorization = `Bearer ${access_token}`;
              originalRequest.headers.Authorization = `Bearer ${access_token}`;

              processQueue(null, access_token);

              return apiClient(originalRequest);
            }
          } catch (refreshError) {
            processQueue(refreshError, null);

            // 刷新失败，清除所有认证信息
            localStorage.removeItem("access-token");
            localStorage.removeItem("refresh-token");
            localStorage.removeItem("token-expires-at");
            localStorage.removeItem("auth-storage");

            // 重定向到登录页面
            window.location.href = "/login";

            return Promise.reject(refreshError);
          } finally {
            isRefreshing = false;
          }
        } else {
          // 没有refresh token，直接清除认证信息
          localStorage.removeItem("access-token");
          localStorage.removeItem("refresh-token");
          localStorage.removeItem("token-expires-at");
          localStorage.removeItem("auth-storage");

          window.location.href = "/login";
        }
      }
    }

    // 统一错误处理
    const apiError: APIError = {
      code: error.response?.data?.code || error.response?.status || 500,
      message: error.response?.data?.message || error.message || "请求失败",
      error: error.response?.data?.error,
    };

    return Promise.reject(apiError);
  }
);

// 通用 API 请求方法
export class ApiService {
  /**
   * GET 请求
   */
  static async get<T>(
    url: string,
    params?: Record<string, unknown>
  ): Promise<APIResponse<T>> {
    const response = await apiClient.get<APIResponse<T>>(url, { params });
    return response.data;
  }

  /**
   * POST 请求
   */
  static async post<T>(url: string, data?: unknown): Promise<APIResponse<T>> {
    const response = await apiClient.post<APIResponse<T>>(url, data);
    return response.data;
  }

  /**
   * PUT 请求
   */
  static async put<T>(url: string, data?: unknown): Promise<APIResponse<T>> {
    const response = await apiClient.put<APIResponse<T>>(url, data);
    return response.data;
  }

  /**
   * DELETE 请求
   */
  static async delete<T>(url: string): Promise<APIResponse<T>> {
    const response = await apiClient.delete<APIResponse<T>>(url);
    return response.data;
  }
}

export default apiClient;
