"use client";

import { useEffect, useRef } from "react";
import { usePathname, useSearchParams } from "next/navigation";
import NProgress from "nprogress";
import "nprogress/nprogress.css";

/**
 * 顶部加载进度条配置
 */
NProgress.configure({
  showSpinner: false, // 不显示右上角的旋转图标
  trickleSpeed: 200, // 自动递增间隔
  minimum: 0.08, // 最小百分比
  easing: "ease", // 动画效果
  speed: 200, // 递增进度条的速度
});

/**
 * 顶部加载进度条组件
 * 
 * 功能：
 * - 监听路由变化，自动显示/隐藏进度条
 * - 页面切换时提供视觉反馈
 * - 提升用户体验
 * 
 * 使用方式：
 * 在根布局中引入即可，无需传递任何 props
 * 
 * @example
 * ```tsx
 * // app/layout.tsx
 * import { TopLoadingBar } from "@/components/common";
 * 
 * export default function RootLayout({ children }) {
 *   return (
 *     <html>
 *       <body>
 *         <TopLoadingBar />
 *         {children}
 *       </body>
 *     </html>
 *   );
 * }
 * ```
 */
export function TopLoadingBar() {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const prevPathname = useRef<string | null>(null);

  useEffect(() => {
    // 首次加载时记录路径
    if (prevPathname.current === null) {
      prevPathname.current = pathname;
      return;
    }

    // 路径变化时，先启动进度条
    if (prevPathname.current !== pathname) {
      NProgress.start();
      prevPathname.current = pathname;
    }

    // 页面加载完成后隐藏进度条
    NProgress.done();
  }, [pathname, searchParams]);

  useEffect(() => {
    // 监听链接点击事件
    const handleAnchorClick = (event: MouseEvent) => {
      const target = event.target as HTMLElement;
      const anchor = target.closest("a");
      
      if (anchor && anchor.href) {
        const currentUrl = window.location.href;
        const targetUrl = anchor.href;
        
        // 如果是站内链接且不是当前页面，启动进度条
        if (
          targetUrl.startsWith(window.location.origin) &&
          targetUrl !== currentUrl &&
          !anchor.target // 不是新窗口打开
        ) {
          NProgress.start();
        }
      }
    };

    // 监听浏览器前进后退
    const handlePopState = () => {
      NProgress.start();
    };

    document.addEventListener("click", handleAnchorClick);
    window.addEventListener("popstate", handlePopState);

    return () => {
      document.removeEventListener("click", handleAnchorClick);
      window.removeEventListener("popstate", handlePopState);
      NProgress.remove();
    };
  }, []);

  return null;
}

/**
 * 自定义样式的顶部加载进度条
 * 
 * 提供自定义颜色和高度
 */
interface CustomTopLoadingBarProps {
  /**
   * 进度条颜色
   * @default "#3b82f6" (blue-500)
   */
  color?: string;
  /**
   * 进度条高度
   * @default "3px"
   */
  height?: string;
}

export function CustomTopLoadingBar({
  color = "#3b82f6",
  height = "3px",
}: CustomTopLoadingBarProps) {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const prevPathname = useRef<string | null>(null);

  useEffect(() => {
    // 首次加载时记录路径
    if (prevPathname.current === null) {
      prevPathname.current = pathname;
      return;
    }

    // 路径变化时，先启动进度条
    if (prevPathname.current !== pathname) {
      NProgress.start();
      prevPathname.current = pathname;
    }

    // 页面加载完成后隐藏进度条
    NProgress.done();
  }, [pathname, searchParams]);

  useEffect(() => {
    // 动态注入自定义样式
    const style = document.createElement("style");
    style.innerHTML = `
      #nprogress .bar {
        background: ${color} !important;
        height: ${height} !important;
      }
      #nprogress .peg {
        box-shadow: 0 0 10px ${color}, 0 0 5px ${color} !important;
      }
    `;
    document.head.appendChild(style);

    // 监听链接点击事件
    const handleAnchorClick = (event: MouseEvent) => {
      const target = event.target as HTMLElement;
      const anchor = target.closest("a");
      
      if (anchor && anchor.href) {
        const currentUrl = window.location.href;
        const targetUrl = anchor.href;
        
        if (
          targetUrl.startsWith(window.location.origin) &&
          targetUrl !== currentUrl &&
          !anchor.target
        ) {
          NProgress.start();
        }
      }
    };

    const handlePopState = () => {
      NProgress.start();
    };

    document.addEventListener("click", handleAnchorClick);
    window.addEventListener("popstate", handlePopState);

    return () => {
      document.head.removeChild(style);
      document.removeEventListener("click", handleAnchorClick);
      window.removeEventListener("popstate", handlePopState);
      NProgress.remove();
    };
  }, [color, height]);

  return null;
}
