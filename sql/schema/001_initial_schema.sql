-- +goose Up
CREATE TABLE categories
(
    id         UUID PRIMARY KEY,
    name       VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products
(
    id          UUID PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    sku         VARCHAR(50)  NOT NULL UNIQUE,
    price       REAL         NOT NULL,
    stock       INT NOT NULl,
    description TEXT,
    category_id UUID NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE INDEX idx_products_created_at_id ON products (created_at, id);

CREATE TABLE product_images
(
    id         UUID PRIMARY KEY,
    product_id UUID      NOT NULL REFERENCES products (id) ON DELETE CASCADE,
    image_type VARCHAR(50),
    image_url  VARCHAR(250),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE product_images;
DROP TABLE products;
DROP TABLE categories;