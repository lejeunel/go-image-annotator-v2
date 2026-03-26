-- +goose Up

CREATE TABLE IF NOT EXISTS labels (
    id varchar(36),
    name varchar(30),
    description text,
    PRIMARY KEY (id)
);
-- +goose Down

DROP TABLE labels;
