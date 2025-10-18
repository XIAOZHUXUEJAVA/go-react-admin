import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Eye } from "lucide-react";
import { AuditLog } from "@/types/audit";
import dayjs from "dayjs";

interface AuditLogTableProps {
  logs: AuditLog[];
  loading?: boolean;
  onView: (log: AuditLog) => void;
}

export const AuditLogTable: React.FC<AuditLogTableProps> = ({
  logs,
  loading,
  onView,
}) => {
  // 获取状态码对应的颜色
  const getStatusColor = (status: number) => {
    if (status >= 200 && status < 300) return "bg-green-100 text-green-800";
    if (status >= 300 && status < 400) return "bg-blue-100 text-blue-800";
    if (status >= 400 && status < 500) return "bg-yellow-100 text-yellow-800";
    if (status >= 500) return "bg-red-100 text-red-800";
    return "bg-gray-100 text-gray-800";
  };

  // 获取 HTTP 方法对应的颜色
  const getMethodColor = (method: string) => {
    switch (method) {
      case "GET":
        return "bg-blue-100 text-blue-800";
      case "POST":
        return "bg-green-100 text-green-800";
      case "PUT":
        return "bg-yellow-100 text-yellow-800";
      case "DELETE":
        return "bg-red-100 text-red-800";
      case "PATCH":
        return "bg-purple-100 text-purple-800";
      default:
        return "bg-gray-100 text-gray-800";
    }
  };

  // 格式化时间
  const formatDate = (dateString: string) => {
    try {
      return dayjs(dateString).format("YYYY-MM-DD HH:mm:ss");
    } catch {
      return dateString;
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-12">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  return (
    <div className="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>ID</TableHead>
            <TableHead>用户</TableHead>
            <TableHead>操作</TableHead>
            <TableHead>资源</TableHead>
            <TableHead>方法</TableHead>
            <TableHead>路径</TableHead>
            <TableHead>状态</TableHead>
            <TableHead>IP地址</TableHead>
            <TableHead>响应时间</TableHead>
            <TableHead>创建时间</TableHead>
            <TableHead className="text-right">操作</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {logs.length === 0 ? (
            <TableRow>
              <TableCell colSpan={11} className="text-center py-8">
                暂无日志数据
              </TableCell>
            </TableRow>
          ) : (
            logs.map((log) => (
              <TableRow key={log.id}>
                <TableCell className="font-medium">{log.id}</TableCell>
                <TableCell>
                  <div className="flex flex-col">
                    <span className="font-medium">{log.username}</span>
                    <span className="text-xs text-muted-foreground">
                      ID: {log.user_id}
                    </span>
                  </div>
                </TableCell>
                <TableCell>
                  <span className="text-sm">{log.action}</span>
                </TableCell>
                <TableCell>
                  <span className="text-sm">{log.resource || "-"}</span>
                </TableCell>
                <TableCell>
                  <Badge className={getMethodColor(log.method)} variant="outline">
                    {log.method}
                  </Badge>
                </TableCell>
                <TableCell>
                  <span className="text-xs text-muted-foreground max-w-[200px] truncate block">
                    {log.path}
                  </span>
                </TableCell>
                <TableCell>
                  <Badge className={getStatusColor(log.status)} variant="outline">
                    {log.status}
                  </Badge>
                </TableCell>
                <TableCell>
                  <span className="text-sm">{log.ip}</span>
                </TableCell>
                <TableCell>
                  <span className="text-sm">{log.duration}ms</span>
                </TableCell>
                <TableCell>
                  <span className="text-sm">{formatDate(log.created_at)}</span>
                </TableCell>
                <TableCell className="text-right">
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={() => onView(log)}
                  >
                    <Eye className="h-4 w-4" />
                  </Button>
                </TableCell>
              </TableRow>
            ))
          )}
        </TableBody>
      </Table>
    </div>
  );
};
