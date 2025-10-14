import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
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
import { Textarea } from "@/components/ui/textarea";
import { Checkbox } from "@/components/ui/checkbox";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { DictItem, CreateDictItemRequest } from "@/types/dict";

interface DictItemModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSubmit: (data: CreateDictItemRequest) => void;
  loading: boolean;
  dictItem?: DictItem | null;
  dictTypeCode: string;
}

interface FormData {
  dict_type_code: string;
  label: string;
  value: string;
  extra: string;
  description: string;
  status: string;
  sort_order: number;
  is_default: boolean;
}

export const DictItemModal: React.FC<DictItemModalProps> = ({
  open,
  onOpenChange,
  onSubmit,
  loading,
  dictItem,
  dictTypeCode,
}) => {
  const [isDefault, setIsDefault] = useState(false);

  const {
    register,
    handleSubmit,
    reset,
    setValue,
    watch,
    formState: { errors },
  } = useForm<FormData>({
    defaultValues: {
      dict_type_code: dictTypeCode,
      label: "",
      value: "",
      extra: "",
      description: "",
      status: "active",
      sort_order: 0,
      is_default: false,
    },
  });

  const status = watch("status");

  useEffect(() => {
    if (dictItem) {
      setValue("dict_type_code", dictItem.dict_type_code);
      setValue("label", dictItem.label);
      setValue("value", dictItem.value);
      setValue(
        "extra",
        dictItem.extra ? JSON.stringify(dictItem.extra, null, 2) : ""
      );
      setValue("description", dictItem.description || "");
      setValue("status", dictItem.status);
      setValue("sort_order", dictItem.sort_order);
      setValue("is_default", dictItem.is_default);
      setIsDefault(dictItem.is_default);
    } else {
      reset({
        dict_type_code: dictTypeCode,
        label: "",
        value: "",
        extra: "",
        description: "",
        status: "active",
        sort_order: 0,
        is_default: false,
      });
      setIsDefault(false);
    }
  }, [dictItem, dictTypeCode, setValue, reset]);

  const onFormSubmit = (data: FormData) => {
    let extraData: Record<string, unknown> | undefined;
    
    if (data.extra.trim()) {
      try {
        extraData = JSON.parse(data.extra);
      } catch {
        return;
      }
    }

    const submitData: CreateDictItemRequest = {
      dict_type_code: data.dict_type_code,
      label: data.label,
      value: data.value,
      extra: extraData,
      description: data.description,
      status: data.status,
      sort_order: data.sort_order,
      is_default: isDefault,
    };

    onSubmit(submitData);
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>
            {dictItem ? "编辑字典项" : "添加字典项"}
          </DialogTitle>
          <DialogDescription>
            {dictItem ? "修改字典项信息" : "填写以下信息创建新的字典项"}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit(onFormSubmit)} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="label">
              显示值 <span className="text-red-500">*</span>
            </Label>
            <Input
              id="label"
              placeholder="例如：启用"
              {...register("label", {
                required: "请输入显示值",
                minLength: { value: 1, message: "显示值至少1个字符" },
                maxLength: { value: 100, message: "显示值最多100个字符" },
              })}
            />
            {errors.label && (
              <p className="text-sm text-red-500">{errors.label.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="value">
              字典值 <span className="text-red-500">*</span>
            </Label>
            <Input
              id="value"
              placeholder="例如：active"
              disabled={!!dictItem}
              {...register("value", {
                required: "请输入字典值",
                minLength: { value: 1, message: "字典值至少1个字符" },
                maxLength: { value: 100, message: "字典值最多100个字符" },
              })}
            />
            {errors.value && (
              <p className="text-sm text-red-500">{errors.value.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="extra">扩展值（JSON格式）</Label>
            <Textarea
              id="extra"
              placeholder='例如：{"color": "green", "icon": "check"}'
              rows={4}
              className="font-mono text-sm"
              {...register("extra", {
                validate: (value) => {
                  if (!value.trim()) return true;
                  try {
                    JSON.parse(value);
                    return true;
                  } catch {
                    return "请输入有效的JSON格式";
                  }
                },
              })}
            />
            {errors.extra && (
              <p className="text-sm text-red-500">{errors.extra.message}</p>
            )}
            <p className="text-xs text-muted-foreground">
              可选，用于存储额外信息如颜色、图标等
            </p>
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">描述</Label>
            <Textarea
              id="description"
              placeholder="字典项的描述信息"
              rows={2}
              {...register("description", {
                maxLength: { value: 255, message: "描述最多255个字符" },
              })}
            />
            {errors.description && (
              <p className="text-sm text-red-500">
                {errors.description.message}
              </p>
            )}
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="status">状态</Label>
              <Select
                value={status}
                onValueChange={(value) => setValue("status", value)}
              >
                <SelectTrigger>
                  <SelectValue placeholder="选择状态" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="active">启用</SelectItem>
                  <SelectItem value="inactive">禁用</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="space-y-2">
              <Label htmlFor="sort_order">排序</Label>
              <Input
                id="sort_order"
                type="number"
                placeholder="0"
                {...register("sort_order", {
                  valueAsNumber: true,
                  min: { value: 0, message: "排序值不能小于0" },
                })}
              />
              {errors.sort_order && (
                <p className="text-sm text-red-500">
                  {errors.sort_order.message}
                </p>
              )}
            </div>
          </div>

          <div className="flex items-center space-x-2">
            <Checkbox
              id="is_default"
              checked={isDefault}
              onCheckedChange={(checked) => {
                setIsDefault(checked as boolean);
                setValue("is_default", checked as boolean);
              }}
            />
            <Label
              htmlFor="is_default"
              className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
            >
              设为默认值
            </Label>
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
            <Button type="submit" disabled={loading}>
              {loading ? "提交中..." : dictItem ? "更新" : "创建"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
};
