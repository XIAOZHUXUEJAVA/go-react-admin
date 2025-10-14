"use client";

import { useEffect, useState, createContext, useContext } from "react";
import { useRouter, usePathname } from "next/navigation";
import { useAuthStore } from "@/stores/authStore";
import { isProtectedRoute, isAuthRoute } from "@/lib/auth";
import { LoadingSpinner } from "@/components/common/LoadingSpinner";

// 创建权限上下文
interface AuthContextType {
  requireRole: (role: string) => boolean;
  hasPermission: (permission: string) => boolean;
}

const AuthContext = createContext<AuthContextType | null>(null);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within AuthGuard");
  }
  return context;
};

interface AuthGuardProps {
  children: React.ReactNode;
  requiredRole?: string;
  fallbackPath?: string;
}

/**
 * 统一的认证守卫组件
 * 处理全局路由保护、角色权限检查和自动重定向
 */
export function AuthGuard({
  children,
  requiredRole,
  fallbackPath = "/login",
}: AuthGuardProps) {
  const router = useRouter();
  const pathname = usePathname();
  const { isAuthenticated, isLoading, checkAuth, user } = useAuthStore();
  const [isInitialized, setIsInitialized] = useState(false);

  // 初始化认证状态
  useEffect(() => {
    const initAuth = async () => {
      await checkAuth();
      setIsInitialized(true);
    };

    initAuth();
  }, [checkAuth]);

  // 处理路由重定向
  useEffect(() => {
    if (!isInitialized || isLoading) {
      return;
    }

    const isProtected = isProtectedRoute(pathname);
    const isAuth = isAuthRoute(pathname);
    // 首页重定向逻辑
    if (pathname === "/") {
      if (isAuthenticated) {
        router.push("/dashboard");
      } else {
        router.push("/login");
      }
      return;
    }

    // 如果用户已认证且在认证页面，重定向到 dashboard
    if (isAuth && isAuthenticated) {
      router.push("/dashboard");
      return;
    }

    // 如果是认证路由且用户未认证，允许访问
    if (isAuth && !isAuthenticated) {
      return;
    }

    // 认证检查
    if (isProtected && !isAuthenticated) {
      router.push(fallbackPath);
      return;
    }

    // 角色权限检查
    if (isAuthenticated && requiredRole && user?.role !== requiredRole) {
      router.push("/unauthorized");
      return;
    }
  }, [
    isAuthenticated,
    isLoading,
    pathname,
    router,
    isInitialized,
    requiredRole,
    user,
    fallbackPath,
  ]);

  // 权限检查函数
  const requireRole = (role: string): boolean => {
    return isAuthenticated && user?.role === role;
  };

  const hasPermission = (permission: string): boolean => {
    // 这里可以扩展更复杂的权限逻辑
    return isAuthenticated && user?.role === "admin";
  };

  // 显示加载状态
  if (!isInitialized || isLoading) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <LoadingSpinner text="正在加载..." />
      </div>
    );
  }

  // 权限不足时不渲染内容
  if (isAuthenticated && requiredRole && user?.role !== requiredRole) {
    return null;
  }

  const authContextValue: AuthContextType = {
    requireRole,
    hasPermission,
  };

  return (
    <AuthContext.Provider value={authContextValue}>
      {children}
    </AuthContext.Provider>
  );
}
