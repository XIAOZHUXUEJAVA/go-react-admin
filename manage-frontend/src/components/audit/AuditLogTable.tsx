import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { MethodBadge, StatusCodeBadge } from "@/components/ui";
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
                  <MethodBadge method={log.method as "GET" | "POST" | "PUT" | "DELETE" | "PATCH"} />
                </TableCell>
                <TableCell>
                  <span className="text-xs text-muted-foreground max-w-[200px] truncate block">
                    {log.path}
                  </span>
                </TableCell>
                <TableCell>
                  <StatusCodeBadge statusCode={log.status} />
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
