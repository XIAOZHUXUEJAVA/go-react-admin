import { NextRequest } from "next/server";
import {
  parseJWTPayload,
  isValidJWTFormat,
  isTokenExpired,
} from "./tokenUtils";

/**
 * 认证工具函数 - 专注于路由保护和请求处理
 */

/**
 * 从请求中获取 token
 */
export function getTokenFromRequest(request: NextRequest): string | null {
  // 从 Authorization header 获取
  const authHeader = request.headers.get("authorization");
  if (authHeader && authHeader.startsWith("Bearer ")) {
    return authHeader.substring(7);
  }

  // 从 cookie 获取
  const tokenCookie = request.cookies.get("auth-token");
  if (tokenCookie) {
    return tokenCookie.value;
  }

  return null;
}

/**
 * 检查是否为受保护的路由
 */
export function isProtectedRoute(pathname: string): boolean {
  const protectedRoutes = ["/dashboard", "/users", "/profile", "/settings"];

  // 精确匹配首页
  if (pathname === "/") {
    return true;
  }

  // 其他路由使用 startsWith 匹配
  return protectedRoutes.some((route) => pathname.startsWith(route));
}

/**
 * 检查是否为认证路由（登录、注册等）
 */
export function isAuthRoute(pathname: string): boolean {
  const authRoutes = ["/login", "/register", "/forgot-password"];
  return authRoutes.includes(pathname);
}

// JWT相关函数现在从 tokenUtils 导入，避免重复定义
// 可以直接使用: isValidJWTFormat, parseJWTPayload, isTokenExpired

// 重新导出以保持向后兼容性
export { isValidJWTFormat, parseJWTPayload, isTokenExpired };
