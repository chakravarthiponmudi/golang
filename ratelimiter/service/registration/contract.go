package registration

import (
	"database/sql"
	"log"

	"github.com/chakra/ratelimiter/library"
)

type APIGroup string

type Contract struct {
	id             int64
	User           string
	Group          APIGroup
	AllowedRequest int64
	Window         int16
}

type Contracts []Contract

type API struct {
	contract Contract
	apipath  string
}

type ContractCRUD interface {
	addContract() error
	deleteContract() bool
	updateContract() bool
	addAPI(apipath string) error
	getContractByNameAndGroup(user string, group APIGroup) (ContractCRUD, error)
	getObject() *Contract
}

type ContractsGetter interface {
	getContractsByUser(user string) error
}

func (c *Contract) getObject() *Contract {
	return c
}
func (c *Contract) addContract() error {
	sqlStatement := `
	INSERT INTO contract (clientname, clientgroup,allowedlimit, windowinminutes)
	VALUES($1, $2,$3,$4)
	`

	_, err := library.GetDBConnection().Exec(sqlStatement, c.User, c.Group, c.AllowedRequest, c.Window)

	return err
}

func (c *Contract) deleteContract() bool {
	sqlStatement := `DELETE FROM contract where contractid = $1`

	_, err := library.GetDBConnection().Exec(sqlStatement, c.id)

	if err != nil {
		log.Println("Deletion of a contract failed", err)
		return false
	}

	return true
}

func (c *Contract) updateContract() bool {
	sqlStatement := `
	UPDATE contract SET clientname = $1, 
	clientgroup = $2,
	allowedlimit= $3,
	windowinminutes = $4
	where contract
	`

	_, err := library.GetDBConnection().Exec(sqlStatement, c.User, c.Group, c.AllowedRequest, c.Window)

	if err != nil {
		log.Println("update of a contract failed", err)
		return false
	}

	return true
}

func (c *Contract) addAPI(apipath string) error {
	sqlStatement := `
	INSERT INTO api (api, contractid, clientgroup)
	VALUES($1, $2,$3)
	`
	_, err := library.GetDBConnection().Exec(sqlStatement, apipath, c.id, c.Group)

	return err
}

func (cc *Contract) getContractByNameAndGroup(user string, group APIGroup) (ContractCRUD, error) {
	sqlStatement := `
	SELECT contractid, clientname, clientgroup,allowedlimit, windowinminutes FROM contract where clientname = $1 and clientgroup = $2
	`

	c := new(Contract)
	row := library.GetDBConnection().QueryRow(sqlStatement, user, group)
	switch err := row.Scan(&c.id, &c.User, &c.Group, &c.AllowedRequest, &c.Window); err {
	case sql.ErrNoRows:
		log.Println("No Rows found")
		return c, err
	case nil:
		return c, nil
	default:
		log.Panic("getContractByNameAndGroup", err)
		return c, err
	}
}

func (c *Contracts) getContractsByUser(user string) error {
	sqlStatement := `
	SELECT contractid, clientname, clientgroup,allowedlimit, windowinminutes FROM contract where clientname = $1
	`
	rows, err := library.GetDBConnection().Query(sqlStatement, user)

	if err != nil {
		log.Println("Selection of the data from contract failed", err)
		return err
	}

	defer rows.Close()

	var contracts []Contract
	for rows.Next() {
		var c Contract
		err = rows.Scan(&c.id, &c.User, &c.Group, &c.AllowedRequest, &c.Window)
		contracts = append(contracts, c)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Iteration of the data from contract failed", err)
		return err
	}
	*c = contracts
	return nil
}
