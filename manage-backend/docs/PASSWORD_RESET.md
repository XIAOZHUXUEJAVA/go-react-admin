# 密码重置功能文档

## 功能概述

本系统实现了完整的密码重置功能，包括：
1. 用户请求密码重置（通过邮箱）
2. 系统生成安全Token并发送邮件
3. 用户点击邮件中的链接验证Token
4. 用户输入新密码完成重置

## 技术实现

### 1. 数据库表结构

已创建 `password_reset_tokens` 表，包含以下字段：
- `id`: 主键
- `user_id`: 用户ID（外键关联users表）
- `email`: 用户邮箱
- `token`: UUID格式的重置Token
- `expires_at`: Token过期时间（默认1小时）
- `used_at`: Token使用时间（NULL表示未使用）
- `ip_address`: 请求IP地址
- `user_agent`: 用户代理信息
- `created_at`: 创建时间

### 2. API接口

#### 2.1 请求密码重置
```
POST /api/auth/forgot-password
Content-Type: application/json

{
  "email": "user@example.com"
}
```

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "如果该邮箱存在，重置链接已发送，请查收邮件"
  }
}
```

#### 2.2 验证重置Token
```
POST /api/auth/verify-reset-token
Content-Type: application/json

{
  "token": "uuid-token-string"
}
```

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "valid": true,
    "email": "user@example.com"
  }
}
```

#### 2.3 重置密码
```
POST /api/auth/reset-password
Content-Type: application/json

{
  "token": "uuid-token-string",
  "new_password": "newPassword123"
}
```

**响应：**
```json
{
  "code": 200,
  "message": "密码重置成功，请使用新密码登录",
  "data": null
}
```

## 配置说明

### 1. 邮件配置 (config.yaml)

```yaml
email:
  smtp_host: "smtp.qq.com"           # SMTP服务器地址
  smtp_port: 465                      # SMTP端口 (465=SSL, 587=TLS)
  username: "your-email@qq.com"       # 邮箱账号
  password: "your-authorization-code" # 邮箱授权码（不是登录密码）
  from_name: "Go Manage System"       # 发件人名称
  from_address: "your-email@qq.com"   # 发件人邮箱
```

### 2. 密码重置配置

```yaml
password_reset:
  token_expire_minutes: 60            # Token有效期（分钟）
  frontend_url: "http://localhost:3000" # 前端地址
```

## 常见邮箱服务器配置

### QQ邮箱
```yaml
smtp_host: "smtp.qq.com"
smtp_port: 465  # SSL
username: "your-qq-email@qq.com"
password: "授权码"  # 在QQ邮箱设置中生成
```

### 163邮箱
```yaml
smtp_host: "smtp.163.com"
smtp_port: 465  # SSL
username: "your-email@163.com"
password: "授权码"
```

### Gmail
```yaml
smtp_host: "smtp.gmail.com"
smtp_port: 587  # TLS
username: "your-email@gmail.com"
password: "应用专用密码"
```

### 腾讯企业邮箱
```yaml
smtp_host: "smtp.exmail.qq.com"
smtp_port: 465  # SSL
username: "your-email@company.com"
password: "邮箱密码"
```

## 安全特性

1. **Token安全**
   - 使用UUID生成随机Token
   - Token有效期限制（默认1小时）
   - Token一次性使用
   - 记录请求IP和User-Agent

2. **防止邮箱枚举**
   - 无论邮箱是否存在，都返回相同的成功消息
   - 避免攻击者通过响应判断邮箱是否存在

3. **审计日志**
   - 记录所有密码重置请求
   - 记录成功的密码重置操作
   - 包含IP地址和时间戳

4. **用户状态检查**
   - 禁用账户无法重置密码
   - 确保只有活跃用户可以重置

## 前端集成示例

### 1. 忘记密码页面

```typescript
// ForgotPasswordForm.tsx
const handleSubmit = async (email: string) => {
  try {
    const response = await fetch('/api/auth/forgot-password', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email })
    });
    
    if (response.ok) {
      // 显示成功消息
      alert('重置链接已发送到您的邮箱');
    }
  } catch (error) {
    console.error('请求失败', error);
  }
};
```

### 2. 重置密码页面

```typescript
// ResetPasswordForm.tsx
const handleReset = async (token: string, newPassword: string) => {
  try {
    // 1. 先验证Token
    const verifyResponse = await fetch('/api/auth/verify-reset-token', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ token })
    });
    
    if (!verifyResponse.ok) {
      alert('重置链接无效或已过期');
      return;
    }
    
    // 2. 重置密码
    const resetResponse = await fetch('/api/auth/reset-password', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ token, new_password: newPassword })
    });
    
    if (resetResponse.ok) {
      alert('密码重置成功，请登录');
      // 跳转到登录页
      window.location.href = '/login';
    }
  } catch (error) {
    console.error('重置失败', error);
  }
};
```

## 邮件模板

系统使用精美的HTML邮件模板，包含：
- 响应式设计，支持移动端
- 渐变色按钮
- 安全提示信息
- 纯文本备用内容

## 测试建议

1. **功能测试**
   - 测试正常的密码重置流程
   - 测试Token过期场景
   - 测试Token重复使用
   - 测试不存在的邮箱

2. **安全测试**
   - 验证Token的随机性
   - 测试过期时间是否生效
   - 验证审计日志记录

3. **邮件测试**
   - 使用Mailtrap等测试工具
   - 验证邮件格式和内容
   - 测试不同邮箱客户端的显示效果

## 故障排查

### 邮件发送失败

1. **检查SMTP配置**
   - 确认SMTP服务器地址和端口
   - 验证邮箱账号和授权码
   - 检查SSL/TLS设置

2. **查看日志**
   ```bash
   # 查看应用日志
   tail -f logs/app.log | grep "send_reset_email"
   ```

3. **常见错误**
   - `535 Authentication failed`: 授权码错误
   - `Connection timeout`: 网络或防火墙问题
   - `550 Mailbox not found`: 收件人邮箱不存在

### Token验证失败

1. 检查Token是否过期
2. 确认Token是否已使用
3. 验证数据库中的Token记录

## 维护建议

1. **定期清理过期Token**
   ```sql
   DELETE FROM password_reset_tokens 
   WHERE expires_at < NOW() - INTERVAL '7 days';
   ```

2. **监控邮件发送**
   - 监控邮件发送成功率
   - 设置发送失败告警

3. **审计日志分析**
   - 定期检查异常的重置请求
   - 识别可能的攻击行为

## 依赖包

- `github.com/google/uuid`: UUID生成
- `gopkg.in/gomail.v2`: 邮件发送

确保已安装：
```bash
go get github.com/google/uuid
go get gopkg.in/gomail.v2
```
