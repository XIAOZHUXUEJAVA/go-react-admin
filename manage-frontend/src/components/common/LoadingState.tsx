import React from "react";
import { Loader2 } from "lucide-react";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

type LoadingMode = "spinner" | "text" | "skeleton" | "table-skeleton";

interface LoadingStateProps {
  /**
   * 加载状态展示模式
   * - spinner: 简单的旋转图标
   * - text: 旋转图标 + 文字说明
   * - skeleton: 骨架屏（通用）
   * - table-skeleton: 表格骨架屏
   */
  mode?: LoadingMode;
  /**
   * 加载文字提示
   * @default "加载中..."
   */
  text?: string;
  /**
   * 骨架屏行数（仅在 table-skeleton 模式下有效）
   * @default 5
   */
  rows?: number;
  /**
   * 表格列配置（仅在 table-skeleton 模式下有效）
   * 定义每列的标题和骨架屏宽度
   */
  columns?: Array<{
    header: string;
    skeletonWidth?: string;
    skeletonHeight?: string;
    skeletonClassName?: string;
  }>;
  /**
   * 自定义容器类名
   */
  className?: string;
  /**
   * 容器内边距
   * @default "py-8"
   */
  padding?: string;
}

/**
 * 统一的加载状态组件
 * 
 * 支持多种展示模式：
 * - spinner: 简单的旋转图标（适用于简单列表）
 * - text: 旋转图标 + 文字说明（适用于需要提示的场景）
 * - skeleton: 通用骨架屏（适用于卡片等）
 * - table-skeleton: 表格骨架屏（适用于表格数据）
 * 
 * @example
 * ```tsx
 * // 简单 spinner
 * <LoadingState mode="spinner" />
 * 
 * // 带文字的 spinner
 * <LoadingState mode="text" text="正在加载数据..." />
 * 
 * // 表格骨架屏
 * <LoadingState
 *   mode="table-skeleton"
 *   rows={5}
 *   columns={[
 *     { header: "用户", skeletonWidth: "w-32" },
 *     { header: "邮箱", skeletonWidth: "w-40" },
 *     { header: "角色", skeletonWidth: "w-20" },
 *   ]}
 * />
 * ```
 */
export const LoadingState: React.FC<LoadingStateProps> = ({
  mode = "spinner",
  text = "加载中...",
  rows = 5,
  columns = [],
  className = "",
  padding = "py-8",
}) => {
  // Spinner 模式：简单的旋转图标
  if (mode === "spinner") {
    return (
      <div className={`flex items-center justify-center ${padding} ${className}`}>
        <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
      </div>
    );
  }

  // Text 模式：旋转图标 + 文字
  if (mode === "text") {
    return (
      <div className={`text-center ${padding} ${className}`}>
        <Loader2 className="h-8 w-8 animate-spin text-muted-foreground mx-auto" />
        <p className="mt-2 text-sm text-muted-foreground">{text}</p>
      </div>
    );
  }

  // Skeleton 模式：通用骨架屏
  if (mode === "skeleton") {
    return (
      <div className={`space-y-4 ${padding} ${className}`}>
        {Array.from({ length: rows }).map((_, index) => (
          <div key={index} className="space-y-2">
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-4 w-3/4" />
          </div>
        ))}
      </div>
    );
  }

  // Table Skeleton 模式：表格骨架屏
  if (mode === "table-skeleton") {
    return (
      <div className={`rounded-md border ${className}`}>
        <Table>
          <TableHeader>
            <TableRow>
              {columns.map((col, index) => (
                <TableHead key={index}>{col.header}</TableHead>
              ))}
            </TableRow>
          </TableHeader>
          <TableBody>
            {Array.from({ length: rows }).map((_, rowIndex) => (
              <TableRow key={rowIndex}>
                {columns.map((col, colIndex) => (
                  <TableCell key={colIndex}>
                    <Skeleton
                      className={
                        col.skeletonClassName ||
                        `h-${col.skeletonHeight || "4"} ${
                          col.skeletonWidth || "w-24"
                        }`
                      }
                    />
                  </TableCell>
                ))}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    );
  }

  return null;
};

/**
 * 表格加载骨架屏组件（预设配置）
 * 
 * 为常见的表格场景提供预设的骨架屏配置
 */
interface TableLoadingSkeletonProps {
  /**
   * 骨架屏行数
   * @default 5
   */
  rows?: number;
  /**
   * 列配置
   */
  columns: Array<{
    header: string;
    /**
     * 骨架屏内容配置
     * - 可以是简单的宽度类名，如 "w-32"
     * - 或者是复杂的 JSX 元素
     */
    skeleton: string | React.ReactNode;
  }>;
  /**
   * 自定义类名
   */
  className?: string;
}

export const TableLoadingSkeleton: React.FC<TableLoadingSkeletonProps> = ({
  rows = 5,
  columns,
  className = "",
}) => {
  return (
    <div className={`rounded-md border ${className}`}>
      <Table>
        <TableHeader>
          <TableRow>
            {columns.map((col, index) => (
              <TableHead key={index}>{col.header}</TableHead>
            ))}
          </TableRow>
        </TableHeader>
        <TableBody>
          {Array.from({ length: rows }).map((_, rowIndex) => (
            <TableRow key={rowIndex}>
              {columns.map((col, colIndex) => (
                <TableCell key={colIndex}>
                  {typeof col.skeleton === "string" ? (
                    <Skeleton className={`h-4 ${col.skeleton}`} />
                  ) : (
                    col.skeleton
                  )}
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};
