import { useState, useEffect, useCallback } from "react";
import { roleApi } from "@/api";
import { Role, RoleListParams } from "@/types/role";
import { PaginationInfo, APIError } from "@/types/common";

// 角色列表 Hook 的状态类型
interface UseRolesState {
  roles: Role[];
  pagination: PaginationInfo | null;
  loading: boolean;
  error: APIError | null;
}

// 角色列表 Hook 的返回类型
interface UseRolesReturn extends UseRolesState {
  fetchRoles: (params?: RoleListParams) => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * 角色列表数据获取 Hook
 */
export const useRoles = (initialParams?: RoleListParams): UseRolesReturn => {
  const [state, setState] = useState<UseRolesState>({
    roles: [],
    pagination: null,
    loading: false,
    error: null,
  });

  const [currentParams, setCurrentParams] = useState<RoleListParams>(
    initialParams || { page: 1, page_size: 10 }
  );

  const fetchRoles = useCallback(
    async (params?: RoleListParams) => {
      const queryParams = params || currentParams;
      setCurrentParams(queryParams);

      setState((prev) => ({ ...prev, loading: true, error: null }));

      try {
        const response = await roleApi.getRoles(queryParams);

        if (response.code === 200 && response.data) {
          setState((prev) => ({
            ...prev,
            roles: response.data as Role[],
            pagination: response.pagination || null,
            loading: false,
          }));
        } else {
          throw new Error(response.message || "获取角色列表失败");
        }
      } catch (error) {
        const apiError = error as APIError;
        setState((prev) => ({
          ...prev,
          roles: [],
          pagination: null,
          loading: false,
          error: apiError,
        }));
      }
    },
    [currentParams]
  );

  const refetch = useCallback(() => {
    return fetchRoles(currentParams);
  }, [fetchRoles, currentParams]);

  useEffect(() => {
    fetchRoles();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return {
    ...state,
    fetchRoles,
    refetch,
  };
};

// 单个角色 Hook 的状态类型
interface UseRoleState {
  role: Role | null;
  loading: boolean;
  error: APIError | null;
}

// 单个角色 Hook 的返回类型
interface UseRoleReturn extends UseRoleState {
  fetchRole: (id: number) => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * 单个角色数据获取 Hook
 */
export const useRole = (id?: number): UseRoleReturn => {
  const [state, setState] = useState<UseRoleState>({
    role: null,
    loading: false,
    error: null,
  });

  const [currentId, setCurrentId] = useState<number | undefined>(id);

  const fetchRole = useCallback(async (roleId: number) => {
    setCurrentId(roleId);
    setState((prev) => ({ ...prev, loading: true, error: null }));

    try {
      const response = await roleApi.getRoleById(roleId);

      if (response.code === 200 && response.data) {
        setState((prev) => ({
          ...prev,
          role: response.data as Role,
          loading: false,
        }));
      } else {
        throw new Error(response.message || "获取角色信息失败");
      }
    } catch (error) {
      const apiError = error as APIError;
      setState((prev) => ({
        ...prev,
        role: null,
        loading: false,
        error: apiError,
      }));
    }
  }, []);

  const refetch = useCallback(() => {
    if (currentId) {
      return fetchRole(currentId);
    }
    return Promise.resolve();
  }, [fetchRole, currentId]);

  useEffect(() => {
    if (id) {
      fetchRole(id);
    }
  }, [id, fetchRole]);

  return {
    ...state,
    fetchRole,
    refetch,
  };
};

// 所有角色 Hook（不分页）
interface UseAllRolesState {
  roles: Role[];
  loading: boolean;
  error: APIError | null;
}

interface UseAllRolesReturn extends UseAllRolesState {
  fetchAllRoles: () => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * 获取所有角色（不分页）Hook
 */
export const useAllRoles = (): UseAllRolesReturn => {
  const [state, setState] = useState<UseAllRolesState>({
    roles: [],
    loading: false,
    error: null,
  });

  const fetchAllRoles = useCallback(async () => {
    setState((prev) => ({ ...prev, loading: true, error: null }));

    try {
      const response = await roleApi.getAllRoles();

      if (response.code === 200 && response.data) {
        setState((prev) => ({
          ...prev,
          roles: response.data as Role[],
          loading: false,
        }));
      } else {
        throw new Error(response.message || "获取角色列表失败");
      }
    } catch (error) {
      const apiError = error as APIError;
      setState((prev) => ({
        ...prev,
        roles: [],
        loading: false,
        error: apiError,
      }));
    }
  }, []);

  const refetch = useCallback(() => {
    return fetchAllRoles();
  }, [fetchAllRoles]);

  useEffect(() => {
    fetchAllRoles();
  }, [fetchAllRoles]);

  return {
    ...state,
    fetchAllRoles,
    refetch,
  };
};
