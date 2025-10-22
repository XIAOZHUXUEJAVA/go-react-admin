"use client";

import React, { useState } from "react";
import { useDictTypes, useDictItems } from "@/hooks/useDict";
import {
  DictType,
  DictItem,
  CreateDictTypeRequest,
  CreateDictItemRequest,
} from "@/types/dict";
import { DashboardHeader } from "@/components/layout/dashboard-header";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Plus, RefreshCw, Book, ArrowLeft } from "lucide-react";
import {
  DictTypeTable,
  DictTypeModal,
  DictItemTable,
  DictItemModal,
} from "@/components/features/system/dict";
import { dictApi } from "@/api/dict";
import { toast } from "sonner";
import { getErrorMessage } from "@/lib/error-handler";
import { PagePermissionGuard, PermissionButton } from "@/components/auth";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

export default function DictManagePage() {
  const {
    dictTypes,
    pagination: typePagination,
    loading: typeLoading,
    error: typeError,
    fetchDictTypes,
    refetch: refetchTypes,
  } = useDictTypes({
    page: 1,
    page_size: 10,
  });

  const [selectedDictType, setSelectedDictType] = useState<DictType | null>(
    null
  );
  const [isTypeModalOpen, setIsTypeModalOpen] = useState(false);
  const [isItemModalOpen, setIsItemModalOpen] = useState(false);
  const [editingType, setEditingType] = useState<DictType | null>(null);
  const [editingItem, setEditingItem] = useState<DictItem | null>(null);
  const [isCreating, setIsCreating] = useState(false);
  const [isUpdating, setIsUpdating] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const [searchTerm, setSearchTerm] = useState("");
  const [statusFilter, setStatusFilter] = useState<string>("all");
  const [viewMode, setViewMode] = useState<"types" | "items">("types");

  const {
    dictItems,
    pagination: itemPagination,
    loading: itemLoading,
    fetchDictItems,
    refetch: refetchItems,
  } = useDictItems({
    page: 1,
    page_size: 10,
    dict_type_code: selectedDictType?.code,
  });

  const handlePageChange = (page: number) => {
    if (viewMode === "types") {
      fetchDictTypes({
        page,
        page_size: typePagination?.page_size || 10,
        status: statusFilter === "all" ? undefined : statusFilter,
        keyword: searchTerm,
      });
    } else {
      fetchDictItems({
        page,
        page_size: itemPagination?.page_size || 10,
        dict_type_code: selectedDictType?.code,
        status: statusFilter === "all" ? undefined : statusFilter,
      });
    }
  };

  const filteredDictTypes =
    dictTypes?.filter((type) => {
      const matchesSearch =
        type.code.toLowerCase().includes(searchTerm.toLowerCase()) ||
        type.name.toLowerCase().includes(searchTerm.toLowerCase());
      const matchesStatus =
        statusFilter === "all" || type.status === statusFilter;
      return matchesSearch && matchesStatus;
    }) || [];

  const handleCreateType = async (data: CreateDictTypeRequest) => {
    setIsCreating(true);
    try {
      const response = await dictApi.createDictType(data);
      if (response.code === 201) {
        toast.success("字典类型创建成功");
        setIsTypeModalOpen(false);
        refetchTypes();
      } else {
        toast.error(response.message || "创建字典类型失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "创建字典类型失败，请稍后重试"));
    } finally {
      setIsCreating(false);
    }
  };

  const handleUpdateType = async (data: CreateDictTypeRequest) => {
    if (!editingType) return;
    setIsUpdating(true);
    try {
      const response = await dictApi.updateDictType(editingType.id, data);
      if (response.code === 200) {
        toast.success("字典类型更新成功");
        setIsTypeModalOpen(false);
        setEditingType(null);
        refetchTypes();
      } else {
        toast.error(response.message || "更新字典类型失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "更新字典类型失败，请稍后重试"));
    } finally {
      setIsUpdating(false);
    }
  };

  const handleDeleteType = async (dictType: DictType) => {
    setIsDeleting(true);
    try {
      const response = await dictApi.deleteDictType(dictType.id);
      if (response.code === 200) {
        toast.success("字典类型删除成功");
        refetchTypes();
      } else {
        toast.error(response.message || "删除字典类型失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "删除字典类型失败，请稍后重试"));
    } finally {
      setIsDeleting(false);
    }
  };

  const handleManageItems = (dictType: DictType) => {
    setSelectedDictType(dictType);
    setViewMode("items");
    fetchDictItems({
      page: 1,
      page_size: 10,
      dict_type_code: dictType.code,
    });
  };

  const handleCreateItem = async (data: CreateDictItemRequest) => {
    setIsCreating(true);
    try {
      const response = await dictApi.createDictItem(data);
      if (response.code === 201) {
        toast.success("字典项创建成功");
        setIsItemModalOpen(false);
        refetchItems();
      } else {
        toast.error(response.message || "创建字典项失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "创建字典项失败，请稍后重试"));
    } finally {
      setIsCreating(false);
    }
  };

  const handleUpdateItem = async (data: CreateDictItemRequest) => {
    if (!editingItem) return;
    setIsUpdating(true);
    try {
      const response = await dictApi.updateDictItem(editingItem.id, data);
      if (response.code === 200) {
        toast.success("字典项更新成功");
        setIsItemModalOpen(false);
        setEditingItem(null);
        refetchItems();
      } else {
        toast.error(response.message || "更新字典项失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "更新字典项失败，请稍后重试"));
    } finally {
      setIsUpdating(false);
    }
  };

  const handleDeleteItem = async (dictItem: DictItem) => {
    setIsDeleting(true);
    try {
      const response = await dictApi.deleteDictItem(dictItem.id);
      if (response.code === 200) {
        toast.success("字典项删除成功");
        refetchItems();
      } else {
        toast.error(response.message || "删除字典项失败");
      }
    } catch (error) {
      toast.error(getErrorMessage(error, "删除字典项失败，请稍后重试"));
    } finally {
      setIsDeleting(false);
    }
  };

  const breadcrumbs = [
    { label: "Dashboard", href: "/dashboard" },
    { label: "系统管理" },
    { label: "字典管理" },
  ];

  const headerActions = (
    <div className="flex items-center gap-2">
      {viewMode === "items" && (
        <Button
          variant="outline"
          size="sm"
          onClick={() => {
            setViewMode("types");
            setSelectedDictType(null);
          }}
        >
          <ArrowLeft className="h-4 w-4" />
          返回
        </Button>
      )}
      <Button
        variant="outline"
        size="sm"
        onClick={viewMode === "types" ? refetchTypes : refetchItems}
        disabled={viewMode === "types" ? typeLoading : itemLoading}
      >
        <RefreshCw
          className={`h-4 w-4 ${
            (viewMode === "types" ? typeLoading : itemLoading)
              ? "animate-spin"
              : ""
          }`}
        />
        刷新
      </Button>
      {viewMode === "types" ? (
        <PermissionButton
          permission="dict_type:create"
          size="sm"
          onClick={() => {
            setEditingType(null);
            setIsTypeModalOpen(true);
          }}
          noPermissionTooltip="您没有创建字典类型的权限"
        >
          <Plus className="h-4 w-4" />
          添加字典类型
        </PermissionButton>
      ) : (
        <PermissionButton
          permission="dict_item:create"
          size="sm"
          onClick={() => {
            setEditingItem(null);
            setIsItemModalOpen(true);
          }}
          noPermissionTooltip="您没有创建字典项的权限"
        >
          <Plus className="h-4 w-4" />
          添加字典项
        </PermissionButton>
      )}
    </div>
  );

  return (
    <PagePermissionGuard permission="dict_type:read">
      <DashboardHeader breadcrumbs={breadcrumbs} actions={headerActions} />

      <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            {viewMode === "items" && (
              <Button
                variant="ghost"
                size="icon"
                onClick={() => {
                  setViewMode("types");
                  setSelectedDictType(null);
                }}
                className="h-8 w-8"
              >
                <ArrowLeft className="h-5 w-5" />
              </Button>
            )}
            <div>
              <h1 className="text-3xl font-bold tracking-tight">
                {viewMode === "types"
                  ? "字典管理"
                  : `字典项管理 - ${selectedDictType?.name}`}
              </h1>
              <p className="text-muted-foreground">
                {viewMode === "types"
                  ? "管理系统中的所有字典类型"
                  : `管理 ${selectedDictType?.code} 的字典项`}
              </p>
            </div>
          </div>
        </div>

        <div className="flex items-center gap-4">
          <div className="flex-1">
            <Input
              placeholder="搜索字典..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="max-w-sm"
            />
          </div>
          <Select value={statusFilter} onValueChange={setStatusFilter}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="状态筛选" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="all">全部状态</SelectItem>
              <SelectItem value="active">启用</SelectItem>
              <SelectItem value="inactive">禁用</SelectItem>
            </SelectContent>
          </Select>
        </div>

        <Card>
          <CardHeader>
            <CardTitle>
              {viewMode === "types" ? "字典类型列表" : "字典项列表"}
            </CardTitle>
            <CardDescription>
              {viewMode === "types"
                ? `显示 ${filteredDictTypes.length} 个字典类型`
                : `显示 ${dictItems?.length || 0} 个字典项`}
            </CardDescription>
          </CardHeader>
          <CardContent>
            {viewMode === "types" ? (
              <>
                {typeError ? (
                  <div className="text-center py-8">
                    <p className="text-red-500">
                      加载失败: {typeError.message}
                    </p>
                    <Button onClick={refetchTypes} className="mt-2">
                      重试
                    </Button>
                  </div>
                ) : filteredDictTypes.length === 0 && !typeLoading ? (
                  <div className="text-center py-8">
                    <Book className="mx-auto h-12 w-12 text-gray-400" />
                    <h3 className="mt-2 text-sm font-medium text-gray-900">
                      暂无字典类型
                    </h3>
                    <p className="mt-1 text-sm text-gray-500">
                      没有找到匹配的字典类型
                    </p>
                  </div>
                ) : (
                  <DictTypeTable
                    dictTypes={filteredDictTypes}
                    loading={typeLoading || isDeleting}
                    onEdit={(type) => {
                      setEditingType(type);
                      setIsTypeModalOpen(true);
                    }}
                    onDelete={handleDeleteType}
                    onManageItems={handleManageItems}
                  />
                )}
              </>
            ) : (
              <DictItemTable
                dictItems={dictItems || []}
                loading={itemLoading || isDeleting}
                onEdit={(item) => {
                  setEditingItem(item);
                  setIsItemModalOpen(true);
                }}
                onDelete={handleDeleteItem}
              />
            )}
          </CardContent>
        </Card>

        {((viewMode === "types" &&
          typePagination &&
          filteredDictTypes.length > 0) ||
          (viewMode === "items" &&
            itemPagination &&
            dictItems &&
            dictItems.length > 0)) && (
          <Card>
            <CardContent className="pt-6">
              <div className="flex items-center justify-between">
                <div className="text-sm text-muted-foreground">
                  {viewMode === "types"
                    ? `显示第 ${
                        (typePagination!.page - 1) * typePagination!.page_size +
                        1
                      } - ${Math.min(
                        typePagination!.page * typePagination!.page_size,
                        typePagination!.total
                      )} 条，共 ${typePagination!.total} 条记录`
                    : `显示第 ${
                        (itemPagination!.page - 1) * itemPagination!.page_size +
                        1
                      } - ${Math.min(
                        itemPagination!.page * itemPagination!.page_size,
                        itemPagination!.total
                      )} 条，共 ${itemPagination!.total} 条记录`}
                </div>
                <div className="flex items-center gap-2">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() =>
                      handlePageChange(
                        (viewMode === "types"
                          ? typePagination
                          : itemPagination)!.page - 1
                      )
                    }
                    disabled={
                      (viewMode === "types" ? typePagination : itemPagination)!
                        .page <= 1
                    }
                  >
                    上一页
                  </Button>
                  <div className="flex items-center gap-1">
                    {Array.from(
                      {
                        length: Math.min(
                          5,
                          (viewMode === "types"
                            ? typePagination
                            : itemPagination)!.total_pages
                        ),
                      },
                      (_, i) => {
                        const pageNum = i + 1;
                        return (
                          <Button
                            key={pageNum}
                            variant={
                              (viewMode === "types"
                                ? typePagination
                                : itemPagination)!.page === pageNum
                                ? "default"
                                : "outline"
                            }
                            size="sm"
                            onClick={() => handlePageChange(pageNum)}
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
                    onClick={() =>
                      handlePageChange(
                        (viewMode === "types"
                          ? typePagination
                          : itemPagination)!.page + 1
                      )
                    }
                    disabled={
                      (viewMode === "types" ? typePagination : itemPagination)!
                        .page >=
                      (viewMode === "types" ? typePagination : itemPagination)!
                        .total_pages
                    }
                  >
                    下一页
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>
        )}

        <DictTypeModal
          open={isTypeModalOpen}
          onOpenChange={setIsTypeModalOpen}
          onSubmit={editingType ? handleUpdateType : handleCreateType}
          loading={isCreating || isUpdating}
          dictType={editingType}
        />

        {selectedDictType && (
          <DictItemModal
            open={isItemModalOpen}
            onOpenChange={setIsItemModalOpen}
            onSubmit={editingItem ? handleUpdateItem : handleCreateItem}
            loading={isCreating || isUpdating}
            dictItem={editingItem}
            dictTypeCode={selectedDictType.code}
          />
        )}
      </div>
    </PagePermissionGuard>
  );
}
