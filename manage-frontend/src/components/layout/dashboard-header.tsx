"use client";

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Separator } from "@/components/ui/separator";
import { SidebarTrigger } from "@/components/ui/sidebar";

interface BreadcrumbItem {
  label: string;
  href?: string;
}

interface DashboardHeaderProps {
  breadcrumbs: BreadcrumbItem[];
  actions?: React.ReactNode;
}

/**
 * Dashboard 页面头部组件
 * 提供面包屑导航和操作按钮区域
 */
export function DashboardHeader({
  breadcrumbs,
  actions,
}: DashboardHeaderProps) {
  return (
    <header className="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12">
      <div className="flex items-center gap-2 px-4 flex-1">
        <SidebarTrigger className="-ml-1" />
        <Separator
          orientation="vertical"
          className="mr-2 data-[orientation=vertical]:h-4"
        />
        <Breadcrumb>
          <BreadcrumbList>
            {breadcrumbs.map((item, index) => (
              <div key={index} className="flex items-center">
                {index > 0 && (
                  <BreadcrumbSeparator className="hidden md:block" />
                )}
                <BreadcrumbItem
                  className={index === 0 ? "hidden md:block" : ""}
                >
                  {item.href ? (
                    <BreadcrumbLink href={item.href}>
                      {item.label}
                    </BreadcrumbLink>
                  ) : (
                    <BreadcrumbPage>{item.label}</BreadcrumbPage>
                  )}
                </BreadcrumbItem>
              </div>
            ))}
          </BreadcrumbList>
        </Breadcrumb>
      </div>

      {/* 操作按钮区域 */}
      {actions && (
        <div className="flex items-center space-x-4 px-4">{actions}</div>
      )}
    </header>
  );
}
