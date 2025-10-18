import * as React from "react";
import { cva, type VariantProps } from "class-variance-authority";
import { cn } from "@/lib/utils";

const methodBadgeVariants = cva(
  "inline-flex items-center justify-center rounded-md px-2 py-0.5 text-xs font-medium transition-colors",
  {
    variants: {
      method: {
        GET: "bg-blue-100 text-blue-800 hover:bg-blue-200",
        POST: "bg-green-100 text-green-800 hover:bg-green-200",
        PUT: "bg-yellow-100 text-yellow-800 hover:bg-yellow-200",
        DELETE: "bg-red-100 text-red-800 hover:bg-red-200",
        PATCH: "bg-purple-100 text-purple-800 hover:bg-purple-200",
      },
    },
    defaultVariants: {
      method: "GET",
    },
  }
);

export interface MethodBadgeProps
  extends React.HTMLAttributes<HTMLSpanElement>,
    VariantProps<typeof methodBadgeVariants> {
  method: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
}

const MethodBadge = React.forwardRef<HTMLSpanElement, MethodBadgeProps>(
  ({ className, method, children, ...props }, ref) => {
    return (
      <span
        ref={ref}
        className={cn(methodBadgeVariants({ method }), className)}
        {...props}
      >
        {children || method}
      </span>
    );
  }
);

MethodBadge.displayName = "MethodBadge";

export { MethodBadge, methodBadgeVariants };
