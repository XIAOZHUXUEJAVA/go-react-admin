import { ApiService } from "@/lib/api";
import { APIResponse } from "@/types/api";

export interface Menu {
  id: number;
  parent_id: number | null;
  name: string;
  title: string;
  path: string;
  component: string;
  icon: string;
  order_num: number;
  type: string;
  permission_code: string;
  visible: boolean;
  status: string;
  created_at: string;
  updated_at: string;
  children?: Menu[];
}

export interface CreateMenuRequest {
  parent_id?: number | null;
  name: string;
  title: string;
  path: string;
  component?: string;
  icon?: string;
  order_num: number;
  type: string;
  permission_code?: string;
  visible: boolean;
}

export interface UpdateMenuRequest {
  parent_id?: number | null;
  name?: string;
  title?: string;
  path?: string;
  component?: string;
  icon?: string;
  order_num?: number;
  type?: string;
  permission_code?: string;
  visible?: boolean;
  status?: string;
}

export const menuApi = {
  // 获取菜单树
  getMenuTree: (): Promise<APIResponse<Menu[]>> => 
    ApiService.get<Menu[]>("/menus/tree"),

  // 获取可见菜单树
  getVisibleMenuTree: (): Promise<APIResponse<Menu[]>> => 
    ApiService.get<Menu[]>("/menus/tree/visible"),

  // 获取用户菜单树
  getUserMenuTree: (): Promise<APIResponse<Menu[]>> => 
    ApiService.get<Menu[]>("/menus/user"),

  // 获取菜单详情
  getMenu: (id: number): Promise<APIResponse<Menu>> => 
    ApiService.get<Menu>(`/menus/${id}`),

  // 创建菜单
  createMenu: (data: CreateMenuRequest): Promise<APIResponse<Menu>> =>
    ApiService.post<Menu>("/menus", data),

  // 更新菜单
  updateMenu: (id: number, data: UpdateMenuRequest): Promise<APIResponse<Menu>> =>
    ApiService.put<Menu>(`/menus/${id}`, data),

  // 删除菜单
  deleteMenu: (id: number): Promise<APIResponse<void>> => 
    ApiService.delete(`/menus/${id}`),

  // 批量更新菜单顺序
  updateMenuOrder: (menus: { id: number; order_num: number; parent_id: number | null }[]): Promise<APIResponse<void>> =>
    ApiService.put("/menus/order", { menus }),
};
