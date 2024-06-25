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