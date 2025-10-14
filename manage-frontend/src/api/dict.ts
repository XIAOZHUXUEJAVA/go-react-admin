/**
 * 字典管理相关 API
 */
import { ApiService } from "@/lib/api";
import { APIResponse } from "@/types/common";
import {
  DictType,
  DictItem,
  CreateDictTypeRequest,
  UpdateDictTypeRequest,
  CreateDictItemRequest,
  UpdateDictItemRequest,
  DictTypeQueryParams,
  DictItemQueryParams,
} from "@/types/dict";

export const dictApi = {
  // ==================== 字典类型 API ====================

  /**
   * 获取字典类型列表（分页）
   */
  getDictTypes: async (
    params?: DictTypeQueryParams
  ): Promise<APIResponse<DictType[]>> => {
    return ApiService.get<DictType[]>("/dict-types", params);
  },

  /**
   * 获取所有字典类型（不分页）
   */
  getAllDictTypes: async (): Promise<APIResponse<DictType[]>> => {
    return ApiService.get<DictType[]>("/dict-types/all");
  },

  /**
   * 根据ID获取字典类型详情
   */
  getDictTypeById: async (id: number): Promise<APIResponse<DictType>> => {
    return ApiService.get<DictType>(`/dict-types/${id}`);
  },

  /**
   * 创建字典类型
   */
  createDictType: async (
    data: CreateDictTypeRequest
  ): Promise<APIResponse<DictType>> => {
    return ApiService.post<DictType>("/dict-types", data);
  },

  /**
   * 更新字典类型
   */
  updateDictType: async (
    id: number,
    data: UpdateDictTypeRequest
  ): Promise<APIResponse<DictType>> => {
    return ApiService.put<DictType>(`/dict-types/${id}`, data);
  },

  /**
   * 删除字典类型
   */
  deleteDictType: async (id: number): Promise<APIResponse<void>> => {
    return ApiService.delete<void>(`/dict-types/${id}`);
  },

  // ==================== 字典项 API ====================

  /**
   * 获取字典项列表（分页）
   */
  getDictItems: async (
    params?: DictItemQueryParams
  ): Promise<APIResponse<DictItem[]>> => {
    return ApiService.get<DictItem[]>("/dict-items", params);
  },

  /**
   * 根据字典类型代码获取字典项
   */
  getDictItemsByType: async (
    typeCode: string,
    activeOnly = true
  ): Promise<APIResponse<DictItem[]>> => {
    return ApiService.get<DictItem[]>(`/dict-items/by-type/${typeCode}`, {
      active_only: activeOnly,
    });
  },

  /**
   * 根据ID获取字典项详情
   */
  getDictItemById: async (id: number): Promise<APIResponse<DictItem>> => {
    return ApiService.get<DictItem>(`/dict-items/${id}`);
  },

  /**
   * 创建字典项
   */
  createDictItem: async (
    data: CreateDictItemRequest
  ): Promise<APIResponse<DictItem>> => {
    return ApiService.post<DictItem>("/dict-items", data);
  },

  /**
   * 更新字典项
   */
  updateDictItem: async (
    id: number,
    data: UpdateDictItemRequest
  ): Promise<APIResponse<DictItem>> => {
    return ApiService.put<DictItem>(`/dict-items/${id}`, data);
  },

  /**
   * 删除字典项
   */
  deleteDictItem: async (id: number): Promise<APIResponse<void>> => {
    return ApiService.delete<void>(`/dict-items/${id}`);
  },
} as const;
