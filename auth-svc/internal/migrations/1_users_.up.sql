-- +goose Up

-- Users table
CREATE TABLE IF NOT EXISTS users (
    user_id UUID DEFAULT uuid_generate_v4(),
    user_name VARCHAR(50),
    full_name VARCHAR(350) NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    address VARCHAR(530) NOT NULL DEFAULT '',
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    refresh_token VARCHAR(255), -- Added column for refresh token
    PRIMARY KEY (user_id)
);

-- Roles table
CREATE TABLE IF NOT EXISTS roles (
    role_id UUID DEFAULT uuid_generate_v4(),
    role_name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (role_id)
);

-- UserRoles table (many-to-many relationship between Users and Roles)
CREATE TABLE IF NOT EXISTS user_roles (
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    shop_id UUID NOT NULL,  -- Assuming shop-specific roles
    PRIMARY KEY (user_id, role_id, shop_id),
    CONSTRAINT fk_user_roles_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
    CONSTRAINT fk_user_roles_role_id FOREIGN KEY (role_id) REFERENCES roles(role_id)
);

-- Add unique constraint for user_name only if it's not NULL
CREATE UNIQUE INDEX idx_unique_user_name ON users (user_name) WHERE user_name IS NOT NULL;

-- +goose Down

DROP TABLE user_roles;
DROP TABLE roles;
DROP TABLE users;
