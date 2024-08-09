-- +goose Up
-- +goose StatementBegin
CREATE TABLE `sys_org` (
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, -- 自增主键
    `parent_id` INTEGER NOT NULL DEFAULT 0, -- 父级ID
    `code` TEXT NOT NULL, -- 机构编码
    `name` TEXT NOT NULL, -- 机构名称
    `description` TEXT, -- 机构描述
    `sort` INTEGER NOT NULL DEFAULT 0, -- 排序
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `created_by` INTEGER DEFAULT 0, -- 创建人ID 
    `updated_by` INTEGER DEFAULT 0 -- 更新人ID
);

INSERT INTO `sys_org` (`code`, `name`, `description`) VALUES ('root', 'Root', 'Root organization');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE `sys_org`;
-- +goose StatementEnd