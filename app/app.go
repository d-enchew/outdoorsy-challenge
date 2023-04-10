package app

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"outdoorsy/handlers"
	"outdoorsy/repositories"
	"outdoorsy/services"
)

type app struct {
	router *mux.Router
}

const logsFileName = "./app.log"

// Initialize performs actions related to the application startup - connection to the database and
// builds the services objects
// panics when connection to any of the servers cannot be made
func Initialize(port, connectionString string) {
	a := &app{}
	logsFile, err := os.OpenFile(logsFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Unable to access logs file")
	}
	defer logsFile.Close()

	logger := logrus.New()
	logger.SetOutput(io.MultiWriter(os.Stdout, logsFile))

	repository, err := repositories.InitializeRepository(connectionString)
	if err != nil {
		panic("Error on DB connection.")
	}
	httpHandler := handlers.InitializeHandler(services.InitializeService(repository, logger))
	a.router = createRouter(httpHandler)

	logger.Info("Starting listening on port: ", port)
	err = http.ListenAndServe(":"+port, a.router)
	if err != nil {
		panic("Error on application startup")
	}
}

// createRouter creates the routing and registers all routes
func createRouter(httpHandler handlers.HTTPHandler) *mux.Router {
	router := mux.NewRouter()

	router.
		Methods(http.MethodGet).
		Path("/rentals/{id}").
		HandlerFunc(httpHandler.Get)
	router.
		Methods(http.MethodGet).
		Path("/rentals").
		HandlerFunc(httpHandler.List)

	return router
}
