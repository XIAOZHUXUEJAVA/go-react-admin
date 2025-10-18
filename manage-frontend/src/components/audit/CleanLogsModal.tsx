import React, { useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { AlertTriangle } from "lucide-react";

interface CleanLogsModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onConfirm: (days: number) => void;
  loading?: boolean;
}

export const CleanLogsModal: React.FC<CleanLogsModalProps> = ({
  open,
  onOpenChange,
  onConfirm,
  loading,
}) => {
  const [days, setDays] = useState<number>(90);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (days > 0) {
      onConfirm(days);
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>清理审计日志</DialogTitle>
          <DialogDescription>
            删除指定天数之前的审计日志记录
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit}>
          <div className="space-y-4 py-4">
            {/* 警告提示 */}
            <div className="flex items-start gap-3 p-3 bg-yellow-50 border border-yellow-200 rounded-lg">
              <AlertTriangle className="h-5 w-5 text-yellow-600 mt-0.5" />
              <div className="flex-1">
                <p className="text-sm font-medium text-yellow-800">
                  警告：此操作不可恢复
                </p>
                <p className="text-xs text-yellow-700 mt-1">
                  删除的日志将无法恢复，请谨慎操作
                </p>
              </div>
            </div>

            {/* 天数输入 */}
            <div className="space-y-2">
              <Label htmlFor="days">保留天数</Label>
              <Input
                id="days"
                type="number"
                min="1"
                value={days}
                onChange={(e) => setDays(parseInt(e.target.value) || 1)}
                placeholder="输入保留天数"
                required
              />
              <p className="text-xs text-muted-foreground">
                将删除 {days} 天之前的所有审计日志
              </p>
            </div>

            {/* 预设选项 */}
            <div className="space-y-2">
              <Label>快速选择</Label>
              <div className="grid grid-cols-3 gap-2">
                <Button
                  type="button"
                  variant="outline"
                  size="sm"
                  onClick={() => setDays(30)}
                >
                  30天
                </Button>
                <Button
                  type="button"
                  variant="outline"
                  size="sm"
                  onClick={() => setDays(90)}
                >
                  90天
                </Button>
                <Button
                  type="button"
                  variant="outline"
                  size="sm"
                  onClick={() => setDays(180)}
                >
                  180天
                </Button>
              </div>
            </div>
          </div>

          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={() => onOpenChange(false)}
              disabled={loading}
            >
              取消
            </Button>
            <Button type="submit" variant="destructive" disabled={loading}>
              {loading ? "清理中..." : "确认清理"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
};
