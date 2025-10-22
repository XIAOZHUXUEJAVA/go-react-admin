import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { DictType, CreateDictTypeRequest } from "@/types/dict";
import { FormDialog } from "@/components/common";

interface DictTypeModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSubmit: (data: CreateDictTypeRequest) => void;
  loading: boolean;
  dictType?: DictType | null;
}

interface FormData {
  code: string;
  name: string;
  description: string;
  status: string;
  sort_order: number;
}

export const DictTypeModal: React.FC<DictTypeModalProps> = ({
  open,
  onOpenChange,
  onSubmit,
  loading,
  dictType,
}) => {
  const form = useForm<FormData>({
    defaultValues: {
      code: "",
      name: "",
      description: "",
      status: "active",
      sort_order: 0,
    },
  });

  const {
    register,
    reset,
    setValue,
    watch,
    formState: { errors },
  } = form;

  const status = watch("status");

  useEffect(() => {
    if (dictType) {
      setValue("code", dictType.code);
      setValue("name", dictType.name);
      setValue("description", dictType.description || "");
      setValue("status", dictType.status);
      setValue("sort_order", dictType.sort_order);
    } else {
      reset({
        code: "",
        name: "",
        description: "",
        status: "active",
        sort_order: 0,
      });
    }
  }, [dictType, setValue, reset]);

  const onFormSubmit = (data: FormData) => {
    onSubmit(data);
  };

  return (
    <FormDialog
      open={open}
      onOpenChange={onOpenChange}
      title={dictType ? "编辑字典类型" : "添加字典类型"}
      description={
        dictType
          ? "修改字典类型信息"
          : "填写以下信息创建新的字典类型"
      }
      form={form}
      onSubmit={onFormSubmit}
      loading={loading}
      submitText={dictType ? "更新" : "创建"}
      maxWidth="sm:max-w-[500px]"
      resetOnClose={false}
    >
      <div className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="code">
              代码 <span className="text-red-500">*</span>
            </Label>
            <Input
              id="code"
              placeholder="例如：user_status"
              disabled={!!dictType}
              {...register("code", {
                required: "请输入字典类型代码",
                minLength: { value: 2, message: "代码至少2个字符" },
                maxLength: { value: 50, message: "代码最多50个字符" },
                pattern: {
                  value: /^[a-z][a-z0-9_]*$/,
                  message: "代码只能包含小写字母、数字和下划线，且以字母开头",
                },
              })}
            />
            {errors.code && (
              <p className="text-sm text-red-500">{errors.code.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="name">
              名称 <span className="text-red-500">*</span>
            </Label>
            <Input
              id="name"
              placeholder="例如：用户状态"
              {...register("name", {
                required: "请输入字典类型名称",
                minLength: { value: 2, message: "名称至少2个字符" },
                maxLength: { value: 100, message: "名称最多100个字符" },
              })}
            />
            {errors.name && (
              <p className="text-sm text-red-500">{errors.name.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">描述</Label>
            <Textarea
              id="description"
              placeholder="字典类型的描述信息"
              rows={3}
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

      </div>
    </FormDialog>
  );
};
