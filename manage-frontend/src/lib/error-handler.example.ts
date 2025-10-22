/**
 * 错误处理工具使用示例
 * 这个文件展示了如何在不同场景下使用 error-handler
 *
 * 注意：这是示例文件，包含未使用的函数和变量，这是正常的
 * eslint-disable 用于忽略示例代码的 lint 警告
 */

/* eslint-disable @typescript-eslint/no-unused-vars */

import {
  ErrorHandler,
  getErrorMessage,
  parseError,
  withRetry,
} from "./errorHandler";
import { toast } from "sonner";
import { CreateUserRequest, UpdateUserRequest, User } from "@/types/api";

// ============================================
// 示例 1: 基础用法 - 简单获取错误消息
// ============================================
async function example1_BasicUsage() {
  try {
    // 某个可能抛出错误的操作
    await someApiCall();
  } catch (error) {
    // 最简单的用法：直接获取错误消息
    const message = getErrorMessage(error, "创建用户失败");
    toast.error(message);
  }
}

// ============================================
// 示例 2: 详细错误信息 - 获取完整的错误对象
// ============================================
async function example2_DetailedError() {
  try {
    await someApiCall();
  } catch (error) {
    // 获取标准化的错误对象
    const standardError = parseError(error, {
      defaultMessage: "操作失败",
      logError: true, // 在控制台输出错误
    });

    // 可以根据错误类型做不同处理
    switch (standardError.type) {
      case "API_ERROR":
        toast.error(standardError.message);
        break;
      case "NETWORK_ERROR":
        toast.error("网络连接失败，请检查网络");
        break;
      case "TIMEOUT_ERROR":
        toast.error("请求超时，请重试");
        break;
      default:
        toast.error(standardError.message);
    }

    // 可以访问错误码
    if (standardError.code === 401) {
      // 跳转到登录页
      window.location.href = "/login";
    }
  }
}

// ============================================
// 示例 3: React 组件中使用（推荐用法）
// ============================================
function Example3_ReactComponent() {
  const handleCreateUser = async (userData: CreateUserRequest) => {
    try {
      const response = await userApi.createUser(userData);
      toast.success("用户创建成功");
      return response;
    } catch (error) {
      // 一行代码搞定错误处理
      toast.error(getErrorMessage(error, "创建用户失败"));
      throw error; // 如果需要，可以继续抛出
    }
  };

  const handleUpdateUser = async (
    userId: number,
    userData: UpdateUserRequest
  ) => {
    try {
      const response = await userApi.updateUser(userId, userData);
      toast.success("用户更新成功");
      return response;
    } catch (error) {
      toast.error(getErrorMessage(error, "更新用户失败"));
    }
  };

  const handleDeleteUser = async (userId: number) => {
    try {
      const response = await userApi.deleteUser(userId);
      toast.success("用户删除成功");
      return response;
    } catch (error) {
      toast.error(getErrorMessage(error, "删除用户失败"));
    }
  };

  return null; // JSX...
}

// ============================================
// 示例 4: 使用类型守卫进行精细控制
// ============================================
async function example4_TypeGuards() {
  try {
    await someApiCall();
  } catch (error) {
    // 使用类型守卫判断错误类型
    if (ErrorHandler.isAPIError(error)) {
      // 这里 TypeScript 知道 error 是 APIError 类型
      console.log("API 错误码:", error.code);
      console.log("API 错误消息:", error.message);
    } else if (ErrorHandler.isNetworkError(error)) {
      toast.error("网络连接失败");
    } else if (ErrorHandler.isTimeoutError(error)) {
      toast.error("请求超时");
    } else {
      toast.error(getErrorMessage(error));
    }
  }
}

// ============================================
// 示例 5: 带重试机制的请求
// ============================================
async function example5_WithRetry() {
  try {
    // 自动重试最多 3 次，每次间隔 1 秒
    const data = await withRetry(
      () => userApi.getUsers({ page: 1, pageSize: 10 }),
      {
        maxRetries: 3,
        retryDelay: 1000,
        // 自定义重试条件
        shouldRetry: (error) => {
          return error.type === "NETWORK_ERROR" || error.code === 503;
        },
      }
    );
    return data;
  } catch (error) {
    toast.error(getErrorMessage(error, "获取用户列表失败"));
  }
}

