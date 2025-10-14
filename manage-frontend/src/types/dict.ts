// 字典类型
export interface DictType {
  id: number;
  code: string;
  name: string;
  description: string;
  status: string;
  sort_order: number;
  is_system: boolean;
  created_at: string;
  updated_at: string;
}

// 字典项
export interface DictItem {
  id: number;
  dict_type_code: string;
  label: string;
  value: string;
  extra?: Record<string, unknown>;
  description: string;
  status: string;
  sort_order: number;
  is_default: boolean;
  is_system: boolean;
  created_at: string;
  updated_at: string;
}

// 创建字典类型请求
export interface CreateDictTypeRequest {
  code: string;
  name: string;
  description?: string;
  status?: string;
  sort_order?: number;
}

// 更新字典类型请求
export interface UpdateDictTypeRequest {
  name?: string;
  description?: string;
  status?: string;
  sort_order?: number;
}

// 创建字典项请求
export interface CreateDictItemRequest {
  dict_type_code: string;
  label: string;
  value: string;
  extra?: Record<string, unknown>;
  description?: string;
  status?: string;
  sort_order?: number;
  is_default?: boolean;
}

// 更新字典项请求
export interface UpdateDictItemRequest {
  label?: string;
  extra?: Record<string, unknown>;
  description?: string;
  status?: string;
  sort_order?: number;
  is_default?: boolean;
}

// 字典类型查询参数
export interface DictTypeQueryParams extends Record<string, unknown> {
  page?: number;
  page_size?: number;
  status?: string;
  keyword?: string;
}

// 字典项查询参数
export interface DictItemQueryParams extends Record<string, unknown> {
  page?: number;
  page_size?: number;
  dict_type_code?: string;
  status?: string;
}
