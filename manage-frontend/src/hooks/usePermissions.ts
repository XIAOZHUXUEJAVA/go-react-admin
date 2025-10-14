import { useState, useEffect, useCallback } from "react";
import { permissionApi } from "@/api";
import { Permission, PermissionTree } from "@/types/permission";
import { APIError } from "@/types/common";

// 所有权限 Hook（不分页）
interface UseAllPermissionsState {
  permissions: Permission[];
  loading: boolean;
  error: APIError | null;
}

interface UseAllPermissionsReturn extends UseAllPermissionsState {
  fetchAllPermissions: () => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * 获取所有权限（不分页）Hook
 */
export const useAllPermissions = (): UseAllPermissionsReturn => {
  const [state, setState] = useState<UseAllPermissionsState>({
    permissions: [],
    loading: false,
    error: null,
  });

  const fetchAllPermissions = useCallback(async () => {
    setState((prev) => ({ ...prev, loading: true, error: null }));

    try {
      const response = await permissionApi.getAllPermissions();

      if (response.code === 200 && response.data) {
        setState((prev) => ({
          ...prev,
          permissions: response.data as Permission[],
          loading: false,
        }));
      } else {
        throw new Error(response.message || "获取权限列表失败");
      }
    } catch (error) {
      const apiError = error as APIError;
      setState((prev) => ({
        ...prev,
        permissions: [],
        loading: false,
        error: apiError,
      }));
    }
  }, []);

  const refetch = useCallback(() => {
    return fetchAllPermissions();
  }, [fetchAllPermissions]);

  useEffect(() => {
    fetchAllPermissions();
  }, [fetchAllPermissions]);

  return {
    ...state,
    fetchAllPermissions,
    refetch,
  };
};

// 权限树 Hook
interface UsePermissionTreeState {
  permissionTree: PermissionTree[];
  loading: boolean;
  error: APIError | null;
}

interface UsePermissionTreeReturn extends UsePermissionTreeState {
  fetchPermissionTree: () => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * 获取权限树 Hook
 */
export const usePermissionTree = (): UsePermissionTreeReturn => {
  const [state, setState] = useState<UsePermissionTreeState>({
    permissionTree: [],
    loading: false,
    error: null,
  });

  const fetchPermissionTree = useCallback(async () => {
    setState((prev) => ({ ...prev, loading: true, error: null }));

    try {
      const response = await permissionApi.getPermissionTree();

      if (response.code === 200 && response.data) {
        setState((prev) => ({
          ...prev,
          permissionTree: response.data as PermissionTree[],
          loading: false,
        }));
      } else {
        throw new Error(response.message || "获取权限树失败");
      }
    } catch (error) {
      const apiError = error as APIError;
      setState((prev) => ({
        ...prev,
        permissionTree: [],
        loading: false,
        error: apiError,
      }));
    }
  }, []);

  const refetch = useCallback(() => {
    return fetchPermissionTree();
  }, [fetchPermissionTree]);

  useEffect(() => {
    fetchPermissionTree();
  }, [fetchPermissionTree]);

  return {
    ...state,
    fetchPermissionTree,
    refetch,
  };
};
