import { useState, useEffect, useCallback } from "react";
import { auditApi } from "@/api";
import { AuditLog, AuditLogQuery } from "@/types/audit";
import { APIError } from "@/types/api";

// 分页信息类型
interface PaginationInfo {
  page: number;
  page_size: number;
  total: number;
  total_pages: number;
}

// 审计日志列表 Hook 的状态类型
interface UseAuditLogsState {
  logs: AuditLog[];
  pagination: PaginationInfo | null;
  loading: boolean;
  error: APIError | null;
}

// 审计日志列表 Hook 的返回类型
interface UseAuditLogsReturn extends UseAuditLogsState {
  fetchLogs: (params?: AuditLogQuery) => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * 审计日志列表数据获取 Hook
 */
export const useAuditLogs = (
  initialParams?: AuditLogQuery
): UseAuditLogsReturn => {
  const [state, setState] = useState<UseAuditLogsState>({
    logs: [],
    pagination: null,
    loading: false,
    error: null,
  });

  const [currentParams, setCurrentParams] = useState<AuditLogQuery>(
    initialParams || { page: 1, page_size: 10 }
  );

  const fetchLogs = useCallback(
    async (params?: AuditLogQuery) => {
      const queryParams = params || currentParams;
      setCurrentParams(queryParams);

      setState((prev) => ({ ...prev, loading: true, error: null }));

      try {
        const response = await auditApi.queryAuditLogs(queryParams);

        if (response.code === 200 && response.data) {
          setState((prev) => ({
            ...prev,
            logs: response.data as AuditLog[],
            pagination: response.pagination || null,
            loading: false,
          }));
        } else {
          throw new Error(response.message || "获取审计日志列表失败");
        }
      } catch (error) {
        const apiError = error as APIError;
        setState((prev) => ({
          ...prev,
          logs: [],
          pagination: null,
          loading: false,
          error: apiError,
        }));
      }
    },
    [currentParams]
  );

  const refetch = useCallback(() => {
    return fetchLogs(currentParams);
  }, [fetchLogs, currentParams]);

  useEffect(() => {
    fetchLogs();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return {
    ...state,
    fetchLogs,
    refetch,
  };
};

// 单个审计日志 Hook 的状态类型
interface UseAuditLogState {
  log: AuditLog | null;
  loading: boolean;
  error: APIError | null;
}

// 单个审计日志 Hook 的返回类型
interface UseAuditLogReturn extends UseAuditLogState {
  fetchLog: (id: number) => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * 单个审计日志数据获取 Hook
 */
export const useAuditLog = (id?: number): UseAuditLogReturn => {
  const [state, setState] = useState<UseAuditLogState>({
    log: null,
    loading: false,
    error: null,
  });

  const [currentId, setCurrentId] = useState<number | undefined>(id);

  const fetchLog = useCallback(async (logId: number) => {
    setCurrentId(logId);
    setState((prev) => ({ ...prev, loading: true, error: null }));

    try {
      const response = await auditApi.getAuditLogById(logId);

      if (response.code === 200 && response.data) {
        setState((prev) => ({
          ...prev,
          log: response.data as AuditLog,
          loading: false,
        }));
      } else {
        throw new Error(response.message || "获取审计日志信息失败");
      }
    } catch (error) {
      const apiError = error as APIError;
      setState((prev) => ({
        ...prev,
        log: null,
        loading: false,
        error: apiError,
      }));
    }
  }, []);

  const refetch = useCallback(() => {
    if (currentId) {
      return fetchLog(currentId);
    }
    return Promise.resolve();
  }, [fetchLog, currentId]);

  useEffect(() => {
    if (id) {
      fetchLog(id);
    }
  }, [id, fetchLog]);

  return {
    ...state,
    fetchLog,
    refetch,
  };
};
