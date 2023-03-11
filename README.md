# Wallet
This is simple wallet REST APIs that contain these endpoints:
- Init Account
- Enable Wallet
- View Balance
- View Transactions
- Deposit Money
- Withdraw Money
- Disable Wallet

## How to Use

### Dependencies
This repo is developed with **Docker**, so you need to install docker first before continue.

These are dependencies that will be run in docker:
1. Golang for programming language. *You may need to install Go if you want to make changes of the code.*
2. Postgres for database.

### Build & Run

Use `docker-compose up` to build and run this service. *You also may use `docker-compose build` and `docker-compose down` for development purpose*
URL will be accessable via http://localhost

### Documentation 
Complete documentation can be found [here](./DOCUMENTATION.md) 
