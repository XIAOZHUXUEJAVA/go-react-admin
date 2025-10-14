import React from "react";
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
import { Edit, Trash2, List } from "lucide-react";
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
            <TableRow>
              <TableCell colSpan={8} className="text-center py-8">
                暂无数据
              </TableCell>
            </TableRow>
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
                      <AlertDialog>
                        <AlertDialogTrigger asChild>
                          <PermissionButton
                            permission="dict_type:delete"
                            variant="ghost"
                            size="sm"
                            noPermissionTooltip="您没有删除字典类型的权限"
                          >
                            <Trash2 className="h-4 w-4 text-red-500" />
                          </PermissionButton>
                        </AlertDialogTrigger>
                        <AlertDialogContent>
                          <AlertDialogHeader>
                            <AlertDialogTitle>确认删除</AlertDialogTitle>
                            <AlertDialogDescription>
                              确定要删除字典类型 &quot;{dictType.name}&quot;
                              吗？此操作不可恢复。
                            </AlertDialogDescription>
                          </AlertDialogHeader>
                          <AlertDialogFooter>
                            <AlertDialogCancel>取消</AlertDialogCancel>
                            <AlertDialogAction
                              onClick={() => onDelete(dictType)}
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