// ============================================
// 示例 6: 批量操作错误处理
// ============================================
async function example6_BatchOperations() {
  const userIds = [1, 2, 3, 4, 5];

  // 使用 Promise.allSettled 执行批量操作
  const results = await Promise.allSettled(
    userIds.map((id) => userApi.deleteUser(id))
  );

  // 使用工具函数处理结果
  const { successes, errors } = ErrorHandler.handleSettled(results);

  if (successes.length > 0) {
    toast.success(`成功删除 ${successes.length} 个用户`);
  }

  if (errors.length > 0) {
    toast.error(`${errors.length} 个用户删除失败`);
    // 可以显示详细的错误信息
    errors.forEach((error) => {
      console.error(error.message);
    });
  }
}

// ============================================
// 示例 7: 自定义错误消息映射
// ============================================
async function example7_CustomMessages() {
  try {
    await userApi.createUser({
      username: "test",
      email: "test@example.com",
      password: "password123",
    });
  } catch (error) {
    const standardError = parseError(error, {
      defaultMessage: "创建用户失败",
      customMessages: {
        409: "用户名已存在，请使用其他用户名",
        422: "用户信息格式不正确，请检查输入",
        400: "请求参数错误，请检查必填项",
      },
    });

    toast.error(standardError.message);
  }
}

// ============================================
// 示例 8: 创建可复用的错误处理器
// ============================================
const handleUserOperation = <T>(operationName: string) => {
  return async (fn: () => Promise<T>): Promise<T> => {
    try {
      const result = await fn();
      toast.success(`${operationName}成功`);
      return result;
    } catch (error) {
      const message = getErrorMessage(error, `${operationName}失败`);
      toast.error(message);
      throw error;
    }
  };
};

// 使用
async function example8_ReusableHandler() {
  const handleCreate = handleUserOperation("创建用户");
  const handleUpdate = handleUserOperation("更新用户");
  const handleDelete = handleUserOperation("删除用户");

  await handleCreate(() =>
    userApi.createUser({
      username: "test",
      email: "test@example.com",
      password: "password",
    })
  );
  await handleUpdate(() =>
    userApi.updateUser(1, {
      id: 1,
      username: "test2",
      email: "test2@example.com",
    })
  );
  await handleDelete(() => userApi.deleteUser(1));
}

// ============================================
// 示例 9: 在 Zustand Store 中使用
// ============================================
interface UserStore {
  users: User[];
  loading: boolean;
  error: string | null;
  fetchUsers: () => Promise<void>;
}

const useUserStore = () => {
  // 模拟 zustand store
  const fetchUsers = async () => {
    try {
      // set({ loading: true, error: null });
      const response = await userApi.getUsers({ page: 1, pageSize: 10 });
      // set({ users: response.data, loading: false });
    } catch (error) {
      const errorMessage = getErrorMessage(error, "获取用户列表失败");
      // set({ error: errorMessage, loading: false });
      toast.error(errorMessage);
    }
  };

  return { fetchUsers };
};

// ============================================
// 示例 10: 全局错误边界（配合 React Error Boundary）
// ============================================
function GlobalErrorHandler(error: Error, errorInfo: React.ErrorInfo) {
  const standardError = parseError(error, {
    defaultMessage: "应用程序发生错误",
    logError: true,
  });

  // 上报错误到监控服务
  // reportErrorToService(standardError);

  // 显示用户友好的错误消息
  toast.error(standardError.message);

  // 可以记录 errorInfo 用于调试
  console.error("Component stack:", errorInfo.componentStack);
}

// ============================================
// 模拟的 API 调用（仅用于示例）
// ============================================
const someApiCall = async (): Promise<void> => {
  throw new Error("API Error");
};

const userApi = {
  createUser: async (data: CreateUserRequest) => ({ code: 201, data }),
  updateUser: async (id: number, data: UpdateUserRequest) => ({
    code: 200,
    data,
  }),
  deleteUser: async (id: number) => ({ code: 200, data: null }),
  getUsers: async (params: { page: number; pageSize: number }) => ({
    code: 200,
    data: [] as User[],
  }),
};

export {
  example1_BasicUsage,
  example2_DetailedError,
  Example3_ReactComponent,
  example4_TypeGuards,
  example5_WithRetry,
  example6_BatchOperations,
  example7_CustomMessages,
  example8_ReusableHandler,
  useUserStore,
  GlobalErrorHandler,
};
