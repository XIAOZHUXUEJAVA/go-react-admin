/**
 * 菜单相关 API
 */
import { ApiService } from "@/lib/api";
import { APIResponse } from "@/types/common";
import { Menu, CreateMenuRequest, UpdateMenuRequest } from "@/types/menu";

export const menuApi = {
  /**
   * 获取完整菜单树
   */
  getMenuTree: async (): Promise<APIResponse<Menu[]>> => {
    return ApiService.get<Menu[]>("/menus/tree");
  },

  /**
   * 获取可见菜单树
   */
  getVisibleMenuTree: async (): Promise<APIResponse<Menu[]>> => {
    return ApiService.get<Menu[]>("/menus/tree/visible");
  },

  /**
   * 获取当前用户的菜单树
   */
  getUserMenuTree: async (): Promise<APIResponse<Menu[]>> => {
    return ApiService.get<Menu[]>("/menus/user");
  },

  /**
   * 根据ID获取菜单详情
   */
  getMenuById: async (id: number): Promise<APIResponse<Menu>> => {
    return ApiService.get<Menu>(`/menus/${id}`);
  },

  /**
   * 创建新菜单
   */
  createMenu: async (data: CreateMenuRequest): Promise<APIResponse<Menu>> => {
    return ApiService.post<Menu>("/menus", data);
  },

  /**
   * 更新菜单信息
   */
  updateMenu: async (
    id: number,
    data: UpdateMenuRequest
  ): Promise<APIResponse<Menu>> => {
    return ApiService.put<Menu>(`/menus/${id}`, data);
  },

  /**
   * 删除菜单
   */
  deleteMenu: async (id: number): Promise<APIResponse<void>> => {
    return ApiService.delete<void>(`/menus/${id}`);
  },
};
