# Outdoorsy challenge task - REST Server

# Description & Requirements
- The purpose of the application is to retrieve data for rentals, saved in a database.
- Handle search for a specific rental by id
- Handle list of rentals by various searching criteria
- The application features two logging mechanisms - in `app.log` file and the standard output
- The application will panic when:
    - Connection to the DB failed
    - Cannot open the log file
    - Fails to start the HTTP server (e.g. port is busy)

# Project structure
### app
* `app.go` - contains logic for the application startup - connecting to the database, initiating the services and and repositories
### handlers
* `rentals.go` - Contains the handlers (controllers) which are called upon endpoint hit.
    * Includes an interface `Handler` that defines all functions and implements it using Gorilla mux. The purpose of the interface is to make it possible to easily switch between different implementations of HTTP requests handling - gin-gonic, etc
    * Each handler's purpose is:
        * to validate parameters (if any)
        * to call the respective service in order to receive the result
        * transform the model to a response structure, encode it and return it
        * handle HTTP errors - missing/invalid parameters, internal server errors
### services
* `rentals.go` - Contains the business logic of the application. The implementation of the services is achieved through an interface and external logic is added through dependency injection.

### repositories
* `rentals.go` - Contains the interface `Repositorer` which define the base functions, used to access a single repository of any kind.
    * Also, contains the implementation of the interface, the structure `repository`, which is a PostgresDB implementation of the interface
    * The idea is to easily make it possible to switch between different repositories - db, file, in-memory, etc

### models
* `rentals.go` - contains model objects, used to transmit data inside the application between different packages and to differentiate the type from request/response/dto


### responses
* `rentals.go` - contains objects, used to be encoded and to transmit data back to the client

# Environment variables
* API_PORT
    * Description - the port, the application should run on
    * Default value - `1212`
* DB_CONNECTION_URL
    * Description - the connection string, used to access a Postgres database
    * Default val - *Used only for debug and testing purposes*

# Running
The application can both run locally and under a docker container.
### To run locally:
* Navigate to the base folder of the application
    * `go run main.go` or `go run .`
* In order to build an executable:
    * `go build`
        * Run for Linux/MacOS
            * `./outdoorsy`
        * Run for Windows
            * `.\outdoorsy.exe`

### To run in container
* Navigate to the base folder of the application
* Build the image
    * `docker build --tag outdoorsy .`
* Run the image
    * `docker run  -p {port}:{port} outdoorsy`
    * Replace `{port}` with the used port, in order for the container to expose it

# Testing
* The project includes integration tests in order to test the REST endpoints and make sure they return correct data. In order to run them, make sure the application is running.
* Navigate to the base folder of the application:
* `go test`
