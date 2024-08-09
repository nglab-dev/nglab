-- +goose Up
-- +goose StatementBegin
CREATE TABLE `sys_role` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键ID
  `name` VARCHAR(255) NOT NULL, -- 角色名称
  `remark` VARCHAR(255) NOT NULL, -- 角色描述
  `enabled` INTEGER DEFAULT 1, -- 是否启用
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
  `create_by` INTEGER NOT NULL DEFAULT 0, -- 创建人ID
  `update_by` INTEGER NOT NULL DEFAULT 0 -- 更新人ID
);

INSERT INTO `sys_role` (`id`, `name`, `remark`, `enabled`)
VALUES (1, 'admin', '超级管理员', 1);

INSERT INTO `sys_role` (`id`, `name`, `remark`, `enabled`, `create_by`, `update_by`)
VALUES (2, 'user', '普通用户', 1, 1, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `sys_role`;
-- +goose StatementEnd
