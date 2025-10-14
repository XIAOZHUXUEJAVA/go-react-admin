-- ==========================================
-- 回滚 RBAC 权限体系
-- ==========================================

-- 删除表的顺序很重要，需要先删除有外键依赖的表

-- 1. 删除 Casbin 策略表
DROP TABLE IF EXISTS casbin_rule;

-- 2. 删除角色-权限关联表
DROP TABLE IF EXISTS role_permissions;

-- 3. 删除用户-角色关联表
DROP TABLE IF EXISTS user_roles;

-- 4. 删除菜单表
DROP TABLE IF EXISTS menus;

-- 4. 删除权限表
DROP TABLE IF EXISTS permissions;

-- 5. 删除角色表
DROP TABLE IF EXISTS roles;

-- 注意：不删除 users 表的 role 字段，因为我们保留它作为备份
