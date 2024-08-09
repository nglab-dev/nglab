-- +goose Up
-- +goose StatementBegin
CREATE TABLE `sys_config` (
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, -- 自增主键
    `config_key` TEXT NOT NULL, -- 配置键
    `config_value` TEXT NOT NULL, -- 配置值
    `description` TEXT, -- 配置描述
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `created_by` TEXT NOT NULL DEFAULT 0, -- 创建人
    `updated_by` TEXT NOT NULL DEFAULT 0-- 更新人
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sys_config;
-- +goose StatementEnd
