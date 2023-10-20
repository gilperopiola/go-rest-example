# Go REST Example

**`Hey kiddo!`** Tired of those same old boilerplate repositories that claim they're basically perfect and simple and elegant and modular and many other lies? sheesh. 🤮 

Here you have **EASY TO FOLLOW CODE** without having to learn an entire framework. Is it perfect? Yes. Well no, but it's robust and solid enough for a large microservice[^1] without having much complexity, and without sacrificing much. `SOLID and Clean Architecture are there, Logging & Monitoring, Auth, Errors, Tests, Docker, User Management`.

|User Management|Unit Tests|JWT Auth|
|:-------------:|:-------------:|:-------------:|
|**Gin Routing**|**Easy Config**|**Logrus Logging**|
|**Auth**|**MySQL Storage**|**Design Patterns**|
|**Best Practices**|**Great Error Handling**|**Dependency Injection**|
|**NewRelic Monitoring**|**Dockerized**|**Modular Interfaces**|
|**Inversion of Control**|**Integration Tests**|**Abstracted DB Models**|

## Installation and Usage

You'll need to have **Go 1.x.x** installed, as well as **Docker**, **Docker Compose** and a **MySQL Database** running. Good thing it's dockerized.[^3]

1. Clone, access folder, install deps:
   ```bash
   git clone https://github.com/gilperopiola/go-rest-example.git
   cd go-rest-example
   go mod download
   ```

2. Set up your environment variables. You should copy the **.env_example** file and name it **.env**, and then just make it yours.

3. Build & Run in Docker:
   ```bash
   docker-compose build
   docker-compose up
   ```

Fix any issue that might arise and `kiddo, you're GTG!`

## Project Architecture

On `cmd/main.go` we have the app's entrypoint and the initialization of the dependencies.

Then it's divided into 3 layers: `Transport, Service and Repository`.

`Transport` handles routing, middleware & auth, request validations, error management and logging.

`Service` handles the core business logic. It receives the requests created by the previous layer, calls the next layer, and then returns the appropiate responses.

`Repository` handles connections with external dependencies, such as the database. Talks in terms of Models.

## Request Lifecycle

Let's assume this is the lifecycle of a DELETE User request.

1. `**Entrypoint & Token Validation**`: 

   First, router.go receives the HTTP request, matches it with the URL and validates the token using auth.
   
2. `**Endpoint Call, Request Handling**`: 

   In `endpoints.go` and `handler.go` the corresponding function is called and the DeleteUserRequest is made and validated.
   
3. `**Service Layer & Models**`: 

   The call to `service_users.go` is made, and there the DeleteUserRequest is transformed to a UserModel

4. `**Repository Layer**`: 
   
   The UserModel makes all operations it needs to make, and then calls .Delete on itself. The DeleteUser method on the Repository Layer is called.

5. `**Database**`: 
   
   The Repository marks the user as deleted on the database. A UserModel with the new state of the user is returned.

6. `**Backtracking & Response Handling**`: 
   
   And then it's just going back the layers and returning the HTTP Response. If there's an error it goes through the `http_errors_mapper.go`.

**HTTP Request -> Router -> Handler -> Request -> Service -> Model -> Repository -> Response -> HTTP Response**[^2]

## Contributing and License

I don't care. Do what you will. I think there's a `LICENSE` file tho.

---

# Thanks for reading! 🐿️

[^1]: If you are aiming for a small or medium microservice, you should check out this repo's [v1.2.0](https://github.com/gilperopiola/go-rest-example/releases/tag/v1.2.0).

[^2]: or if you wanna go deep

**`HTTP Request -> Router -> Middleware -> Handler -> Request -> Service -> Model -> Repository -> Model -> Service -> Response -> Handler -> HTTP Response`**.

[^3]: If you don't want to use Docker, you can just run a MySQL database and point to it in `config.go`.