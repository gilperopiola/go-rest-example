# Go REST Example

**`Hey Kidd0!`** If you write (or wanna write) **HTTP REST APIs** in Golang this is your lucky day. 

Here you have **`Easy To Follow Code`** without having to learn an entire framework. Is it perfect? **Yes**. Well no, but it's pretty easy to understand and modify and extend. It has that **`Simplicity`** most other boilerplates don't, it has a fuckload of features and it's actually quite solid and robust, perfect for a large microservice. If your project is smaller, you should check out [v1.2.0](https://github.com/gilperopiola/go-rest-example/releases/tag/v1.2.0). 

## Features

|User Management|Easy Config|JWT Auth|MySQL Storage
|:-------------:|:-------------:|:-------------:|:-------------:|
|**Gin Routing**|**Unit Tests**|**Logrus Logging**|**24/7 Support 🤥**|
|**Clean Architecture**|**SOLID**|**Design Patterns**|**Prometheus Metrics**|
|**Best Practices**|**Great Error Handling**|**Dependency Injection**|**Profiling**|
|**New Relic Monitoring**|**Fully Dockerized**|**Modular Interfaces**|**E2E Testing**|
|**Inversion of Control**|**Postman Collection**|**Abstracted DB Models**|**Free Forever**|
|**Performance Optimized**|**Custom Log Formatter**|**Easy to Pick Up**|**Easy to Master**|

## Interested? ;)

Here below lies the installation and configuration guide (it's pretty short, I swear), and on the [**`README2.md`**](https://github.com/gilperopiola/go-rest-example/blob/master/README2.md) file you will find a small guide that will help you on this beautiful journey you've embarked.

## Installation and Usage

All you require to run this is having **Go 1.x.x** and a **MySQL Database** running. You can either run that manually or through **Docker**.

1. Clone repo, access folder, install dependencies:
   ```bash
   git clone https://github.com/gilperopiola/go-rest-example.git
   cd go-rest-example
   go mod download
   ```

2. Set up your environment variables. You should copy the **.env_example** file and name it **.env**. Then it's yours to manage.

3. Build & Run in Docker:
   ```bash
   make run
   ```

4. Build & Run locally:
   ```bash
   make run-local
   ```

The only potential issue you could have is the DB connection, but you'll fix it. I know. 🌈

## Contributing and License

**I don't care**. Do what you will. I think there's a `LICENSE` file in here, who reads those anyways.

---

# Thanks for your time! 🐿️
