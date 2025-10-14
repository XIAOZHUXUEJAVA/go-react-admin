/**
 * 权限相关 API
 */
import { ApiService } from "@/lib/api";
import { APIResponse } from "@/types/common";
import {
  Permission,
  PermissionTree,
  CreatePermissionRequest,
  UpdatePermissionRequest,
  PermissionListParams,
} from "@/types/permission";

export const permissionApi = {
  /**
   * 获取权限列表（分页）
   */
  getPermissions: async (
    params?: PermissionListParams
  ): Promise<APIResponse<Permission[]>> => {
    return ApiService.get<Permission[]>("/permissions", params as Record<string, unknown>);
  },

  /**
   * 获取所有权限（不分页）
   */
  getAllPermissions: async (): Promise<APIResponse<Permission[]>> => {
    return ApiService.get<Permission[]>("/permissions/all");
  },

  /**
   * 获取权限树
   */
  getPermissionTree: async (): Promise<APIResponse<PermissionTree[]>> => {
    return ApiService.get<PermissionTree[]>("/permissions/tree");
  },

  /**
   * 根据ID获取权限详情
   */
  getPermissionById: async (id: number): Promise<APIResponse<Permission>> => {
    return ApiService.get<Permission>(`/permissions/${id}`);
  },

  /**
   * 创建新权限
   */
  createPermission: async (
    data: CreatePermissionRequest
  ): Promise<APIResponse<Permission>> => {
    return ApiService.post<Permission>("/permissions", data);
  },

  /**
   * 更新权限信息
   */
  updatePermission: async (
    id: number,
    data: UpdatePermissionRequest
  ): Promise<APIResponse<Permission>> => {
    return ApiService.put<Permission>(`/permissions/${id}`, data);
  },

  /**
   * 删除权限
   */
  deletePermission: async (id: number): Promise<APIResponse<void>> => {
    return ApiService.delete<void>(`/permissions/${id}`);
  },

  /**
   * 根据资源类型获取权限
   */
  getPermissionsByResource: async (
    resource: string
  ): Promise<APIResponse<Permission[]>> => {
    return ApiService.get<Permission[]>(`/permissions/resource/${resource}`);
  },

  /**
   * 根据权限类型获取权限
   */
  getPermissionsByType: async (
    type: string
  ): Promise<APIResponse<Permission[]>> => {
    return ApiService.get<Permission[]>(`/permissions/type/${type}`);
  },

  /**
   * 获取当前用户的权限
   */
  getUserPermissions: async (): Promise<APIResponse<{ roles: string[]; permissions: string[] }>> => {
    return ApiService.get<{ roles: string[]; permissions: string[] }>("/users/permissions");
  },
};
