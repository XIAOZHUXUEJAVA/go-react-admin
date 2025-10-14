import { usePermissionStore } from "@/stores/permissionStore";

/**
 * 权限检查 Hook
 * 提供便捷的权限检查方法
 */
export const usePermission = () => {
  const {
    permissions,
    hasPermission,
    hasAnyPermission,
    hasAllPermissions,
    isLoaded,
  } = usePermissionStore();

  return {
    permissions,
    hasPermission,
    hasAnyPermission,
    hasAllPermissions,
    isLoaded,
  };
};

/**
 * 检查单个权限的 Hook
 * @param permissionCode 权限代码
 * @returns 是否有该权限
 */
export const useHasPermission = (permissionCode: string): boolean => {
  const { hasPermission } = usePermissionStore();
  return hasPermission(permissionCode);
};

/**
 * 检查多个权限（任意一个）的 Hook
 * @param permissionCodes 权限代码数组
 * @returns 是否有任意一个权限
 */
export const useHasAnyPermission = (permissionCodes: string[]): boolean => {
  const { hasAnyPermission } = usePermissionStore();
  return hasAnyPermission(permissionCodes);
};

/**
 * 检查多个权限（全部）的 Hook
 * @param permissionCodes 权限代码数组
 * @returns 是否有全部权限
 */
export const useHasAllPermissions = (permissionCodes: string[]): boolean => {
  const { hasAllPermissions } = usePermissionStore();
  return hasAllPermissions(permissionCodes);
};
