# Receipt Processor Challenge Maaz

 This API provides functionality to process receipts, calculate points based on specific rules, and retrieve health status. 

 ## Features

- **POST /receipts/process**: Process a receipt by validating and storing it with a unique ID.
- **GET /receipts/:id/points**: Retrieve points for a specific receipt.
- **GET /receipts/health**: Check API health status.

## Project Structure

- **handlers.go**: Defines HTTP handlers for receipt processing, point calculation, and health check.
- **main.go**: Sets up routes and starts the server.
- **models.go**: Contains data models (`Item`, `Receipt`).
- **utils.go**: Contains helper functions for receipt validation, points calculation, and formatting.
- **Dockerfile**: Used for containerizing the application.

## Prerequisites

- **Go**: version 1.20+
- **Docker** (optional): If running the application in a Docker container.

## Installation and Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/maazahmed87/fetch-rewards-maaz
   cd fetch-rewards-maaz
   ```

2. **Install Dependencies**
    ```bash
    go mod tidy
    ```

## Running the Application

### Running Locally
1. **Run the server**:
    ```bash
    go run .
    ```
2. **Access the API**:
 The server will start on http://localhost:8080

### Running with Docker
1. **Build the Docker Image:**
    ```bash
    docker build -t receipt-processor-solution-api-maaz .
    ```

2. **Run the Docker Container:**
    ```bash
    docker run -p 8080:8080 receipt-processor-solution-api-maaz
    ```

## References
Here are some references and resources used for guidance in building this project:

1. [Gin Web Framework Documentation](https://gin-gonic.com/docs/)  
   Official documentation for the Gin framework, which provides an extensive guide on setting up routes, handling requests, and building APIs in Go. Like every language, I found using an external framework easier to code than using the standard library.

2. [Dockerizing Go Applications](https://docs.docker.com/language/golang/build-images/)  
   A guide from Docker's official documentation on containerizing Go applications, covering Dockerfile structure, multi-stage builds, and best practices.

3. [Writing Go Modules and Using go.mod](https://golang.org/doc/modules/managing-dependencies)  
   An overview of Go modules, dependency management, and usage of `go.mod` for managing dependencies in a Go project.

4. [Go Time Parsing and Formatting](https://golang.org/pkg/time/#Parse)  
   Documentation on Go’s `time` package, particularly useful for time parsing and formatting, used in date and time validation for this project. It was a surprising revelation on how Go differed in its Time format parsing.

5. [Go Code Style and Conventions](https://github.com/eleniums/code-conventions/blob/master/go/style.md)  
   A style guide for Go code conventions, which offers best practices for writing clean, maintainable Go code.