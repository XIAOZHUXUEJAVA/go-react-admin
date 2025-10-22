import React from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Badge } from "@/components/ui/badge";
import { Separator } from "@/components/ui/separator";
import { AuditLog } from "@/types/audit";
import { formatDateWithSeconds } from "@/lib/date";

interface AuditLogDetailModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  log: AuditLog | null;
}

export const AuditLogDetailModal: React.FC<AuditLogDetailModalProps> = ({
  open,
  onOpenChange,
  log,
}) => {
  if (!log) return null;

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

  // 格式化 JSON
  const formatJSON = (jsonString: string) => {
    try {
      const parsed = JSON.parse(jsonString);
      return JSON.stringify(parsed, null, 2);
    } catch {
      return jsonString;
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-3xl max-h-[80vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>审计日志详情</DialogTitle>
          <DialogDescription>查看审计日志的详细信息</DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          {/* 基本信息 */}
          <div>
            <h3 className="text-sm font-semibold mb-2">基本信息</h3>
            <div className="grid grid-cols-2 gap-4 text-sm">
              <div>
                <span className="text-muted-foreground">日志ID:</span>
                <span className="ml-2 font-medium">{log.id}</span>
              </div>
              <div>
                <span className="text-muted-foreground">用户ID:</span>
                <span className="ml-2 font-medium">{log.user_id}</span>
              </div>
              <div>
                <span className="text-muted-foreground">用户名:</span>
                <span className="ml-2 font-medium">{log.username}</span>
              </div>
              <div>
                <span className="text-muted-foreground">创建时间:</span>
                <span className="ml-2 font-medium">
                  {formatDateWithSeconds(log.created_at)}
                </span>
              </div>
            </div>
          </div>

          <Separator />

          {/* 请求信息 */}
          <div>
            <h3 className="text-sm font-semibold mb-2">请求信息</h3>
            <div className="grid grid-cols-2 gap-4 text-sm">
              <div>
                <span className="text-muted-foreground">操作:</span>
                <span className="ml-2 font-medium">{log.action}</span>
              </div>
              <div>
                <span className="text-muted-foreground">资源:</span>
                <span className="ml-2 font-medium">{log.resource || "-"}</span>
              </div>
              <div>
                <span className="text-muted-foreground">资源ID:</span>
                <span className="ml-2 font-medium">{log.resource_id || "-"}</span>
              </div>
              <div>
                <span className="text-muted-foreground">HTTP方法:</span>
                <Badge className={`ml-2 ${getMethodColor(log.method)}`} variant="outline">
                  {log.method}
                </Badge>
              </div>
              <div className="col-span-2">
                <span className="text-muted-foreground">请求路径:</span>
                <span className="ml-2 font-medium break-all">{log.path}</span>
              </div>
            </div>
          </div>

          <Separator />

          {/* 响应信息 */}
          <div>
            <h3 className="text-sm font-semibold mb-2">响应信息</h3>
            <div className="grid grid-cols-2 gap-4 text-sm">
              <div>
                <span className="text-muted-foreground">状态码:</span>
                <Badge className={`ml-2 ${getStatusColor(log.status)}`} variant="outline">
                  {log.status}
                </Badge>
              </div>
              <div>
                <span className="text-muted-foreground">响应时间:</span>
                <span className="ml-2 font-medium">{log.duration}ms</span>
              </div>
              {log.error_msg && (
                <div className="col-span-2">
                  <span className="text-muted-foreground">错误信息:</span>
                  <div className="mt-1 p-2 bg-red-50 border border-red-200 rounded text-red-800">
                    {log.error_msg}
                  </div>
                </div>
              )}
            </div>
          </div>

          <Separator />

          {/* 客户端信息 */}
          <div>
            <h3 className="text-sm font-semibold mb-2">客户端信息</h3>
            <div className="grid grid-cols-1 gap-4 text-sm">
              <div>
                <span className="text-muted-foreground">IP地址:</span>
                <span className="ml-2 font-medium">{log.ip}</span>
              </div>
              <div>
                <span className="text-muted-foreground">User Agent:</span>
                <div className="mt-1 p-2 bg-gray-50 border border-gray-200 rounded break-all">
                  {log.user_agent}
                </div>
              </div>
            </div>
          </div>

          {/* 请求体 */}
          {log.request_body && (
            <>
              <Separator />
              <div>
                <h3 className="text-sm font-semibold mb-2">请求体</h3>
                <pre className="p-3 bg-gray-50 border border-gray-200 rounded text-xs overflow-x-auto">
                  {formatJSON(log.request_body)}
                </pre>
              </div>
            </>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
};
