package registration

import (
	"fmt"
	"log"
)

func CreateContract(user string, group string, allowedRequest int64, window int16) *Contract {
	var apigroup = APIGroup(group)

	c := &Contract{
		User:           user,
		Group:          apigroup,
		AllowedRequest: allowedRequest,
		Window:         window,
	}
	return c
}

func RegisterContract(c ContractCRUD) bool {
	contract := c.getObject()
	existingContract, _ := c.getContractByNameAndGroup(contract.User, contract.Group)
	if existingContract.getObject().id != 0 {
		log.Println("It seems to be duplicate contract")
		return false
	}
	fmt.Printf("Register the user %s under the group %s with the limit %d per %d minute(s)\n", contract.User, contract.Group, contract.AllowedRequest, contract.Window)
	return c.addContract()
}

func GetContractByName(contracts ContractsGetter, user string) {

	contracts.getContractsByUser(user)
	fmt.Println("contracts", contracts)
}

func GetContractByNameAndGroup(c ContractCRUD, user string, group APIGroup) (Contract, error) {

	contract, err := c.getContractByNameAndGroup(user, group)

	return *contract.getObject(), err
}

func AddApi(api string, contract ContractCRUD) bool {
	return contract.addAPI(api)
}
