import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

/**
 * Next.js 中间件
 * 简化版本，主要处理服务端重定向
 */
export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;

  // 跳过客户端路由，让 AuthGuard 处理
  if (pathname.startsWith("/_next") || pathname.includes(".")) {
    return NextResponse.next();
  }

  // 只处理明确的服务端重定向需求

  return NextResponse.next();
}

// 配置中间件匹配的路径
export const config = {
  matcher: [
    /*
     * 匹配所有路径除了:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     * - public files (public directory)
     */
    "/((?!api|_next/static|_next/image|favicon.ico|.*\\.(?:svg|png|jpg|jpeg|gif|webp)$).*)",
  ],
};
