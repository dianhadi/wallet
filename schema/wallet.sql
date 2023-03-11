CREATE TABLE wallets (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    token VARCHAR(42) UNIQUE,
    balance BIGINT,
    status VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    enabled_at TIMESTAMP,
    disabled_at TIMESTAMP
);

CREATE TABLE transactions (
    id VARCHAR(36) PRIMARY KEY,
    wallet_id VARCHAR(36) NOT NULL,
    type VARCHAR(10),
    reference_id VARCHAR(36) UNIQUE,
    amount BIGINT,
    balance_before BIGINT,
    balance_after BIGINT,
    deposited_by VARCHAR(36),
    deposited_at TIMESTAMP,
    withdrawn_by VARCHAR(36),
    withdrawn_at TIMESTAMP,
    FOREIGN KEY (wallet_id) REFERENCES Wallets(id)
);