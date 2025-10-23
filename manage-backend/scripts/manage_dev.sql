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

 Date: 22/10/2025 14:47:38
*/


-- ----------------------------
-- Sequence structure for audit_logs_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."audit_logs_id_seq";
CREATE SEQUENCE "manage_dev"."audit_logs_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for casbin_rule_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."casbin_rule_id_seq";
CREATE SEQUENCE "manage_dev"."casbin_rule_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for dict_items_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."dict_items_id_seq";
CREATE SEQUENCE "manage_dev"."dict_items_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for dict_types_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."dict_types_id_seq";
CREATE SEQUENCE "manage_dev"."dict_types_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for menus_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."menus_id_seq";
CREATE SEQUENCE "manage_dev"."menus_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for migration_records_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."migration_records_id_seq";
CREATE SEQUENCE "manage_dev"."migration_records_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for permissions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."permissions_id_seq";
CREATE SEQUENCE "manage_dev"."permissions_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for role_permissions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."role_permissions_id_seq";
CREATE SEQUENCE "manage_dev"."role_permissions_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for roles_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."roles_id_seq";
CREATE SEQUENCE "manage_dev"."roles_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_roles_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."user_roles_id_seq";
CREATE SEQUENCE "manage_dev"."user_roles_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "manage_dev"."users_id_seq";
CREATE SEQUENCE "manage_dev"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for audit_logs
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."audit_logs";
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

-- ----------------------------
-- Records of audit_logs
-- ----------------------------
INSERT INTO "manage_dev"."audit_logs" VALUES (1, 0, '', '用户登录', 'auth', '', 'POST', '/api/v1/auth/login', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36', 200, '', '{"username":"admin","password":"admin123","captcha_id":"2Bbtz3hAHh8mkZfgCGHD","captcha_code":"6143"}', 76, '2025-10-13 15:39:59.177212+08', NULL);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS "manage_dev"."roles";
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
DROP TABLE IF EXISTS "manage_dev"."user_roles";
CREATE TABLE "manage_dev"."user_roles" (
  "id" int8 NOT NULL DEFAULT nextval('"manage_dev".user_roles_id_seq'::regclass),
  "user_id" int8 NOT NULL,
  "role_id" int8 NOT NULL,
  "assigned_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" int8
)
;

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
DROP TABLE IF EXISTS "manage_dev"."users";
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

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "manage_dev"."users" VALUES (1, 'admin', 'admin@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'admin', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."users" VALUES (2, 'user1', 'user1@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'user', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."users" VALUES (3, 'user2', 'user2@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'user', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."users" VALUES (4, 'manager', 'manager@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'manager', 'active', NULL, NULL, NULL);
INSERT INTO "manage_dev"."users" VALUES (8, 'xiaozhu1', 'xiaozhulzqwork1@2925.com', '$2a$10$nt1Fsp2rwehcpqvd79hgpuAJrv8KsNqV55NsavoIyVrhJneyUO6R2', 'user', 'active', '2025-10-20 14:27:29.616307+08', '2025-10-20 14:27:29.616307+08', '2025-10-20 14:27:45.467468+08');
INSERT INTO "manage_dev"."users" VALUES (10, 'xiaozhu', 'xiaozhulzqwork1@2925.com', '$2a$10$BGOO/MjN4ixzzfSBLjeVvOb/q5mSmrnoWYLB6p9.MzCSHGJ2rWaaW', 'user', 'active', '2025-10-20 14:32:45.566734+08', '2025-10-20 14:32:45.566734+08', NULL);
INSERT INTO "manage_dev"."users" VALUES (11, '测试新增用户1', 'xiaozhulzqwork5@2925.com', '$2a$10$CHTHg196qYa/qUTGALrTfO14cOXXTLOuvi2x155mwpTmGyijLUC3K', 'user', 'active', '2025-10-20 14:48:45.111729+08', '2025-10-20 14:48:52.662225+08', NULL);
INSERT INTO "manage_dev"."users" VALUES (12, '测试信的埃塞挨打', 'xiaozhulzqgame11@2925.com', '$2a$10$YxWwc6G4vIE2HvKOjAnz2ud1guRqPTRCYfeliVO0po8w6YsSszgbm', 'user', 'active', '2025-10-20 15:10:15.762983+08', '2025-10-20 15:10:26.242874+08', NULL);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "manage_dev"."audit_logs_id_seq"
OWNED BY "manage_dev"."audit_logs"."id";
SELECT setval('"manage_dev"."audit_logs_id_seq"', 1553, true);

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
