# Transaction Service - Digital Wallet Microservices

The **Transaction Service** is a core component of the Digital Wallet Microservices project. It is responsible for recording, managing, and auditing all financial transactions within the system, such as deposits, withdrawals, and transfers. This service ensures a robust audit trail and supports reporting and analytics, providing transparency and accountability for wallet operations.

## Overview

The Transaction Service handles:
- **Recording Transactions:** Capturing details of every wallet operation (deposits, withdrawals, transfers) in a dedicated collection.
- **Auditing and Reporting:** Maintaining a complete log for auditing purposes and enabling data analytics on transaction flows.
- **Decoupled Processing:** Operating independently to ensure that wallet balance updates and transaction logging are handled separately for better scalability and reliability.

The service is designed to work seamlessly with other microservices in the ecosystem, such as the Wallet Service and User Service, by receiving events or direct API calls that trigger transaction logging.

## Technologies Used

- **Go:** The primary programming language for building the service.
- **Gin:** A high-performance HTTP web framework for exposing RESTful endpoints.
- **MongoDB:** The NoSQL database used for storing transaction records.
- **Docker:** Containerization to ensure consistent deployment across environments.
- **Messaging Systems (Optional):** Integration with systems like Kafka, RabbitMQ, or SQS for asynchronous event processing.
- **Environment Variables:** For secure and flexible configuration management.

## Features

- **Transaction Logging:** Securely record every financial operation with details including wallet IDs, transaction amount, type, and timestamp.
- **RESTful API Endpoints:** Expose endpoints for creating and retrieving transaction records.
- **Event-Driven Updates:** Optionally subscribe to messaging queues to process transactions asynchronously.
- **Robust Error Handling:** Ensure data integrity with proper error management and logging.
- **Audit Trail:** Provide a complete history of operations to support auditing and reporting needs.

## Getting Started

### Prerequisites

- Go 1.18+ installed on your machine.
- A running MongoDB instance.
- Docker (optional, for containerization).
- (Optional) A messaging system (e.g., Kafka, RabbitMQ) if integrating asynchronous events.

### Setup Instructions

1. **Clone the Repository**

   ```bash
   git clone https://github.com/BrooitsFeiskJR/digital-wallet-micro-services.git
   cd digital-wallet-micro-services/transaction-service
   ```
2. **Configure Environment Variables**
    ```dotenv
    PORT=8082
    MONGO_URI=mongodb://localhost:27017
    MONGO_DB=transaction_db
    ```
3. **Install Dependencies**

    Download the necessary dependencies:
    ```bash
    go mod tidy
    ```
## Running with Docker
To run the service in a containerized environment:
1. **Build the docker image**
    ```bash
    docker compose up --build
    ```
## API Endpoints
Below is a summary of the main endpoints provided by the Transaction Service:

- __POST /transactions__

    Create a new transaction record.

    ___(Expected payload includes wallet IDs, transaction amount, type, etc.)___

- __GET /transactions/{id}__

    Retrieve details of a specific transaction by its ID.

- __GET /transactions?wallet_id={walletID}__

    Retrieve all transactions related to a specific wallet.

For detailed API documentation and request/response examples, please refer to the API specification or Swagger/OpenAPI documentation.

## Contributing
Contributions to improve functionality, security, and performance are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to your branch (`git push origin feature/your-feature`).
5. Open a pull request describing your changes.

Please ensure that your contributions adhere to the project's coding standards and include appropriate tests.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contact
For questions, suggestions, or issues, please reach out at tontech.dev@outlook.com
