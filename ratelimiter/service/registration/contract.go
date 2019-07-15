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

type ContractCUD interface {
	addContract() bool
	deleteContract() bool
	updateContract() bool
	addAPI(apipath string) bool

	getObject() *Contract
}

type ContractGetter interface {
	getContractByNameAndGroup(user string, group APIGroup) error
}

type ContractsGetter interface {
	getContractsByUser(user string) error
}

func (c *Contract) getObject() *Contract {
	return c
}
func (c *Contract) addContract() bool {
	sqlStatement := `
	INSERT INTO contract (clientname, clientgroup,allowedlimit, windowinminutes)
	VALUES($1, $2,$3,$4)
	`

	_, err := library.GetDBConnection().Exec(sqlStatement, c.User, c.Group, c.AllowedRequest, c.Window)

	if err != nil {
		log.Println("Insertion of a contract failed", err)
		return false
	}

	return true
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

func (c *Contract) getContractByNameAndGroup(user string, group APIGroup) error {
	sqlStatement := `
	SELECT contractid, clientname, clientgroup,allowedlimit, windowinminutes FROM contract where clientname = $1 and clientgroup = $2
	`

	row := library.GetDBConnection().QueryRow(sqlStatement, user, group)
	switch err := row.Scan(&c.id, &c.User, &c.Group, &c.AllowedRequest, &c.Window); err {
	case sql.ErrNoRows:
		log.Println("No Rows found")
		return err
	case nil:
		return nil
	default:
		log.Panic("getContractByNameAndGroup", err)
		return err
	}
}

func (c *Contract) addAPI(apipath string) bool {
	sqlStatement := `
	INSERT INTO api (api, contractid, clientgroup)
	VALUES($1, $2,$3)
	`
	_, err := library.GetDBConnection().Exec(sqlStatement, apipath, c.id, c.Group)

	if err != nil {
		log.Println("Insertion of a contract failed", err)
		return false
	}

	return true
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
