"use client";

import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Input } from "@/components/ui/input";
import { ScrollArea } from "@/components/ui/scroll-area";
import * as LucideIcons from "lucide-react";

interface IconPickerProps {
  value: string;
  onChange: (icon: string) => void;
}

// 常用图标列表
const commonIcons = [
  "LayoutDashboard",
  "Users",
  "Shield",
  "Key",
  "Settings2",
  "FileText",
  "Database",
  "Server",
  "Activity",
  "BarChart",
  "PieChart",
  "TrendingUp",
  "Calendar",
  "Mail",
  "Bell",
  "Search",
  "Filter",
  "Download",
  "Upload",
  "Trash2",
  "Edit",
  "Plus",
  "Minus",
  "Check",
  "X",
  "ChevronRight",
  "ChevronLeft",
  "ChevronUp",
  "ChevronDown",
  "Menu",
  "MoreVertical",
  "MoreHorizontal",
  "Home",
  "Folder",
  "File",
  "Image",
  "Video",
  "Music",
  "Code",
  "Terminal",
  "Package",
  "Layers",
  "Grid",
  "List",
  "Table",
  "Columns",
  "Layout",
  "Sidebar",
  "PanelLeft",
  "PanelRight",
];

export function IconPicker({ value, onChange }: IconPickerProps) {
  const [open, setOpen] = useState(false);
  const [search, setSearch] = useState("");

  // 获取图标组件
  const getIconComponent = (iconName: string) => {
    const Icon = LucideIcons[iconName as keyof typeof LucideIcons] as React.ComponentType<{ className?: string }>;
    return Icon ? <Icon className="h-4 w-4" /> : null;
  };

  // 过滤图标
  const filteredIcons = commonIcons.filter((icon) =>
    icon.toLowerCase().includes(search.toLowerCase())
  );

  const SelectedIcon = LucideIcons[value as keyof typeof LucideIcons] as React.ComponentType<{ className?: string }> | undefined;

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button variant="outline" className="w-full justify-start">
          {SelectedIcon && <SelectedIcon className="h-4 w-4 mr-2" />}
          {value || "选择图标"}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-80 p-0" align="start">
        <div className="p-2 border-b">
          <Input
            placeholder="搜索图标..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
        </div>
        <ScrollArea className="h-64">
          <div className="grid grid-cols-6 gap-2 p-2">
            {filteredIcons.map((iconName) => {
              const IconComponent = getIconComponent(iconName);
              return (
                <Button
                  key={iconName}
                  variant={value === iconName ? "default" : "ghost"}
                  size="sm"
                  className="h-10 w-10 p-0"
                  onClick={() => {
                    onChange(iconName);
                    setOpen(false);
                  }}
                  title={iconName}
                >
                  {IconComponent}
                </Button>
              );
            })}
          </div>
        </ScrollArea>
      </PopoverContent>
    </Popover>
  );
}
