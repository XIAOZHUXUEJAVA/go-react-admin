/**
 * 用户相关 API
 */
import { ApiService } from "@/lib/api";
import { APIResponse } from "@/types/common";
import {
  User,
  CreateUserRequest,
  UpdateUserRequest,
  UserQueryParams,
  CheckAvailabilityRequest,
  CheckAvailabilityResponse,
} from "@/types/user";

export const userApi = {
  /**
   * 获取用户列表
   */
  getUsers: async (params?: UserQueryParams): Promise<APIResponse<User[]>> => {
    return ApiService.get<User[]>("/users", params);
  },

  /**
   * 根据ID获取用户详情
   */
  getUserById: async (id: number): Promise<APIResponse<User>> => {
    return ApiService.get<User>(`/users/${id}`);
  },

  /**
   * 创建新用户
   */
  createUser: async (data: CreateUserRequest): Promise<APIResponse<User>> => {
    return ApiService.post<User>("/users", data);
  },

  /**
   * 更新用户信息
   */
  updateUser: async (
    id: number,
    data: UpdateUserRequest
  ): Promise<APIResponse<User>> => {
    return ApiService.put<User>(`/users/${id}`, data);
  },

  /**
   * 删除用户
   */
  deleteUser: async (id: number): Promise<APIResponse<void>> => {
    return ApiService.delete<void>(`/users/${id}`);
  },

  /**
   * 批量删除用户
   */
  deleteUsers: async (ids: number[]): Promise<APIResponse<void>> => {
    return ApiService.post<void>("/users/batch-delete", { ids });
  },

  /**
   * 检查用户名或邮箱可用性
   */
  checkAvailability: async (
    data: CheckAvailabilityRequest
  ): Promise<APIResponse<CheckAvailabilityResponse>> => {
    return ApiService.post<CheckAvailabilityResponse>(
      "/users/check-availability",
      data
    );
  },

  /**
   * 获取当前用户信息
   */
  getCurrentUser: async (): Promise<APIResponse<User>> => {
    return ApiService.get<User>("/users/profile");
  },

  /**
   * 更新当前用户信息
   */
  updateCurrentUser: async (
    data: Partial<UpdateUserRequest>
  ): Promise<APIResponse<User>> => {
    return ApiService.put<User>("/users/profile", data);
  },

  /**
   * 检查用户名可用性
   */
  checkUsernameAvailable: async (
    username: string
  ): Promise<APIResponse<CheckAvailabilityResponse>> => {
    return userApi.checkAvailability({ username });
  },

  /**
   * 检查邮箱可用性
   */
  checkEmailAvailable: async (
    email: string
  ): Promise<APIResponse<CheckAvailabilityResponse>> => {
    return userApi.checkAvailability({ email });
  },
} as const;
