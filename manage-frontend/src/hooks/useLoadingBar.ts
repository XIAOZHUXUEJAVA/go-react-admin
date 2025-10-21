"use client";

import { useCallback } from "react";
import NProgress from "nprogress";

/**
 * 加载进度条 Hook
 * 
 * 提供手动控制进度条的方法，适用于：
 * - 异步数据加载
 * - 表单提交
 * - 文件上传
 * - 其他耗时操作
 * 
 * @example
 * ```tsx
 * const { start, done, set } = useLoadingBar();
 * 
 * const handleSubmit = async () => {
 *   start();
 *   try {
 *     await api.submit();
 *     done();
 *   } catch (error) {
 *     done();
 *   }
 * };
 * ```
 */
export function useLoadingBar() {
  const start = useCallback(() => {
    NProgress.start();
  }, []);

  const done = useCallback(() => {
    NProgress.done();
  }, []);

  const set = useCallback((n: number) => {
    NProgress.set(n);
  }, []);

  const inc = useCallback((amount?: number) => {
    NProgress.inc(amount);
  }, []);

  return {
    /**
     * 开始显示进度条
     */
    start,
    /**
     * 完成并隐藏进度条
     */
    done,
    /**
     * 设置进度条到指定百分比 (0-1)
     */
    set,
    /**
     * 增加进度条百分比
     */
    inc,
  };
}

/**
 * 包装异步函数，自动显示/隐藏进度条
 * 
 * @example
 * ```tsx
 * const { start, done, wrap } = useLoadingBar();
 * 
 * const handleFetch = wrap(async () => {
 *   const data = await fetchData();
 *   return data;
 * });
 * ```
 */
export function useLoadingBarWrapper() {
  const { start, done } = useLoadingBar();

  const wrap = useCallback(
    <T>(fn: () => Promise<T>) => {
      return async (): Promise<T> => {
        start();
        try {
          const result = await fn();
          done();
          return result;
        } catch (error) {
          done();
          throw error;
        }
      };
    },
    [start, done]
  );

  return {
    start,
    done,
    /**
     * 包装异步函数，自动处理进度条
     */
    wrap,
  };
}
