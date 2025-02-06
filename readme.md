# Transaction Service

This project is a RESTful web service for storing and retrieving transaction information using a relational database such as PostgreSQL or MySQL.

## Project Structure

```
crud-transaction-main/
├── cmd/server                # Main application entry point
├── config                    # Application configuration files
├── db                        # Database connection and initialization
├── handlers                  # API handlers for creating and retrieving transactions
├── models                    # Data models
├── router                    # API routing and logging setup
├── makefile                  # Build and task automation
└── application.yml.sample    # Sample configuration file
```

## Prerequisites

- Go 1.16 or later
- PostgreSQL or MySQL for database storage

## Build and Run Instructions

### Setting Up the Environment

```bash
make setup
```

### Linting the Code

```bash
make lint
```

### Building the Application

```bash
make build
```

The executable will be created at `out/transaction-service`.

### Running the Application

```bash
make run
```

### Running Tests

```bash
make test
```

### Generating Test Coverage Report

```bash
make test.cover
```

### Generating Cobertura Coverage Report

```bash
make test.report
```

The coverage report will be generated in `coverage/coverage.html` and `coverage/coverage.xml`.

## Endpoints

### 1. Create Transaction

- **PUT** `/transactionservice/transaction/{transaction_id}`
- Request Body:

  ```json
  {
    "amount": 10000,
    "type": "shopping",
    "parent_id": 10
  }
  ```

### 2. Get Transaction

- **GET** `/transactionservice/transaction/{transaction_id}`

### 3. Get Transaction

- **GET** `/transactionservice/types/{type}`

### 4. Get Transaction

- **GET** `/transactionservice/sum/{transaction_id}`

## Notes

- Ensure that the database connection details in `application.yml` are correctly configured before running the service.
- Use `make help` to see all available commands.
