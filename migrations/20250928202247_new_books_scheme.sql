-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS books;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS books;
-- +goose StatementEnd
