-- +goose Up

-- Insert data into categories table
INSERT INTO categories (id, name, created_at, updated_at)
VALUES (gen_random_uuid(), 'Electronics', NOW(), NOW()),
       (gen_random_uuid(), 'Books', NOW(), NOW()),
       (gen_random_uuid(), 'Clothing', NOW(), NOW());

-- Insert data into products table
INSERT INTO products (id, name, sku, price, stock, description, category_id, created_at, updated_at)
VALUES (gen_random_uuid(), 'Laptop', 'SKU123', 999.99, 50, 'A high-performance laptop for all your needs.',
        (SELECT id FROM categories WHERE name = 'Electronics'), NOW(), NOW()),
       (gen_random_uuid(), 'Smartphone', 'SKU124', 699.99, 150, 'A latest model smartphone.',
        (SELECT id FROM categories WHERE name = 'Electronics'), NOW(), NOW()),
       (gen_random_uuid(), 'Fiction Book', 'SKU125', 19.99, 100, 'A best-selling fiction book.',
        (SELECT id FROM categories WHERE name = 'Books'), NOW(), NOW()),
       (gen_random_uuid(), 'T-Shirt', 'SKU126', 15.99, 200, 'A comfortable cotton t-shirt.',
        (SELECT id FROM categories WHERE name = 'Clothing'), NOW(), NOW());

-- Insert data into product_images table
INSERT INTO product_images (id, product_id, image_type, image_url, created_at, updated_at)
VALUES (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU123'), 'thumbnail',
        'http://localhost:8080/images/annie-spratt-fbAnIjhrOL4-unsplash.jpg', NOW(), NOW()),
       (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU124'), 'thumbnail',
        'http://localhost:8080/images/austin-chan-ukzHlkoz1IE-unsplash.jpg', NOW(), NOW()),
       (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU125'), 'thumbnail',
        'http://localhost:8080/images/carson-arias-7ZO3R1wOdmI-unsplash.jpg', NOW(), NOW()),
       (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU126'), 'thumbnail',
        'http://localhost:8080/images/chris-lawton-5lHz5WhosQE-unsplash.jpg', NOW(), NOW());

-- Add additional images (non-thumbnail)
INSERT INTO product_images (id, product_id, image_type, image_url, created_at, updated_at)
VALUES (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU123'), 'main',
        'http://localhost:8080/images/david-kovalenko-G85VuTpw6jg-unsplash.jpg', NOW(), NOW()),
       (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU124'), 'main',
        'http://localhost:8080/images/ian-dooley-hpTH5b6mo2s-unsplash.jpg', NOW(), NOW()),
       (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU125'), 'main',
        'http://localhost:8080/images/ian-dooley-TLD6iCOlyb0-unsplash.jpg', NOW(), NOW()),
       (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU126'), 'main',
        'http://localhost:8080/images/katie-moum-iRMUDX0kyOc-unsplash.jpg', NOW(), NOW()),
       (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU123'), 'main',
        'http://localhost:8080/images/laura-vinck-Hyu76loQLdk-unsplash.jpg', NOW(), NOW()),
       (gen_random_uuid(), (SELECT id FROM products WHERE sku = 'SKU124'), 'main',
        'http://localhost:8080/images/mike-dorner-sf_1ZDA1YFw-unsplash.jpg', NOW(), NOW());

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

-- Delete data from product_images table
DELETE FROM product_images;

-- Delete data from products table
DELETE FROM products;

-- Delete data from categories table
DELETE FROM categories;
