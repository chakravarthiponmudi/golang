package registration

import (
	"fmt"
	"testing"
)

type mockContract struct {
}

func (m *mockContract) addContract() bool {
	return true
}

func (m *mockContract) deleteContract() bool {
	return true
}

func (m *mockContract) updateContract() bool {
	return true
}

func (m *mockContract) addAPI(apipath string) bool {
	return true
}

func (m *mockContract) getContractByNameAndGroup(user string, group APIGroup) (ContractCRUD, error) {
	fmt.Println("calling mockedgetContractByNameAndGroup by", m)
	return &mockContract{}, nil
}

func (m *mockContract) getObject() *Contract {
	return new(Contract)
}

func TestRegisterContract(t *testing.T) {
	var testContract = new(mockContract)
	result := RegisterContract(testContract)
	if !result {
		t.Error(result)
	}
}
