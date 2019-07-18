package registration

import (
	"testing"
)

type compositeResult struct {
	c ContractCRUD
	e error
}

type mockContract struct {
	testCase                        string
	testdata                        Contract
	addContractResult               bool
	deleteContractResult            bool
	updateContractResult            bool
	addAPIResult                    bool
	getContractByNameAndGroupResult compositeResult
	expectedResult                  bool
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
	return m.getContractByNameAndGroupResult.c, m.getContractByNameAndGroupResult.e
}

func (m *mockContract) getObject() *Contract {
	return &m.testdata
}

func TestRegisterContract(t *testing.T) {

	var testContracts []mockContract
	testContracts = []mockContract{
		mockContract{
			testCase:          "All data is fine",
			testdata:          Contract{},
			addContractResult: true,
			getContractByNameAndGroupResult: compositeResult{
				e: nil,
				c: &Contract{},
			},
			expectedResult: true,
		},
		mockContract{
			testCase:          "When a contract already exists",
			testdata:          Contract{},
			addContractResult: true,
			getContractByNameAndGroupResult: compositeResult{
				e: nil,
				c: &Contract{
					id: 1,
				},
			},
			expectedResult: false,
		},
	}

	for _, test := range testContracts {
		actualResult := RegisterContract(&test)
		if actualResult != test.expectedResult {
			t.Error(test.testCase)
		}
	}
}

func TestAddAPI(t *testing.T) {

	var testContracts []mockContract
	testContracts = []mockContract{
		mockContract{
			testCase:     "All data is fine",
			testdata:     Contract{},
			addAPIResult: true,
			getContractByNameAndGroupResult: compositeResult{
				e: nil,
				c: &Contract{},
			},
			expectedResult: true,
		},
		mockContract{
			testCase:     "When a API for a contract and group already exists",
			testdata:     Contract{},
			addAPIResult: false,
			getContractByNameAndGroupResult: compositeResult{
				e: nil,
				c: &Contract{
					id: 1,
				},
			},
			expectedResult: false,
		},
	}

	for _, test := range testContracts {
		actualResult := AddApi("/test/path", &test)
		if actualResult != test.expectedResult {
			t.Error(test.testCase)
		}
	}
}
