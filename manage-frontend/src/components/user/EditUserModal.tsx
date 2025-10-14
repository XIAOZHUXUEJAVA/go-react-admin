"use client";

import React, { useState, useEffect } from "react";
import { User } from "@/types/api";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useAllRoles } from "@/hooks/useRoles";
import { Loader2 } from "lucide-react";

interface EditUserModalProps {
  user: User | null;
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSave: (user: User) => void;
}

export function EditUserModal({
  user,
  open,
  onOpenChange,
  onSave,
}: EditUserModalProps) {
  const [editingUser, setEditingUser] = useState<User | null>(null);
  const { roles, loading: rolesLoading } = useAllRoles();

  // 当用户或模态框状态改变时，重置编辑状态
  useEffect(() => {
    if (user && open) {
      setEditingUser({ ...user });
    }
  }, [user, open]);

  const handleSave = () => {
    if (editingUser) {
      onSave(editingUser);
      onOpenChange(false);
    }
  };

  const handleCancel = () => {
    setEditingUser(null);
    onOpenChange(false);
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-md">
        <DialogHeader>
          <DialogTitle>编辑用户</DialogTitle>
          <DialogDescription>修改用户的基本信息</DialogDescription>
        </DialogHeader>
        {editingUser && (
          <div className="grid gap-4">
            <div className="space-y-2">
              <Label htmlFor="edit-username">用户名</Label>
              <Input
                id="edit-username"
                value={editingUser.username}
                onChange={(e) =>
                  setEditingUser({
                    ...editingUser,
                    username: e.target.value,
                  })
                }
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="edit-email">邮箱</Label>
              <Input
                id="edit-email"
                type="email"
                value={editingUser.email}
                onChange={(e) =>
                  setEditingUser({
                    ...editingUser,
                    email: e.target.value,
                  })
                }
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="edit-role">角色</Label>
              <Select
                value={editingUser.role}
                onValueChange={(value) =>
                  setEditingUser({
                    ...editingUser,
                    role: value,
                  })
                }
                disabled={rolesLoading}
              >
                <SelectTrigger>
                  <SelectValue placeholder={rolesLoading ? "加载中..." : "选择角色"} />
                </SelectTrigger>
                <SelectContent>
                  {rolesLoading ? (
                    <div className="flex items-center justify-center p-2">
                      <Loader2 className="h-4 w-4 animate-spin" />
                    </div>
                  ) : roles.length > 0 ? (
                    roles
                      .filter((role) => role.status === "active")
                      .map((role) => (
                        <SelectItem key={role.id} value={role.code}>
                          {role.name}
                        </SelectItem>
                      ))
                  ) : (
                    <div className="p-2 text-sm text-muted-foreground">
                      暂无可用角色
                    </div>
                  )}
                </SelectContent>
              </Select>
            </div>
            <div className="space-y-2">
              <Label htmlFor="edit-status">状态</Label>
              <Select
                value={editingUser.status}
                onValueChange={(value) =>
                  setEditingUser({
                    ...editingUser,
                    status: value,
                  })
                }
              >
                <SelectTrigger>
                  <SelectValue placeholder="选择状态" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="active">活跃</SelectItem>
                  <SelectItem value="inactive">非活跃</SelectItem>
                  <SelectItem value="pending">待审核</SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>
        )}
        <DialogFooter>
          <Button variant="outline" onClick={handleCancel}>
            取消
          </Button>
          <Button onClick={handleSave}>保存更改</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
