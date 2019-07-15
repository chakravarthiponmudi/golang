package registration

import "fmt"

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

func RegisterContract(c ContractCUD) bool {
	contract := c.getObject()
	fmt.Printf("Register the user %s under the group %s with the limit %d per %d minute(s)\n", contract.User, contract.Group, contract.AllowedRequest, contract.Window)
	return c.addContract()
}

func GetContractByName(contracts ContractsGetter, user string) {

	contracts.getContractsByUser(user)
	fmt.Println("contracts", contracts)
}

func GetContractByNameAndGroup(c ContractGetter, user string, group APIGroup) error {

	err := c.getContractByNameAndGroup(user, group)

	return err
}

func AddApi(api string, contract ContractCUD) bool {
	return contract.addAPI(api)
}
