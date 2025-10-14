import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Shield, CheckCircle, XCircle, Users } from "lucide-react";
import { Role } from "@/types/role";
import { PaginationInfo } from "@/types/common";

interface RoleStatsCardsProps {
  roles: Role[];
  pagination: PaginationInfo | null;
}

export const RoleStatsCards: React.FC<RoleStatsCardsProps> = ({
  roles,
  pagination,
}) => {
  const totalRoles = pagination?.total || roles.length;
  const activeRoles = roles.filter((role) => role.status === "active").length;
  const inactiveRoles = roles.filter((role) => role.status === "inactive")
    .length;
  const systemRoles = roles.filter((role) => role.is_system).length;

  const stats = [
    {
      title: "总角色数",
      value: totalRoles,
      icon: Shield,
      description: "系统中的角色总数",
      color: "text-blue-600",
      bgColor: "bg-blue-50",
    },
    {
      title: "启用角色",
      value: activeRoles,
      icon: CheckCircle,
      description: "当前启用的角色",
      color: "text-green-600",
      bgColor: "bg-green-50",
    },
    {
      title: "禁用角色",
      value: inactiveRoles,
      icon: XCircle,
      description: "当前禁用的角色",
      color: "text-red-600",
      bgColor: "bg-red-50",
    },
    {
      title: "系统角色",
      value: systemRoles,
      icon: Users,
      description: "系统预设角色",
      color: "text-purple-600",
      bgColor: "bg-purple-50",
    },
  ];

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
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
