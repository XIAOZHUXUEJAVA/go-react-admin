"use client";

import React, { useEffect } from "react";
import { useRouter } from "next/navigation";
import { usePermission } from "@/hooks/usePermission";
import { Loader2 } from "lucide-react";

interface PagePermissionGuardProps {
  /**
   * 需要的权限代码
   */
  permission?: string;
  /**
   * 需要的权限代码数组（任意一个）
   */
  anyPermissions?: string[];
  /**
   * 需要的权限代码数组（全部）
   */
  allPermissions?: string[];
  /**
   * 页面内容
   */
  children: React.ReactNode;
  /**
   * 加载时显示的内容
   */
  loadingFallback?: React.ReactNode;
}

/**
 * 页面级权限守卫组件
 * 用于保护整个页面，无权限时自动跳转到 403 页面
 *
 * @example
 * // 在页面组件中使用
 * export default function UsersPage() {
 *   return (
 *     <PagePermissionGuard permission="user:read">
 *       <div>用户管理页面内容</div>
 *     </PagePermissionGuard>
 *   );
 * }
 *
 * @example
 * // 需要任意一个权限
 * <PagePermissionGuard anyPermissions={["user:read", "user:create"]}>
 *   <div>页面内容</div>
 * </PagePermissionGuard>
 *
 * @example
 * // 需要全部权限
 * <PagePermissionGuard allPermissions={["user:read", "user:delete"]}>
 *   <div>页面内容</div>
 * </PagePermissionGuard>
 */
export const PagePermissionGuard: React.FC<PagePermissionGuardProps> = ({
  permission,
  anyPermissions,
  allPermissions,
  children,
  loadingFallback,
}) => {
  const router = useRouter();
  const { hasPermission, hasAnyPermission, hasAllPermissions, isLoaded } =
    usePermission();

  useEffect(() => {
    // 等待权限数据加载完成
    if (!isLoaded) {
      return;
    }

    // 检查权限
    let hasAccess = true;

    if (permission) {
      hasAccess = hasPermission(permission);
    } else if (anyPermissions && anyPermissions.length > 0) {
      hasAccess = hasAnyPermission(anyPermissions);
    } else if (allPermissions && allPermissions.length > 0) {
      hasAccess = hasAllPermissions(allPermissions);
    }

    // 无权限时跳转到 403 页面
    if (!hasAccess) {
      router.replace("/forbidden");
    }
  }, [
    permission,
    anyPermissions,
    allPermissions,
    hasPermission,
    hasAnyPermission,
    hasAllPermissions,
    isLoaded,
    router,
  ]);

  // 权限未加载完成时显示加载状态
  if (!isLoaded) {
    return (
      loadingFallback || (
        <div className="flex items-center justify-center min-h-screen">
          <div className="text-center space-y-4">
            <Loader2 className="h-8 w-8 animate-spin mx-auto text-primary" />
            <p className="text-sm text-muted-foreground">正在验证权限...</p>
          </div>
        </div>
      )
    );
  }

  // 检查权限
  let hasAccess = true;

  if (permission) {
    hasAccess = hasPermission(permission);
  } else if (anyPermissions && anyPermissions.length > 0) {
    hasAccess = hasAnyPermission(anyPermissions);
  } else if (allPermissions && allPermissions.length > 0) {
    hasAccess = hasAllPermissions(allPermissions);
  }

  // 有权限时渲染页面内容
  // 无权限时返回 null（因为会在 useEffect 中跳转）
  return hasAccess ? <>{children}</> : null;
};
