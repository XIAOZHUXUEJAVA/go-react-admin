-- ==========================================
-- 字典管理系统表结构
-- ==========================================

-- 1. 字典类型表
CREATE TABLE IF NOT EXISTS dict_types (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    sort_order INT DEFAULT 0,
    is_system BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX IF NOT EXISTS idx_dict_types_code ON dict_types(code);
CREATE INDEX IF NOT EXISTS idx_dict_types_status ON dict_types(status);
CREATE INDEX IF NOT EXISTS idx_dict_types_sort_order ON dict_types(sort_order);
CREATE INDEX IF NOT EXISTS idx_dict_types_deleted_at ON dict_types(deleted_at);

COMMENT ON TABLE dict_types IS '字典类型表';
COMMENT ON COLUMN dict_types.code IS '字典类型代码';
COMMENT ON COLUMN dict_types.name IS '字典类型名称';
COMMENT ON COLUMN dict_types.description IS '字典类型描述';
COMMENT ON COLUMN dict_types.status IS '状态：active-启用, inactive-禁用';
COMMENT ON COLUMN dict_types.sort_order IS '排序标记';
COMMENT ON COLUMN dict_types.is_system IS '是否系统内置';

-- 2. 字典项表
CREATE TABLE IF NOT EXISTS dict_items (
    id BIGSERIAL PRIMARY KEY,
    dict_type_code VARCHAR(50) NOT NULL,
    label VARCHAR(100) NOT NULL,
    value VARCHAR(100) NOT NULL,
    extra JSONB,
    description VARCHAR(255),
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    sort_order INT DEFAULT 0,
    is_default BOOLEAN DEFAULT false,
    is_system BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    
    UNIQUE(dict_type_code, value),
    FOREIGN KEY (dict_type_code) REFERENCES dict_types(code) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_dict_items_dict_type_code ON dict_items(dict_type_code);
CREATE INDEX IF NOT EXISTS idx_dict_items_value ON dict_items(value);
CREATE INDEX IF NOT EXISTS idx_dict_items_status ON dict_items(status);
CREATE INDEX IF NOT EXISTS idx_dict_items_sort_order ON dict_items(sort_order);
CREATE INDEX IF NOT EXISTS idx_dict_items_is_default ON dict_items(is_default);
CREATE INDEX IF NOT EXISTS idx_dict_items_deleted_at ON dict_items(deleted_at);
CREATE INDEX IF NOT EXISTS idx_dict_items_extra ON dict_items USING GIN (extra);

COMMENT ON TABLE dict_items IS '字典项表';
COMMENT ON COLUMN dict_items.dict_type_code IS '关联字典类型代码';
COMMENT ON COLUMN dict_items.label IS '显示值';
COMMENT ON COLUMN dict_items.value IS '字典值';
COMMENT ON COLUMN dict_items.extra IS '扩展值（JSON）';
COMMENT ON COLUMN dict_items.status IS '状态：active-启用, inactive-禁用';
COMMENT ON COLUMN dict_items.sort_order IS '排序标记';
COMMENT ON COLUMN dict_items.is_default IS '是否默认值';
COMMENT ON COLUMN dict_items.is_system IS '是否系统内置';

-- ==========================================
-- 初始化字典类型数据
-- ==========================================

-- 用户状态字典
INSERT INTO dict_types (code, name, description, status, sort_order, is_system) VALUES
('user_status', '用户状态', '用户账户状态枚举', 'active', 1, true);

INSERT INTO dict_items (dict_type_code, label, value, extra, description, status, sort_order, is_default, is_system) VALUES
('user_status', '启用', 'active', '{"color": "green", "badge": "success"}', '用户账户正常可用', 'active', 1, true, true),
('user_status', '禁用', 'inactive', '{"color": "red", "badge": "error"}', '用户账户被禁用', 'active', 2, false, true),
('user_status', '待审核', 'pending', '{"color": "orange", "badge": "warning"}', '用户账户待审核', 'active', 3, false, true),
('user_status', '已锁定', 'locked', '{"color": "gray", "badge": "default"}', '用户账户被锁定', 'active', 4, false, true);

-- 性别字典
INSERT INTO dict_types (code, name, description, status, sort_order, is_system) VALUES
('gender', '性别', '用户性别枚举', 'active', 2, true);

INSERT INTO dict_items (dict_type_code, label, value, extra, description, status, sort_order, is_default, is_system) VALUES
('gender', '男', 'male', '{"icon": "male"}', '男性', 'active', 1, false, true),
('gender', '女', 'female', '{"icon": "female"}', '女性', 'active', 2, false, true),
('gender', '未知', 'unknown', '{"icon": "question"}', '未知或不愿透露', 'active', 3, true, true);

-- 通用状态字典
INSERT INTO dict_types (code, name, description, status, sort_order, is_system) VALUES
('common_status', '通用状态', '通用的启用/禁用状态', 'active', 3, true);

INSERT INTO dict_items (dict_type_code, label, value, extra, description, status, sort_order, is_default, is_system) VALUES
('common_status', '启用', 'active', '{"color": "green"}', '启用状态', 'active', 1, true, true),
('common_status', '禁用', 'inactive', '{"color": "red"}', '禁用状态', 'active', 2, false, true);

-- 是否标识字典
INSERT INTO dict_types (code, name, description, status, sort_order, is_system) VALUES
('yes_no', '是否标识', '通用的是/否标识', 'active', 4, true);

INSERT INTO dict_items (dict_type_code, label, value, extra, description, status, sort_order, is_default, is_system) VALUES
('yes_no', '是', 'yes', '{"color": "green"}', '是', 'active', 1, false, true),
('yes_no', '否', 'no', '{"color": "gray"}', '否', 'active', 2, true, true);

-- ==========================================
-- 为字典管理添加权限
-- ==========================================

-- 字典类型权限
INSERT INTO permissions (name, code, resource, action, path, method, description, type, status) VALUES
('查看字典类型', 'dict_type:read', 'dict_type', 'read', '/dict-types', 'GET', '查看字典类型列表', 'api', 'active'),
('创建字典类型', 'dict_type:create', 'dict_type', 'create', '/dict-types', 'POST', '创建新字典类型', 'api', 'active'),
('更新字典类型', 'dict_type:update', 'dict_type', 'update', '/dict-types/:id', 'PUT', '更新字典类型信息', 'api', 'active'),
('删除字典类型', 'dict_type:delete', 'dict_type', 'delete', '/dict-types/:id', 'DELETE', '删除字典类型', 'api', 'active');

-- 字典项权限
INSERT INTO permissions (name, code, resource, action, path, method, description, type, status) VALUES
('查看字典项', 'dict_item:read', 'dict_item', 'read', '/dict-items', 'GET', '查看字典项列表', 'api', 'active'),
('创建字典项', 'dict_item:create', 'dict_item', 'create', '/dict-items', 'POST', '创建新字典项', 'api', 'active'),
('更新字典项', 'dict_item:update', 'dict_item', 'update', '/dict-items/:id', 'PUT', '更新字典项信息', 'api', 'active'),
('删除字典项', 'dict_item:delete', 'dict_item', 'delete', '/dict-items/:id', 'DELETE', '删除字典项', 'api', 'active');

-- ==========================================
-- 为超级管理员分配字典管理权限
-- ==========================================

-- 获取 admin 角色的 ID 和字典权限的 ID，然后分配
INSERT INTO role_permissions (role_id, permission_id, assigned_at)
SELECT r.id, p.id, CURRENT_TIMESTAMP
FROM roles r
CROSS JOIN permissions p
WHERE r.code = 'admin'
  AND p.resource IN ('dict_type', 'dict_item')
  AND NOT EXISTS (
    SELECT 1 FROM role_permissions rp 
    WHERE rp.role_id = r.id AND rp.permission_id = p.id
  );

-- ==========================================
-- 添加字典管理菜单（仅超级管理员可见）
-- ==========================================

-- 插入字典管理菜单（通过 permission_code 关联权限）
INSERT INTO menus (parent_id, name, title, path, component, icon, type, status, order_num, permission_code, visible, created_at, updated_at)
VALUES (
    (SELECT id FROM menus WHERE name = 'system' LIMIT 1),
    'dict',
    '字典管理',
    '/system/dict',
    '@/app/system/dict/page',
    'book',
    'menu',
    'active',
    50,
    'dict_type:read',
    true,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- ==========================================
-- 添加 Casbin 策略（字典管理权限）
-- ==========================================

-- 为 admin 角色添加字典管理的 Casbin 策略
INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
SELECT 'p', 'admin', '/api/v1/dict-types', '(GET)|(POST)', '', '', ''
WHERE NOT EXISTS (
    SELECT 1 FROM casbin_rule 
    WHERE ptype = 'p' AND v0 = 'admin' AND v1 = '/api/v1/dict-types'
);

INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
SELECT 'p', 'admin', '/api/v1/dict-types/:id', '(GET)|(PUT)|(DELETE)', '', '', ''
WHERE NOT EXISTS (
    SELECT 1 FROM casbin_rule 
    WHERE ptype = 'p' AND v0 = 'admin' AND v1 = '/api/v1/dict-types/:id'
);

INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
SELECT 'p', 'admin', '/api/v1/dict-items', '(GET)|(POST)', '', '', ''
WHERE NOT EXISTS (
    SELECT 1 FROM casbin_rule 
    WHERE ptype = 'p' AND v0 = 'admin' AND v1 = '/api/v1/dict-items'
);

INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
SELECT 'p', 'admin', '/api/v1/dict-items/:id', '(GET)|(PUT)|(DELETE)', '', '', ''
WHERE NOT EXISTS (
    SELECT 1 FROM casbin_rule 
    WHERE ptype = 'p' AND v0 = 'admin' AND v1 = '/api/v1/dict-items/:id'
);

INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
SELECT 'p', 'admin', '/api/v1/dict-items/by-type/:code', 'GET', '', '', ''
WHERE NOT EXISTS (
    SELECT 1 FROM casbin_rule 
    WHERE ptype = 'p' AND v0 = 'admin' AND v1 = '/api/v1/dict-items/by-type/:code'
);
