# Wallet

## Database Schema 

### Wallets Table

The `Wallets` table contains information about each wallet that a user has created. :

| Column Name | Data Type | Description |
|-------------|----------|-------------|
| wallet_id   | VARCHAR(36) | Primary key for the table, with the format of a UUID (e.g. `ea0212d3-abd6-406f-8c67-868e814a2436`). |
| user_id     | VARCHAR(36) | Foreign key referencing the `id` column of the existing `Users` table. |
| token       | VARCHAR(50) | Unique token that can be used to authenticate requests to the wallet API. |
| balance     | BIGINT   | The current balance of the wallet, in the smallest unit of currency (e.g. cents). |
| status     | VARCHAR(10)  | Indicates whether the wallet is currently enabled or disabled. |
| created_at  | TIMESTAMP | The date and time that the wallet was created. |
| enabled_at  | TIMESTAMP | The date and time that the wallet was enabled. |
| disabled_at  | TIMESTAMP | The date and time that the wallet was disabled. |

### Transactions Table

The `Transactions` table contains information about each transaction that occurs within a wallet.

| Column Name   | Data Type | Description |
|---------------|-----------|-------------|
| id            | VARCHAR(36) | Primary key for the table, with the format of a UUID (e.g. `a73d4a37-4f4b-4d2c-85c9-7dcf6cddc51d`). |
| wallet_id     | VARCHAR(36) | Foreign key referencing the `wallet_id` column of the `Wallets` table. |
| type          | VARCHAR(10) | Indicates whether the transaction is a deposit or withdrawal. |
| reference_id  | VARCHAR(36) | A unique identifier for the transaction, with the format of UUID (e.g. `ea0212d3-abd6-406f-8c67-868e814a2436`). |
| amount        | BIGINT    | The amount of the transaction, in the smallest unit of currency (e.g. cents). |
| balance_before | BIGINT    | The balance of the wallet before the transaction was processed. |
| balance_after | BIGINT    | The balance of the wallet after the transaction was processed. |
| deposited_by            | VARCHAR(36) | User that doing deposit transaction, with the format of a UUID (e.g. `a73d4a37-4f4b-4d2c-85c9-7dcf6cddc51d`). |
| deposited_at    | TIMESTAMP | The date and time that the deposit transaction was processed. |
| withdrawn_by            | VARCHAR(36) | User that doing withdrawal transaction, with the format of a UUID (e.g. `a73d4a37-4f4b-4d2c-85c9-7dcf6cddc51d`). |
| withdrawn_at    | TIMESTAMP | The date and time that the withdrawal transaction was processed. |

## Code Structure/Design
```
- config/
- internal/
    - entity/
    - handler/
    - repo/
- pkg/
- schema/
- main.go
```

### config/
This directory contains configuration files used by the application.

### internal/
This directory contains internal packages used by the application.

### entity/
This package defines the data models used by the application.

### handler/
This package contains the business logic of the application.

### repo/
This package handles the interaction with the database or other data stores.

### pkg/
This directory contains public packages that can be used by other applications.

### schema/
This directory contains database schema files.

### main.go
This file contains the main entry point of the application.


