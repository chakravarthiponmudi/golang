package throttler

import (
	"log"
	"time"

	"github.com/chakra/ratelimiter/library"

	"github.com/chakra/ratelimiter/service/registration"
)

//TODO: Creating a map which is local, should be moved to the redis cluster later

var windowMap map[registration.Contract]*library.Window

//TODO: api path are to compare as a regex , becuase that's what is going to be in realtime..
func getContract(client string, api string) (registration.Contract, error) {
	contract, err := registration.GetContractByNameAndPath(new(registration.Contract), client, api)
	if err != nil {
		log.Println("getContract failed while fetching for ", client, "and path ", api)
	}

	return contract, err
}

func getSliderForContract(contract registration.Contract) *library.Window {
	return windowMap[contract]
}

func createNewSlider(contract registration.Contract) *library.Window {
	windowMap[contract] = library.CreateWindow(10, (1 * time.Second), contract.AllowedRequest)
	return nil
}

//TODO: The initial naive approach is to build in a in memory sliding window counter first. later this is to be
//moved to the distributed cluster for HA and scalability
func CheckAPI(client string, api string) bool {
	contract, err := getContract(client, api)
	if err != nil {
		return false
	}
	slider := getSliderForContract(contract)
	if slider == nil {
		slider = createNewSlider(contract)
	}

	if slider.IsLimitExceeded() {
		return false
	}
	return true
}
