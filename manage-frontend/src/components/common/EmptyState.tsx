import React from "react";
import { LucideIcon } from "lucide-react";
import { cn } from "@/lib/utils";

interface EmptyStateProps {
  /**
   * 图标组件
   */
  icon?: LucideIcon;
  /**
   * 主标题
   */
  title?: string;
  /**
   * 描述文本
   */
  description?: string;
  /**
   * 自定义操作按钮
   */
  action?: React.ReactNode;
  /**
   * 自定义类名
   */
  className?: string;
  /**
   * 图标大小
   * @default "default"
   */
  iconSize?: "sm" | "default" | "lg";
}

/**
 * 通用空状态组件
 * 
 * @example
 * ```tsx
 * // 基本用法
 * <EmptyState
 *   icon={Users}
 *   title="暂无用户数据"
 *   description="还没有创建任何用户"
 * />
 * 
 * // 带操作按钮
 * <EmptyState
 *   icon={FileText}
 *   title="暂无数据"
 *   action={<Button>创建新数据</Button>}
 * />
 * 
 * // 在表格中使用
 * {data.length === 0 ? (
 *   <TableRow>
 *     <TableCell colSpan={columns.length}>
 *       <EmptyState icon={Users} title="暂无用户数据" />
 *     </TableCell>
 *   </TableRow>
 * ) : (
 *   // ... 数据行
 * )}
 * ```
 */
export const EmptyState: React.FC<EmptyStateProps> = ({
  icon: Icon,
  title = "暂无数据",
  description,
  action,
  className,
  iconSize = "default",
}) => {
  const iconSizeClasses = {
    sm: "h-8 w-8",
    default: "h-12 w-12",
    lg: "h-16 w-16",
  };

  return (
    <div
      className={cn(
        "flex flex-col items-center justify-center py-8 text-center",
        className
      )}
    >
      {Icon && (
        <Icon
          className={cn(
            "mb-3 text-muted-foreground opacity-50",
            iconSizeClasses[iconSize]
          )}
        />
      )}
      <p className="text-base font-medium text-muted-foreground mb-1">
        {title}
      </p>
      {description && (
        <p className="text-sm text-muted-foreground/80 max-w-sm">
          {description}
        </p>
      )}
      {action && <div className="mt-4">{action}</div>}
    </div>
  );
};

/**
 * 表格空状态组件（专门用于表格中）
 * 
 * @example
 * ```tsx
 * <TableBody>
 *   {data.length === 0 ? (
 *     <TableEmptyState
 *       colSpan={5}
 *       icon={Users}
 *       title="暂无用户数据"
 *     />
 *   ) : (
 *     data.map(item => <TableRow key={item.id}>...</TableRow>)
 *   )}
 * </TableBody>
 * ```
 */
interface TableEmptyStateProps extends EmptyStateProps {
  /**
   * 表格列数（用于 colSpan）
   */
  colSpan: number;
}

export const TableEmptyState: React.FC<TableEmptyStateProps> = ({
  colSpan,
  ...emptyStateProps
}) => {
  return (
    <>
      <tr>
        <td colSpan={colSpan} className="p-0">
          <EmptyState {...emptyStateProps} />
        </td>
      </tr>
    </>
  );
};
