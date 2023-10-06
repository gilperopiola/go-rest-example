# Go REST Example

**Hey kiddo!** If you're reading this, take a few seconds to be mindful of your body and your surroundings, feeling your breath as it enters and exits your lungs. We are now out of the auto-pilot state we were in, and are fully ready to continue.

This repository is just a template or example of a **medium-sized microservice** that exposes a simple yet elegant and extensible **HTTP REST API**.

Using Gin, it provides a clean and layered architecture complete with User Management, JWT Authentication, MySQL Storage, Easy Configuration and much more.

Go's best practices and standards are followed most of the time, using techniques as Dependency Injection and Table Driven Tests and applying the holy proverbs 🙏.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation and Usage](#installation-and-usage)
- [Project Architecture](#project-architecture)
- [Request Flow Overview](#request-flow-overview)
- [Contributing and License](#contributing-and-license)

## Prerequisites

- Go 1.x
- A suitable database (MySQL, PostgreSQL, etc)
- Having a keyboard idk

## Installation and Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/gilperopiola/go-rest-example.git
   ```

2. Navigate to the project directory:
   ```bash
   cd go-rest-example
   ```

3. Install the required dependencies:
   ```bash
   go mod download
   ```

4. Set up your database and update the `config.json` file as you see fit.

5. Build the project:
   ```bash
   go build -o go-rest-example
   ```

6. Run the built binary:
    ```bash
    ./go-rest-example
    ```

This will start the server and you'll hopefully be able to hit the endpoints.

You can also use Docker :)

## Project Architecture

| Package | Description |
|---------|-------------|
| **cmd** | Contains the main application entry point. |
| **pkg > auth** | Handles authentication logic. |
| **pkg > codec** | Responsible for encoding and decoding tasks. |
| **pkg > entities** | Defines various entities and custom errors. |
| **pkg > repository** | Manages database connections and operations. |
| **pkg > service** | Contains the core business logic. |
| **pkg > transport** | Sets up routes and manages the transport layer. |
| **pkg > utils** | Houses utility functions used across the project. |

## Request Flow Overview

1. **Entry Point**: 
   
   All requests start at `router.go`, where the URL is matched.

2. **Token Validation**: 
   
   For private endpoints, the request is routed through `token_validation.go` for authentication.

3. **Endpoint Matching**: 
   
   The corresponding function in `transport_endpoints.go` is invoked based on the matched URL.

4. **Request Handling**: 
   
   This function then calls `HandleRequest` from `transport.go`.

5. **Request Typification & Validation**: 
   
   Inside `transport.go`, the appropriate function from `transport_requests.go` is called. This step involves typifying and validating the incoming request.

6. **Service Layer Invocation**: 
   
   Upon successful validation, `HandleRequest` invokes the corresponding method in `service.go`.

7. **Service Implementation**: 
   
   The actual implementation of these methods resides in `service_xxx.go`, organized by domain.

8. **Entity to Model Conversion**: 
   
   The service layer operates using **entities**. If needed, the `codec` is used to convert these entities into **models** before interacting with the `repository` layer.

9. **Repository Layer**: 
   
   Here, the matching `repository_xxx.go` file is invoked. This layer handles database operations and returns data or errors as required.

10. **Backtracking & Response Handling**: 
   
   The system then retraces its steps. If the repository layer returned a model, it's converted back to an entity in `service_xxx.go`. This file also handles any errors or continues execution if there are none.

11. **Finalizing Response**: 
   
   The process returns to `HandleRequest` in `transport.go`, where the service's response is processed. If there's an error, it's mapped using `errors_mapper.go`.

12. **Sending HTTP Response**: 
   
   Finally, the HTTP response is sent back to the client.

## Contributing and License

### Contributing

We welcome contributions! If you find a bug or want to add a new feature, feel free to create an issue or submit a pull request.

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your changes and commit them.
4. Push your changes to your fork.
5. Submit a pull request.

### License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

# 🐿️

