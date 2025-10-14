// 通用类型定义

// API 响应类型
export interface APIResponse<T = unknown> {
  code: number;
  message: string;
  data?: T;
  error?: string;
  pagination?: PaginationInfo;
}

// 分页信息 - 匹配后端响应格式
export interface PaginationInfo {
  page: number;
  page_size: number;
  total: number;
  total_pages: number;
}

// 分页查询参数
export interface PaginationParams extends Record<string, unknown> {
  page?: number;
  pageSize?: number;
}

// API 错误类型
export interface APIError {
  code: number;
  message: string;
  error?: string;
}

// 通用查询参数
export interface BaseQueryParams extends PaginationParams {
  search?: string;
  sortBy?: string;
  sortOrder?: "asc" | "desc";
}

// 操作结果
export interface OperationResult {
  success: boolean;
  message?: string;
  data?: unknown;
}
