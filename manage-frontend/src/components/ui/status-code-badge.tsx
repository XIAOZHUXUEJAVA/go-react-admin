import * as React from "react";
import { cn } from "@/lib/utils";

export interface StatusCodeBadgeProps
  extends React.HTMLAttributes<HTMLSpanElement> {
  statusCode: number;
}

const getStatusCodeColor = (status: number): string => {
  if (status >= 200 && status < 300) return "bg-green-100 text-green-800";
  if (status >= 300 && status < 400) return "bg-blue-100 text-blue-800";
  if (status >= 400 && status < 500) return "bg-yellow-100 text-yellow-800";
  if (status >= 500) return "bg-red-100 text-red-800";
  return "bg-gray-100 text-gray-800";
};

const StatusCodeBadge = React.forwardRef<HTMLSpanElement, StatusCodeBadgeProps>(
  ({ className, statusCode, children, ...props }, ref) => {
    return (
      <span
        ref={ref}
        className={cn(
          "inline-flex items-center justify-center rounded-md px-2 py-0.5 text-xs font-medium transition-colors",
          getStatusCodeColor(statusCode),
          className
        )}
        {...props}
      >
        {children || statusCode}
      </span>
    );
  }
);

StatusCodeBadge.displayName = "StatusCodeBadge";

export { StatusCodeBadge };
