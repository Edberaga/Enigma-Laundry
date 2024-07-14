# Enigma Laundry Management System

This is a simple Laundry Management System written in Go, utilizing PostgreSQL as the database. The application allows for the management of customers, services, and transactions.

## Features

- View, insert, update, and delete customers
- View, insert, update, and delete services
- View and insert transactions

## Prerequisites

- Go 1.22 or higher
- PostgreSQL

## Setup

### Database Setup

1. Install PostgreSQL and create a new database:
   ```sh
   sudo apt-get install postgresql postgresql-contrib
   sudo -u postgres psql
   ```
   ```sql
   CREATE DATABASE enigma_laundry_challenge;
   ```

2. Create the necessary tables in the enigma_laundry_challenge database:
```sql
CREATE TABLE customer (
    nama_customer VARCHAR(100) NOT NULL,
    no_HP VARCHAR(15) NOT NULL,
    total_fee DECIMAL DEFAULT 0,
    PRIMARY KEY (nama_customer)
);

CREATE TABLE service (
    nama VARCHAR(100) NOT NULL,
    satuan VARCHAR(10) NOT NULL,
    harga DECIMAL NOT NULL,
    PRIMARY KEY (nama)
);

CREATE TABLE tx_laundry (
    nomor INT NOT NULL,
    tanggal_masuk DATE NOT NULL,
    tanggal_selesai DATE NOT NULL,
    pelayanan VARCHAR(100) NOT NULL,
    jumlah INT NOT NULL,
    satuan VARCHAR(10) NOT NULL,
    harga DECIMAL NOT NULL,
    total DECIMAL NOT NULL,
    penerima VARCHAR(100) NOT NULL,
    PRIMARY KEY (nomor),
    FOREIGN KEY (pelayanan) REFERENCES service(nama),
    FOREIGN KEY (penerima) REFERENCES customer(nama_customer)
);
```

## Go Setup
1. Clone the repository:

```sh
git clone https://github.com/yourusername/enigma-laundry.git
cd enigma-laundry
```

2. Initialize the Go module:
```sh
go mod tidy
Set up the PostgreSQL connection details in main.go:
```

3. Set up your PostgreSQL connection details in main.go:
```go
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "yourpassword"
    dbname   = "enigma_laundry_challenge"
)
```

## Running the Application

1. Start the application:

```sh
go run main.go
```
2. Follow the on-screen menu to manage customers, services, and transactions.

## Code Structure
- main.go: Contains the main function and menu logic
- entity/: Contains the data structures for Customer, Service, and Transaction
- function/: Contains the functions for database operations

## Example Usage
1. To view customers, select option 1.
2. To insert a new customer, select option 2 and follow the prompts.
3. To update an existing customer, select option 3 and follow the prompts.
4. To delete a customer, select option 4 and follow the prompts.
5. To view services, select option 5.
6. To insert a new service, select option 6 and follow the prompts.
7. To update an existing service, select option 7 and follow the prompts.
8. To delete a service, select option 8 and follow the prompts.
9. To view transactions, select option 9.
10. To insert a new transaction, select option 10 and follow the prompts.

## Using the Application
User may create customer and service data with the application or may initialize the dummy data that I have write at SQL queries/DML.sql file.

Happy Coding!