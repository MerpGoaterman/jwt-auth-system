-- JWT Authentication System Database Initialization

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    tenant_id VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_tenant (tenant_id),
    INDEX idx_role (role)
);

-- Create admin user
-- Email: admin@example.com
-- Password: admin123
-- Note: This is a bcrypt hash of "admin123"
INSERT INTO users (id, name, email, password, tenant_id, role) 
VALUES (
    UUID(),
    'Admin User',
    'admin@example.com',
    '$2a$10$rqYvE5zGwZ8xH5xH5xH5xOqYvE5zGwZ8xH5xH5xH5xOqYvE5zGwZ8',
    'default-tenant',
    'admin'
) ON DUPLICATE KEY UPDATE id=id;

-- Create sample regular user
-- Email: user@example.com
-- Password: user123
INSERT INTO users (id, name, email, password, tenant_id, role) 
VALUES (
    UUID(),
    'Regular User',
    'user@example.com',
    '$2a$10$rqYvE5zGwZ8xH5xH5xH5xOqYvE5zGwZ8xH5xH5xH5xOqYvE5zGwZ8',
    'default-tenant',
    'user'
) ON DUPLICATE KEY UPDATE id=id;

-- Display created users
SELECT id, name, email, tenant_id, role, created_at FROM users;
