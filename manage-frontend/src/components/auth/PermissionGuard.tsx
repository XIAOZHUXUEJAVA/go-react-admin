"use client";

import React from "react";
import { usePermission } from "@/hooks/usePermission";

interface PermissionGuardProps {
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
   * 有权限时渲染的内容
   */
  children: React.ReactNode;
  /**
   * 无权限时渲染的内容（可选）
   */
  fallback?: React.ReactNode;
}

/**
 * 权限守卫组件
 * 根据用户权限条件渲染子组件
 *
 * @example
 * // 单个权限
 * <PermissionGuard permission="user:create">
 *   <Button>创建用户</Button>
 * </PermissionGuard>
 *
 * @example
 * // 任意一个权限
 * <PermissionGuard anyPermissions={["user:create", "user:update"]}>
 *   <Button>编辑用户</Button>
 * </PermissionGuard>
 *
 * @example
 * // 全部权限
 * <PermissionGuard allPermissions={["user:create", "user:delete"]}>
 *   <Button>批量操作</Button>
 * </PermissionGuard>
 *
 * @example
 * // 带 fallback
 * <PermissionGuard
 *   permission="user:delete"
 *   fallback={<span>无权限</span>}
 * >
 *   <Button>删除用户</Button>
 * </PermissionGuard>
 */
export const PermissionGuard: React.FC<PermissionGuardProps> = ({
  permission,
  anyPermissions,
  allPermissions,
  children,
  fallback = null,
}) => {
  const { hasPermission, hasAnyPermission, hasAllPermissions } =
    usePermission();

  // 检查权限
  let hasAccess = true;

  if (permission) {
    hasAccess = hasPermission(permission);
  } else if (anyPermissions && anyPermissions.length > 0) {
    hasAccess = hasAnyPermission(anyPermissions);
  } else if (allPermissions && allPermissions.length > 0) {
    hasAccess = hasAllPermissions(allPermissions);
  }

  return hasAccess ? <>{children}</> : <>{fallback}</>;
};
