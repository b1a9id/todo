-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE task MODIFY id INT UNSIGNED AUTO_INCREMENT NOT NULL;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE task MODIFY id INT UNSIGNED NOT NULL;
