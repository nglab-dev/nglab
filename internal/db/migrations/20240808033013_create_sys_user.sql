-- +goose Up
-- +goose StatementBegin

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

INSERT INTO `sys_user` (`username`, `nickname`, `gender`, `phone`, `password`, `email`, `avatar_url`, `enabled`)
VALUES ('admin', 'admin', 1, '13811111111', '$2a$10$05SaFGfrDvckPCV54xTNWezvpzL1JzJn8uwYAoXvuQR.Fe.0rqogy', 'admin@localhost.com', 'https://unpkg.com/outeres@0.0.5/demo/000.jpg', 1);

INSERT INTO `sys_user` (`username`, `nickname`, `gender`, `phone`, `password`, `email`, `avatar_url`, `enabled`)
VALUES ('user', 'user', 1, '13811112223', '$2a$10$DM5nXuZ.x.bExWcfqCbp0OAwRwFUwMmJ5cSZ0W5V2iMMV6Mr2LwOS', 'user@localhost.com', 'https://unpkg.com/outeres@0.0.5/demo/000.jpg', 1);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sys_user;
-- +goose StatementEnd
