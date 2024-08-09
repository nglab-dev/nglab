-- +goose Up
-- +goose StatementBegin
CREATE TABLE `base_dict` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT, -- 自增主键
    `type` INTEGER NOT NULL, -- 字典类型
    `code` TEXT NOT NULL, -- 字典编码
    `name` TEXT NOT NULL, -- 字典名称
    `remark` TEXT, -- 备注
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `created_by` INTEGER DEFAULT 0, -- 创建人
    `updated_by` INTEGER DEFAULT 0 -- 更新人
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `base_dict`;
-- +goose StatementEnd
