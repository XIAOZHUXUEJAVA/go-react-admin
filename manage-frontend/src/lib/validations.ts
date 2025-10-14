import { z } from "zod";

// 基础验证规则
export const baseValidations = {
  // 用户名验证
  username: z
    .string()
    .min(3, "用户名至少3个字符")
    .max(50, "用户名最多50个字符")
    .regex(
      /^[a-zA-Z0-9_\u4e00-\u9fa5]+$/,
      "用户名只能包含字母、数字、下划线和中文"
    ),

  // 邮箱验证
  email: z
    .string()
    .min(1, "邮箱地址不能为空")
    .email("请输入有效的邮箱地址")
    .max(255, "邮箱地址过长"),

  // 密码验证 - 增强版
  password: z
    .string()
    .min(8, "密码至少8个字符")
    .max(128, "密码最多128个字符")
    .regex(/^(?=.*[a-z])/, "密码必须包含小写字母")
    .regex(/^(?=.*[A-Z])/, "密码必须包含大写字母")
    .regex(/^(?=.*\d)/, "密码必须包含数字")
    .regex(/^(?=.*[!@#$%^&*(),.?":{}|<>])/, "密码必须包含特殊字符"),

  // 简单密码验证（用于当前系统）
  simplePassword: z
    .string()
    .min(6, "密码至少6个字符")
    .max(128, "密码最多128个字符"),

  // 角色验证
  role: z.enum(["admin", "user", "moderator"]),

  // 状态验证
  status: z.enum(["active", "inactive", "pending"]),

  // ID 验证
  id: z.number().positive("ID必须是正整数"),

  // 可选的字符串字段
  optionalString: z.string().optional(),

  // 非空字符串
  requiredString: z.string().min(1, "此字段不能为空"),
};

// 用户相关验证 schemas
export const userSchemas = {
  // 创建用户
  create: z.object({
    username: baseValidations.username,
    email: baseValidations.email,
    password: baseValidations.simplePassword, // 使用简单密码规则
    role: baseValidations.role.default("user"),
  }),

  // 更新用户
  update: z.object({
    username: baseValidations.username,
    email: baseValidations.email,
    role: baseValidations.role,
    status: baseValidations.status,
  }),

  // 用户登录
  login: z.object({
    username: z.string().min(1, "用户名不能为空"),
    password: z.string().min(1, "密码不能为空"),
  }),

  // 用户注册
  register: z
    .object({
      username: baseValidations.username,
      email: baseValidations.email,
      password: baseValidations.simplePassword,
      confirmPassword: z.string(),
    })
    .refine((data) => data.password === data.confirmPassword, {
      message: "两次输入的密码不一致",
      path: ["confirmPassword"],
    }),

  // 修改密码
  changePassword: z
    .object({
      currentPassword: z.string().min(1, "请输入当前密码"),
      newPassword: baseValidations.simplePassword,
      confirmPassword: z.string(),
    })
    .refine((data) => data.newPassword === data.confirmPassword, {
      message: "两次输入的新密码不一致",
      path: ["confirmPassword"],
    }),

  // 批量删除用户
  bulkDelete: z.object({
    ids: z.array(baseValidations.id).min(1, "请至少选择一个用户"),
    confirmText: z.string().refine((val) => val === "DELETE", {
      message: "请输入 DELETE 确认删除",
    }),
  }),

  // 用户搜索
  search: z.object({
    keyword: z.string().max(100, "搜索关键词过长"),
    role: baseValidations.role.optional(),
    status: baseValidations.status.optional(),
  }),
};

// 分页验证
export const paginationSchema = z.object({
  page: z.number().min(1, "页码必须大于0").default(1),
  pageSize: z
    .number()
    .min(1, "每页数量必须大于0")
    .max(100, "每页最多100条")
    .default(10),
});

// 通用验证 schemas
export const commonSchemas = {
  // 分页参数
  pagination: paginationSchema,

  // ID 参数
  idParam: z.object({
    id: baseValidations.id,
  }),

  // 确认操作
  confirmation: z.object({
    confirmed: z.boolean().refine((val) => val === true, {
      message: "请确认此操作",
    }),
  }),

  // 文件上传
  fileUpload: z
    .object({
      file: z.instanceof(File, { message: "请选择文件" }),
      maxSize: z.number().default(5 * 1024 * 1024), // 5MB
    })
    .refine((data) => data.file.size <= data.maxSize, {
      message: "文件大小不能超过5MB",
      path: ["file"],
    }),
};

// 表单验证辅助函数
export const validationHelpers = {
  // 检查用户名格式
  isValidUsername: (username: string): boolean => {
    return baseValidations.username.safeParse(username).success;
  },

  // 检查邮箱格式
  isValidEmail: (email: string): boolean => {
    return baseValidations.email.safeParse(email).success;
  },

  // 检查密码强度
  getPasswordStrength: (password: string): "weak" | "medium" | "strong" => {
    if (password.length < 6) return "weak";

    let score = 0;
    if (/[a-z]/.test(password)) score++;
    if (/[A-Z]/.test(password)) score++;
    if (/\d/.test(password)) score++;
    if (/[!@#$%^&*(),.?":{}|<>]/.test(password)) score++;

    if (score <= 2) return "weak";
    if (score === 3) return "medium";
    return "strong";
  },

  // 格式化验证错误
  formatValidationError: (error: z.ZodError): Record<string, string> => {
    const formattedErrors: Record<string, string> = {};

    error.issues.forEach((err: any) => {
      const path = err.path.join(".");
      formattedErrors[path] = err.message;
    });

    return formattedErrors;
  },

  // 验证并返回错误信息
  validateField: <T>(schema: z.ZodSchema<T>, value: unknown): string | null => {
    const result = schema.safeParse(value);
    if (!result.success) {
      return result.error.issues[0]?.message || "验证失败";
    }
    return null;
  },
};

// 自定义验证规则
export const customValidations = {
  // 中国手机号验证
  chinesePhone: z.string().regex(/^1[3-9]\d{9}$/, "请输入有效的手机号码"),

  // 身份证号验证
  chineseIdCard: z
    .string()
    .regex(
      /^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$/,
      "请输入有效的身份证号"
    ),

  // URL 验证
  url: z.string().url("请输入有效的URL地址"),

  // 颜色值验证
  hexColor: z
    .string()
    .regex(/^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$/, "请输入有效的颜色值"),

  // 日期范围验证
  dateRange: z
    .object({
      startDate: z.date(),
      endDate: z.date(),
    })
    .refine((data) => data.startDate <= data.endDate, {
      message: "开始日期不能晚于结束日期",
      path: ["endDate"],
    }),
};

// 导出类型
export type CreateUserFormData = z.infer<typeof userSchemas.create>;
export type UpdateUserFormData = z.infer<typeof userSchemas.update>;
export type LoginFormData = z.infer<typeof userSchemas.login>;
export type RegisterFormData = z.infer<typeof userSchemas.register>;
export type ChangePasswordFormData = z.infer<typeof userSchemas.changePassword>;
export type BulkDeleteFormData = z.infer<typeof userSchemas.bulkDelete>;
export type SearchFormData = z.infer<typeof userSchemas.search>;
export type PaginationFormData = z.infer<typeof paginationSchema>;
