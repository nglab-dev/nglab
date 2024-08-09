-- +goose Up
-- +goose StatementBegin
CREATE TABLE `sys_config` (
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, -- 自增主键
    `config_key` TEXT NOT NULL, -- 配置键
    `config_value` TEXT NOT NULL, -- 配置值
    `remark` TEXT, -- 配置描述
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `created_by` TEXT NOT NULL DEFAULT 0, -- 创建人
    `updated_by` TEXT NOT NULL DEFAULT 0-- 更新人
);

INSERT INTO `sys_config` (`config_key`, `config_value`, `remark`, `created_by`, `updated_by`)
VALUES ('app_name', 'nglab', '应用名称', '0', '0');

INSERT INTO `sys_config` (`config_key`, `config_value`, `remark`, `created_by`, `updated_by`)
VALUES ('app_version', 'v1.0.0', '应用版本', '0', '0');

INSERT INTO `sys_config` (`config_key`, `config_value`, `remark`, `created_by`, `updated_by`)
VALUES ('app_logo', 'https://nglab.oss-cn-beijing.aliyuncs.com/nglab_logo.png', '应用Logo', '0', '0');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sys_config;
-- +goose StatementEnd
