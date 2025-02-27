# User Service - Digital Wallet Microservices

The **User Service** is a core component of the Digital Wallet Microservices project. This microservice is responsible for handling all user-related functionalities, including registration, authentication, and profile management. Built in Go, it leverages modern technologies to provide a secure, scalable, and high-performance API for managing user data.

## Overview

The User Service exposes RESTful endpoints to support:
- **User Registration:** Allow new users to create an account.
- **User Login:** Authenticate users and issue JWTs.
- **Profile Management:** Enable users to view and update their personal information.

By using a microservices architecture, each component of the Digital Wallet ecosystem remains modular and maintainable. This service is designed to work seamlessly within the broader ecosystem while being independently deployable.

## Technologies Used

- **Go:** The programming language used for building the service.
- **Gin:** A high-performance HTTP web framework used for routing and middleware.
- **PostgreSQL:** The relational database for persisting user data.
- **JWT (JSON Web Tokens):** For secure user authentication and session management.
- **SQLX:** An extension library for Go’s standard `database/sql` package, simplifying interactions with PostgreSQL.
- **Docker:** Containerization for consistent deployment and easy scalability.
- **Environment Variables:** For managing configuration securely.

## Features

- **User Registration & Login:** Secure endpoints for new user sign-up and login.
- **JWT Authentication:** Secure token generation and validation to protect endpoints.
- **Profile Management:** Endpoints to retrieve and update user profiles.
- **Password Security:** Implementation of password hashing and secure storage.
- **Error Handling & Logging:** Robust error management and logging to facilitate debugging and monitoring.
- **Testing:** Unit and integration tests to ensure code quality and reliability.

## Getting Started

### Prerequisites

- Go 1.18+ installed on your machine.
- A running PostgreSQL instance.
- Docker (optional, for containerization).

### Setup Instructions

1. **Clone the Repository**

```bash
   git clone https://github.com/BrooitsFeiskJR/digital-wallet-micro-services.git
   cd digital-wallet-micro-services/user-service
```
### Configure Environment Variables

Create a `.env` file in the root of the user-service directory with the following content:

```dotenv
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret
```
### Install Dependencies

Run the following command to tidy up and download the necessary dependencies:

```bash
go mod tidy
```
Run the Service

Start the service using:

```bash
go run main.go
```

### Running with Docker
If you prefer to run the service in a containerized environment:

1. __Build the Docker Image__

```bash
docker build -t user-service .
```

2. __Run the Container__

```bash
docker run -p 8080:8080 --env-file .env user-service
```
## API Endpoints
Below is a summary of the main endpoints exposed by the User Service:

- __POST /register__ 
Register a new user with required details.

- __POST /login__
Authenticate the user and return a JWT for secure access.

- __GET /profile__

    Retrieve the profile information of the authenticated user.

    ___(Requires JWT authentication)___   

- __PUT /profile__

    Update the profile information of the authenticated user.
    
    ___(Requires JWT authentication)___

For a comprehensive list of endpoints and request/response schemas, please refer to the API documentation or Swagger/OpenAPI specification (if available).

## Contributing
Contributions to enhance functionality, security, and performance are welcome! Please follow the standard Git workflow:

- Fork the repository.
- Create a feature branch (`git checkout -b feature/your-feature`).
- Commit your changes (`git commit -m 'Add some feature`').
- Push to the branch (`git push origin feature/your-feature`).
- Open a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Contact
For questions, suggestions, or issues, please reach out at [tontech.dev@outlook.com].