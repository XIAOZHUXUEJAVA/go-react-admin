"use client";

import React from "react";
import { Button, buttonVariants } from "@/components/ui/button";
import { usePermission } from "@/hooks/usePermission";
import { type VariantProps } from "class-variance-authority";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

interface PermissionButtonProps
  extends React.ComponentProps<"button">,
    VariantProps<typeof buttonVariants> {
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
   * 无权限时的提示文本
   */
  noPermissionTooltip?: string;
  /**
   * 无权限时是否隐藏按钮（默认禁用）
   */
  hideWhenNoPermission?: boolean;
}

/**
 * 权限控制按钮组件
 * 根据用户权限自动启用/禁用按钮
 *
 * @example
 * // 单个权限
 * <PermissionButton permission="user:create">
 *   创建用户
 * </PermissionButton>
 *
 * @example
 * // 任意一个权限
 * <PermissionButton anyPermissions={["user:create", "user:update"]}>
 *   编辑用户
 * </PermissionButton>
 *
 * @example
 * // 全部权限
 * <PermissionButton allPermissions={["user:create", "user:delete"]}>
 *   批量操作
 * </PermissionButton>
 *
 * @example
 * // 无权限时隐藏
 * <PermissionButton permission="user:delete" hideWhenNoPermission>
 *   删除用户
 * </PermissionButton>
 *
 * @example
 * // 自定义提示
 * <PermissionButton
 *   permission="user:delete"
 *   noPermissionTooltip="您没有删除用户的权限"
 * >
 *   删除用户
 * </PermissionButton>
 */
export const PermissionButton: React.FC<PermissionButtonProps> = ({
  permission,
  anyPermissions,
  allPermissions,
  noPermissionTooltip = "您没有权限执行此操作",
  hideWhenNoPermission = false,
  children,
  disabled,
  ...buttonProps
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

  // 无权限且设置隐藏时，不渲染按钮
  if (!hasAccess && hideWhenNoPermission) {
    return null;
  }

  // 无权限时显示禁用按钮和提示
  if (!hasAccess) {
    return (
      <TooltipProvider>
        <Tooltip>
          <TooltipTrigger asChild>
            <span className="inline-block">
              <Button {...buttonProps} disabled={true}>
                {children}
              </Button>
            </span>
          </TooltipTrigger>
          <TooltipContent>
            <p>{noPermissionTooltip}</p>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>
    );
  }

  // 有权限时正常渲染按钮
  return (
    <Button {...buttonProps} disabled={disabled}>
      {children}
    </Button>
  );
};
