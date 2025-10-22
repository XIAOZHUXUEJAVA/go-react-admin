/**
 * 通用错误处理工具
 * 提供类型安全的错误处理和用户友好的错误消息提取
 */

import { APIError } from "@/types/common";

/**
 * 错误类型枚举
 */
export enum ErrorType {
  API_ERROR = "API_ERROR",
  NETWORK_ERROR = "NETWORK_ERROR",
  VALIDATION_ERROR = "VALIDATION_ERROR",
  TIMEOUT_ERROR = "TIMEOUT_ERROR",
  UNKNOWN_ERROR = "UNKNOWN_ERROR",
}

/**
 * 标准化的错误信息接口
 */
export interface StandardError {
  type: ErrorType;
  message: string;
  code?: number;
  details?: unknown;
  originalError?: unknown;
}

/**
 * 错误处理配置选项
 */
export interface ErrorHandlerOptions {
  /** 默认错误消息 */
  defaultMessage?: string;
  /** 是否在控制台输出错误（开发环境建议开启） */
  logError?: boolean;
  /** 错误消息前缀 */
  messagePrefix?: string;
  /** 自定义错误消息映射 */
  customMessages?: Record<number | string, string>;
}

/**
 * 类型守卫：检查是否为 APIError
 */
export function isAPIError(error: unknown): error is APIError {
  return (
    error !== null &&
    typeof error === "object" &&
    "code" in error &&
    "message" in error &&
    typeof (error as APIError).message === "string"
  );
}

/**
 * 类型守卫：检查是否为标准 Error 对象
 */
export function isError(error: unknown): error is Error {
  return error instanceof Error;
}

/**
 * 类型守卫：检查是否为网络错误
 */
export function isNetworkError(error: unknown): boolean {
  if (isError(error)) {
    return (
      error.message.includes("fetch") ||
      error.message.includes("network") ||
      error.message.includes("Network") ||
      error.name === "NetworkError" ||
      error.name === "TypeError"
    );
  }
  return false;
}

/**
 * 类型守卫：检查是否为超时错误
 */
export function isTimeoutError(error: unknown): boolean {
  if (isError(error)) {
    return (
      error.message.includes("timeout") ||
      error.message.includes("Timeout") ||
      error.name === "TimeoutError"
    );
  }
  if (isAPIError(error)) {
    return error.code === 408 || error.message.includes("超时");
  }
  return false;
}

/**
 * 根据错误码获取友好的错误消息
 */
export function getMessageByCode(code: number): string | null {
  const codeMessages: Record<number, string> = {
    400: "请求参数错误",
    401: "未授权，请重新登录",
    403: "没有权限访问该资源",
    404: "请求的资源不存在",
    408: "请求超时，请稍后重试",
    409: "资源冲突，请检查数据",
    422: "数据验证失败",
    429: "请求过于频繁，请稍后再试",
    500: "服务器内部错误",
    502: "网关错误",
    503: "服务暂时不可用",
    504: "网关超时",
  };

  return codeMessages[code] || null;
}

/**
 * 解析错误并返回标准化的错误信息
 */
export function parseError(
  error: unknown,
  options: ErrorHandlerOptions = {}
): StandardError {
  const {
    defaultMessage = "操作失败，请稍后重试",
    logError = process.env.NODE_ENV === "development",
  } = options;

  // 开发环境下输出错误日志
  if (logError) {
    console.error("Error caught:", error);
  }

  // 1. 处理 APIError
  if (isAPIError(error)) {
    const codeMessage = error.code ? getMessageByCode(error.code) : null;
    return {
      type: ErrorType.API_ERROR,
      message: error.message || codeMessage || defaultMessage,
      code: error.code,
      details: error.error,
      originalError: error,
    };
  }

  // 2. 处理超时错误
  if (isTimeoutError(error)) {
    return {
      type: ErrorType.TIMEOUT_ERROR,
      message: "请求超时，请检查网络连接后重试",
      originalError: error,
    };
  }

  // 3. 处理网络错误
  if (isNetworkError(error)) {
    return {
      type: ErrorType.NETWORK_ERROR,
      message: "网络连接失败，请检查您的网络设置",
      originalError: error,
    };
  }

  // 4. 处理标准 Error 对象
  if (isError(error)) {
    return {
      type: ErrorType.UNKNOWN_ERROR,
      message: error.message || defaultMessage,
      originalError: error,
    };
  }

  // 5. 处理字符串类型的错误
  if (typeof error === "string") {
    return {
      type: ErrorType.UNKNOWN_ERROR,
      message: error || defaultMessage,
      originalError: error,
    };
  }

  // 6. 其他未知错误
  return {
    type: ErrorType.UNKNOWN_ERROR,
    message: defaultMessage,
    originalError: error,
  };
}

