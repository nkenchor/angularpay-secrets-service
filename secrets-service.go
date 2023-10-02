package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"secret/interfaces"
	"secret/routes"
	"secret/sharedinfrastructure/helper"
	"secret/sharedinfrastructure/persistence"
)

func init() {
	fileName := "log/secrets-service.log"
	var w http.ResponseWriter
	var e helper.ErrorBody
	var errorMessage helper.ErrorResponse
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		message := errors.New("logfile:" + err.Error())
		errorMessage.ErrorMessage(e, "500", "unable to open log file", "log.log", message.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.SetOutput(f)
}

func main() {
	address, port, mode, dbhost, dbname := helper.LoadConfig()

	secret, err := persistence.ConnectDB(dbhost, dbname)
	if err != nil {
		fmt.Println(err)
	}

	secretEndPoint := interfaces.NewSecret(secret.Secret)


	fmt.Println("App running on " + address + ":" + port)
	if mode == "dev" {
		r := routes.SetupRouter(port, address, secretEndPoint)
		http.ListenAndServe(":" + port, r)
	}

}
