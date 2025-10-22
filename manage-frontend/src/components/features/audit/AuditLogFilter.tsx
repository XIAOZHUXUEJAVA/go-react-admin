import React from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Search, X } from "lucide-react";
import { AuditLogQuery } from "@/types/audit";

interface AuditLogFilterProps {
  filters: AuditLogQuery;
  onFilterChange: (filters: AuditLogQuery) => void;
  onReset: () => void;
}

export const AuditLogFilter: React.FC<AuditLogFilterProps> = ({
  filters,
  onFilterChange,
  onReset,
}) => {
  const handleInputChange = (field: keyof AuditLogQuery, value: string) => {
    onFilterChange({
      ...filters,
      [field]: value || undefined,
    });
  };

  const handleSelectChange = (field: keyof AuditLogQuery, value: string) => {
    if (value === "all") {
      const newFilters = { ...filters };
      delete newFilters[field];
      onFilterChange(newFilters);
    } else {
      onFilterChange({
        ...filters,
        [field]: field === "status" ? parseInt(value) : value,
      });
    }
  };

  const hasActiveFilters = Object.keys(filters).some(
    (key) =>
      key !== "page" &&
      key !== "page_size" &&
      filters[key as keyof AuditLogQuery] !== undefined
  );

  return (
    <Card>
      <CardContent className="pt-6">
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
          {/* 用户名搜索 */}
          <div className="space-y-2">
            <label className="text-sm font-medium">用户名</label>
            <div className="relative">
              <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="搜索用户名..."
                value={filters.username || ""}
                onChange={(e) => handleInputChange("username", e.target.value)}
                className="pl-8"
              />
            </div>
          </div>

          {/* 操作搜索 */}
          <div className="space-y-2">
            <label className="text-sm font-medium">操作</label>
            <Input
              placeholder="搜索操作..."
              value={filters.action || ""}
              onChange={(e) => handleInputChange("action", e.target.value)}
            />
          </div>

          {/* 资源搜索 */}
          <div className="space-y-2">
            <label className="text-sm font-medium">资源</label>
            <Input
              placeholder="搜索资源..."
              value={filters.resource || ""}
              onChange={(e) => handleInputChange("resource", e.target.value)}
            />
          </div>

          {/* HTTP 方法筛选 */}
          <div className="space-y-2">
            <label className="text-sm font-medium">HTTP 方法</label>
            <Select
              value={filters.method || "all"}
              onValueChange={(value) => handleSelectChange("method", value)}
            >
              <SelectTrigger>
                <SelectValue placeholder="选择方法" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">全部</SelectItem>
                <SelectItem value="GET">GET</SelectItem>
                <SelectItem value="POST">POST</SelectItem>
                <SelectItem value="PUT">PUT</SelectItem>
                <SelectItem value="DELETE">DELETE</SelectItem>
                <SelectItem value="PATCH">PATCH</SelectItem>
              </SelectContent>
            </Select>
          </div>

          {/* 状态码筛选 */}
          <div className="space-y-2">
            <label className="text-sm font-medium">状态码</label>
            <Select
              value={filters.status?.toString() || "all"}
              onValueChange={(value) => handleSelectChange("status", value)}
            >
              <SelectTrigger>
                <SelectValue placeholder="选择状态" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">全部</SelectItem>
                <SelectItem value="200">200 - 成功</SelectItem>
                <SelectItem value="201">201 - 已创建</SelectItem>
                <SelectItem value="400">400 - 错误请求</SelectItem>
                <SelectItem value="401">401 - 未授权</SelectItem>
                <SelectItem value="403">403 - 禁止访问</SelectItem>
                <SelectItem value="404">404 - 未找到</SelectItem>
                <SelectItem value="500">500 - 服务器错误</SelectItem>
              </SelectContent>
            </Select>
          </div>

          {/* 开始时间 */}
          <div className="space-y-2">
            <label className="text-sm font-medium">开始时间</label>
            <Input
              type="datetime-local"
              value={filters.start_time || ""}
              onChange={(e) => handleInputChange("start_time", e.target.value)}
            />
          </div>

          {/* 结束时间 */}
          <div className="space-y-2">
            <label className="text-sm font-medium">结束时间</label>
            <Input
              type="datetime-local"
              value={filters.end_time || ""}
              onChange={(e) => handleInputChange("end_time", e.target.value)}
            />
          </div>

          {/* 重置按钮 */}
          <div className="space-y-2 flex items-end">
            <Button
              variant="outline"
              onClick={onReset}
              disabled={!hasActiveFilters}
              className="w-full"
            >
              <X className="h-4 w-4 mr-2" />
              重置筛选
            </Button>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};
