/**
 * API 层统一导出
 *
 * 使用方式：
 * import { api } from '@/api';
 * const users = await api.user.getUsers();
 *
 * 或者：
 * import { userApi, roleApi, authApi } from '@/api';
 * const users = await userApi.getUsers();
 */

import { userApi } from "./user";
import { authApi } from "./auth";
import { roleApi } from "./role";
import { permissionApi } from "./permission";
import { menuApi } from "./menu";
import { auditApi } from "./audit";
import { dictApi } from "./dict";

// 重新导出各个 API
export { userApi } from "./user";
export { authApi } from "./auth";
export { roleApi } from "./role";
export { permissionApi } from "./permission";
export { menuApi } from "./menu";
export { auditApi } from "./audit";
export { dictApi } from "./dict";

// 统一的 API 对象
export const api = {
  user: userApi,
  auth: authApi,
  role: roleApi,
  permission: permissionApi,
  menu: menuApi,
  audit: auditApi,
  dict: dictApi,
} as const;

// 类型导出，方便其他地方使用
export type Api = typeof api;
export type UserApi = typeof userApi;
export type AuthApi = typeof authApi;
export type RoleApi = typeof roleApi;
export type PermissionApi = typeof permissionApi;
export type MenuApi = typeof menuApi;
export type AuditApi = typeof auditApi;
export type DictApi = typeof dictApi;
