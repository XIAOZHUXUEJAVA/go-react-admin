/**
 * 页面权限配置
 * 集中管理所有页面的权限要求
 */

export interface PagePermissionConfig {
  /**
   * 页面路径
   */
  path: string;
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
   * 页面描述
   */
  description?: string;
}

/**
 * 页面权限映射表
 * 定义每个页面需要的权限
 */
export const PAGE_PERMISSIONS: Record<string, PagePermissionConfig> = {
  // 用户管理
  USERS_LIST: {
    path: "/system/users",
    permission: "user:read",
    description: "用户管理页面 - 查看用户列表",
  },

  // 角色管理
  ROLES_LIST: {
    path: "/system/roles",
    permission: "role:read",
    description: "角色管理页面 - 查看角色列表",
  },

  // 权限管理
  PERMISSIONS_LIST: {
    path: "/system/permissions",
    permission: "permission:read",
    description: "权限管理页面 - 查看权限列表",
  },

  // 菜单管理
  MENUS_LIST: {
    path: "/system/menus",
    permission: "menu:read",
    description: "菜单管理页面 - 查看菜单列表",
  },
};

/**
 * 按钮权限配置
 * 定义常用按钮的权限要求
 */
export const BUTTON_PERMISSIONS = {
  // 用户管理按钮
  USER_CREATE: "user:create",
  USER_UPDATE: "user:update",
  USER_DELETE: "user:delete",

  // 角色管理按钮
  ROLE_CREATE: "role:create",
  ROLE_UPDATE: "role:update",
  ROLE_DELETE: "role:delete",
  ROLE_ASSIGN_PERMISSIONS: "role:update",

  // 权限管理按钮
  PERMISSION_CREATE: "permission:create",
  PERMISSION_UPDATE: "permission:update",
  PERMISSION_DELETE: "permission:delete",

  // 菜单管理按钮
  MENU_CREATE: "menu:create",
  MENU_UPDATE: "menu:update",
  MENU_DELETE: "menu:delete",
} as const;

/**
 * 根据路径获取页面权限配置
 */
export function getPagePermission(path: string): PagePermissionConfig | undefined {
  return Object.values(PAGE_PERMISSIONS).find((config) => config.path === path);
}

/**
 * 检查路径是否需要权限验证
 */
export function requiresPermission(path: string): boolean {
  const config = getPagePermission(path);
  return !!(
    config &&
    (config.permission || config.anyPermissions || config.allPermissions)
  );
}
