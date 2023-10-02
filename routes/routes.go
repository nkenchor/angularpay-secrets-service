package routes

import (
	"errors"
	"log"
	"net/http"
	"os"
	"secret/interfaces"
	"secret/sharedinfrastructure/helper"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(appPort, hostAddress string, secret interfaces.SecretInterface) *chi.Mux {
	var w http.ResponseWriter
	var e helper.ErrorBody
	var errorMessage helper.ErrorResponse

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	fileName := "log/secrets-service.log"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		message := errors.New("logfile:" + err.Error())
		errorMessage.ErrorMessage(e, "500", "unable to open log file", "secret-service.log", message.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	customLogger := log.New(f, "", log.LstdFlags)
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: customLogger, NoColor: true})
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Mount("/secret", secretEndpoint(secret))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(hostAddress + ":" + appPort+"/swagger/doc.json"),
	))

	return router
}

func secretEndpoint(secret interfaces.SecretInterface) http.Handler {
	r := chi.NewRouter()
	r.Post("/entries", secret.CreateSecret)
	r.Put("/entries/{reference}", secret.UpdateSecret)
	r.Delete("/entries/{reference}", secret.DeleteSecret)
	r.Get("/entries/{reference}", secret.GetSecretByRef)
	r.Get("/entries", secret.GetAllSecret)
	r.Post("/entries/{reference}", secret.GetServiceSecretList)
	return r
}
