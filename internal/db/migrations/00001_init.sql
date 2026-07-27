-- +goose Up
CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INT NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
    product VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

-- +goose Down
DROP TABLE cart_items;
DROP TABLE carts;
