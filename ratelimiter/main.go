package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chakra/ratelimiter/library"
	"github.com/chakra/ratelimiter/service/registration"

	"github.com/chakra/ratelimiter/controller"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 20)
	if controller.CheckAPI(r.URL.Path) {
		fmt.Fprintf(w, "%s is allowed", r.URL.Path[1:])
	} else {
		fmt.Fprintf(w, "%s is denied", r.URL.Path[1:])
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	params := r.Form
	allowedRequest, err := strconv.Atoi(params.Get("allowedRequest"))

	if err != nil {
		fmt.Fprintf(w, "Bad input yet to send the proper code")
		return
	}

	window, err := strconv.Atoi(params.Get("window"))

	if err != nil {
		fmt.Fprintf(w, "Bad input yet to send the proper code")
		return
	}
	c := registration.CreateContract(params.Get("user"), params.Get("group"), int64(allowedRequest), int16(window))

	if registration.RegisterContract(c) {
		fmt.Fprintf(w, "Successfully registered")
	} else {
		fmt.Fprintf(w, "Registration Failed")
	}

}

func createServer() *http.Server {
	s := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	return s
}

func getContracts(w http.ResponseWriter, r *http.Request) {
	contract := new(registration.Contracts)
	registration.GetContractByName(contract, "Athena")

}

func addApi(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clientName := r.Form.Get("clientname")
	apipath := r.Form.Get("path")
	clientgroup := registration.APIGroup(r.Form.Get("clientgroup"))

	var (
		contract = new(registration.Contract)
		err      error
	)
	*contract, err = registration.GetContractByNameAndGroup(contract, clientName, clientgroup)

	if err != nil {
		fmt.Fprintf(w, "No Contract found")
		return
	}

	if registration.AddApi(apipath, contract) {
		fmt.Fprintf(w, "Adding API done")
	} else {
		fmt.Fprintf(w, "Adding API failed")
	}

}

func main() {
	library.SetupConnection()
	http.HandleFunc("/", handler)
	http.HandleFunc("/contract", registerHandler)
	http.HandleFunc("/contract/user", getContracts)
	http.HandleFunc("/api", addApi)

	server := createServer()
	log.Fatal(server.ListenAndServe())
}
