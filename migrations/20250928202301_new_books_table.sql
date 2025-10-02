-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books.genres(
    genre_id smallint PRIMARY KEY,
    name text NOT NULL UNIQUE,
    description text
);

CREATE TABLE IF NOT EXISTS books.books(
    book_id uuid PRIMARY KEY,
    author_id uuid NOT NULL,
    title text NOT NULL,
    description text,
    date_released date NOT NULL,
    genre_id smallint NOT NULL REFERENCES books.genres(genre_id) ON DELETE CASCADE,
    created_at timestamp without time zone default current_timestamp,
    rating numeric(5,2) NOT NULL
);

CREATE INDEX books_date_released_idx
    ON books.books USING btree(date_released);

CREATE INDEX books_genre_id_idx
    ON books.books (genre_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS books_genre_id_idx;
DROP INDEX IF EXISTS books_date_released_idx;

DROP TABLE IF EXISTS books.books;
DROP TABLE IF EXISTS books.genres;
-- +goose StatementEnd
