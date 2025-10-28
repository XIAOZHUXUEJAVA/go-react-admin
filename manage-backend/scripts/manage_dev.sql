/*
 Navicat Premium Dump SQL

 Source Server         : monitordb
 Source Server Type    : PostgreSQL
 Source Server Version : 160010 (160010)
 Source Host           : localhost:5432
 Source Catalog        : go_manage_starter
 Source Schema         : manage_dev

 Target Server Type    : PostgreSQL
 Target Server Version : 160010 (160010)
 File Encoding         : 65001

 Date: 27/10/2025 16:31:07
*/


-- ----------------------------
-- Sequence structure for audit_logs_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."audit_logs_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."audit_logs_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for casbin_rule_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."casbin_rule_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."casbin_rule_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for dict_items_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."dict_items_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."dict_items_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for dict_types_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."dict_types_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."dict_types_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for menus_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."menus_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."menus_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for migration_records_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."migration_records_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."migration_records_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for password_reset_tokens_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."password_reset_tokens_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."password_reset_tokens_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for permissions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."permissions_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."permissions_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for role_permissions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."role_permissions_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."role_permissions_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for roles_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."roles_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."roles_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_roles_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."user_roles_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."user_roles_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."users_id_seq" CASCADE;
CREATE SEQUENCE "manage_dev"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for audit_logs
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."audit_logs" CASCADE;
CREATE TABLE "manage_dev"."audit_logs" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".audit_logs_id_seq'::regclass),
  "user_id" int8,
  "username" varchar(50) COLLATE "pg_catalog"."default",
  "action" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "resource" varchar(100) COLLATE "pg_catalog"."default",
  "resource_id" varchar(100) COLLATE "pg_catalog"."default",
  "method" varchar(10) COLLATE "pg_catalog"."default",
  "path" varchar(500) COLLATE "pg_catalog"."default",
  "ip" varchar(50) COLLATE "pg_catalog"."default",
  "user_agent" text COLLATE "pg_catalog"."default",
  "status" int4,
  "error_msg" text COLLATE "pg_catalog"."default",
  "request_body" text COLLATE "pg_catalog"."default",
  "duration" int8,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "manage_dev"."audit_logs"."user_id" IS '用户ID';
COMMENT ON COLUMN "manage_dev"."audit_logs"."username" IS '用户名';
COMMENT ON COLUMN "manage_dev"."audit_logs"."action" IS '操作动作';
COMMENT ON COLUMN "manage_dev"."audit_logs"."resource" IS '资源类型';
COMMENT ON COLUMN "manage_dev"."audit_logs"."resource_id" IS '资源ID';
COMMENT ON COLUMN "manage_dev"."audit_logs"."method" IS 'HTTP方法';
COMMENT ON COLUMN "manage_dev"."audit_logs"."path" IS '请求路径';
COMMENT ON COLUMN "manage_dev"."audit_logs"."ip" IS 'IP地址';
COMMENT ON COLUMN "manage_dev"."audit_logs"."user_agent" IS '用户代理';
COMMENT ON COLUMN "manage_dev"."audit_logs"."status" IS 'HTTP状态码';
COMMENT ON COLUMN "manage_dev"."audit_logs"."error_msg" IS '错误信息';
COMMENT ON COLUMN "manage_dev"."audit_logs"."request_body" IS '请求体';
COMMENT ON COLUMN "manage_dev"."audit_logs"."duration" IS '请求耗时（毫秒）';
COMMENT ON TABLE "manage_dev"."audit_logs" IS '审计日志表';

