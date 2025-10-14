-- ==========================================
-- 回滚字典管理系统
-- ==========================================

-- 1. 删除 Casbin 策略
DELETE FROM casbin_rule 
WHERE ptype = 'p' 
  AND v0 = 'admin' 
  AND v1 LIKE '/api/v1/dict-%';

-- 2. 删除菜单
DELETE FROM menus WHERE name = 'dict';

-- 3. 删除角色权限关联
DELETE FROM role_permissions 
WHERE permission_id IN (
    SELECT id FROM permissions WHERE resource IN ('dict_type', 'dict_item')
);

-- 4. 删除权限
DELETE FROM permissions WHERE resource IN ('dict_type', 'dict_item');

-- 5. 删除字典项表
DROP TABLE IF EXISTS dict_items;

-- 6. 删除字典类型表
DROP TABLE IF EXISTS dict_types;
