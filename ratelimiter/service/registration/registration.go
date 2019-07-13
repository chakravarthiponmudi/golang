package registration

import "fmt"

type APIGroup string

type Contract struct {
	id             int64
	User           string
	Group          APIGroup
	AllowedRequest int64
	Window         int16
}

type API struct {
	contract Contract
	apipath  string
}

func RegisterAPI(c Contract) bool {
	fmt.Printf("Register the user %s under the group %s with the limit %d per %d minute(s)", c.User, c.Group, c.AllowedRequest, c.Window)
	return c.addContract()
}

func GetContractByName(user string) {

	contracts, _ := getContractsByUser(user)
	fmt.Println("contracts", contracts)
}

func GetContractByNameAndGroup(user string, group APIGroup) (Contract, error) {
	contract, err := getContractByNameAndGroup(user, group)
	return contract, err
}

func AddApi(api string, contract Contract) bool {
	return contract.addAPI(api)
}
