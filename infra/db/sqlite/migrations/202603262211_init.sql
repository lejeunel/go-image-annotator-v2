-- +goose Up

CREATE TABLE IF NOT EXISTS labels (
    id varchar(36),
    name varchar(30),
    description text,
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX idx_labels_name ON labels(name);

CREATE TABLE IF NOT EXISTS collections (
    id varchar(36),
    name varchar(30),
    description text,
    PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_collections_name ON collections(name);
-- +goose Down

DROP TABLE labels;
