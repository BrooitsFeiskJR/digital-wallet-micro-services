# Wallet Service - Digital Wallet Microservices

The **Wallet Service** is a key component of the Digital Wallet Microservices project. It manages all wallet-related operations including balance inquiries, deposits, withdrawals, and maintaining a transaction history. This microservice is built in Go using modern frameworks and libraries to ensure performance, scalability, and security.

## Overview

The Wallet Service provides RESTful endpoints to:
- **Check Balance:** Retrieve the current wallet balance for a user.
- **Process Deposits & Withdrawals:** Handle deposit and withdrawal operations, ensuring accurate and secure updates to wallet balances.
- **Transaction History:** Log every wallet transaction for auditing and tracking purposes.

This service integrates with other components of the Digital Wallet ecosystem, such as the User Service and Transaction Service, to deliver a comprehensive financial solution.

## Technologies Used

- **Go:** The primary programming language.
- **Gin:** A high-performance HTTP web framework for building RESTful APIs.
- **PostgreSQL:** The relational database for storing wallet data.
- **SQLX:** An extension to Go’s standard `database/sql` package that simplifies interactions with PostgreSQL.
- **Docker:** Containerization for consistent deployment across environments.
- **Environment Variables:** Secure configuration management.

## Features

- **Balance Inquiry:** Quickly fetch the current wallet balance.
- **Deposit & Withdrawal Processing:** Safely update wallet balances with robust transactional integrity.
- **Transaction Logging:** Maintain a detailed log of all wallet operations.
- **Error Handling:** Comprehensive error management to ensure system reliability.
- **Testing:** Unit and integration tests to uphold code quality and reliability.

## Getting Started

### Prerequisites

- Go 1.18+ installed on your machine.
- A running PostgreSQL instance.
- Docker (optional, for containerization).

### Setup Instructions

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your_username/digital-wallet-micro-services.git
   cd digital-wallet-micro-services/wallet-service
2. **Configure Environment Variables**

Create a `.env` file in the wallet-service directory with content similar to:

```dotenv
PORT=8081
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_wallet_db
```
3. **Install Dependencies**

Download the necessary dependencies:

```bash
go mod tidy
```
Run the Service

4. **Start the service:**

```bash
go run main.go
```
## Running with Docker
To run the service in a containerized environment:

1. **Build the Docker Image**

```bash
docker build -t wallet-service .
```
2. **Run the Container**

```bash
docker run -p 8081:8081 --env-file .env wallet-service
```
## API Endpoints
Here is a summary of the main endpoints provided by the Wallet Service:

- __GET /wallet__

    Retrieve the current wallet balance for a user.

- __POST /wallet/deposit__

    Process a deposit to increase the wallet balance.

- __POST /wallet/withdraw__

    Process a withdrawal to decrease the wallet balance.

For detailed API documentation and request/response examples, please refer to the API specification or additional documentation.

## Contributing
Contributions to enhance functionality, security, and performance are welcome! To contribute:

- Fork the repository.
- Create a feature branch (`git checkout -b feature/your-feature`).
- Commit your changes (`git commit -m 'Add new feature'`).
- Push to the branch (`git push origin feature/your-feature`).
- Open a pull request detailing your changes.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Contact
For questions, suggestions, or issues, please reach out at tontech.dev@outlook.com.