package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

func createContract(user string, group string, allowedRequest int64, window int16) *registration.Contract {
	var apigroup = registration.APIGroup(group)

	c := &registration.Contract{
		User:           user,
		Group:          apigroup,
		AllowedRequest: allowedRequest,
		Window:         window,
	}
	return c
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
	c := createContract(params.Get("user"), params.Get("group"), int64(allowedRequest), int16(window))

	registration.RegisterAPI(*c)
	fmt.Fprintf(w, "Successfully registered")
}

func createServer() *http.Server {
	s := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	return s
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/contract", registerHandler)

	server := createServer()
	log.Fatal(server.ListenAndServe())
}
