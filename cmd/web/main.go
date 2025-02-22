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
	logger *slog.Logger
}

func main() {
	logger := slog.New((slog.NewTextHandler(os.Stdout, nil)))

	cfg := config{
		env: "development",
	}

	app := &application{
		logger: logger,
		config: cfg,
	}

	srv := &http.Server{
		Addr:     ":3000",
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	logger.Info("server is running on port 3000")
	err := srv.ListenAndServe()

	logger.Error(err.Error())
	os.Exit(1)
}
