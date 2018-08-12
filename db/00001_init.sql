-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS task (
    id INT UNSIGNED NOT NULL,
    created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE task;
