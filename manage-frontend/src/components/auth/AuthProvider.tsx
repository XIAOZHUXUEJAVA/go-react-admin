"use client";

import { useEffect } from "react";
import { useAuthStore } from "@/stores/authStore";

/**
 * 认证提供者组件
 * 在应用启动时检查认证状态
 */
export function AuthProvider({ children }: { children: React.ReactNode }) {
  const { checkAuth, token } = useAuthStore();

  useEffect(() => {
    // 如果有 token，检查认证状态
    if (token) {
      checkAuth();
    }
  }, [checkAuth, token]);

  return <>{children}</>;
}
