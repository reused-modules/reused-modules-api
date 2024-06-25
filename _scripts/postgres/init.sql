GRANT ALL PRIVILEGES ON DATABASE product TO dev;

CREATE TABLE transfers (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    account_origin_id VARCHAR NOT NULL,
    account_destination_id VARCHAR NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE accounts (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR NOT NULL,
    cpf VARCHAR UNIQUE NOT NULL,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

INSERT INTO accounts (id, name, cpf, balance, created_at)
VALUES (1, 'abc', 'cpf', 1000, now())

------------------------------------------------------------

CREATE TABLE products (
    id INT PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    sku VARCHAR(50) NOT NULL,
    price BIGINT NOT NULL,
    description VARCHAR,
    category_id INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id INT PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    parent_id INT DEFAULT 0,
    root_parent_id INT DEFAULT 0,
    sub_level INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON COLUMN categories.root_parent_id IS 'ID of category has parent_id=0 & root_parent_id=0 & sub_level=1';
COMMENT ON COLUMN categories.sub_level IS 'Layer of relations. Highest is 1 then 2, 3, ...';