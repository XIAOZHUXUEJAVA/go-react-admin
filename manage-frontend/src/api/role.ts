/**
 * 角色相关 API
 */
import { ApiService } from "@/lib/api";
import { APIResponse } from "@/types/common";
import {
  Role,
  RoleWithPermissions,
  CreateRoleRequest,
  UpdateRoleRequest,
  AssignRolePermissionsRequest,
  AssignUserRolesRequest,
  RoleListParams,
} from "@/types/role";

export const roleApi = {
  /**
   * 获取角色列表（分页）
   */
  getRoles: async (params?: RoleListParams): Promise<APIResponse<Role[]>> => {
    return ApiService.get<Role[]>("/roles", params as Record<string, unknown>);
  },

  /**
   * 获取所有角色（不分页）
   */
  getAllRoles: async (): Promise<APIResponse<Role[]>> => {
    return ApiService.get<Role[]>("/roles/all");
  },

  /**
   * 根据ID获取角色详情
   */
  getRoleById: async (id: number): Promise<APIResponse<Role>> => {
    return ApiService.get<Role>(`/roles/${id}`);
  },

  /**
   * 创建新角色
   */
  createRole: async (data: CreateRoleRequest): Promise<APIResponse<Role>> => {
    return ApiService.post<Role>("/roles", data);
  },

  /**
   * 更新角色信息
   */
  updateRole: async (
    id: number,
    data: UpdateRoleRequest
  ): Promise<APIResponse<Role>> => {
    return ApiService.put<Role>(`/roles/${id}`, data);
  },

  /**
   * 删除角色
   */
  deleteRole: async (id: number): Promise<APIResponse<void>> => {
    return ApiService.delete<void>(`/roles/${id}`);
  },

  /**
   * 获取角色的权限列表
   */
  getRolePermissions: async (
    id: number
  ): Promise<APIResponse<RoleWithPermissions>> => {
    return ApiService.get<RoleWithPermissions>(`/roles/${id}/permissions`);
  },

  /**
   * 为角色分配权限
   */
  assignPermissions: async (
    id: number,
    data: AssignRolePermissionsRequest
  ): Promise<APIResponse<void>> => {
    return ApiService.put<void>(`/roles/${id}/permissions`, data);
  },

  /**
   * 获取用户的角色列表
   */
  getUserRoles: async (userId: number): Promise<APIResponse<Role[]>> => {
    return ApiService.get<Role[]>(`/users/${userId}/roles`);
  },

  /**
   * 为用户分配角色
   */
  assignRolesToUser: async (
    userId: number,
    data: AssignUserRolesRequest
  ): Promise<APIResponse<void>> => {
    return ApiService.put<void>(`/users/${userId}/roles`, data);
  },
};
