package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/antonisgkamitsios/simple-todo-app/internal/data"
	"github.com/go-playground/form/v4"
)

type config struct {
	port int
	env  string
}

type application struct {
	config      config
	logger      *slog.Logger
	models      data.Models
	formDecoder *form.Decoder
}

func main() {
	logger := slog.New((slog.NewTextHandler(os.Stdout, nil)))

	cfg := config{
		env: "development",
	}

	db := data.NewDummyDB()

	formDecoder := form.NewDecoder()

	app := &application{
		logger:      logger,
		config:      cfg,
		models:      data.NewModels(db),
		formDecoder: formDecoder,
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
