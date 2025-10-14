import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";
import { Permission } from "@/types/permission";
import { Menu } from "@/types/menu";
import { permissionApi, menuApi } from "@/api";

/**
 * 权限状态
 */
interface PermissionState {
  permissions: Permission[];
  menus: Menu[];
  userMenus: Menu[];
  isLoading: boolean;
  isLoaded: boolean;
  error: string | null;
}

/**
 * 权限操作
 */
interface PermissionActions {
  loadPermissions: () => Promise<void>;
  loadMenus: () => Promise<void>;
  loadUserMenus: () => Promise<void>;
  hasPermission: (permissionCode: string) => boolean;
  hasAnyPermission: (permissionCodes: string[]) => boolean;
  hasAllPermissions: (permissionCodes: string[]) => boolean;
  clearPermissions: () => void;
  setLoading: (loading: boolean) => void;
}

/**
 * 权限 Store 类型
 */
export type PermissionStore = PermissionState & PermissionActions;

/**
 * 权限状态管理 Store
 * 管理用户权限、菜单等信息
 */
export const usePermissionStore = create<PermissionStore>()(
  persist(
    (set, get) => ({
      // 初始状态
      permissions: [],
      menus: [],
      userMenus: [],
      isLoading: false,
      isLoaded: false,
      error: null,

      // 设置加载状态
      setLoading: (loading: boolean) => {
        set({ isLoading: loading });
      },

      // 加载用户权限
      loadPermissions: async () => {
        try {
          set({ isLoading: true, error: null });

          // 调用获取用户权限的接口，而不是获取所有权限
          const response = await permissionApi.getUserPermissions();

          if (response.code === 200 && response.data) {
            // 将权限代码数组转换为 Permission 对象数组
            const permissionObjects = response.data.permissions.map((code) => ({
              id: 0, // ID 不重要，只用于权限检查
              code,
              name: code,
              resource: code.split(":")[0] || "",
              action: code.split(":")[1] || "",
              path: "",
              method: "",
              description: "",
              type: "api",
              status: "active",
              created_at: "",
              updated_at: "",
            }));

            set({
              permissions: permissionObjects,
              isLoading: false,
              isLoaded: true,
            });
          } else {
            throw new Error(response.message || "加载权限失败");
          }
        } catch (error) {
          set({
            permissions: [],
            isLoading: false,
            isLoaded: false,
            error: error instanceof Error ? error.message : "加载权限失败",
          });
        }
      },

      // 加载所有菜单
      loadMenus: async () => {
        try {
          set({ isLoading: true, error: null });

          const response = await menuApi.getMenuTree();

          if (response.code === 200 && response.data) {
            set({
              menus: response.data,
              isLoading: false,
            });
          } else {
            throw new Error(response.message || "加载菜单失败");
          }
        } catch (error) {
          set({
            menus: [],
            isLoading: false,
            error: error instanceof Error ? error.message : "加载菜单失败",
          });
        }
      },

      // 加载用户菜单
      loadUserMenus: async () => {
        try {
          set({ isLoading: true, error: null });

          const response = await menuApi.getUserMenuTree();

          if (response.code === 200 && response.data) {
            set({
              userMenus: response.data,
              isLoading: false,
            });
          } else {
            throw new Error(response.message || "加载用户菜单失败");
          }
        } catch (error) {
          set({
            userMenus: [],
            isLoading: false,
            error: error instanceof Error ? error.message : "加载用户菜单失败",
          });
        }
      },

      // 检查是否有某个权限
      hasPermission: (permissionCode: string): boolean => {
        const { permissions } = get();
        return permissions.some(
          (p) => p.code === permissionCode && p.status === "active"
        );
      },

      // 检查是否有任意一个权限
      hasAnyPermission: (permissionCodes: string[]): boolean => {
        const { hasPermission } = get();
        return permissionCodes.some((code) => hasPermission(code));
      },

      // 检查是否有所有权限
      hasAllPermissions: (permissionCodes: string[]): boolean => {
        const { hasPermission } = get();
        return permissionCodes.every((code) => hasPermission(code));
      },

      // 清除权限数据
      clearPermissions: () => {
        set({
          permissions: [],
          menus: [],
          userMenus: [],
          isLoading: false,
          isLoaded: false,
          error: null,
        });
      },
    }),
    {
      name: "permission-storage",
      storage: createJSONStorage(() => localStorage),
      // 持久化权限和菜单数据
      partialize: (state) => ({
        permissions: state.permissions,
        menus: state.menus,
        userMenus: state.userMenus,
        isLoaded: state.isLoaded,
      }),
    }
  )
);
