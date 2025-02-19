package main

import (
	"log/slog" // new import
	"net/http"
	"os" // new import
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	// add new logger field
	logger *slog.Logger
}

func main() {
	// initialize logger and pass it to the application struct
	logger := slog.New((slog.NewTextHandler(os.Stdout, nil)))
	app := &application{
		logger: logger,
	}

	srv := &http.Server{
		Addr:    ":3000",
		Handler: app.routes(),
		// here we explicitly provide to the server our logger for logging errors
		ErrorLog: slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	// change fmt to logger.Info
	logger.Info("server is running on port 3000")
	// remove log.Fatal and replace it with waiting on the error and if so call the logger.Error + os.Exit
	err := srv.ListenAndServe()

	logger.Error(err.Error())
	os.Exit(1)
}
