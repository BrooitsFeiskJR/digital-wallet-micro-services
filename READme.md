# Digital Wallet Microservices

The **Digital Wallet Microservices** project is a distributed system designed as a proof-of-concept for building scalable, secure, and maintainable financial applications using a microservices architecture. This project serves as a portfolio piece to demonstrate proficiency in modern backend technologies and best practices in service design.

## Overview

The project is composed of several independent microservices, each responsible for a distinct domain within the digital wallet ecosystem. The main services include:

- **User Service:**  
  Manages user registration, authentication, and profile management.
  
- **Wallet Service:**  
  Handles wallet operations such as deposits, withdrawals, and balance management.
  
- **Transaction Service:**  
  Processes financial transactions and utilizes messaging for asynchronous, reliable communication.
  
- **Notification Service:**  
  Sends real-time notifications and alerts based on system events and transaction outcomes.

Each microservice is implemented as an independent Go module, enabling isolated development, testing, and deployment. Services communicate via RESTful APIs and, where applicable, asynchronous messaging to ensure loose coupling and scalability.

## Architecture

- **Microservices Architecture:**  
  Each service runs independently, allowing you to scale and deploy components as needed.
  
- **RESTful APIs:**  
  Built using the Gin framework, the services expose endpoints that are fast and easy to integrate.
  
- **Database Integration:**  
  PostgreSQL is used for reliable relational data storage, with SQLX simplifying interactions with the database.
  
- **Authentication:**  
  JSON Web Tokens (JWT) are implemented for secure, stateless authentication across services.
  
- **Containerization:**  
  Docker is used to containerize each service, ensuring consistency across different environments. The project is also ready for deployment on orchestration platforms like Kubernetes.
  
- **Asynchronous Communication:**  
  Messaging queues (e.g., Kafka, RabbitMQ, or SQS) are integrated for processing transactions and inter-service communication.

## Technologies Used

- **Go:** The primary programming language for building the microservices.
- **Gin:** A lightweight and high-performance HTTP web framework for creating RESTful APIs.
- **PostgreSQL:** The relational database used for persistent data storage.
- **SQLX:** An extension for Go’s `database/sql` package that simplifies database interactions.
- **JWT:** For secure user authentication and session management.
- **Docker:** For containerizing the services and ensuring environment consistency.
- **Messaging Queues:** (Optional) Tools like Kafka, RabbitMQ, or SQS for enabling asynchronous communication between services.

## Project Structure

```plaintext
digital-wallet-micro-services/
├── notification-service
├── transaction-service
├── user-service
└── wallet-service
```
Each subfolder is an independent Go module with its own go.mod file and specific README documentation detailing its setup, configuration, and usage instructions.

## Getting Started
### Prerequisites
- Go 1.18+ installed on your machine.
- A running PostgreSQL instance (locally or via Docker).
- Docker for containerization.
- A message broker (e.g., Kafka, RabbitMQ) for asynchronous processing.
## Running the Services
To run a service individually (for example, the User Service):

1. Navigate to the service directory:
```bash
cd digital-wallet-micro-services/user-service
```
2. Configure your environment variables by creating a `.env` file.
## Install dependencies:
```bash
go mod tidy
```
4. Start the service:
```bash
go run main.go
```
Repeat similar steps for the other services. Alternatively, you can use Docker to build and run containers for each service.

## Contributing
Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request detailing your changes.
Please ensure your code adheres to the project’s standards and includes appropriate tests.

## License
This project is licensed under the MIT License. See the LICENSE file for further details.

## Contact
For questions, suggestions, or issues, please reach out at tontech.dev@outlook.com