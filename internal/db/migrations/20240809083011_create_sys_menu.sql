-- +goose Up
-- +goose StatementBegin
CREATE TABLE `sys_menu` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `parent_id` INTEGER DEFAULT 0, -- 父菜单ID
    `name` TEXT NOT NULL, -- 菜单名称
    `path` TEXT, -- 路由地址
    `type` INTEGER DEFAULT 0, -- 菜单类型 0:目录 1:菜单 2:按钮
    `icon` TEXT, -- 菜单图标
    `sort` INTEGER DEFAULT 0, -- 排序
    `enabled` INTEGER DEFAULT 1, -- 状态 0:禁用 1:启用
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `created_by` INTEGER NOT NULL DEFAULT 0, -- 创建人
    `updated_by` INTEGER NOT NULL DEFAULT 0 -- 更新人
);

INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `type`, `icon`, `sort`, `enabled`) VALUES (1, 0, '系统管理', '', 0, 'fa fa-cog', 1, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `type`, `icon`, `sort`, `enabled`) VALUES (2, 1, '用户管理', '/sys/user', 1, 'fa fa-user', 1, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `type`, `icon`, `sort`, `enabled`) VALUES (3, 1, '角色管理', '/sys/role', 1, 'fa fa-users', 2, 1);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `type`, `icon`, `sort`, `enabled`) VALUES (4, 1, '菜单管理', '/sys/menu', 1, 'fa fa-th-list', 3, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `sys_menu`;
-- +goose StatementEnd
