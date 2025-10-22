import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Activity, CheckCircle, XCircle, Clock } from "lucide-react";
import { AuditLog } from "@/types/audit";

interface AuditLogStatsCardsProps {
  logs: AuditLog[];
  pagination?: {
    total: number;
  } | null;
}

export const AuditLogStatsCards: React.FC<AuditLogStatsCardsProps> = ({
  logs,
  pagination,
}) => {
  // 计算统计数据
  const totalLogs = pagination?.total || logs.length;
  const successLogs = logs.filter(
    (log) => log.status >= 200 && log.status < 300
  ).length;
  const errorLogs = logs.filter((log) => log.status >= 400).length;
  const avgDuration =
    logs.length > 0
      ? Math.round(
          logs.reduce((sum, log) => sum + (log.duration || 0), 0) / logs.length
        )
      : 0;

  const stats = [
    {
      title: "总日志数",
      value: totalLogs.toLocaleString(),
      icon: Activity,
      description: "系统记录的总日志数",
      color: "text-blue-600",
      bgColor: "bg-blue-100",
    },
    {
      title: "成功请求",
      value: successLogs.toLocaleString(),
      icon: CheckCircle,
      description: "状态码 2xx 的请求",
      color: "text-green-600",
      bgColor: "bg-green-100",
    },
    {
      title: "失败请求",
      value: errorLogs.toLocaleString(),
      icon: XCircle,
      description: "状态码 4xx/5xx 的请求",
      color: "text-red-600",
      bgColor: "bg-red-100",
    },
    {
      title: "平均响应时间",
      value: `${avgDuration}ms`,
      icon: Clock,
      description: "请求的平均处理时间",
      color: "text-purple-600",
      bgColor: "bg-purple-100",
    },
  ];

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      {stats.map((stat, index) => (
        <Card key={index}>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">{stat.title}</CardTitle>
            <div className={`${stat.bgColor} p-2 rounded-lg`}>
              <stat.icon className={`h-4 w-4 ${stat.color}`} />
            </div>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{stat.value}</div>
            <p className="text-xs text-muted-foreground">{stat.description}</p>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
