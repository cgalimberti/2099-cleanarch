# Ordersystem Project

## Overview
The Ordersystem project is a clean architecture implementation for managing orders. It provides a web server, gRPC server, and GraphQL server for interacting with order data. The project is designed to be modular and maintainable, following best practices in Go development.

## Features
- Create and list orders via REST API, gRPC, and GraphQL.
- Event-driven architecture using RabbitMQ for handling order creation events.
- Configuration management for database and server settings.
- Unit tests for use cases to ensure reliability.

## Setup Instructions
1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd ordersystem
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Set up the database:**
   - Configure your database settings in `configs/config.go`.
   - Run database migrations located in the `migrations` directory.

4. **Run the application:**
   ```
   go run cmd/ordersystem/main.go
   ```

## Ports
- **Web Server:** Accessible on the port defined in `configs.WebServerPort`.
- **gRPC Server:** Accessible on the port defined in `configs.GRPCServerPort`.
- **GraphQL Server:** Accessible on the port defined in `configs.GraphQLServerPort`.

## API Endpoints
- **Create Order:** `POST /order`
- **List Orders:** `GET /order`

## Docker Setup
To run the application using Docker, ensure you have Docker installed and run:
```
docker-compose up
```

## Testing
Unit tests for the use cases can be run using:
```
go test ./internal/usecase
```

## Additional Files
- **Dockerfile:** Defines the Docker image for the application.
- **docker-compose.yaml:** Defines services for the application and database.
- **api.http:** Contains HTTP requests for testing the endpoints.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.