/**
 * 获取错误消息（简化版本，直接返回字符串）
 */
export function getErrorMessage(
  error: unknown,
  defaultMessage = "操作失败，请稍后重试"
): string {
  const standardError = parseError(error, { defaultMessage, logError: false });
  return standardError.message;
}

/**
 * 创建错误处理器（高阶函数）
 * 用于包装异步函数，自动处理错误
 */
export function createErrorHandler<
  TArgs extends unknown[],
  TReturn
>(
  fn: (...args: TArgs) => Promise<TReturn>,
  options: ErrorHandlerOptions = {}
): (...args: TArgs) => Promise<TReturn> {
  return async (...args: TArgs) => {
    try {
      return await fn(...args);
    } catch (error) {
      const standardError = parseError(error, options);
      
      // 可以在这里添加全局错误处理逻辑
      // 例如：显示 toast、上报错误等
      
      throw standardError;
    }
  };
}

/**
 * 错误处理装饰器工厂（用于类方法）
 * 注意：需要启用 TypeScript 的 experimentalDecorators
 */
export function HandleError(options: ErrorHandlerOptions = {}) {
  return function <T extends (...args: unknown[]) => Promise<unknown>>(
    target: object,
    propertyKey: string,
    descriptor: TypedPropertyDescriptor<T>
  ): TypedPropertyDescriptor<T> {
    const originalMethod = descriptor.value;

    if (!originalMethod) {
      return descriptor;
    }

    descriptor.value = (async function (this: unknown, ...args: unknown[]) {
      try {
        return await originalMethod.apply(this, args as Parameters<T>);
      } catch (error) {
        const standardError = parseError(error, options);
        throw standardError;
      }
    }) as T;

    return descriptor;
  };
}

/**
 * 批量错误处理（用于 Promise.allSettled）
 */
export function handleSettledResults<T>(
  results: PromiseSettledResult<T>[],
  options: ErrorHandlerOptions = {}
): {
  successes: T[];
  errors: StandardError[];
} {
  const successes: T[] = [];
  const errors: StandardError[] = [];

  results.forEach((result) => {
    if (result.status === "fulfilled") {
      successes.push(result.value);
    } else {
      errors.push(parseError(result.reason, options));
    }
  });

  return { successes, errors };
}

/**
 * 重试机制包装器
 */
export async function withRetry<T>(
  fn: () => Promise<T>,
  options: {
    maxRetries?: number;
    retryDelay?: number;
    shouldRetry?: (error: StandardError) => boolean;
  } = {}
): Promise<T> {
  const {
    maxRetries = 3,
    retryDelay = 1000,
    shouldRetry = (error) =>
      error.type === ErrorType.NETWORK_ERROR ||
      error.type === ErrorType.TIMEOUT_ERROR,
  } = options;

  let lastError: StandardError | null = null;

  for (let attempt = 0; attempt <= maxRetries; attempt++) {
    try {
      return await fn();
    } catch (error) {
      lastError = parseError(error);

      // 如果是最后一次尝试或不应该重试，直接抛出错误
      if (attempt === maxRetries || !shouldRetry(lastError)) {
        throw lastError;
      }

      // 等待后重试
      await new Promise((resolve) => setTimeout(resolve, retryDelay * (attempt + 1)));
    }
  }

  throw lastError;
}

/**
 * 导出默认的错误处理器实例
 */
export const ErrorHandler = {
  parse: parseError,
  getMessage: getErrorMessage,
  isAPIError,
  isError,
  isNetworkError,
  isTimeoutError,
  getMessageByCode,
  createHandler: createErrorHandler,
  handleSettled: handleSettledResults,
  withRetry,
} as const;
