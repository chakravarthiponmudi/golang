package throttler

import (
	"log"

	"github.com/chakra/ratelimiter/service/registration"
)

//TODO: api path are to compare as a regex , becuase that's what is going to be in realtime..
func getContract(client string, api string) (registration.Contract, error) {
	var contractCRUD *registration.Contract
	contract, err := registration.GetContractByNameAndPath(contractCRUD, client, api)
	if err != nil {
		log.Println("getContract failed while fetching for ", client, "and path ", api)
	}

	return contract, err
}

//TODO: The initial naive approach is to build in a in memory sliding window counter first. later this is to be
//moved to the distributed cluster for HA and scalability
func CheckAPI(client string, api string) bool {
	return false
}
