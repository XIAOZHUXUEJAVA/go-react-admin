import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Menu, Eye, CheckCircle, Layers } from "lucide-react";
import { Menu as MenuType } from "@/types/menu";

interface MenuStatsCardsProps {
  menus: MenuType[];
}

export const MenuStatsCards: React.FC<MenuStatsCardsProps> = ({ menus }) => {
  // 递归计算所有菜单（包括子菜单）
  const countAllMenus = (menuList: MenuType[]): number => {
    return menuList.reduce((count, menu) => {
      return count + 1 + (menu.children ? countAllMenus(menu.children) : 0);
    }, 0);
  };

  // 递归统计可见菜单
  const countVisibleMenus = (menuList: MenuType[]): number => {
    return menuList.reduce((count, menu) => {
      const current = menu.visible ? 1 : 0;
      const children = menu.children ? countVisibleMenus(menu.children) : 0;
      return count + current + children;
    }, 0);
  };

  // 递归统计启用菜单
  const countActiveMenus = (menuList: MenuType[]): number => {
    return menuList.reduce((count, menu) => {
      const current = menu.status === "active" ? 1 : 0;
      const children = menu.children ? countActiveMenus(menu.children) : 0;
      return count + current + children;
    }, 0);
  };

  const totalMenus = countAllMenus(menus);
  const visibleMenus = countVisibleMenus(menus);
  const activeMenus = countActiveMenus(menus);
  const topLevelMenus = menus.length;

  const stats = [
    {
      title: "总菜单数",
      value: totalMenus,
      icon: Menu,
      description: "系统中的菜单总数",
      color: "text-blue-600",
      bgColor: "bg-blue-50",
    },
    {
      title: "顶级菜单",
      value: topLevelMenus,
      icon: Layers,
      description: "一级菜单数量",
      color: "text-purple-600",
      bgColor: "bg-purple-50",
    },
    {
      title: "可见菜单",
      value: visibleMenus,
      icon: Eye,
      description: "当前可见的菜单",
      color: "text-green-600",
      bgColor: "bg-green-50",
    },
    {
      title: "启用菜单",
      value: activeMenus,
      icon: CheckCircle,
      description: "当前启用的菜单",
      color: "text-emerald-600",
      bgColor: "bg-emerald-50",
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
