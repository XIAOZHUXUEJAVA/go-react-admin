import React, { useState } from "react";
import { DictType } from "@/types/dict";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Edit, Trash2, List, BookType } from "lucide-react";
import { PermissionButton } from "@/components/auth";
import { DeleteConfirmDialog, TableEmptyState } from "@/components/common";

interface DictTypeTableProps {
  dictTypes: DictType[];
  loading: boolean;
  onEdit: (dictType: DictType) => void;
  onDelete: (dictType: DictType) => void;
  onManageItems: (dictType: DictType) => void;
}

export const DictTypeTable: React.FC<DictTypeTableProps> = ({
  dictTypes,
  loading,
  onEdit,
  onDelete,
  onManageItems,
}) => {
  const [selectedType, setSelectedType] = useState<DictType | null>(null);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);

  const handleDeleteClick = (dictType: DictType) => {
    setSelectedType(dictType);
    setIsDeleteDialogOpen(true);
  };

  const handleDeleteConfirm = () => {
    if (selectedType) {
      onDelete(selectedType);
      setIsDeleteDialogOpen(false);
      setSelectedType(null);
    }
  };
  if (loading) {
    return (
      <div className="text-center py-8">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900 mx-auto"></div>
        <p className="mt-2 text-sm text-gray-500">加载中...</p>
      </div>
    );
  }

  return (
    <div className="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>代码</TableHead>
            <TableHead>名称</TableHead>
            <TableHead>描述</TableHead>
            <TableHead>状态</TableHead>
            <TableHead>排序</TableHead>
            <TableHead>系统内置</TableHead>
            <TableHead>创建时间</TableHead>
            <TableHead className="text-right">操作</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {dictTypes.length === 0 ? (
            <TableEmptyState
              colSpan={8}
              icon={BookType}
              title="暂无字典类型"
              description="还没有创建任何字典类型，点击上方按钮添加新类型"
            />
          ) : (
            dictTypes.map((dictType) => (
              <TableRow key={dictType.id}>
                <TableCell className="font-mono">{dictType.code}</TableCell>
                <TableCell className="font-medium">{dictType.name}</TableCell>
                <TableCell className="text-muted-foreground">
                  {dictType.description || "-"}
                </TableCell>
                <TableCell>
                  <Badge
                    variant={
                      dictType.status === "active" ? "default" : "secondary"
                    }
                  >
                    {dictType.status === "active" ? "启用" : "禁用"}
                  </Badge>
                </TableCell>
                <TableCell>{dictType.sort_order}</TableCell>
                <TableCell>
                  {dictType.is_system ? (
                    <Badge variant="outline">是</Badge>
                  ) : (
                    <span className="text-muted-foreground">否</span>
                  )}
                </TableCell>
                <TableCell className="text-sm text-muted-foreground">
                  {new Date(dictType.created_at).toLocaleString("zh-CN")}
                </TableCell>
                <TableCell className="text-right">
                  <div className="flex items-center justify-end gap-2">
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => onManageItems(dictType)}
                    >
                      <List className="h-4 w-4" />
                    </Button>
                    <PermissionButton
                      permission="dict_type:update"
                      variant="ghost"
                      size="sm"
                      onClick={() => onEdit(dictType)}
                      noPermissionTooltip="您没有编辑字典类型的权限"
                    >
                      <Edit className="h-4 w-4" />
                    </PermissionButton>
                    {!dictType.is_system && (
                      <PermissionButton
                        permission="dict_type:delete"
                        variant="ghost"
                        size="sm"
                        onClick={() => handleDeleteClick(dictType)}
                        noPermissionTooltip="您没有删除字典类型的权限"
                      >
                        <Trash2 className="h-4 w-4 text-red-500" />
                      </PermissionButton>
                    )}
                  </div>
                </TableCell>
              </TableRow>
            ))
          )}
        </TableBody>
      </Table>

      {/* 删除确认对话框 */}
      <DeleteConfirmDialog
        open={isDeleteDialogOpen}
        onOpenChange={setIsDeleteDialogOpen}
        onConfirm={handleDeleteConfirm}
        resourceName={selectedType?.name}
        resourceType="字典类型"
        description={`确定要删除字典类型 "${selectedType?.name}" 吗？此操作不可恢复。`}
      />
    </div>
  );
};
