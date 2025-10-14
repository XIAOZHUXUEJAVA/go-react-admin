/**
 * Token管理工具函数
 * 支持双token（Access Token + Refresh Token）认证系统
 */

/**
 * 验证 JWT token 格式
 */
export function isValidJWTFormat(token: string): boolean {
  const jwtRegex = /^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]*$/;
  return jwtRegex.test(token);
}

// JWT Payload 类型定义
interface JWTPayload {
  user_id: number;
  username: string;
  role: string;
  exp: number; // 过期时间戳
  iat: number; // 签发时间戳
  jti: string; // JWT ID
  [key: string]: unknown; // 允许其他字段
}

// 用户信息类型（从token解析）
interface TokenUserInfo {
  id: number;
  username: string;
  role: string;
  exp: number;
  iat: number;
  jti: string;
}

/**
 * 获取存储的访问token
 * @returns access token字符串或null
 */
export function getAccessToken(): string | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem("access-token");
}

/**
 * 获取存储的刷新token
 * @returns refresh token字符串或null
 */
export function getRefreshToken(): string | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem("refresh-token");
}

/**
 * 获取token过期时间
 * @returns 过期时间戳或null
 */
export function getTokenExpiresAt(): number | null {
  if (typeof window === "undefined") return null;
  const expiresAt = localStorage.getItem("token-expires-at");
  return expiresAt ? parseInt(expiresAt, 10) : null;
}

/**
 * 设置认证tokens
 * @param accessToken - JWT access token
 * @param refreshToken - JWT refresh token
 * @param expiresIn - access token过期时间（秒）
 */
export function setTokens(
  accessToken: string,
  refreshToken: string,
  expiresIn: number
): void {
  if (typeof window === "undefined") return;

  const expiresAt = Date.now() + expiresIn * 1000;

  localStorage.setItem("access-token", accessToken);
  localStorage.setItem("refresh-token", refreshToken);
  localStorage.setItem("token-expires-at", expiresAt.toString());
}

/**
 * 更新访问token
 * @param accessToken - 新的access token
 * @param expiresIn - 过期时间（秒）
 */
export function updateAccessToken(
  accessToken: string,
  expiresIn: number
): void {
  if (typeof window === "undefined") return;

  const expiresAt = Date.now() + expiresIn * 1000;

  localStorage.setItem("access-token", accessToken);
  localStorage.setItem("token-expires-at", expiresAt.toString());
}

/**
 * 移除所有认证tokens
 */
export function removeTokens(): void {
  if (typeof window === "undefined") return;

  localStorage.removeItem("access-token");
  localStorage.removeItem("refresh-token");
  localStorage.removeItem("token-expires-at");
  localStorage.removeItem("auth-storage");
}

/**
 * 检查access token是否有效
 * @returns 是否有效
 */
export function isAccessTokenValid(): boolean {
  const token = getAccessToken();
  const expiresAt = getTokenExpiresAt();

  if (!token || !expiresAt) return false;

  // 检查是否过期（提前30秒判断为过期）
  return Date.now() < expiresAt - 30 * 1000;
}

/**
 * 检查用户是否已认证（有有效的access token或可用的refresh token）
 * @returns 是否已认证
 */
export function isAuthenticated(): boolean {
  const accessToken = getAccessToken();
  const refreshToken = getRefreshToken();

  // 如果有有效的access token，直接返回true
  if (accessToken && isAccessTokenValid()) {
    return true;
  }

  // 如果access token无效但有refresh token，也认为是已认证状态
  // 实际的token刷新会在API调用时自动处理
  return !!refreshToken;
}

/**
 * 检查access token是否即将过期（5分钟内）
 * @returns 是否即将过期
 */
export function isTokenExpiringSoon(): boolean {
  const expiresAt = getTokenExpiresAt();
  if (!expiresAt) return false;

  const fiveMinutesFromNow = Date.now() + 5 * 60 * 1000;
  return expiresAt < fiveMinutesFromNow;
}

/**
 * 检查指定token是否过期（基于JWT payload）
 * @param token - JWT token
 * @returns 是否过期
 */
export function isTokenExpired(token: string): boolean {
  const payload = parseJWTPayload(token);
  if (!payload || !payload.exp) {
    return true;
  }

  const currentTime = Math.floor(Date.now() / 1000);
  return payload.exp < currentTime;
}

/**
 * 解析JWT payload（不验证签名，仅用于客户端显示）
 * @param token - JWT token
 * @returns payload对象或null
 */
export function parseJWTPayload(token: string): JWTPayload | null {
  try {
    const base64Url = token.split(".")[1];
    if (!base64Url) return null;

    const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split("")
        .map((c) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
        .join("")
    );
    return JSON.parse(jsonPayload) as JWTPayload;
  } catch (error) {
    return null;
  }
}

/**
 * 获取当前用户信息（从access token中解析）
 * @returns 用户信息或null
 */
export function getCurrentUserFromToken(): TokenUserInfo | null {
  const accessToken = getAccessToken();
  if (!accessToken) return null;

  const payload = parseJWTPayload(accessToken);
  return payload
    ? {
        id: payload.user_id,
        username: payload.username,
        role: payload.role,
        exp: payload.exp,
        iat: payload.iat,
        jti: payload.jti,
      }
    : null;
}

// 向后兼容的函数别名
export const getToken = getAccessToken;
export const removeToken = removeTokens;
