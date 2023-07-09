# go-rest-example

I made this project on 2018-2019, now on 2023 I refactored the sh!t out of it as it was basic as fvck. Still needs some luv <3

## Architecture

Here's a brief overview of the project's structure:

- `cmd/main.go`: This is the main file that initializes all dependencies and starts the application.

- `pkg/auth`: This package contains the code needed for user authentication.

- `pkg/codec`: This package is responsible for transforming between entities and models.

- `pkg/config`: This package manages the application's configuration.

- `pkg/entities`: This package defines the data structures used in the transport/services layers.

- `pkg/models`: This package defines the data structures used for the database.

- `pkg/repository`: This package contains all the database code, including setup and queries.

- `pkg/service`: This package contains all the business logic code. It accepts and returns entities.

- `pkg/transport`: This package manages the different endpoints and handles requests/responses.

- `pkg/utils`: This package contains utility functions used across the application.

- `config.json`: This file holds the JSON configuration for the application.

## Running the application

To get started with this project, clone the repository and install the necessary dependencies. Then, you can run the application using the command `go run cmd/main.go`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
