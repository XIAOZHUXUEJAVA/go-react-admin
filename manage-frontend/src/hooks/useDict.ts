import { useState, useEffect, useCallback } from "react";
import { dictApi } from "@/api/dict";
import {
  DictType,
  DictItem,
  DictTypeQueryParams,
  DictItemQueryParams,
} from "@/types/dict";
import { PaginationInfo, APIError } from "@/types/api";

// ==================== 字典类型 Hooks ====================

interface UseDictTypesState {
  dictTypes: DictType[];
  pagination: PaginationInfo | null;
  loading: boolean;
  error: APIError | null;
}

interface UseDictTypesReturn extends UseDictTypesState {
  fetchDictTypes: (params?: DictTypeQueryParams) => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * 字典类型列表 Hook
 */
export const useDictTypes = (
  initialParams?: DictTypeQueryParams
): UseDictTypesReturn => {
  const [state, setState] = useState<UseDictTypesState>({
    dictTypes: [],
    pagination: null,
    loading: false,
    error: null,
  });

  const [currentParams, setCurrentParams] = useState<DictTypeQueryParams>(
    initialParams || { page: 1, page_size: 10 }
  );

  const fetchDictTypes = useCallback(
    async (params?: DictTypeQueryParams) => {
      const queryParams = params || currentParams;
      setCurrentParams(queryParams);

      setState((prev) => ({ ...prev, loading: true, error: null }));

      try {
        const response = await dictApi.getDictTypes(queryParams);

        if (response.code === 200 && response.data) {
          setState((prev) => ({
            ...prev,
            dictTypes: response.data as DictType[],
            pagination: response.pagination || null,
            loading: false,
          }));
        } else {
          throw new Error(response.message || "获取字典类型列表失败");
        }
      } catch (error) {
        const apiError = error as APIError;
        setState((prev) => ({
          ...prev,
          dictTypes: [],
          pagination: null,
          loading: false,
          error: apiError,
        }));
      }
    },
    [currentParams]
  );

  const refetch = useCallback(() => {
    return fetchDictTypes(currentParams);
  }, [fetchDictTypes, currentParams]);

  useEffect(() => {
    fetchDictTypes();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return {
    ...state,
    fetchDictTypes,
    refetch,
  };
};

// ==================== 字典项 Hooks ====================

interface UseDictItemsState {
  dictItems: DictItem[];
  pagination: PaginationInfo | null;
  loading: boolean;
  error: APIError | null;
}

interface UseDictItemsReturn extends UseDictItemsState {
  fetchDictItems: (params?: DictItemQueryParams) => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * 字典项列表 Hook
 */
export const useDictItems = (
  initialParams?: DictItemQueryParams
): UseDictItemsReturn => {
  const [state, setState] = useState<UseDictItemsState>({
    dictItems: [],
    pagination: null,
    loading: false,
    error: null,
  });

  const [currentParams, setCurrentParams] = useState<DictItemQueryParams>(
    initialParams || { page: 1, page_size: 10 }
  );

  const fetchDictItems = useCallback(
    async (params?: DictItemQueryParams) => {
      const queryParams = params || currentParams;
      setCurrentParams(queryParams);

      setState((prev) => ({ ...prev, loading: true, error: null }));

      try {
        const response = await dictApi.getDictItems(queryParams);

        if (response.code === 200 && response.data) {
          setState((prev) => ({
            ...prev,
            dictItems: response.data as DictItem[],
            pagination: response.pagination || null,
            loading: false,
          }));
        } else {
          throw new Error(response.message || "获取字典项列表失败");
        }
      } catch (error) {
        const apiError = error as APIError;
        setState((prev) => ({
          ...prev,
          dictItems: [],
          pagination: null,
          loading: false,
          error: apiError,
        }));
      }
    },
    [currentParams]
  );

  const refetch = useCallback(() => {
    return fetchDictItems(currentParams);
  }, [fetchDictItems, currentParams]);

  useEffect(() => {
    fetchDictItems();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return {
    ...state,
    fetchDictItems,
    refetch,
  };
};

/**
 * 根据类型代码获取字典项 Hook
 */
export const useDictItemsByType = (typeCode: string, activeOnly = true) => {
  const [state, setState] = useState<{
    dictItems: DictItem[];
    loading: boolean;
    error: APIError | null;
  }>({
    dictItems: [],
    loading: false,
    error: null,
  });

  const fetchDictItems = useCallback(async () => {
    if (!typeCode) return;

    setState((prev) => ({ ...prev, loading: true, error: null }));

    try {
      const response = await dictApi.getDictItemsByType(typeCode, activeOnly);

      if (response.code === 200 && response.data) {
        setState((prev) => ({
          ...prev,
          dictItems: response.data as DictItem[],
          loading: false,
        }));
      } else {
        throw new Error(response.message || "获取字典项失败");
      }
    } catch (error) {
      const apiError = error as APIError;
      setState((prev) => ({
        ...prev,
        dictItems: [],
        loading: false,
        error: apiError,
      }));
    }
  }, [typeCode, activeOnly]);

  useEffect(() => {
    fetchDictItems();
  }, [fetchDictItems]);

  return {
    ...state,
    refetch: fetchDictItems,
  };
};
