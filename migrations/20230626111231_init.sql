-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name varchar(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE INDEX idx_user_login ON users USING btree (login);

CREATE TABLE IF NOT EXISTS products (
    barcode VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    cost INTEGER NOT NULL,
    user_id BIGINT NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS product_price (
    name VARCHAR(255) PRIMARY KEY,
    product_barcode VARCHAR(255) NOT NULL,
    CONSTRAINT fk_product_barcode FOREIGN KEY(product_barcode) REFERENCES products(barcode) ON UPDATE CASCADE
);
CREATE UNIQUE INDEX idx_product_price ON product_price(name, product_barcode);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users, products, product_price CASCADE;
-- +goose StatementEnd
