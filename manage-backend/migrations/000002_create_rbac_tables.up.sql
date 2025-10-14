-- ==========================================
-- RBAC 权限体系表结构
-- ==========================================

-- 1. 角色表
CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    code VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(255),
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    is_system BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX IF NOT EXISTS idx_roles_code ON roles(code);
CREATE INDEX IF NOT EXISTS idx_roles_status ON roles(status);
CREATE INDEX IF NOT EXISTS idx_roles_deleted_at ON roles(deleted_at);

-- 2. 权限表
CREATE TABLE IF NOT EXISTS permissions (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(100) NOT NULL UNIQUE,
    resource VARCHAR(50) NOT NULL,
    action VARCHAR(50) NOT NULL,
    path VARCHAR(255),
    method VARCHAR(10),
    description VARCHAR(255),
    type VARCHAR(20) DEFAULT 'api',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX IF NOT EXISTS idx_permissions_code ON permissions(code);
CREATE INDEX IF NOT EXISTS idx_permissions_resource_action ON permissions(resource, action);
CREATE INDEX IF NOT EXISTS idx_permissions_type ON permissions(type);
CREATE INDEX IF NOT EXISTS idx_permissions_deleted_at ON permissions(deleted_at);

-- 3. 菜单表
CREATE TABLE IF NOT EXISTS menus (
    id BIGSERIAL PRIMARY KEY,
    parent_id BIGINT NULL,
    name VARCHAR(50) NOT NULL,
    title VARCHAR(100) NOT NULL,
    path VARCHAR(255),
    component VARCHAR(255),
    icon VARCHAR(50),
    order_num INT DEFAULT 0,
    type VARCHAR(20) NOT NULL,
    permission_code VARCHAR(100),
    visible BOOLEAN DEFAULT true,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    
    FOREIGN KEY (parent_id) REFERENCES menus(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_menus_parent_id ON menus(parent_id);
CREATE INDEX IF NOT EXISTS idx_menus_order_num ON menus(order_num);
CREATE INDEX IF NOT EXISTS idx_menus_permission_code ON menus(permission_code);
CREATE INDEX IF NOT EXISTS idx_menus_deleted_at ON menus(deleted_at);

-- 4. 用户-角色关联表
CREATE TABLE IF NOT EXISTS user_roles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    assigned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    assigned_by BIGINT,
    
    UNIQUE(user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_user_roles_user_id ON user_roles(user_id);
CREATE INDEX IF NOT EXISTS idx_user_roles_role_id ON user_roles(role_id);

-- 5. 角色-权限关联表
CREATE TABLE IF NOT EXISTS role_permissions (
    id BIGSERIAL PRIMARY KEY,
    role_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    assigned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(role_id, permission_id),
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX IF NOT EXISTS idx_role_permissions_permission_id ON role_permissions(permission_id);

-- 7. Casbin 策略表
CREATE TABLE IF NOT EXISTS casbin_rule (
    id BIGSERIAL PRIMARY KEY,
    ptype VARCHAR(100),
    v0 VARCHAR(100),
    v1 VARCHAR(100),
    v2 VARCHAR(100),
    v3 VARCHAR(100),
    v4 VARCHAR(100),
    v5 VARCHAR(100)
);

CREATE INDEX IF NOT EXISTS idx_casbin_rule ON casbin_rule(ptype, v0, v1, v2);

-- ==========================================
-- 初始化基础数据
-- ==========================================

-- 插入默认角色
INSERT INTO roles (code, name, description, is_system, status) VALUES
('admin', '超级管理员', '系统最高权限，拥有所有操作权限', true, 'active'),
('manager', '管理员', '管理权限，可以管理用户和基础配置', true, 'active'),
('user', '普通用户', '基础权限，只能访问基本功能', true, 'active')
ON CONFLICT (code) DO NOTHING;

-- 插入基础权限
INSERT INTO permissions (code, name, resource, action, path, method, description, type, status) VALUES
-- 用户管理权限
('user:read', '查看用户', 'user', 'read', '/api/users', 'GET', '查看用户列表和详情', 'api', 'active'),
('user:create', '创建用户', 'user', 'create', '/api/users', 'POST', '创建新用户', 'api', 'active'),
('user:update', '更新用户', 'user', 'update', '/api/users/*', 'PUT', '更新用户信息', 'api', 'active'),
('user:delete', '删除用户', 'user', 'delete', '/api/users/*', 'DELETE', '删除用户', 'api', 'active'),

-- 角色管理权限
('role:read', '查看角色', 'role', 'read', '/api/roles', 'GET', '查看角色列表和详情', 'api', 'active'),
('role:create', '创建角色', 'role', 'create', '/api/roles', 'POST', '创建新角色', 'api', 'active'),
('role:update', '更新角色', 'role', 'update', '/api/roles/*', 'PUT', '更新角色信息', 'api', 'active'),
('role:delete', '删除角色', 'role', 'delete', '/api/roles/*', 'DELETE', '删除角色', 'api', 'active'),

-- 权限管理权限
('permission:read', '查看权限', 'permission', 'read', '/api/permissions', 'GET', '查看权限列表', 'api', 'active'),
('permission:create', '创建权限', 'permission', 'create', '/api/permissions', 'POST', '创建新权限', 'api', 'active'),
('permission:update', '更新权限', 'permission', 'update', '/api/permissions/*', 'PUT', '更新权限信息', 'api', 'active'),
('permission:delete', '删除权限', 'permission', 'delete', '/api/permissions/*', 'DELETE', '删除权限', 'api', 'active'),

-- 菜单管理权限
('menu:read', '查看菜单', 'menu', 'read', '/api/menus', 'GET', '查看菜单列表', 'api', 'active'),
('menu:create', '创建菜单', 'menu', 'create', '/api/menus', 'POST', '创建新菜单', 'api', 'active'),
('menu:update', '更新菜单', 'menu', 'update', '/api/menus/*', 'PUT', '更新菜单信息', 'api', 'active'),
('menu:delete', '删除菜单', 'menu', 'delete', '/api/menus/*', 'DELETE', '删除菜单', 'api', 'active'),

-- 个人信息权限
('profile:read', '查看个人信息', 'profile', 'read', '/api/users/profile', 'GET', '查看自己的个人信息', 'api', 'active'),
('profile:update', '更新个人信息', 'profile', 'update', '/api/users/profile', 'PUT', '更新自己的个人信息', 'api', 'active')
ON CONFLICT (code) DO NOTHING;

-- 插入默认菜单
INSERT INTO menus (name, title, path, component, icon, order_num, type, permission_code, visible, status) VALUES
-- 一级菜单
('dashboard', '工作台', '/dashboard', '@/app/dashboard/page', 'LayoutDashboard', 1, 'menu', NULL, true, 'active'),
('system', '系统管理', '/system', NULL, 'Settings', 100, 'menu', NULL, true, 'active'),

-- 系统管理子菜单
('user-management', '用户管理', '/system/users', '@/app/system/users/page', 'Users', 101, 'menu', 'user:read', true, 'active'),
('role-management', '角色管理', '/system/roles', '@/app/system/roles/page', 'Shield', 102, 'menu', 'role:read', true, 'active'),
('permission-management', '权限管理', '/system/permissions', '@/app/system/permissions/page', 'Key', 103, 'menu', 'permission:read', true, 'active'),
('menu-management', '菜单管理', '/system/menus', '@/app/system/menus/page', 'Menu', 104, 'menu', 'menu:read', true, 'active');

-- 更新菜单父子关系
UPDATE menus SET parent_id = (SELECT id FROM menus WHERE name = 'system' LIMIT 1)
WHERE name IN ('user-management', 'role-management', 'permission-management', 'menu-management');

-- ==========================================
-- 数据迁移：为现有用户分配角色
-- ==========================================

-- 为现有用户根据其 role 字段分配对应角色
INSERT INTO user_roles (user_id, role_id, assigned_at)
SELECT u.id, r.id, CURRENT_TIMESTAMP
FROM users u
JOIN roles r ON r.code = u.role
WHERE u.deleted_at IS NULL
ON CONFLICT (user_id, role_id) DO NOTHING;

-- ==========================================
-- 初始化 Casbin 策略
-- ==========================================

-- 超级管理员拥有所有权限（通配符匹配）
INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES
('p', 'role:admin', '/api/*', '(GET|POST|PUT|DELETE|PATCH)'),
('p', 'role:admin', '/*', '(GET|POST|PUT|DELETE|PATCH)');

-- 管理员权限
INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES
-- 用户管理
('p', 'role:manager', '/api/users', 'GET'),
('p', 'role:manager', '/api/users/*', 'GET'),
('p', 'role:manager', '/api/users', 'POST'),
('p', 'role:manager', '/api/users/*', 'PUT');
-- -- 查看角色和权限
-- ('p', 'role:manager', '/api/roles', 'GET'),
-- ('p', 'role:manager', '/api/roles/*', 'GET'),
-- ('p', 'role:manager', '/api/permissions', 'GET'),
-- ('p', 'role:manager', '/api/menus', 'GET');

-- 普通用户权限
INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES
-- 个人信息
('p', 'role:user', '/api/users/profile', 'GET'),
('p', 'role:user', '/api/users/profile', 'PUT'),
-- 查看用户列表（只读）
('p', 'role:user', '/api/users', 'GET');

-- 用户-角色关系（g 策略）
-- 这些会在用户分配角色时动态添加，这里添加现有用户的关系
INSERT INTO casbin_rule (ptype, v0, v1)
SELECT 'g', CONCAT('user:', u.id), CONCAT('role:', r.code)
FROM users u
JOIN roles r ON r.code = u.role
WHERE u.deleted_at IS NULL;

-- ==========================================
-- 角色权限关联
-- ==========================================

-- 为 admin 角色分配所有权限
INSERT INTO role_permissions (role_id, permission_id)
SELECT 
    (SELECT id FROM roles WHERE code = 'admin'),
    id
FROM permissions
WHERE status = 'active'
ON CONFLICT (role_id, permission_id) DO NOTHING;

-- 为 manager 角色分配用户管理权限
INSERT INTO role_permissions (role_id, permission_id)
SELECT 
    (SELECT id FROM roles WHERE code = 'manager'),
    id
FROM permissions
WHERE code IN ('user:read', 'user:create', 'user:update')
  AND status = 'active'
ON CONFLICT (role_id, permission_id) DO NOTHING;

-- 为 user 角色分配基础权限
INSERT INTO role_permissions (role_id, permission_id)
SELECT 
    (SELECT id FROM roles WHERE code = 'user'),
    id
FROM permissions
WHERE code IN ('profile:read', 'profile:update')
  AND status = 'active'
ON CONFLICT (role_id, permission_id) DO NOTHING;

-- ==========================================
-- 完成
-- ==========================================
-- 注意：菜单显示现在完全基于权限控制
-- 如果菜单有 permission_code，用户需要有对应权限才能看到
-- 如果菜单没有 permission_code，所有登录用户都能看到
