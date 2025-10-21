"use client";

import React from "react";
import { DropdownMenuItem } from "@/components/ui/dropdown-menu";
import { usePermission } from "@/hooks/usePermission";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

interface PermissionDropdownMenuItemProps
  extends React.ComponentProps<typeof DropdownMenuItem> {
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
   * 无权限时是否隐藏菜单项（默认禁用）
   */
  hideWhenNoPermission?: boolean;
}

/**
 * 权限控制下拉菜单项组件
 * 根据用户权限自动启用/禁用菜单项
 *
 * @example
 * ```tsx
 * <DropdownMenu>
 *   <DropdownMenuTrigger>操作</DropdownMenuTrigger>
 *   <DropdownMenuContent>
 *     <PermissionDropdownMenuItem permission="user:edit">
 *       编辑
 *     </PermissionDropdownMenuItem>
 *     <PermissionDropdownMenuItem 
 *       permission="user:delete"
 *       hideWhenNoPermission
 *     >
 *       删除
 *     </PermissionDropdownMenuItem>
 *   </DropdownMenuContent>
 * </DropdownMenu>
 * ```
 */
export const PermissionDropdownMenuItem: React.FC<
  PermissionDropdownMenuItemProps
> = ({
  permission,
  anyPermissions,
  allPermissions,
  noPermissionTooltip = "您没有权限执行此操作",
  hideWhenNoPermission = false,
  children,
  disabled,
  ...menuItemProps
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

  // 无权限且设置隐藏时，不渲染菜单项
  if (!hasAccess && hideWhenNoPermission) {
    return null;
  }

  // 无权限时显示禁用菜单项和提示
  if (!hasAccess) {
    return (
      <TooltipProvider>
        <Tooltip>
          <TooltipTrigger asChild>
            <span>
              <DropdownMenuItem {...menuItemProps} disabled={true}>
                {children}
              </DropdownMenuItem>
            </span>
          </TooltipTrigger>
          <TooltipContent>
            <p>{noPermissionTooltip}</p>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>
    );
  }

  // 有权限时正常渲染菜单项
  return (
    <DropdownMenuItem {...menuItemProps} disabled={disabled}>
      {children}
    </DropdownMenuItem>
  );
};
