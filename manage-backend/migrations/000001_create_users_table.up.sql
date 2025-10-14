CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- 初始化用户数据
-- 注意：以下密码均使用 bcrypt 加密
-- admin - 密码: admin123
-- user1 - 密码: user123
-- user2 - 密码: user123
-- manager - 密码: manager123

INSERT INTO users (username, email, password, role, status) VALUES
('admin', 'admin@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'admin', 'active'),
('user1', 'user1@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'user', 'active'),
('user2', 'user2@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'user', 'active'),
('manager', 'manager@example.com', '$2a$10$6DOUbQEThHbOWY99oTbn2.OQ843xrjduFbgz3RgI7TaKG07baMTq6', 'manager', 'active')
ON CONFLICT (username) DO NOTHING;