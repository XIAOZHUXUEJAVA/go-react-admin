import * as React from "react";
import { cva, type VariantProps } from "class-variance-authority";
import { cn } from "@/lib/utils";

const permissionTypeBadgeVariants = cva(
  "inline-flex items-center justify-center rounded-md px-2 py-0.5 text-xs font-medium transition-colors",
  {
    variants: {
      type: {
        api: "bg-purple-100 text-purple-800 hover:bg-purple-200",
        menu: "bg-green-100 text-green-800 hover:bg-green-200",
        button: "bg-orange-100 text-orange-800 hover:bg-orange-200",
      },
    },
    defaultVariants: {
      type: "api",
    },
  }
);

export interface PermissionTypeBadgeProps
  extends React.HTMLAttributes<HTMLSpanElement>,
    VariantProps<typeof permissionTypeBadgeVariants> {
  type: "api" | "menu" | "button";
}

const PermissionTypeBadge = React.forwardRef<
  HTMLSpanElement,
  PermissionTypeBadgeProps
>(({ className, type, children, ...props }, ref) => {
  return (
    <span
      ref={ref}
      className={cn(permissionTypeBadgeVariants({ type }), className)}
      {...props}
    >
      {children || type.toUpperCase()}
    </span>
  );
});

PermissionTypeBadge.displayName = "PermissionTypeBadge";

export { PermissionTypeBadge, permissionTypeBadgeVariants };
