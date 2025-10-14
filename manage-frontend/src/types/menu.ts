// 菜单相关类型定义

export interface Menu {
  id: number;
  parent_id: number | null;
  name: string;
  title: string;
  path: string;
  component: string;
  icon: string;
  order_num: number;
  type: string; // menu, button
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
  component: string;
  icon: string;
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
  permission_code?: string;
  visible?: boolean;
  status?: string;
}