-- ----------------------------
-- Records of audit_logs
-- ----------------------------
INSERT INTO "manage_dev"."audit_logs" VALUES (1, 0, '', '用户登录', 'auth', '', 'POST', '/api/v1/auth/login', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '{"username":"admin","password":"admin123","captcha_id":"2Bbtz3hAHh8mkZfgCGHD","captcha_code":"6143"}', 76, '2025-10-13 15:39:59.177212+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (2, 1, 'admin', '查询资源: /api/v1/users/permissions', 'users', '', 'GET', '/api/v1/users/permissions', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 33, '2025-10-13 15:39:59.26459+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (3, 1, 'admin', '查询资源: /api/v1/menus/user', 'menus', '', 'GET', '/api/v1/menus/user', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 57, '2025-10-13 15:39:59.292425+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (4, 1, 'admin', '查询资源: /api/v1/users/profile', 'users', '', 'GET', '/api/v1/users/profile', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 14, '2025-10-13 15:40:24.002652+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (5, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 18, '2025-10-13 15:40:24.186799+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (6, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 68, '2025-10-13 15:40:24.234443+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (7, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 39, '2025-10-13 15:40:24.250954+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (8, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 35, '2025-10-13 15:40:24.280801+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (9, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 8, '2025-10-13 15:40:38.988185+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (10, 1, 'admin', '查询资源: /api/v1/users/profile', 'users', '', 'GET', '/api/v1/users/profile', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 7, '2025-10-13 15:40:58.677146+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (11, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 7, '2025-10-13 15:40:58.846634+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (12, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 11, '2025-10-13 15:40:58.848959+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (13, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 20, '2025-10-13 15:40:58.876211+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (14, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 52, '2025-10-13 15:40:58.900474+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (15, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 10, '2025-10-13 15:41:06.75403+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (16, 1, 'admin', '查询资源: /api/v1/users/profile', 'users', '', 'GET', '/api/v1/users/profile', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 9, '2025-10-13 15:41:43.411024+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (17, 1, 'admin', '查询资源: /api/v1/users/profile', 'users', '', 'GET', '/api/v1/users/profile', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 6, '2025-10-13 15:41:50.41691+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (18, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 35, '2025-10-13 15:41:50.620033+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (19, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 28, '2025-10-13 15:41:50.611682+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (20, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 62, '2025-10-13 15:41:50.675075+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (21, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 67, '2025-10-13 15:41:50.701285+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (22, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 6, '2025-10-13 15:42:09.187182+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (23, 1, 'admin', '查询资源: /api/v1/users/profile', 'users', '', 'GET', '/api/v1/users/profile', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 6, '2025-10-13 15:42:29.684941+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (24, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 9, '2025-10-13 15:42:29.852097+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (25, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 11, '2025-10-13 15:42:29.853368+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (26, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 20, '2025-10-13 15:42:29.883599+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (27, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 22, '2025-10-13 15:42:29.88572+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (28, 1, 'admin', '查询资源: /api/v1/users/profile', 'users', '', 'GET', '/api/v1/users/profile', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 10, '2025-10-13 15:44:26.643035+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (29, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 6, '2025-10-13 15:44:27.027774+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (30, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 7, '2025-10-13 15:44:27.028649+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (32, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 10, '2025-10-13 15:44:27.042009+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (33, 1, 'admin', '查询资源: /api/v1/users/profile', 'users', '', 'GET', '/api/v1/users/profile', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 40, '2025-10-13 15:46:22.260869+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (34, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 26, '2025-10-13 15:46:22.557217+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (36, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 28, '2025-10-13 15:46:22.589278+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (39, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 10, '2025-10-13 15:46:40.720554+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (41, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 29, '2025-10-13 15:46:40.770827+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (42, 1, 'admin', '查询资源: /api/v1/dict-types', 'dict-types', '', 'GET', '/api/v1/dict-types', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 101, '2025-10-13 15:46:40.85411+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (43, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 7, '2025-10-13 15:46:42.848799+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (44, 1, 'admin', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 7, '2025-10-13 15:46:55.448293+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (45, 0, '', '用户登出', 'auth', '', 'POST', '/api/v1/auth/logout', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 401, '', '{"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6ImFkbWluIiwianRpIjoiZmEzMGZhMzIzM2Q1NWMyMmIzMGJkNGE4NTRlMmRmN2UiLCJpc3MiOiJnby1tYW5hZ2Utc3RhcnRlciIsInN1YiI6InJlZnJlc2giLCJleHAiOjE3NjI5MzMxOTksIm5iZiI6MTc2MDM0MTE5OSwiaWF0IjoxNzYwMzQxMTk5fQ.5LJFvsfQXoQi-OjDkDQijDX005Rs7EhVyfgzsp7q1C0"}', 0, '2025-10-13 15:47:48.409374+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (46, 0, '', '用户登录', 'auth', '', 'POST', '/api/v1/auth/login', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '{"username":"manager","password":"admin123","captcha_id":"jBCE2hQmO6ouAAJeDBRs","captcha_code":"0304"}', 68, '2025-10-13 15:47:59.794529+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (47, 4, 'manager', '查询资源: /api/v1/menus/user', 'menus', '', 'GET', '/api/v1/menus/user', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 12, '2025-10-13 15:47:59.835614+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (52, 4, 'manager', '查询资源: /api/v1/menus/user', 'menus', '', 'GET', '/api/v1/menus/user', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 8, '2025-10-13 15:48:18.199714+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (53, 4, 'manager', '查询资源: /api/v1/users/profile', 'users', '', 'GET', '/api/v1/users/profile', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 6, '2025-10-13 15:48:28.079064+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (54, 4, 'manager', '查询资源: /api/v1/users/profile', 'users', '', 'GET', '/api/v1/users/profile', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 6, '2025-10-13 15:48:32.537168+08', NULL);
INSERT INTO "manage_dev"."audit_logs" VALUES (56, 4, 'manager', '查询资源: /api/v1/dict-items', 'dict-items', '', 'GET', '/api/v1/dict-items', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '', 7, '2025-10-13 15:48:32.687089+08', NULL);
-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."casbin_rule" CASCADE;
CREATE TABLE "manage_dev"."casbin_rule" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".casbin_rule_id_seq'::regclass),
  "ptype" varchar(100) COLLATE "pg_catalog"."default",
  "v0" varchar(100) COLLATE "pg_catalog"."default",
  "v1" varchar(100) COLLATE "pg_catalog"."default",
  "v2" varchar(100) COLLATE "pg_catalog"."default",
  "v3" varchar(100) COLLATE "pg_catalog"."default",
  "v4" varchar(100) COLLATE "pg_catalog"."default",
  "v5" varchar(100) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "manage_dev"."casbin_rule"."ptype" IS '策略类型（p=策略, g=角色继承）';
COMMENT ON COLUMN "manage_dev"."casbin_rule"."v0" IS '主体（角色或用户）';
COMMENT ON COLUMN "manage_dev"."casbin_rule"."v1" IS '资源路径';
COMMENT ON COLUMN "manage_dev"."casbin_rule"."v2" IS '操作方法';
COMMENT ON TABLE "manage_dev"."casbin_rule" IS 'Casbin权限规则表';

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO "manage_dev"."casbin_rule" VALUES (7, 'p', 'role:user', '/api/users/profile', 'GET', NULL, NULL, NULL);
INSERT INTO "manage_dev"."casbin_rule" VALUES (8, 'p', 'role:user', '/api/users/profile', 'PUT', NULL, NULL, NULL);
INSERT INTO "manage_dev"."casbin_rule" VALUES (9, 'p', 'role:user', '/api/users', 'GET', NULL, NULL, NULL);
INSERT INTO "manage_dev"."casbin_rule" VALUES (10, 'g', 'user:1', 'role:admin', NULL, NULL, NULL, NULL);
INSERT INTO "manage_dev"."casbin_rule" VALUES (11, 'g', 'user:4', 'role:manager', NULL, NULL, NULL, NULL);
INSERT INTO "manage_dev"."casbin_rule" VALUES (12, 'g', 'user:3', 'role:user', NULL, NULL, NULL, NULL);
INSERT INTO "manage_dev"."casbin_rule" VALUES (13, 'g', 'user:2', 'role:user', NULL, NULL, NULL, NULL);
INSERT INTO "manage_dev"."casbin_rule" VALUES (14, 'p', 'admin', '/api/v1/dict-types', '(GET)|(POST)', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (15, 'p', 'admin', '/api/v1/dict-types/:id', '(GET)|(PUT)|(DELETE)', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (16, 'p', 'admin', '/api/v1/dict-items', '(GET)|(POST)', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (17, 'p', 'admin', '/api/v1/dict-items/:id', '(GET)|(PUT)|(DELETE)', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (18, 'p', 'admin', '/api/v1/dict-items/by-type/:code', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (41, 'p', 'role:manager', '/api/users', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (42, 'p', 'role:manager', '/api/users', 'POST', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (43, 'p', 'role:manager', '/api/users/*', 'PUT', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (44, 'p', 'role:admin', '/api/users', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (45, 'p', 'role:admin', '/api/users', 'POST', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (46, 'p', 'role:admin', '/api/users/*', 'PUT', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (47, 'p', 'role:admin', '/api/users/*', 'DELETE', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (48, 'p', 'role:admin', '/api/roles', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (49, 'p', 'role:admin', '/api/roles', 'POST', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (50, 'p', 'role:admin', '/api/roles/*', 'PUT', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (51, 'p', 'role:admin', '/api/roles/*', 'DELETE', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (52, 'p', 'role:admin', '/api/permissions', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (53, 'p', 'role:admin', '/api/permissions', 'POST', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (54, 'p', 'role:admin', '/api/permissions/*', 'PUT', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (55, 'p', 'role:admin', '/api/permissions/*', 'DELETE', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (56, 'p', 'role:admin', '/api/menus', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (57, 'p', 'role:admin', '/api/menus', 'POST', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (58, 'p', 'role:admin', '/api/menus/*', 'PUT', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (59, 'p', 'role:admin', '/api/menus/*', 'DELETE', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (60, 'p', 'role:admin', '/api/users/profile', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (61, 'p', 'role:admin', '/api/users/profile', 'PUT', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (62, 'p', 'role:admin', '/dict-types', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (63, 'p', 'role:admin', '/dict-types', 'POST', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (64, 'p', 'role:admin', '/dict-types/:id', 'PUT', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (65, 'p', 'role:admin', '/dict-types/:id', 'DELETE', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (66, 'p', 'role:admin', '/dict-items', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (67, 'p', 'role:admin', '/dict-items', 'POST', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (68, 'p', 'role:admin', '/dict-items/:id', 'PUT', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (69, 'p', 'role:admin', '/dict-items/:id', 'DELETE', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (70, 'p', 'role:admin', '/audit-logs', 'GET', '', '', '');
INSERT INTO "manage_dev"."casbin_rule" VALUES (71, 'p', 'role:admin', '/audit-logs/clean', 'POST', '', '', '');

-- ----------------------------
-- Table structure for dict_items
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."dict_items" CASCADE;
CREATE TABLE "manage_dev"."dict_items" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".dict_items_id_seq'::regclass),
  "dict_type_code" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "label" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "value" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "extra" jsonb,
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'active'::character varying,
  "sort_order" int4 DEFAULT 0,
  "is_default" bool DEFAULT false,
  "is_system" bool DEFAULT false,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "manage_dev"."dict_items"."dict_type_code" IS '关联字典类型代码';
COMMENT ON COLUMN "manage_dev"."dict_items"."label" IS '显示值';
COMMENT ON COLUMN "manage_dev"."dict_items"."value" IS '字典值';
COMMENT ON COLUMN "manage_dev"."dict_items"."extra" IS '扩展值（JSON）';
COMMENT ON COLUMN "manage_dev"."dict_items"."status" IS '状态：active-启用, inactive-禁用';
COMMENT ON COLUMN "manage_dev"."dict_items"."sort_order" IS '排序标记';
COMMENT ON COLUMN "manage_dev"."dict_items"."is_default" IS '是否默认值';
COMMENT ON COLUMN "manage_dev"."dict_items"."is_system" IS '是否系统内置';
COMMENT ON TABLE "manage_dev"."dict_items" IS '字典项表';

-- ----------------------------
-- Records of dict_items
-- ----------------------------
INSERT INTO "manage_dev"."dict_items" VALUES (1, 'user_status', '启用', 'active', '{"badge": "success", "color": "green"}', '用户账户正常可用', 'active', 1, 't', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (2, 'user_status', '禁用', 'inactive', '{"badge": "error", "color": "red"}', '用户账户被禁用', 'active', 2, 'f', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (3, 'user_status', '待审核', 'pending', '{"badge": "warning", "color": "orange"}', '用户账户待审核', 'active', 3, 'f', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (4, 'user_status', '已锁定', 'locked', '{"badge": "default", "color": "gray"}', '用户账户被锁定', 'active', 4, 'f', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (5, 'gender', '男', 'male', '{"icon": "male"}', '男性', 'active', 1, 'f', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (6, 'gender', '女', 'female', '{"icon": "female"}', '女性', 'active', 2, 'f', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (7, 'gender', '未知', 'unknown', '{"icon": "question"}', '未知或不愿透露', 'active', 3, 't', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (8, 'common_status', '启用', 'active', '{"color": "green"}', '启用状态', 'active', 1, 't', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (9, 'common_status', '禁用', 'inactive', '{"color": "red"}', '禁用状态', 'active', 2, 'f', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (10, 'yes_no', '是', 'yes', '{"color": "green"}', '是', 'active', 1, 'f', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_items" VALUES (11, 'yes_no', '否', 'no', '{"color": "gray"}', '否', 'active', 2, 't', 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);

-- ----------------------------
-- Table structure for dict_types
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."dict_types" CASCADE;
CREATE TABLE "manage_dev"."dict_types" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".dict_types_id_seq'::regclass),
  "code" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'active'::character varying,
  "sort_order" int4 DEFAULT 0,
  "is_system" bool DEFAULT false,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "manage_dev"."dict_types"."code" IS '字典类型代码';
COMMENT ON COLUMN "manage_dev"."dict_types"."name" IS '字典类型名称';
COMMENT ON COLUMN "manage_dev"."dict_types"."description" IS '字典类型描述';
COMMENT ON COLUMN "manage_dev"."dict_types"."status" IS '状态：active-启用, inactive-禁用';
COMMENT ON COLUMN "manage_dev"."dict_types"."sort_order" IS '排序标记';
COMMENT ON COLUMN "manage_dev"."dict_types"."is_system" IS '是否系统内置';
COMMENT ON TABLE "manage_dev"."dict_types" IS '字典类型表';

-- ----------------------------
-- Records of dict_types
-- ----------------------------
INSERT INTO "manage_dev"."dict_types" VALUES (1, 'user_status', '用户状态', '用户账户状态枚举', 'active', 1, 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_types" VALUES (2, 'gender', '性别', '用户性别枚举', 'active', 2, 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_types" VALUES (3, 'common_status', '通用状态', '通用的启用/禁用状态', 'active', 3, 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);
INSERT INTO "manage_dev"."dict_types" VALUES (4, 'yes_no', '是否标识', '通用的是/否标识', 'active', 4, 't', '2025-10-13 15:38:32.332429+08', '2025-10-13 15:38:32.332429+08', NULL);

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."menus" CASCADE;
CREATE TABLE "manage_dev"."menus" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".menus_id_seq'::regclass),
  "parent_id" int8,
  "name" text COLLATE "pg_catalog"."default" NOT NULL,
  "title" text COLLATE "pg_catalog"."default" NOT NULL,
  "path" text COLLATE "pg_catalog"."default",
  "component" text COLLATE "pg_catalog"."default",
  "icon" text COLLATE "pg_catalog"."default",
  "order_num" int8 DEFAULT 0,
  "type" text COLLATE "pg_catalog"."default" NOT NULL,
  "permission_code" text COLLATE "pg_catalog"."default",
  "visible" bool DEFAULT true,
  "status" text COLLATE "pg_catalog"."default" DEFAULT 'active'::text,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "manage_dev"."menus"."parent_id" IS '父菜单ID';
COMMENT ON COLUMN "manage_dev"."menus"."name" IS '菜单名称（路由名）';
COMMENT ON COLUMN "manage_dev"."menus"."title" IS '菜单标题';
COMMENT ON COLUMN "manage_dev"."menus"."path" IS '菜单路径';
COMMENT ON COLUMN "manage_dev"."menus"."component" IS '组件路径';
COMMENT ON COLUMN "manage_dev"."menus"."icon" IS '菜单图标';
COMMENT ON COLUMN "manage_dev"."menus"."order_num" IS '排序号';
COMMENT ON COLUMN "manage_dev"."menus"."type" IS '菜单类型';
COMMENT ON COLUMN "manage_dev"."menus"."permission_code" IS '权限代码';
COMMENT ON COLUMN "manage_dev"."menus"."visible" IS '是否可见';
COMMENT ON COLUMN "manage_dev"."menus"."status" IS '状态';
COMMENT ON TABLE "manage_dev"."menus" IS '菜单表';

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO "manage_dev"."menus" VALUES (1, NULL, 'dashboard', '工作台', '/dashboard', '@/app/dashboard/page', 'LayoutDashboard', 1, 'menu', NULL, 't', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."menus" VALUES (3, 2, 'user-management', '用户管理', '/system/users', '@/app/system/users/page', 'Users', 101, 'menu', 'user:read', 't', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."menus" VALUES (4, 2, 'role-management', '角色管理', '/system/roles', '@/app/system/roles/page', 'Shield', 102, 'menu', 'role:read', 't', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."menus" VALUES (5, 2, 'permission-management', '权限管理', '/system/permissions', '@/app/system/permissions/page', 'Key', 103, 'menu', 'permission:read', 't', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."menus" VALUES (6, 2, 'menu-management', '菜单管理', '/system/menus', '@/app/system/menus/page', 'Menu', 104, 'menu', 'menu:read', 't', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."menus" VALUES (7, 2, 'dict', '字典管理', '/system/dict', '@/app/system/dict/page', 'Server', 50, 'menu', 'dict_type:read', 't', 'active', '2025-10-13 15:38:32.332429+08', '2025-10-15 16:57:24.042573+08', NULL);
INSERT INTO "manage_dev"."menus" VALUES (2, NULL, 'system', '系统管理', '/system', '', 'Settings', 2, 'menu', 'system:read', 't', 'active', '0001-01-01 08:05:43+08:05:43', '2025-10-16 09:19:47.805042+08', NULL);
INSERT INTO "manage_dev"."menus" VALUES (10, 9, 'test-menu-child', '测试删除子菜单', '/test/children', '', 'FileText', 2, 'menu', '', 't', 'active', '2025-10-16 14:14:35.192076+08', '2025-10-16 14:14:35.192076+08', '2025-10-16 14:14:43.42447+08');
INSERT INTO "manage_dev"."menus" VALUES (9, NULL, 'test-menu', '测试删除顶级菜单', '/test', '', 'FileText', 4, 'menu', '', 't', 'active', '2025-10-16 14:13:10.470075+08', '2025-10-16 14:13:10.470075+08', '2025-10-16 14:14:47.147965+08');
INSERT INTO "manage_dev"."menus" VALUES (11, NULL, 'test-page-new', '测试菜单', '/test/test', '', 'FileText', 0, 'menu', '', 't', 'active', '2025-10-17 09:57:18.327522+08', '2025-10-17 09:57:25.235616+08', '2025-10-17 09:57:55.868942+08');
INSERT INTO "manage_dev"."menus" VALUES (8, NULL, 'logs', '日志管理', '/logs', '@/app/logs/page', 'Calendar', 3, 'menu', 'logs:read', 't', 'active', '2025-10-16 09:19:25.049774+08', '2025-10-18 10:50:47.258426+08', NULL);

-- ----------------------------
-- Table structure for migration_records
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."migration_records" CASCADE;
CREATE TABLE "manage_dev"."migration_records" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".migration_records_id_seq'::regclass),
  "migration_id" text COLLATE "pg_catalog"."default" NOT NULL,
  "executed_at" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "manage_dev"."migration_records"."migration_id" IS '迁移ID';
COMMENT ON COLUMN "manage_dev"."migration_records"."executed_at" IS '执行时间';
COMMENT ON TABLE "manage_dev"."migration_records" IS '数据库迁移记录表';

-- ----------------------------
-- Records of migration_records
-- ----------------------------
INSERT INTO "manage_dev"."migration_records" VALUES (1, '001_create_users_table', '2025-10-13 15:38:32.24265+08');
INSERT INTO "manage_dev"."migration_records" VALUES (2, '002_create_rbac_tables', '2025-10-13 15:38:32.290381+08');
INSERT INTO "manage_dev"."migration_records" VALUES (3, '003_create_audit_logs_table', '2025-10-13 15:38:32.324239+08');
INSERT INTO "manage_dev"."migration_records" VALUES (4, '004_create_dict_tables', '2025-10-13 15:38:32.384307+08');

-- ----------------------------
-- Table structure for password_reset_tokens
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."password_reset_tokens" CASCADE;
CREATE TABLE "manage_dev"."password_reset_tokens" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".password_reset_tokens_id_seq'::regclass),
  "user_id" int8 NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "token" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "expires_at" timestamptz(6) NOT NULL,
  "used_at" timestamptz(6),
  "ip_address" varchar(50) COLLATE "pg_catalog"."default",
  "user_agent" text COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "manage_dev"."password_reset_tokens"."token" IS '重置Token（UUID格式）';
COMMENT ON COLUMN "manage_dev"."password_reset_tokens"."expires_at" IS 'Token过期时间（建议1小时）';
COMMENT ON COLUMN "manage_dev"."password_reset_tokens"."used_at" IS 'Token使用时间（NULL表示未使用）';
COMMENT ON TABLE "manage_dev"."password_reset_tokens" IS '密码重置Token表';

-- ----------------------------
-- Records of password_reset_tokens
-- ----------------------------
INSERT INTO "manage_dev"."password_reset_tokens" VALUES (1, 2, 'xiaozhulzq@2925.com', 'ff9eb11c-0fad-40c5-8a54-0424199eedd5', '2025-10-23 11:20:08.011024+08', '2025-10-23 10:21:08.586758+08', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', '2025-10-23 10:20:08.011689+08');
INSERT INTO "manage_dev"."password_reset_tokens" VALUES (7, 2, 'xiaozhulzq@2925.com', '641c861d-a98b-4a27-a9c8-3465d555306c', '2025-10-24 11:54:20.319569+08', NULL, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', '2025-10-24 10:54:20.320521+08');

-- ----------------------------
-- Table structure for permissions
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."permissions" CASCADE;
CREATE TABLE "manage_dev"."permissions" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".permissions_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default" NOT NULL,
  "code" text COLLATE "pg_catalog"."default" NOT NULL,
  "resource" text COLLATE "pg_catalog"."default" NOT NULL,
  "action" text COLLATE "pg_catalog"."default" NOT NULL,
  "path" text COLLATE "pg_catalog"."default",
  "method" text COLLATE "pg_catalog"."default",
  "description" text COLLATE "pg_catalog"."default",
  "type" text COLLATE "pg_catalog"."default" DEFAULT 'api'::text,
  "status" text COLLATE "pg_catalog"."default" DEFAULT 'active'::text,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "manage_dev"."permissions"."name" IS '权限名称';
COMMENT ON COLUMN "manage_dev"."permissions"."code" IS '权限代码';
COMMENT ON COLUMN "manage_dev"."permissions"."resource" IS '资源类型';
COMMENT ON COLUMN "manage_dev"."permissions"."action" IS '操作动作';
COMMENT ON COLUMN "manage_dev"."permissions"."path" IS 'API路径';
COMMENT ON COLUMN "manage_dev"."permissions"."method" IS 'HTTP方法';
COMMENT ON COLUMN "manage_dev"."permissions"."type" IS '权限类型';
COMMENT ON COLUMN "manage_dev"."permissions"."status" IS '状态';
COMMENT ON TABLE "manage_dev"."permissions" IS '权限表';

-- ----------------------------
-- Records of permissions
-- ----------------------------
INSERT INTO "manage_dev"."permissions" VALUES (1, '查看用户', 'user:read', 'user', 'read', '/api/users', 'GET', '查看用户列表和详情', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (2, '创建用户', 'user:create', 'user', 'create', '/api/users', 'POST', '创建新用户', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (3, '更新用户', 'user:update', 'user', 'update', '/api/users/*', 'PUT', '更新用户信息', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (4, '删除用户', 'user:delete', 'user', 'delete', '/api/users/*', 'DELETE', '删除用户', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (5, '查看角色', 'role:read', 'role', 'read', '/api/roles', 'GET', '查看角色列表和详情', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (6, '创建角色', 'role:create', 'role', 'create', '/api/roles', 'POST', '创建新角色', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (7, '更新角色', 'role:update', 'role', 'update', '/api/roles/*', 'PUT', '更新角色信息', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (8, '删除角色', 'role:delete', 'role', 'delete', '/api/roles/*', 'DELETE', '删除角色', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (9, '查看权限', 'permission:read', 'permission', 'read', '/api/permissions', 'GET', '查看权限列表', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (10, '创建权限', 'permission:create', 'permission', 'create', '/api/permissions', 'POST', '创建新权限', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (11, '更新权限', 'permission:update', 'permission', 'update', '/api/permissions/*', 'PUT', '更新权限信息', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (12, '删除权限', 'permission:delete', 'permission', 'delete', '/api/permissions/*', 'DELETE', '删除权限', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (13, '查看菜单', 'menu:read', 'menu', 'read', '/api/menus', 'GET', '查看菜单列表', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (15, '更新菜单', 'menu:update', 'menu', 'update', '/api/menus/*', 'PUT', '更新菜单信息', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (16, '删除菜单', 'menu:delete', 'menu', 'delete', '/api/menus/*', 'DELETE', '删除菜单', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (17, '查看个人信息', 'profile:read', 'profile', 'read', '/api/users/profile', 'GET', '查看自己的个人信息', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (18, '更新个人信息', 'profile:update', 'profile', 'update', '/api/users/profile', 'PUT', '更新自己的个人信息', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (19, '查看字典类型', 'dict_type:read', 'dict_type', 'read', '/dict-types', 'GET', '查看字典类型列表', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (20, '创建字典类型', 'dict_type:create', 'dict_type', 'create', '/dict-types', 'POST', '创建新字典类型', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (21, '更新字典类型', 'dict_type:update', 'dict_type', 'update', '/dict-types/:id', 'PUT', '更新字典类型信息', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (22, '删除字典类型', 'dict_type:delete', 'dict_type', 'delete', '/dict-types/:id', 'DELETE', '删除字典类型', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (23, '查看字典项', 'dict_item:read', 'dict_item', 'read', '/dict-items', 'GET', '查看字典项列表', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (24, '创建字典项', 'dict_item:create', 'dict_item', 'create', '/dict-items', 'POST', '创建新字典项', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (25, '更新字典项', 'dict_item:update', 'dict_item', 'update', '/dict-items/:id', 'PUT', '更新字典项信息', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (26, '删除字典项', 'dict_item:delete', 'dict_item', 'delete', '/dict-items/:id', 'DELETE', '删除字典项', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (27, '查看系统管理', 'system:read', 'system', 'read', '/system', '', '', 'menu', 'active', '2025-10-15 17:01:01.730473+08', '2025-10-15 17:01:01.730473+08', NULL);
INSERT INTO "manage_dev"."permissions" VALUES (28, '日志管理', 'logs:read', 'logs', 'read', '/logs', '', '', 'menu', 'active', '2025-10-16 09:17:15.716852+08', '2025-10-16 09:17:15.716852+08', NULL);
INSERT INTO "manage_dev"."permissions" VALUES (29, '查看日志', 'logs:get', 'logs', 'get', '/audit-logs', 'GET', '', 'api', 'active', '2025-10-18 10:56:28.334156+08', '2025-10-18 10:56:28.334156+08', NULL);
INSERT INTO "manage_dev"."permissions" VALUES (30, '清理日志', 'logs:clean', 'logs', 'clean', '/audit-logs/clean', 'POST', '', 'api', 'active', '2025-10-18 11:00:15.620183+08', '2025-10-18 11:00:15.620183+08', NULL);
INSERT INTO "manage_dev"."permissions" VALUES (14, '创建菜单', 'menu:create', 'menu', 'create', '/api/menus', 'POST', '创建新菜单', 'api', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."permissions" VALUES (31, '测试权限', 'test:test', 'test', 'test', '', '', '', 'api', 'active', '2025-10-20 11:17:57.572313+08', '2025-10-20 11:17:57.572313+08', '2025-10-20 11:18:09.200203+08');
INSERT INTO "manage_dev"."permissions" VALUES (32, '测试测出', 'test:test2', 'test', 'test2', '', '', '', 'api', 'active', '2025-10-20 11:57:14.8564+08', '2025-10-20 11:57:14.8564+08', '2025-10-20 11:57:37.114172+08');
INSERT INTO "manage_dev"."permissions" VALUES (33, '测试测试测试1', 'test:test8', 'test', 'test8', '', '', '', 'api', 'active', '2025-10-20 15:58:51.694298+08', '2025-10-20 15:59:07.841682+08', '2025-10-20 15:59:13.278167+08');

-- ----------------------------
-- Table structure for role_permissions
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."role_permissions" CASCADE;
CREATE TABLE "manage_dev"."role_permissions" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".role_permissions_id_seq'::regclass),
  "role_id" int8 NOT NULL,
  "permission_id" int8 NOT NULL,
  "assigned_at" timestamptz(6)
)
;
COMMENT ON COLUMN "manage_dev"."role_permissions"."role_id" IS '角色ID';
COMMENT ON COLUMN "manage_dev"."role_permissions"."permission_id" IS '权限ID';
COMMENT ON COLUMN "manage_dev"."role_permissions"."assigned_at" IS '分配时间';
COMMENT ON TABLE "manage_dev"."role_permissions" IS '角色权限关联表';

-- ----------------------------
-- Records of role_permissions
-- ----------------------------
INSERT INTO "manage_dev"."role_permissions" VALUES (22, 3, 17, NULL);
INSERT INTO "manage_dev"."role_permissions" VALUES (23, 3, 18, NULL);
INSERT INTO "manage_dev"."role_permissions" VALUES (55, 2, 27, '2025-10-15 17:14:38.170666+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (56, 2, 2, '2025-10-15 17:14:38.170666+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (57, 2, 1, '2025-10-15 17:14:38.170666+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (58, 2, 3, '2025-10-15 17:14:38.170666+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (59, 1, 24, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (60, 1, 26, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (61, 1, 23, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (62, 1, 25, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (63, 1, 20, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (64, 1, 22, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (65, 1, 19, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (66, 1, 21, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (67, 1, 14, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (68, 1, 16, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (69, 1, 13, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (70, 1, 15, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (71, 1, 10, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (72, 1, 12, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (73, 1, 9, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (74, 1, 11, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (75, 1, 17, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (76, 1, 18, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (77, 1, 6, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (78, 1, 8, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (79, 1, 5, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (80, 1, 7, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (81, 1, 27, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (82, 1, 2, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (83, 1, 4, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (84, 1, 1, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (85, 1, 3, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (86, 1, 28, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (87, 1, 29, '2025-10-18 11:21:25.296831+08');
INSERT INTO "manage_dev"."role_permissions" VALUES (88, 1, 30, '2025-10-18 11:21:25.296831+08');

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."roles" CASCADE;
CREATE TABLE "manage_dev"."roles" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".roles_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default" NOT NULL,
  "code" text COLLATE "pg_catalog"."default" NOT NULL,
  "description" text COLLATE "pg_catalog"."default",
  "status" text COLLATE "pg_catalog"."default" DEFAULT 'active'::text,
  "is_system" bool DEFAULT false,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "manage_dev"."roles"."name" IS '角色名称';
COMMENT ON COLUMN "manage_dev"."roles"."code" IS '角色代码';
COMMENT ON COLUMN "manage_dev"."roles"."description" IS '角色描述';
COMMENT ON COLUMN "manage_dev"."roles"."status" IS '状态';
COMMENT ON COLUMN "manage_dev"."roles"."is_system" IS '是否系统角色';
COMMENT ON TABLE "manage_dev"."roles" IS '角色表';

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO "manage_dev"."roles" VALUES (1, '超级管理员', 'admin', '系统最高权限，拥有所有操作权限', 'active', 't', NULL, NULL, NULL);
INSERT INTO "manage_dev"."roles" VALUES (2, '管理员', 'manager', '管理权限，可以管理用户和基础配置', 'active', 't', NULL, NULL, NULL);
INSERT INTO "manage_dev"."roles" VALUES (3, '普通用户', 'user', '基础权限，只能访问基本功能', 'active', 't', NULL, NULL, NULL);
INSERT INTO "manage_dev"."roles" VALUES (4, '测试角色', 'testrole', '', 'active', 'f', '2025-10-20 11:07:16.933402+08', '2025-10-20 11:07:16.933402+08', '2025-10-20 11:07:23.762151+08');
INSERT INTO "manage_dev"."roles" VALUES (6, '测试角色1', 'testrole', '', 'active', 'f', '2025-10-20 15:46:20.338954+08', '2025-10-20 15:46:27.440498+08', NULL);

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."user_roles" CASCADE;
CREATE TABLE "manage_dev"."user_roles" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".user_roles_id_seq'::regclass),
  "user_id" int8 NOT NULL,
  "role_id" int8 NOT NULL,
  "assigned_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" int8
)
;
COMMENT ON COLUMN "manage_dev"."user_roles"."user_id" IS '用户ID';
COMMENT ON COLUMN "manage_dev"."user_roles"."role_id" IS '角色ID';
COMMENT ON COLUMN "manage_dev"."user_roles"."assigned_at" IS '分配时间';
COMMENT ON COLUMN "manage_dev"."user_roles"."assigned_by" IS '分配人';
COMMENT ON TABLE "manage_dev"."user_roles" IS '用户角色关联表';

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO "manage_dev"."user_roles" VALUES (1, 1, 1, '2025-10-13 15:38:32.253633+08', NULL);
INSERT INTO "manage_dev"."user_roles" VALUES (2, 4, 2, '2025-10-13 15:38:32.253633+08', NULL);
INSERT INTO "manage_dev"."user_roles" VALUES (3, 3, 3, '2025-10-13 15:38:32.253633+08', NULL);
INSERT INTO "manage_dev"."user_roles" VALUES (4, 2, 3, '2025-10-13 15:38:32.253633+08', NULL);
INSERT INTO "manage_dev"."user_roles" VALUES (5, 5, 3, '2025-10-20 11:06:50.634507+08', 0);
INSERT INTO "manage_dev"."user_roles" VALUES (6, 8, 3, '2025-10-20 14:27:29.637082+08', 0);
INSERT INTO "manage_dev"."user_roles" VALUES (7, 10, 3, '2025-10-20 14:32:45.587809+08', 0);
INSERT INTO "manage_dev"."user_roles" VALUES (8, 11, 3, '2025-10-20 14:48:45.129333+08', 0);
INSERT INTO "manage_dev"."user_roles" VALUES (9, 12, 3, '2025-10-20 15:10:15.775768+08', 0);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."users" CASCADE;
CREATE TABLE "manage_dev"."users" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".users_id_seq'::regclass),
  "username" text COLLATE "pg_catalog"."default" NOT NULL,
  "email" text COLLATE "pg_catalog"."default" NOT NULL,
  "password" text COLLATE "pg_catalog"."default" NOT NULL,
  "role" text COLLATE "pg_catalog"."default" DEFAULT 'user'::text,
  "status" text COLLATE "pg_catalog"."default" DEFAULT 'active'::text,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "manage_dev"."users"."username" IS '用户名';
COMMENT ON COLUMN "manage_dev"."users"."email" IS '邮箱';
COMMENT ON COLUMN "manage_dev"."users"."password" IS '密码（加密）';
COMMENT ON COLUMN "manage_dev"."users"."role" IS '角色';
COMMENT ON COLUMN "manage_dev"."users"."status" IS '状态';
COMMENT ON TABLE "manage_dev"."users" IS '用户表';

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "manage_dev"."users" VALUES (1, 'admin', 'admin@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'admin', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."users" VALUES (3, 'user2', 'user2@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'user', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."users" VALUES (8, 'xiaozhu1', 'xiaozhulzqwork1@2925.com', '$2a$10$nt1Fsp2rwehcpqvd79hgpuAJrv8KsNqV55NsavoIyVrhJneyUO6R2', 'user', 'active', '2025-10-20 14:27:29.616307+08', '2025-10-20 14:27:29.616307+08', '2025-10-20 14:27:45.467468+08');
INSERT INTO "manage_dev"."users" VALUES (10, 'xiaozhu', 'xiaozhulzqwork1@2925.com', '$2a$10$BGOO/MjN4ixzzfSBLjeVvOb/q5mSmrnoWYLB6p9.MzCSHGJ2rWaaW', 'user', 'active', '2025-10-20 14:32:45.566734+08', '2025-10-20 14:32:45.566734+08', NULL);
INSERT INTO "manage_dev"."users" VALUES (11, '测试新增用户1', 'xiaozhulzqwork5@2925.com', '$2a$10$CHTHg196qYa/qUTGALrTfO14cOXXTLOuvi2x155mwpTmGyijLUC3K', 'user', 'active', '2025-10-20 14:48:45.111729+08', '2025-10-20 14:48:52.662225+08', NULL);
INSERT INTO "manage_dev"."users" VALUES (12, '测试信的埃塞挨打', 'xiaozhulzqgame11@2925.com', '$2a$10$YxWwc6G4vIE2HvKOjAnz2ud1guRqPTRCYfeliVO0po8w6YsSszgbm', 'user', 'active', '2025-10-20 15:10:15.762983+08', '2025-10-20 15:10:26.242874+08', NULL);
INSERT INTO "manage_dev"."users" VALUES (4, 'manager', 'xiaozhulzq@2925.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'manager', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."users" VALUES (2, 'user1', 'xiaozhulzq@2925.com', '$2a$10$cYuTag0oSlnO8O/N4mSaIO8Fedq9n3bRpXB71XcgkiZoyrxAzJI5O', 'user', 'active', NULL, '2025-10-23 10:21:08.559969+08', NULL);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."audit_logs_id_seq"
OWNED BY "manage_dev"."audit_logs"."id";
SELECT setval('"manage_dev"."audit_logs_id_seq"', 1658, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."casbin_rule_id_seq"
OWNED BY "manage_dev"."casbin_rule"."id";
SELECT setval('"manage_dev"."casbin_rule_id_seq"', 71, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."dict_items_id_seq"
OWNED BY "manage_dev"."dict_items"."id";
SELECT setval('"manage_dev"."dict_items_id_seq"', 11, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."dict_types_id_seq"
OWNED BY "manage_dev"."dict_types"."id";
SELECT setval('"manage_dev"."dict_types_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."menus_id_seq"
OWNED BY "manage_dev"."menus"."id";
SELECT setval('"manage_dev"."menus_id_seq"', 11, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."migration_records_id_seq"
OWNED BY "manage_dev"."migration_records"."id";
SELECT setval('"manage_dev"."migration_records_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."password_reset_tokens_id_seq"
OWNED BY "manage_dev"."password_reset_tokens"."id";
SELECT setval('"manage_dev"."password_reset_tokens_id_seq"', 7, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."permissions_id_seq"
OWNED BY "manage_dev"."permissions"."id";
SELECT setval('"manage_dev"."permissions_id_seq"', 33, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."role_permissions_id_seq"
OWNED BY "manage_dev"."role_permissions"."id";
SELECT setval('"manage_dev"."role_permissions_id_seq"', 88, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."roles_id_seq"
OWNED BY "manage_dev"."roles"."id";
SELECT setval('"manage_dev"."roles_id_seq"', 6, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."user_roles_id_seq"
OWNED BY "manage_dev"."user_roles"."id";
SELECT setval('"manage_dev"."user_roles_id_seq"', 9, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."users_id_seq"
OWNED BY "manage_dev"."users"."id";
SELECT setval('"manage_dev"."users_id_seq"', 12, true);

-- ----------------------------
-- Indexes structure for table audit_logs
-- ----------------------------
CREATE INDEX "idx_audit_logs_action" ON "manage_dev"."audit_logs" USING btree (
  "action" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_audit_logs_created_at" ON "manage_dev"."audit_logs" USING btree (
  "created_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_audit_logs_deleted_at" ON "manage_dev"."audit_logs" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_audit_logs_resource" ON "manage_dev"."audit_logs" USING btree (
  "resource" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_audit_logs_status" ON "manage_dev"."audit_logs" USING btree (
  "status" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "idx_audit_logs_user_id" ON "manage_dev"."audit_logs" USING btree (
  "user_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table audit_logs
-- ----------------------------
ALTER TABLE "manage_dev"."audit_logs" ADD CONSTRAINT "audit_logs_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table casbin_rule
-- ----------------------------
CREATE INDEX "idx_casbin_rule" ON "manage_dev"."casbin_rule" USING btree (
  "ptype" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v0" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v1" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v2" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table casbin_rule
-- ----------------------------
ALTER TABLE "manage_dev"."casbin_rule" ADD CONSTRAINT "casbin_rule_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table dict_items
-- ----------------------------
CREATE INDEX "idx_dict_items_deleted_at" ON "manage_dev"."dict_items" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_dict_items_dict_type_code" ON "manage_dev"."dict_items" USING btree (
  "dict_type_code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_dict_items_extra" ON "manage_dev"."dict_items" USING gin (
  "extra" "pg_catalog"."jsonb_ops"
);
CREATE INDEX "idx_dict_items_is_default" ON "manage_dev"."dict_items" USING btree (
  "is_default" "pg_catalog"."bool_ops" ASC NULLS LAST
);
CREATE INDEX "idx_dict_items_sort_order" ON "manage_dev"."dict_items" USING btree (
  "sort_order" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "idx_dict_items_status" ON "manage_dev"."dict_items" USING btree (
  "status" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_dict_items_value" ON "manage_dev"."dict_items" USING btree (
  "value" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table dict_items
-- ----------------------------
ALTER TABLE "manage_dev"."dict_items" ADD CONSTRAINT "dict_items_dict_type_code_value_key" UNIQUE ("dict_type_code", "value");

-- ----------------------------
-- Primary Key structure for table dict_items
-- ----------------------------
ALTER TABLE "manage_dev"."dict_items" ADD CONSTRAINT "dict_items_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table dict_types
-- ----------------------------
CREATE INDEX "idx_dict_types_code" ON "manage_dev"."dict_types" USING btree (
  "code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_dict_types_deleted_at" ON "manage_dev"."dict_types" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_dict_types_sort_order" ON "manage_dev"."dict_types" USING btree (
  "sort_order" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "idx_dict_types_status" ON "manage_dev"."dict_types" USING btree (
  "status" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table dict_types
-- ----------------------------
ALTER TABLE "manage_dev"."dict_types" ADD CONSTRAINT "dict_types_code_key" UNIQUE ("code");

-- ----------------------------
-- Primary Key structure for table dict_types
-- ----------------------------
ALTER TABLE "manage_dev"."dict_types" ADD CONSTRAINT "dict_types_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table menus
-- ----------------------------
CREATE INDEX "idx_menus_deleted_at" ON "manage_dev"."menus" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_menus_order_num" ON "manage_dev"."menus" USING btree (
  "order_num" "pg_catalog"."int8_ops" ASC NULLS LAST
);
CREATE INDEX "idx_menus_parent_id" ON "manage_dev"."menus" USING btree (
  "parent_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);
CREATE INDEX "idx_menus_permission_code" ON "manage_dev"."menus" USING btree (
  "permission_code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table menus
-- ----------------------------
ALTER TABLE "manage_dev"."menus" ADD CONSTRAINT "menus_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table migration_records
-- ----------------------------
CREATE UNIQUE INDEX "idx_migration_records_migration_id" ON "manage_dev"."migration_records" USING btree (
  "migration_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table migration_records
-- ----------------------------
ALTER TABLE "manage_dev"."migration_records" ADD CONSTRAINT "migration_records_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table password_reset_tokens
-- ----------------------------
CREATE INDEX "idx_password_reset_tokens_email" ON "manage_dev"."password_reset_tokens" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_password_reset_tokens_expires_at" ON "manage_dev"."password_reset_tokens" USING btree (
  "expires_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_password_reset_tokens_token" ON "manage_dev"."password_reset_tokens" USING btree (
  "token" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_password_reset_tokens_user_id" ON "manage_dev"."password_reset_tokens" USING btree (
  "user_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table password_reset_tokens
-- ----------------------------
ALTER TABLE "manage_dev"."password_reset_tokens" ADD CONSTRAINT "password_reset_tokens_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table permissions
-- ----------------------------
CREATE UNIQUE INDEX "idx_permissions_code" ON "manage_dev"."permissions" USING btree (
  "code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_permissions_deleted_at" ON "manage_dev"."permissions" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_permissions_resource_action" ON "manage_dev"."permissions" USING btree (
  "resource" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "action" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_permissions_type" ON "manage_dev"."permissions" USING btree (
  "type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table permissions
-- ----------------------------
ALTER TABLE "manage_dev"."permissions" ADD CONSTRAINT "permissions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table role_permissions
-- ----------------------------
CREATE UNIQUE INDEX "idx_role_permission" ON "manage_dev"."role_permissions" USING btree (
  "role_id" "pg_catalog"."int8_ops" ASC NULLS LAST,
  "permission_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);
CREATE INDEX "idx_role_permissions_permission_id" ON "manage_dev"."role_permissions" USING btree (
  "permission_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);
CREATE INDEX "idx_role_permissions_role_id" ON "manage_dev"."role_permissions" USING btree (
  "role_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table role_permissions
-- ----------------------------
ALTER TABLE "manage_dev"."role_permissions" ADD CONSTRAINT "role_permissions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table roles
-- ----------------------------
CREATE INDEX "idx_roles_code" ON "manage_dev"."roles" USING btree (
  "code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_roles_deleted_at" ON "manage_dev"."roles" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_roles_name" ON "manage_dev"."roles" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_roles_status" ON "manage_dev"."roles" USING btree (
  "status" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table roles
-- ----------------------------
ALTER TABLE "manage_dev"."roles" ADD CONSTRAINT "roles_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_roles
-- ----------------------------
CREATE UNIQUE INDEX "idx_user_role" ON "manage_dev"."user_roles" USING btree (
  "user_id" "pg_catalog"."int8_ops" ASC NULLS LAST,
  "role_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);
CREATE INDEX "idx_user_roles_role_id" ON "manage_dev"."user_roles" USING btree (
  "role_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);
CREATE INDEX "idx_user_roles_user_id" ON "manage_dev"."user_roles" USING btree (
  "user_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_roles
-- ----------------------------
ALTER TABLE "manage_dev"."user_roles" ADD CONSTRAINT "user_roles_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE INDEX "idx_users_deleted_at" ON "manage_dev"."users" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_users_email" ON "manage_dev"."users" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "idx_users_username" ON "manage_dev"."users" USING btree (
  "username" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "manage_dev"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table dict_items
-- ----------------------------
ALTER TABLE "manage_dev"."dict_items" ADD CONSTRAINT "dict_items_dict_type_code_fkey" FOREIGN KEY ("dict_type_code") REFERENCES "manage_dev"."dict_types" ("code") ON DELETE CASCADE ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table password_reset_tokens
-- ----------------------------
ALTER TABLE "manage_dev"."password_reset_tokens" ADD CONSTRAINT "fk_password_reset_tokens_user" FOREIGN KEY ("user_id") REFERENCES "manage_dev"."users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
