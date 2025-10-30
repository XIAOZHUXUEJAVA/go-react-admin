"use client";

import * as React from "react";
import { ThemeProvider as NextThemesProvider } from "next-themes";

type ThemeProviderProps = Parameters<typeof NextThemesProvider>[0];

/**
 * 主题提供者组件
 * 封装 next-themes 的 ThemeProvider
 */
export function ThemeProvider({ children, ...props }: ThemeProviderProps) {
  return <NextThemesProvider {...props}>{children}</NextThemesProvider>;
}
