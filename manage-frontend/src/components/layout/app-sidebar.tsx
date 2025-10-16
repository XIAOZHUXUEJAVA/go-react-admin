"use client";

import * as React from "react";
import {
  AudioWaveform,
  BookOpen,
  Bot,
  Command,
  Frame,
  GalleryVerticalEnd,
  Map,
  PieChart,
  Settings2,
  SquareTerminal,
  Users,
  LayoutDashboard,
  Database,
  Shield,
  FileText,
  Key,
  Server,
  Activity,
  BarChart,
  TrendingUp,
  Calendar,
  Mail,
  Bell,
  Search,
  Filter,
  Download,
  Upload,
  Trash2,
  Edit,
  Plus,
  Minus,
  Check,
  X,
  ChevronRight,
  ChevronLeft,
  ChevronUp,
  ChevronDown,
  Menu as MenuIcon,
  MoreVertical,
  MoreHorizontal,
  Home,
  Folder,
  File,
  Image,
  Video,
  Music,
  Code,
  Terminal,
  Package,
  Layers,
  Grid,
  List,
  Table,
  Columns,
  Layout,
  Sidebar as SidebarIcon,
  PanelLeft,
  PanelRight,
  Settings,
} from "lucide-react";

import { NavMain } from "./nav-main";
import { NavProjects } from "./nav-projects";
import { NavUser } from "./nav-user";
import { TeamSwitcher } from "./team-switcher";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar";
import { useAuthStore } from "@/stores/authStore";
import { usePermissionStore } from "@/stores/permissionStore";
import { Menu } from "@/types/menu";
import { LucideIcon } from "lucide-react";

// This is sample data.
const data = {
  teams: [
    {
      name: "Acme Inc",
      logo: GalleryVerticalEnd,
      plan: "Enterprise",
    },
    {
      name: "Acme Corp.",
      logo: AudioWaveform,
      plan: "Startup",
    },
    {
      name: "Evil Corp.",
      logo: Command,
      plan: "Free",
    },
  ],
  navMain: [
    {
      title: "仪表板",
      url: "/dashboard",
      icon: LayoutDashboard,
      isActive: true,
      items: [
        {
          title: "概览",
          url: "/dashboard",
        },
        {
          title: "统计",
          url: "/dashboard/analytics",
        },
        {
          title: "报告",
          url: "/dashboard/reports",
        },
      ],
    },
  ],
  projects: [
    {
      name: "个人站点",
      url: "https://piggyblog.xyz/",
      icon: Frame,
    },
    {
      name: "测试热重载",
      url: "https://piggyblog.xyz/",
      icon: Frame,
    },
  ],
};

// 图标映射表
const iconMap: Record<string, LucideIcon> = {
  LayoutDashboard,
  Users,
  Shield,
  Key,
  Settings,
  Settings2,
  FileText,
  Database,
  Server,
  Activity,
  BarChart,
  PieChart,
  TrendingUp,
  Calendar,
  Mail,
  Bell,
  Search,
  Filter,
  Download,
  Upload,
  Trash2,
  Edit,
  Plus,
  Minus,
  Check,
  X,
  ChevronRight,
  ChevronLeft,
  ChevronUp,
  ChevronDown,
  Menu: MenuIcon,
  MoreVertical,
  MoreHorizontal,
  Home,
  Folder,
  File,
  Image,
  Video,
  Music,
  Code,
  Terminal,
  Package,
  Layers,
  Grid,
  List,
  Table,
  Columns,
  Layout,
  Sidebar: SidebarIcon,
  PanelLeft,
  PanelRight,
  Frame,
  Map,
  Bot,
  BookOpen,
  SquareTerminal,
};

// 侧边栏导航项类型
interface NavItem {
  title: string;
  url: string;
  icon: LucideIcon;
  isActive: boolean;
  items: Array<{
    title: string;
    url: string;
    icon?: LucideIcon;
  }>;
}

// 将后端菜单转换为侧边栏格式
function convertMenuToNavItem(menu: Menu): NavItem {
  return {
    title: menu.title,
    url: menu.path,
    icon: iconMap[menu.icon] || LayoutDashboard,
    isActive: false,
    items:
      menu.children
        ?.filter((child) => child.visible && child.status === "active")
        .sort((a, b) => a.order_num - b.order_num)
        .map((child) => ({
          title: child.title,
          url: child.path,
          icon: child.icon ? iconMap[child.icon] : undefined,
        })) || [],
  };
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const { user } = useAuthStore();
  const { userMenus, isLoaded } = usePermissionStore();

  // 构建用户数据，如果没有登录用户则使用默认值
  const userData = {
    name: user?.username || "Guest",
    email: user?.email || "guest@example.com",
    avatar: user?.avatar || "/avatars/default.jpg",
  };

  // 动态菜单：如果权限已加载且有用户菜单，使用动态菜单；否则使用静态菜单
  const navMainItems = React.useMemo(() => {
    if (isLoaded && userMenus.length > 0) {
      // 只显示顶级菜单（parent_id 为 null）
      return userMenus
        .filter(
          (menu) =>
            menu.parent_id === null && menu.visible && menu.status === "active"
        )
        .sort((a, b) => a.order_num - b.order_num)
        .map(convertMenuToNavItem);
    }
    // 使用静态菜单作为后备
    return data.navMain;
  }, [isLoaded, userMenus]);

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <TeamSwitcher teams={data.teams} />
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={navMainItems} />
        <NavProjects projects={data.projects} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={userData} />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}
