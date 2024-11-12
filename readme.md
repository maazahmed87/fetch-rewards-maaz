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

## Future Improvements

While this application demonstrates modularity and ease of readability, several improvements and enhancements could be implemented to bring it up to production-level quality. These improvements would help ensure scalability, reliability, security, and overall robustness in a real-world environment.

Since this exercise is part of the interview process for a Backend Engineer position at Fetch Rewards, I focused more on core programming principles and thoughtful design. While I understand that the goal is to create production-ready code, I’ve aimed to present this project as an extension of my understanding of programming fundamentals and feature functionality. The following improvements represent enhancements that would be implemented to make the project fully production-ready:

1. **Persistent Storage**  
   Currently, the application uses in-memory storage for receipts, which is suitable for testing but limits scalability and causes data loss on restart. Implementing a persistent database would enable durable data storage and better scalability.

2. **Rate Limiting**  
   To safeguard against abuse, adding rate limiting would restrict the number of requests each user or IP can make over a defined period. This would prevent excessive load and protect the service from potential DoS attacks. 

3. **Security Enhancements**
   - **Authentication and Authorization**: Adding authentication would ensure secure access control, only allowing authorized users to access the API.
   - **Input Validation and Sanitization**: Strengthening validation and sanitization on all user inputs would minimize risks of injection attacks and data tampering.
   - **HTTPS**: Enforcing HTTPS ensures encrypted communication.

4. **Logging and Monitoring**  
   Introducing logging and monitoring solutions would provide visibility into application performance, errors, and system health. This would also enable real-time alerts, helping with proactive maintenance and issue resolution.

5. **Error Handling and Retry Logic**  
   Improved error handling and retry mechanisms would make the application more resilient to unexpected issues, ensuring a smoother and more reliable experience for end users.

6. **Testing and CI/CD Pipeline**  
   Enhancing test coverage with unit and integration tests and setting up a CI/CD pipeline would promote code quality and ensure consistent, reliable deployments.

7. **Container Orchestration and Scalability**  
   Deploying the application on a container orchestration platform like Kubernetes would enable automatic scaling, load balancing, and better resource management, making the service more adaptable to high traffic and ensuring greater resilience.

By listing these suggestions, I demonstrate my understanding of the additional steps needed for a production environment, while highlighting the foundational aspects of coding and design I focused on in this project.
