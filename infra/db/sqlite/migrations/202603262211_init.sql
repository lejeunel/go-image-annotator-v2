-- +goose Up

CREATE TABLE IF NOT EXISTS labels (
    id varchar(36),
    name varchar(30) not null unique,
    description text,
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX idx_labels_name ON labels(name);

CREATE TABLE IF NOT EXISTS collections (
    id varchar(36),
    name varchar(30) not null unique,
    description text,
    PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_collections_name ON collections(name);

CREATE TABLE IF NOT EXISTS images_collections (
  image_id varchar(36),
  collection_id varchar(36) REFERENCES collections(id),
  PRIMARY KEY (image_id, collection_id)
);

-- +goose Down

DROP TABLE labels;
DROP TABLE collections;
DROP TABLE images;
