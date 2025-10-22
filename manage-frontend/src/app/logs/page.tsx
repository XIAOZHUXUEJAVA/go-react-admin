"use client";

import React, { useState } from "react";
import { useAuditLogs } from "@/hooks/useAuditLogs";
import { AuditLog, AuditLogQuery } from "@/types/audit";
import { DashboardHeader } from "@/components/layout/dashboard-header";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { FileText, RefreshCw, Trash2 } from "lucide-react";
import {
  AuditLogStatsCards,
  AuditLogFilter,
  AuditLogTable,
  AuditLogDetailModal,
  CleanLogsModal,
} from "@/components/features/audit";
import { auditApi } from "@/api";
import { toast } from "sonner";
import { getErrorMessage } from "@/lib/errorHandler";
import { PagePermissionGuard, PermissionButton } from "@/components/auth";

/**
 * 日志管理页面
 */
export default function LogsManagePage() {
  const { logs, pagination, loading, error, fetchLogs, refetch } = useAuditLogs(
    {
      page: 1,
      page_size: 10,
    }
  );

  const [selectedLog, setSelectedLog] = useState<AuditLog | null>(null);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isCleanModalOpen, setIsCleanModalOpen] = useState(false);
  const [isCleaning, setIsCleaning] = useState(false);
  const [filters, setFilters] = useState<AuditLogQuery>({
    page: 1,
    page_size: 10,
  });

  // 处理查看日志详情
  const handleViewLog = (log: AuditLog) => {
    setSelectedLog(log);
    setIsDetailModalOpen(true);
  };

  // 处理分页
  const handlePageChange = (page: number) => {
    const newFilters = {
      ...filters,
      page,
    };
    setFilters(newFilters);
    fetchLogs(newFilters);
  };

  // 处理筛选变更
  const handleFilterChange = (newFilters: AuditLogQuery) => {
    const updatedFilters = {
      ...newFilters,
      page: 1, // 重置到第一页
      page_size: filters.page_size,
    };
    setFilters(updatedFilters);
    fetchLogs(updatedFilters);
  };

  // 重置筛选
  const handleResetFilters = () => {
    const defaultFilters: AuditLogQuery = {
      page: 1,
      page_size: 10,
    };
    setFilters(defaultFilters);
    fetchLogs(defaultFilters);
  };

  // 处理清理日志
  const handleCleanLogs = async (days: number) => {
    setIsCleaning(true);
    try {
      const response = await auditApi.cleanOldLogs(days);
      if (response.code === 200) {
        toast.success(`成功清理 ${days} 天前的审计日志`);
        setIsCleanModalOpen(false);
        refetch(); // 刷新日志列表
      } else {
        toast.error(response.message || "清理日志失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "清理日志失败，请稍后重试"));
    } finally {
      setIsCleaning(false);
    }
  };

  // 面包屑导航配置
  const breadcrumbs = [
    { label: "Dashboard", href: "/dashboard" },
    { label: "日志管理" },
  ];

  // 头部操作按钮
  const headerActions = (
    <div className="flex items-center gap-2">
      <Button variant="outline" size="sm" onClick={refetch} disabled={loading}>
        <RefreshCw className={`h-4 w-4 ${loading ? "animate-spin" : ""}`} />
        刷新
      </Button>
      <PermissionButton
        permission="logs:clean"
        variant="destructive"
        size="sm"
        onClick={() => setIsCleanModalOpen(true)}
        noPermissionTooltip="您没有清理日志的权限"
      >
        <Trash2 className="h-4 w-4" />
        清理日志
      </PermissionButton>
    </div>
  );

  return (
    <PagePermissionGuard permission="logs:read">
      <DashboardHeader breadcrumbs={breadcrumbs} actions={headerActions} />

      {/* 主要内容区域 */}
      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        {/* 页面标题 */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">日志管理</h1>
            <p className="text-muted-foreground">查看和管理系统审计日志记录</p>
          </div>
        </div>

        {/* 统计卡片 */}
        <AuditLogStatsCards logs={logs} pagination={pagination} />

        {/* 筛选器 */}
        <AuditLogFilter
          filters={filters}
          onFilterChange={handleFilterChange}
          onReset={handleResetFilters}
        />

        {/* 日志表格 */}
        <Card>
          <CardHeader>
            <CardTitle>审计日志列表</CardTitle>
            <CardDescription>
              {pagination
                ? `显示第 ${
                    (pagination.page - 1) * pagination.page_size + 1
                  } - ${Math.min(
                    pagination.page * pagination.page_size,
                    pagination.total
                  )} 条，共 ${pagination.total} 条记录`
                : `显示 ${logs.length} 条记录`}
            </CardDescription>
          </CardHeader>
          <CardContent>
            {error ? (
              <div className="text-center py-8">
                <p className="text-red-500">加载失败: {error.message}</p>
                <Button onClick={refetch} className="mt-2">
                  重试
                </Button>
              </div>
            ) : logs.length === 0 && !loading ? (
              <div className="text-center py-8">
                <FileText className="mx-auto h-12 w-12 text-gray-400" />
                <h3 className="mt-2 text-sm font-medium text-gray-900">
                  暂无日志数据
                </h3>
                <p className="mt-1 text-sm text-gray-500">
                  没有找到匹配的审计日志
                </p>
              </div>
            ) : (
              <AuditLogTable
                logs={logs}
                loading={loading}
                onView={handleViewLog}
              />
            )}
          </CardContent>
        </Card>

        {/* 分页 */}
        {pagination && logs.length > 0 && (
          <Card>
            <CardContent className="pt-6">
              <div className="flex items-center justify-between">
                <div className="text-sm text-muted-foreground">
                  显示第 {(pagination.page - 1) * pagination.page_size + 1} -{" "}
                  {Math.min(
                    pagination.page * pagination.page_size,
                    pagination.total
                  )}{" "}
                  条，共 {pagination.total} 条记录
                </div>
                <div className="flex items-center gap-2">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => handlePageChange(pagination.page - 1)}
                    disabled={pagination.page <= 1 || loading}
                  >
                    上一页
                  </Button>
                  <div className="flex items-center gap-1">
                    {Array.from(
                      { length: Math.min(5, pagination.total_pages) },
                      (_, i) => {
                        let pageNum: number;
                        if (pagination.total_pages <= 5) {
                          pageNum = i + 1;
                        } else if (pagination.page <= 3) {
                          pageNum = i + 1;
                        } else if (
                          pagination.page >=
                          pagination.total_pages - 2
                        ) {
                          pageNum = pagination.total_pages - 4 + i;
                        } else {
                          pageNum = pagination.page - 2 + i;
                        }

                        return (
                          <Button
                            key={pageNum}
                            variant={
                              pagination.page === pageNum
                                ? "default"
                                : "outline"
                            }
                            size="sm"
                            onClick={() => handlePageChange(pageNum)}
                            disabled={loading}
                          >
                            {pageNum}
                          </Button>
                        );
                      }
                    )}
                  </div>
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => handlePageChange(pagination.page + 1)}
                    disabled={
                      pagination.page >= pagination.total_pages || loading
                    }
                  >
                    下一页
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>
        )}

        {/* 日志详情模态框 */}
        <AuditLogDetailModal
          open={isDetailModalOpen}
          onOpenChange={setIsDetailModalOpen}
          log={selectedLog}
        />

        {/* 清理日志模态框 */}
        <CleanLogsModal
          open={isCleanModalOpen}
          onOpenChange={setIsCleanModalOpen}
          onConfirm={handleCleanLogs}
          loading={isCleaning}
        />
      </div>
    </PagePermissionGuard>
  );
}
