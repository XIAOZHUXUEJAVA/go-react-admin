import { useState, useCallback, useMemo, useEffect } from "react";
import { userApi } from "@/api";
import { CheckAvailabilityResponse } from "@/types/user";
import { APIResponse } from "@/types/common";

// 带取消功能的防抖函数
function debounceWithCancel<T extends (...args: string[]) => Promise<void>>(
  func: T,
  wait: number
): {
  debouncedFn: (...args: Parameters<T>) => void;
  cancel: () => void;
} {
  let timeout: NodeJS.Timeout | undefined;

  const debouncedFn = (...args: Parameters<T>) => {
    if (timeout) {
      clearTimeout(timeout);
    }
    timeout = setTimeout(() => func(...args), wait);
  };

  const cancel = () => {
    if (timeout) {
      clearTimeout(timeout);
      timeout = undefined;
    }
  };

  return { debouncedFn, cancel };
}

export interface AvailabilityState {
  isChecking: boolean;
  isAvailable: boolean | null;
  message: string;
}

export interface UseAvailabilityCheckReturn {
  usernameAvailability: AvailabilityState;
  emailAvailability: AvailabilityState;
  checkUsernameAvailability: (username: string) => void;
  checkEmailAvailability: (email: string) => void;
  resetAvailabilityCheck: () => void;
}

// 验证类型枚举
enum AvailabilityType {
  USERNAME = "username",
  EMAIL = "email",
}

// 错误信息映射
const ERROR_MESSAGES = {
  [AvailabilityType.USERNAME]: {
    network: "网络连接失败，请检查网络后重试",
    server: "服务器暂时无法响应，请稍后重试",
    unknown: "检查用户名时出现未知错误，请重试",
  },
  [AvailabilityType.EMAIL]: {
    network: "网络连接失败，请检查网络后重试",
    server: "服务器暂时无法响应，请稍后重试",
    unknown: "检查邮箱时出现未知错误，请重试",
  },
} as const;

/**
 * 用户可用性检查 Hook - 专注于数据库可用性检查
 *
 * 职责：
 * - 检查用户名/邮箱在数据库中是否已被使用
 * - 提供防抖功能，避免频繁API调用
 * - 管理检查状态和错误处理
 *
 * 不负责：
 * - 格式验证（交给表单验证处理）
 * - 长度检查（交给表单验证处理）
 * - 其他业务逻辑验证
 */
export const useAvailabilityCheck = (
  excludeUserId?: number
): UseAvailabilityCheckReturn => {
  const [usernameAvailability, setUsernameAvailability] =
    useState<AvailabilityState>({
      isChecking: false,
      isAvailable: null,
      message: "",
    });

  const [emailAvailability, setEmailAvailability] = useState<AvailabilityState>(
    {
      isChecking: false,
      isAvailable: null,
      message: "",
    }
  );

  // 获取友好的错误信息
  const getErrorMessage = useCallback(
    (error: unknown, type: AvailabilityType): string => {
      if (error && typeof error === "object" && "code" in error) {
        const errorCode = (error as { code: string }).code;
        if (errorCode === "NETWORK_ERROR" || errorCode === "ERR_NETWORK") {
          return ERROR_MESSAGES[type].network;
        }
        if (errorCode >= "500" && errorCode < "600") {
          return ERROR_MESSAGES[type].server;
        }
      }
      return ERROR_MESSAGES[type].unknown;
    },
    []
  );

  // 抽象的可用性检查逻辑
  const createAvailabilityChecker = useCallback(
    (
      type: AvailabilityType,
      apiCall: (
        value: string
      ) => Promise<APIResponse<CheckAvailabilityResponse>>,
      setState: React.Dispatch<React.SetStateAction<AvailabilityState>>,
      availableMessage: string,
      unavailableMessage: string
    ) => {
      return async (value: string): Promise<void> => {
        // 如果值为空，重置状态
        if (!value.trim()) {
          setState({
            isChecking: false,
            isAvailable: null,
            message: "",
          });
          return;
        }

        // 延迟设置 loading 状态，避免短时间内的抖动
        const loadingTimer = setTimeout(() => {
          setState((prev) => ({
            ...prev,
            isChecking: true,
          }));
        }, 200);

        try {
          const response = await apiCall(value);
          clearTimeout(loadingTimer);

          // 根据类型获取正确的结果
          const result = response.data?.[type];
          const isAvailable = Boolean(result?.available);
          const message =
            result?.message ||
            (isAvailable ? availableMessage : unavailableMessage);

          setState({
            isChecking: false,
            isAvailable,
            message,
          });
        } catch (error) {
          clearTimeout(loadingTimer);
          setState({
            isChecking: false,
            isAvailable: null,
            message: getErrorMessage(error, type),
          });
        }
      };
    },
    [getErrorMessage]
  );

  // 使用 useMemo 创建防抖函数
  const { debouncedFn: debouncedCheckUsername, cancel: cancelUsernameCheck } =
    useMemo(() => {
      const checker = createAvailabilityChecker(
        AvailabilityType.USERNAME,
        userApi.checkUsernameAvailable,
        setUsernameAvailability,
        "用户名可用",
        "用户名已被使用"
      );
      return debounceWithCancel(checker, 500);
    }, [createAvailabilityChecker]);

  const { debouncedFn: debouncedCheckEmail, cancel: cancelEmailCheck } =
    useMemo(() => {
      const checker = createAvailabilityChecker(
        AvailabilityType.EMAIL,
        userApi.checkEmailAvailable,
        setEmailAvailability,
        "邮箱可用",
        "邮箱已被使用"
      );
      return debounceWithCancel(checker, 500);
    }, [createAvailabilityChecker]);

  // 清理函数
  useEffect(() => {
    return () => {
      cancelUsernameCheck();
      cancelEmailCheck();
    };
  }, [cancelUsernameCheck, cancelEmailCheck]);

  const checkUsernameAvailability = useCallback(
    (username: string) => {
      debouncedCheckUsername(username);
    },
    [debouncedCheckUsername]
  );

  const checkEmailAvailability = useCallback(
    (email: string) => {
      debouncedCheckEmail(email);
    },
    [debouncedCheckEmail]
  );

  const resetAvailabilityCheck = useCallback(() => {
    // 取消所有进行中的检查
    cancelUsernameCheck();
    cancelEmailCheck();

    setUsernameAvailability({
      isChecking: false,
      isAvailable: null,
      message: "",
    });
    setEmailAvailability({
      isChecking: false,
      isAvailable: null,
      message: "",
    });
  }, [cancelUsernameCheck, cancelEmailCheck]);

  return {
    usernameAvailability,
    emailAvailability,
    checkUsernameAvailability,
    checkEmailAvailability,
    resetAvailabilityCheck,
  };
};
