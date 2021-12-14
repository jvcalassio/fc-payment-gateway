CREATE TABLE transactions(
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    account_id VARCHAR(255) NOT NULL,
    amount DOUBLE NOT NULL,
    status VARCHAR(255) NOT NULL,
    error_message VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);