// 认证相关类型定义

export interface LoginRequest {
  username: string;
  password: string;
  captcha_id?: string;
  captcha_code?: string;
}

export interface LoginResponse {
  access_token: string;
  refresh_token: string;
  expires_in: number; // Access token expiration in seconds
  refresh_expires_in: number; // Refresh token expiration in seconds
  token_type: string; // Always "Bearer"
  user: AuthUser;
}

// 刷新token请求
export interface RefreshTokenRequest {
  refresh_token: string;
}

// 刷新token响应
export interface RefreshTokenResponse {
  access_token: string;
  expires_in: number;
  token_type: string;
}

// 登出请求
export interface LogoutRequest {
  refresh_token?: string;
}

export interface AuthUser {
  id: number;
  username: string;
  email: string;
  role: string;
  status: string;
  avatar?: string;
  created_at: string;
  updated_at: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
  role?: string;
}

export interface AuthState {
  user: AuthUser | null;
  accessToken: string | null;
  refreshToken: string | null;
  tokenExpiresAt: number | null; // Access token expiration timestamp
  isAuthenticated: boolean;
  isLoading: boolean;
}

export interface AuthActions {
  login: (credentials: LoginRequest) => Promise<void>;
  logout: () => void;
  register: (data: RegisterRequest) => Promise<void>;
  checkAuth: () => Promise<void>;
  setLoading: (loading: boolean) => void;
}

// 验证码相关类型
export interface CaptchaResponse {
  captcha_id: string;
  captcha_data: string; // base64 图片数据
}

export interface CaptchaState {
  captchaId: string | null;
  captchaImage: string | null;
  isLoading: boolean;
  error: string | null;
}

export interface CaptchaActions {
  generateCaptcha: () => Promise<void>;
  refreshCaptcha: () => Promise<void>;
  clearCaptcha: () => void;
  setLoading: (loading: boolean) => void;
  setError: (error: string | null) => void;
}

export type CaptchaStore = CaptchaState & CaptchaActions;

export type AuthStore = AuthState & AuthActions;
