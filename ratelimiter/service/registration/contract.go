package registration

import (
	"database/sql"
	"log"

	"github.com/chakra/ratelimiter/library"
)

func (c Contract) addContract() bool {
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

func getContractsByUser(user string) ([]Contract, error) {
	sqlStatement := `
	SELECT contractid, clientname, clientgroup,allowedlimit, windowinminutes FROM contract where clientname = $1
	`
	rows, err := library.GetDBConnection().Query(sqlStatement, user)

	if err != nil {
		log.Println("Selection of the data from contract failed", err)
		return nil, err
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
		return nil, err
	}
	return contracts, nil
}

func getContractByNameAndGroup(user string, group APIGroup) (Contract, error) {
	sqlStatement := `
	SELECT contractid, clientname, clientgroup,allowedlimit, windowinminutes FROM contract where clientname = $1 and clientgroup = $2
	`

	row := library.GetDBConnection().QueryRow(sqlStatement, user, group)
	var c Contract
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

func (c Contract) addAPI(apipath string) bool {
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
