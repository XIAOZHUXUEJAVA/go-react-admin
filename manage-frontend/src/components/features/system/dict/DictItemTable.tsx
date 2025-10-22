import React, { useState } from "react";
import { DictItem } from "@/types/dict";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { Edit, Trash2, Star, BookOpen } from "lucide-react";
import { PermissionButton } from "@/components/auth";
import { DeleteConfirmDialog, TableEmptyState, LoadingState } from "@/components/common";

interface DictItemTableProps {
  dictItems: DictItem[];
  loading: boolean;
  onEdit: (dictItem: DictItem) => void;
  onDelete: (dictItem: DictItem) => void;
}

export const DictItemTable: React.FC<DictItemTableProps> = ({
  dictItems,
  loading,
  onEdit,
  onDelete,
}) => {
  const [selectedItem, setSelectedItem] = useState<DictItem | null>(null);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);

  const handleDeleteClick = (item: DictItem) => {
    setSelectedItem(item);
    setIsDeleteDialogOpen(true);
  };

  const handleDeleteConfirm = () => {
    if (selectedItem) {
      onDelete(selectedItem);
      setIsDeleteDialogOpen(false);
      setSelectedItem(null);
    }
  };
  if (loading) {
    return <LoadingState mode="text" text="加载字典项..." />;
  }

  const renderExtra = (extra?: Record<string, unknown>) => {
    if (!extra || Object.keys(extra).length === 0) return "-";
    
    return (
      <div className="space-y-1">
        {Object.entries(extra).map(([key, value]) => (
          <div key={key} className="text-xs">
            <span className="font-medium">{key}:</span>{" "}
            <span className="text-muted-foreground">
              {typeof value === "object" ? JSON.stringify(value) : String(value)}
            </span>
          </div>
        ))}
      </div>
    );
  };

  return (
    <div className="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>显示值</TableHead>
            <TableHead>字典值</TableHead>
            <TableHead>扩展值</TableHead>
            <TableHead>描述</TableHead>
            <TableHead>状态</TableHead>
            <TableHead>排序</TableHead>
            <TableHead>默认值</TableHead>
            <TableHead>系统内置</TableHead>
            <TableHead className="text-right">操作</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {dictItems.length === 0 ? (
            <TableEmptyState
              colSpan={9}
              icon={BookOpen}
              title="暂无字典项数据"
              description="请先选择一个字典类型，然后添加字典项"
            />
          ) : (
            dictItems.map((dictItem) => (
              <TableRow key={dictItem.id}>
                <TableCell className="font-medium">{dictItem.label}</TableCell>
                <TableCell className="font-mono">{dictItem.value}</TableCell>
                <TableCell className="max-w-xs">
                  {renderExtra(dictItem.extra)}
                </TableCell>
                <TableCell className="text-muted-foreground">
                  {dictItem.description || "-"}
                </TableCell>
                <TableCell>
                  <Badge
                    variant={
                      dictItem.status === "active" ? "default" : "secondary"
                    }
                  >
                    {dictItem.status === "active" ? "启用" : "禁用"}
                  </Badge>
                </TableCell>
                <TableCell>{dictItem.sort_order}</TableCell>
                <TableCell>
                  {dictItem.is_default && (
                    <Star className="h-4 w-4 text-yellow-500 fill-yellow-500" />
                  )}
                </TableCell>
                <TableCell>
                  {dictItem.is_system ? (
                    <Badge variant="outline">是</Badge>
                  ) : (
                    <span className="text-muted-foreground">否</span>
                  )}
                </TableCell>
                <TableCell className="text-right">
                  <div className="flex items-center justify-end gap-2">
                    <PermissionButton
                      permission="dict_item:update"
                      variant="ghost"
                      size="sm"
                      onClick={() => onEdit(dictItem)}
                      noPermissionTooltip="您没有编辑字典项的权限"
                    >
                      <Edit className="h-4 w-4" />
                    </PermissionButton>
                    {!dictItem.is_system && (
                      <PermissionButton
                        permission="dict_item:delete"
                        variant="ghost"
                        size="sm"
                        onClick={() => handleDeleteClick(dictItem)}
                        noPermissionTooltip="您没有删除字典项的权限"
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
        resourceName={selectedItem?.label}
        resourceType="字典项"
        description={`确定要删除字典项 "${selectedItem?.label}" 吗？此操作不可恢复。`}
      />
    </div>
  );
};
