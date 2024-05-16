-- +goose Up
CREATE TABLE categories
(
    id         UUID PRIMARY KEY,
    name       VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE products
(
    id          UUID PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    sku         VARCHAR(50)  NOT NULL UNIQUE,
    price       REAL         NOT NULL,
    stock       REAL,
    description TEXT,
    category_id UUID NOT NULL,
    created_at  TIMESTAMP    NOT NULL,
    updated_at  TIMESTAMP    NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE TABLE product_images
(
    id         UUID PRIMARY KEY,
    product_id UUID      NOT NULL REFERENCES products (id) ON DELETE CASCADE,
    image_type VARCHAR(50),
    image_url  VARCHAR(250),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE product_images;
DROP TABLE products;
DROP TABLE categories;