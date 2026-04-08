-- +goose Up
ALTER TABLE collections
ADD COLUMN created_at DATETIME;

-- +goose Down
CREATE TABLE collections_new (
    id varchar(36),
    name varchar(30) NOT NULL UNIQUE,
    description TEXT,
    PRIMARY KEY (id)
);

INSERT INTO collections_new (id, name, description)
SELECT id, name, description
FROM collections;

DROP TABLE collections;

ALTER TABLE collections_new RENAME TO collections;
