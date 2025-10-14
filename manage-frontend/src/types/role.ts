// 角色相关类型定义

export interface Role {
  id: number;
  name: string;
  code: string;
  description: string;
  status: string;
  is_system: boolean;
  created_at: string;
  updated_at: string;
}

export interface RoleWithPermissions extends Role {
  permissions: Permission[];
}

export interface Permission {
  id: number;
  name: string;
  code: string;
  resource: string;
  action: string;
  path: string;
  method: string;
  description: string;
  type: string;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface CreateRoleRequest {
  name: string;
  code: string;
  description: string;
}

export interface UpdateRoleRequest {
  name?: string;
  description?: string;
  status?: string;
}

export interface AssignRolePermissionsRequest {
  permission_ids: number[];
}

export interface AssignUserRolesRequest {
  role_ids: number[];
}

export interface RoleListParams {
  page?: number;
  page_size?: number;
}

export interface RoleListResponse {
  data: Role[];
  pagination: {
    page: number;
    page_size: number;
    total: number;
    total_pages: number;
  };
}
