package controller

import "github.com/chakra/ratelimiter/service/registration"

//CheckAPI Function to validate the request on the throttle limit
func CheckAPI(url string) bool {
	return true
}

func buildContract() *registration.Contract {
	c := &registration.Contract{}
	return c
}

//RegisterContract - Allows the registration of the API
func RegisterContract(url string) {

	c := buildContract()
	registration.RegisterContract(c)
}
