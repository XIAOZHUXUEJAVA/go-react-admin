/**
 * 审计日志相关 API
 */
import { ApiService } from "@/lib/api";
import { APIResponse } from "@/types/common";
import {
  AuditLog,
  AuditLogQuery,
} from "@/types/audit";

export const auditApi = {
  /**
   * 查询审计日志列表
   */
  queryAuditLogs: async (
    params?: AuditLogQuery
  ): Promise<APIResponse<AuditLog[]>> => {
    return ApiService.get<AuditLog[]>("/audit-logs", params as Record<string, unknown>);
  },

  /**
   * 根据ID获取审计日志详情
   */
  getAuditLogById: async (id: number): Promise<APIResponse<AuditLog>> => {
    return ApiService.get<AuditLog>(`/audit-logs/${id}`);
  },

  /**
   * 清理旧的审计日志
   */
  cleanOldLogs: async (days: number): Promise<APIResponse<void>> => {
    return ApiService.post<void>(`/audit-logs/clean?days=${days}`);
  },
};
