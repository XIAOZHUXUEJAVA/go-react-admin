import React from "react";
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
import { Edit, Trash2, Star } from "lucide-react";
import { PermissionButton } from "@/components/auth";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";

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
  if (loading) {
    return (
      <div className="text-center py-8">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900 mx-auto"></div>
        <p className="mt-2 text-sm text-gray-500">加载中...</p>
      </div>
    );
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
            <TableRow>
              <TableCell colSpan={9} className="text-center py-8">
                暂无数据
              </TableCell>
            </TableRow>
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
                      <AlertDialog>
                        <AlertDialogTrigger asChild>
                          <PermissionButton
                            permission="dict_item:delete"
                            variant="ghost"
                            size="sm"
                            noPermissionTooltip="您没有删除字典项的权限"
                          >
                            <Trash2 className="h-4 w-4 text-red-500" />
                          </PermissionButton>
                        </AlertDialogTrigger>
                        <AlertDialogContent>
                          <AlertDialogHeader>
                            <AlertDialogTitle>确认删除</AlertDialogTitle>
                            <AlertDialogDescription>
                              确定要删除字典项 &quot;{dictItem.label}&quot;
                              吗？此操作不可恢复。
                            </AlertDialogDescription>
                          </AlertDialogHeader>
                          <AlertDialogFooter>
                            <AlertDialogCancel>取消</AlertDialogCancel>
                            <AlertDialogAction
                              onClick={() => onDelete(dictItem)}
                              className="bg-red-500 hover:bg-red-600"
                            >
                              删除
                            </AlertDialogAction>
                          </AlertDialogFooter>
                        </AlertDialogContent>
                      </AlertDialog>
                    )}
                  </div>
                </TableCell>
              </TableRow>
            ))
          )}
        </TableBody>
      </Table>
    </div>
  );
};
