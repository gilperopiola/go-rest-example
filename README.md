# Go REST Example

**`Hey kiddo!`** Now that you're here, please take a few seconds to breathe. In & out. Slowly... Gently... Yeah, right there!

Try to be mindful of your body and your surroundings... are you comfortable? Are you tense? **Untense those muscles!** 

...And now that your auto-pilot mode is turned off, let's dive in! 🐙

---

This project is a template ~ example ~ boilerplate of a `large-sized microservice` exposing an elegant and extensible `HTTP REST API`. If you're going for a simpler microservice you should check out [v1.2.0](https://github.com/gilperopiola/go-rest-example/releases/tag/v1.2.0), it's much simpler.

Using **Gin**, it provides a **clean and layered architecture** complete with:

|User Management|Unit Tests|JWT Auth|
|:-------------:|:-------------:|:-------------:|
|**Mature Error Handling**|**Easy Config**|**Logrus Logging**|
|**NewRelic Monitoring**|**MySQL Storage**|**Design Patterns**|
|**Follows Best Practices**|**Inversion of Control**|**Dependency Injection**|
|**Entities & DB Models**|**Dockerized**|**Modular Interfaces**|

## Table of Contents

- [Installation and Usage](#installation-and-usage)
- [Project Architecture](#project-architecture)
- [Request Flow Overview](#request-flow-overview)
- [Contributing and License](#contributing-and-license)

## Installation and Usage

You'll need to have **Go 1.x.x** installed, as well as a **MySQL Database** running.

1. Clone the repository:
   ```bash
   git clone https://github.com/gilperopiola/go-rest-example.git
   ```

2. Navigate to the project directory and install dependencies:
   ```bash
   cd go-rest-example
   go mod download
   ```

3. Set up your database _(you can use XAMPP)_ and your environment variables. **You should read the `config.go` file before going further**.

4. Run the project:
   ```bash
   go run cmd/main.go
   ```

This will start the server and you'll (hopefully 🛐) be able to hit the endpoints.

**You can also use Docker :)**

## Project Architecture

On `cmd/main.go` we have the app's entrypoint as well as the initialization of the dependencies.

Then, it's basically divided into 3 layers: `Transport, Service and Repository`.

`Transport` handles routing, validations, error management and logging.

`Service` handles the business logic, it receives the requests created by the previous layer, calls the next one, and returns the appropiate responses. The `handlers` package is also part of this layer.

`Repository` handles connections with external dependencies, such as the database.

| Package | 👀 |
|---------|-------------|
| `auth` | Token creation and validation. |
| `common` | Things used across the project. Logger, requests & responses. |
| `config` | Self explanatory. |
| `entities` | Our structs, uncoupled of the database models. |
| `errors` | Our custom errors. |
| `handlers` | Interfaces used to interact with our models. Part of the service layer. |
| `models` | Our database models. |

## Request Flow Overview

1. **Entrypoint & Token Validation**: 
   
   All requests start at `router.go`, where the URL is matched with a transport function. For private endpoints, the request is routed through `token_validation.go` for authentication. Each token contains a Role, some endpoints require a specific Role to work.

2. **Endpoint Call, Request Handling**: 
   
   The corresponding function in `transport_endpoints.go` is invoked based on the matched URL. This function then calls `HandleRequest` from `transport.go`. This function orchestrates everything.

3. **Request Typification & Validation**: 
   
   Inside `transport.go`, the appropriate function from `transport_requests.go` is called. This step involves typifying and validating the incoming request, returning it back to the previous file.

4. **Service Layer Invocation**: 
   
   Upon successful validation, `HandleRequest` invokes the corresponding method in `service.go`. The actual implementation of these methods resides in `service_xxx.go`, divided by domain.

5. **Requests to Models to Handlers**: 
   
   The service layer receives requests and transforms them into **models**. With a model, the handler is then created and the business functions can be called. These will interact with the last layer, the Repository.

6. **Repository Layer**: 
   
   Here, the matching `repository_xxx.go` file is invoked. This layer handles outside dependencies such as database operations. Then, the backtracking begins.

7. **Backtracking & Response Handling**: 
   
   The system then retraces its steps. If the repository layer returned a model, it's converted back to an entity in `service_xxx.go`. This file also handles any errors or continues execution if there are none.

8. **Finalizing Response**: 
   
   The process returns to `HandleRequest` in `transport.go`, where the service's response is processed. If there's an error, it's mapped and logged using `errors_mapper.go`. The HTTP response is User Friendly, but the logs contain the whole stacktrace of the request.

9. **Voilá**: 
   
   And voilá 🌞

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

# Thanks for reading! 🐿️

