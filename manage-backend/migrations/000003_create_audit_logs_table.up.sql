-- ==========================================
-- 审计日志表
-- ==========================================

CREATE TABLE IF NOT EXISTS audit_logs (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT,
    username VARCHAR(50),
    action VARCHAR(255) NOT NULL,
    resource VARCHAR(100),
    resource_id VARCHAR(100),
    method VARCHAR(10),
    path VARCHAR(500),
    ip VARCHAR(50),
    user_agent TEXT,
    status INT,
    error_msg TEXT,
    request_body TEXT,
    duration BIGINT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

-- 创建索引以提高查询性能
CREATE INDEX IF NOT EXISTS idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX IF NOT EXISTS idx_audit_logs_action ON audit_logs(action);
CREATE INDEX IF NOT EXISTS idx_audit_logs_resource ON audit_logs(resource);
CREATE INDEX IF NOT EXISTS idx_audit_logs_created_at ON audit_logs(created_at);
CREATE INDEX IF NOT EXISTS idx_audit_logs_status ON audit_logs(status);
CREATE INDEX IF NOT EXISTS idx_audit_logs_deleted_at ON audit_logs(deleted_at);

-- 添加外键约束（可选）
-- ALTER TABLE audit_logs ADD CONSTRAINT fk_audit_logs_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
