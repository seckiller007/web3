-- employees 表
CREATE TABLE IF NOT EXISTS employees (
                                         id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                                         name       VARCHAR(64)   NOT NULL,
    department VARCHAR(64)   NOT NULL,
    salary     DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    created_at DATETIME      NULL,
    updated_at DATETIME      NULL
    );

-- 一点测试数据
INSERT INTO employees (name, department, salary, created_at, updated_at) VALUES
                                                                             ('张三', '技术部', 25000.00, NOW(), NOW()),
                                                                             ('李四', '技术部', 30000.00, NOW(), NOW()),
                                                                             ('王五', '市场部', 22000.00, NOW(), NOW())
    ON DUPLICATE KEY UPDATE salary=VALUES(salary);

CREATE TABLE IF NOT EXISTS books (
                                     id     BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                                     title  VARCHAR(128)  NOT NULL,
    author VARCHAR(64)   NOT NULL,
    price  DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    created_at DATETIME NULL,
    updated_at DATETIME NULL
    );

-- 可选：一些测试数据
INSERT INTO books (title, author, price, created_at, updated_at) VALUES
                                                                     ('Go 语言实战', 'John', 88.00, NOW(), NOW()),
                                                                     ('数据库系统概念', 'Abraham', 129.00, NOW(), NOW()),
                                                                     ('算法导论', 'CLRS', 99.50, NOW(), NOW()),
                                                                     ('你不知道的JS', 'Kyle', 45.00, NOW(), NOW())
    ON DUPLICATE KEY UPDATE price=VALUES(price);
