import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Key, Globe, Menu, MousePointer, CheckCircle, XCircle } from "lucide-react";
import { Permission } from "@/types/permission";

interface PermissionStatsCardsProps {
  permissions: Permission[];
}

export const PermissionStatsCards: React.FC<PermissionStatsCardsProps> = ({
  permissions,
}) => {
  const totalPermissions = permissions.length;
  const apiPermissions = permissions.filter((p) => p.type === "api").length;
  const menuPermissions = permissions.filter((p) => p.type === "menu").length;
  const buttonPermissions = permissions.filter((p) => p.type === "button").length;
  const activePermissions = permissions.filter((p) => p.status === "active").length;
  const inactivePermissions = permissions.filter((p) => p.status === "inactive").length;

  const stats = [
    {
      title: "总权限数",
      value: totalPermissions,
      icon: Key,
      description: "系统中的权限总数",
      color: "text-blue-600",
      bgColor: "bg-blue-50",
    },
    {
      title: "API 权限",
      value: apiPermissions,
      icon: Globe,
      description: "接口访问权限",
      color: "text-purple-600",
      bgColor: "bg-purple-50",
    },
    {
      title: "菜单权限",
      value: menuPermissions,
      icon: Menu,
      description: "菜单访问权限",
      color: "text-green-600",
      bgColor: "bg-green-50",
    },
    {
      title: "按钮权限",
      value: buttonPermissions,
      icon: MousePointer,
      description: "按钮操作权限",
      color: "text-orange-600",
      bgColor: "bg-orange-50",
    },
    {
      title: "启用权限",
      value: activePermissions,
      icon: CheckCircle,
      description: "当前启用的权限",
      color: "text-emerald-600",
      bgColor: "bg-emerald-50",
    },
    {
      title: "禁用权限",
      value: inactivePermissions,
      icon: XCircle,
      description: "当前禁用的权限",
      color: "text-red-600",
      bgColor: "bg-red-50",
    },
  ];

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
      {stats.map((stat) => {
        const Icon = stat.icon;
        return (
          <Card key={stat.title}>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">
                {stat.title}
              </CardTitle>
              <div className={`${stat.bgColor} p-2 rounded-lg`}>
                <Icon className={`h-4 w-4 ${stat.color}`} />
              </div>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{stat.value}</div>
              <p className="text-xs text-muted-foreground mt-1">
                {stat.description}
              </p>
            </CardContent>
          </Card>
        );
      })}
    </div>
  );
};
