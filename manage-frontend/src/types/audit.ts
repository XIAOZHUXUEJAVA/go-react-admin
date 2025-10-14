// 审计日志相关类型定义

export interface AuditLog {
  id: number;
  user_id: number;
  username: string;
  action: string;
  resource: string;
  resource_id: string;
  method: string;
  path: string;
  ip: string;
  user_agent: string;
  status: number;
  error_msg: string;
  request_body: string;
  duration: number;
  created_at: string;
}

export interface AuditLogQuery {
  user_id?: number;
  username?: string;
  action?: string;
  resource?: string;
  method?: string;
  status?: number;
  start_time?: string;
  end_time?: string;
  page?: number;
  page_size?: number;
}

export interface AuditLogListResponse {
  data: AuditLog[];
  pagination: {
    page: number;
    page_size: number;
    total: number;
    total_pages: number;
  };
}
