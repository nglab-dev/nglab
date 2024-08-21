-- +goose Up
-- +goose StatementBegin

-- sys_org 系统机构表
CREATE TABLE `sys_org` (
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, -- 自增主键
    `parent_id` INTEGER NOT NULL DEFAULT 0, -- 父级ID
    `code` TEXT NOT NULL, -- 机构编码
    `name` TEXT NOT NULL, -- 机构名称
    `remark` TEXT, -- 机构描述
    `sort` INTEGER NOT NULL DEFAULT 0, -- 排序
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `created_by` INTEGER NOT NULL DEFAULT 0, -- 创建人ID 
    `updated_by` INTEGER NOT NULL DEFAULT 0 -- 更新人ID
);

INSERT INTO `sys_org` (`code`, `name`, `remark`) VALUES ('test', '测试机构', '测试机构');

-- sys_config 系统配置表
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

-- sys_user 系统用户表
CREATE TABLE `sys_user` (
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `username` TEXT NOT NULL,
  `nickname` TEXT NOT NULL,
  `gender` INTEGER NOT NULL,
  `phone` TEXT,
  `password` TEXT NOT NULL,
  `email` TEXT UNIQUE NOT NULL,
  `avatar_url` TEXT,
  `enabled` INT NOT NULL, -- 0: disabled, 1: enabled
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `created_by` INTEGER DEFAULT 0,
  `updated_by` INTEGER DEFAULT 0
);
INSERT INTO `sys_user` (`ID`, `username`, `nickname`, `gender`, `phone`, `password`, `email`, `avatar_url`, `enabled`)
VALUES (1, 'admin', 'admin', 1, '13811111111', '$2a$10$05SaFGfrDvckPCV54xTNWezvpzL1JzJn8uwYAoXvuQR.Fe.0rqogy', 'admin@localhost.com', 'https://unpkg.com/outeres@0.0.5/demo/000.jpg', 1);

INSERT INTO `sys_user` (`ID`, `username`, `nickname`, `gender`, `phone`, `password`, `email`, `avatar_url`, `enabled`)
VALUES (2, 'user', 'user', 1, '13811112223', '$2a$10$DM5nXuZ.x.bExWcfqCbp0OAwRwFUwMmJ5cSZ0W5V2iMMV6Mr2LwOS', 'user@localhost.com', 'https://unpkg.com/outeres@0.0.5/demo/000.jpg', 1);

-- sys_role 系统角色表
CREATE TABLE `sys_role` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键ID
  `name` VARCHAR(255) NOT NULL, -- 角色名称
  `remark` VARCHAR(255) NOT NULL, -- 角色描述
  `sort` INTEGER NOT NULL DEFAULT 0, -- 排序
  `enabled` INTEGER DEFAULT 1, -- 是否启用
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 更新时间
  `created_by` INTEGER NOT NULL DEFAULT 0, -- 创建人ID
  `updated_by` INTEGER NOT NULL DEFAULT 0 -- 更新人ID
);

INSERT INTO `sys_role` (`id`, `name`, `remark`, `enabled`) VALUES (1, 'admin', '超级管理员', 1);
INSERT INTO `sys_role` (`id`, `name`, `remark`, `enabled`) VALUES (2, 'user', '普通用户', 1);

-- sys_user_role 用户角色关系表
CREATE TABLE `sys_user_role` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键ID
  `user_id` INTEGER NOT NULL, -- 用户ID
  `role_id` INTEGER NOT NULL, -- 角色ID
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 更新时间
  `created_by` INTEGER NOT NULL DEFAULT 0, -- 创建人ID
  `updated_by` INTEGER NOT NULL DEFAULT 0 -- 更新人ID
);

INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (1, 1, 1);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`) VALUES (2, 2, 2);

-- sys_menu 系统菜单表
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

-- sys_role_menu 角色菜单关系表
CREATE TABLE `sys_role_menu` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `role_id` INTEGER NOT NULL, -- 角色ID
    `menu_id` INTEGER NOT NULL, -- 菜单ID
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `created_by` INTEGER NOT NULL DEFAULT 0, -- 创建人
    `updated_by` INTEGER NOT NULL DEFAULT 0 -- 更新人
);

INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (1, 1, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (2, 1, 2);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (3, 1, 3);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (4, 1, 4);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (5, 2, 1);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (6, 2, 2);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (7, 2, 3);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`) VALUES (8, 2, 4);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `sys_config`;
DROP TABLE `sys_org`;
DROP TABLE `sys_user`;
DROP TABLE `sys_role`;
DROP TABLE `sys_user_role`;
DROP TABLE `sys_menu`;
DROP TABLE `sys_role_menu`;
-- +goose StatementEnd
