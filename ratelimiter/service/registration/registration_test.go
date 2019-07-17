package registration

import (
	"fmt"
	"testing"
)

type mockContract struct {
	testdata                        Contract
	addContractResult               bool
	deleteContractResult            bool
	updateContractResult            bool
	addAPIResult                    bool
	getContractByNameAndGroupResult struct {
		c ContractCRUD
		e error
	}
}

func (m *mockContract) addContract() bool {
	return m.addContractResult
}

func (m *mockContract) deleteContract() bool {
	return m.deleteContractResult
}

func (m *mockContract) updateContract() bool {
	return m.updateContractResult
}

func (m *mockContract) addAPI(apipath string) bool {
	return m.addAPIResult
}

func (m *mockContract) getContractByNameAndGroup(user string, group APIGroup) (ContractCRUD, error) {
	fmt.Println("calling mockedgetContractByNameAndGroup by", m)
	return m.getContractByNameAndGroupResult.c, m.getContractByNameAndGroupResult.e
}

func (m *mockContract) getObject() *Contract {
	return &m.testdata
}

func TestRegisterContract(t *testing.T) {
	var testContract = new(mockContract)
	testContract.testdata = Contract{}
	testContract.addContractResult = true
	testContract.getContractByNameAndGroupResult.e = nil
	testContract.getContractByNameAndGroupResult.c = &testContract.testdata
	result := RegisterContract(testContract)
	if !result {
		t.Error(result)
	}
}
