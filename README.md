# Go REST Example

**Hey kiddo!** If you're reading this, take a few seconds to be mindful of your body and your surroundings, feeling your breath as it enters and exits your lungs. We are now out of the auto-pilot state we were in, and are fully ready to continue.

This repository is just a template or example of a **medium-sized microservice** that exposes a simple yet elegant and extensible **HTTP REST API**.

Using Gin, it provides a clean and layered architecture complete with User Management, JWT Authentication, MySQL Storage, Easy Configuration and much more.

Go's best practices and standards are followed most of the time, using techniques as Dependency Injection and Table Driven Tests and applying the holy proverbs 🙏.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation and Usage](#installation-and-usage)
- [Project Architecture](#project-architecture)
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

