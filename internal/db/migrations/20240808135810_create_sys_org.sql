-- +goose Up
-- +goose StatementBegin
CREATE TABLE `sys_org` (
    `id` INTEGER NOT NULL,
    `parent_id` INTEGER NOT NULL DEFAULT 0,
    `code` TEXT NOT NULL,
    `name` TEXT NOT NULL,
    `description` TEXT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

INSERT INTO `sys_org` (`code`, `name`, `description`) VALUES ('root', 'Root', 'Root organization');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE `sys_org`;

-- +goose StatementEnd