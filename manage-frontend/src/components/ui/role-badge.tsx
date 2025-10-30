import * as React from "react";
import { cva, type VariantProps } from "class-variance-authority";
import { cn } from "@/lib/utils";

const roleBadgeVariants = cva(
  "inline-flex items-center justify-center rounded-md px-2 py-0.5 text-xs font-medium transition-colors",
  {
    variants: {
      role: {
        admin: "bg-red-100 text-red-800 hover:bg-red-200",
        manager: "bg-blue-100 text-blue-800 hover:bg-blue-200",
        user: "bg-green-100 text-green-800 hover:bg-green-200",
      },
    },
    defaultVariants: {
      role: "user",
    },
  }
);

export interface RoleBadgeProps
  extends React.HTMLAttributes<HTMLSpanElement>,
    VariantProps<typeof roleBadgeVariants> {
  role: "admin" | "manager" | "user";
}

const RoleBadge = React.forwardRef<HTMLSpanElement, RoleBadgeProps>(
  ({ className, role, children, ...props }, ref) => {
    return (
      <span
        ref={ref}
        className={cn(roleBadgeVariants({ role }), className)}
        {...props}
      >
        {children || role}
      </span>
    );
  }
);

RoleBadge.displayName = "RoleBadge";

export { RoleBadge, roleBadgeVariants };
