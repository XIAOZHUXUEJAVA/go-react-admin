// 权限相关类型定义

export interface Permission {
  id: number;
  name: string;
  code: string;
  resource: string;
  action: string;
  path: string;
  method: string;
  description: string;
  type: string; // api, menu, button
  status: string;
  created_at: string;
  updated_at: string;
}

export interface PermissionTree {
  resource: string;
  permissions: Permission[];
}

export interface CreatePermissionRequest {
  name: string;
  code: string;
  resource: string;
  action: string;
  path?: string;
  method?: string;
  description?: string;
  type: string;
}

export interface UpdatePermissionRequest {
  name?: string;
  description?: string;
  path?: string;
  method?: string;
  status?: string;
}

export interface PermissionListParams {
  page?: number;
  page_size?: number;
}

export interface PermissionListResponse {
  data: Permission[];
  pagination: {
    page: number;
    page_size: number;
    total: number;
    total_pages: number;
  };
}
