package registration

import "fmt"

type API string

type APIGroup string

type APISettings struct {
	group APIGroup
	api   API
}

type Contract struct {
	User           string
	Group          APIGroup
	AllowedRequest int64
	Window         int16
}

func RegisterAPI(c Contract) {
	fmt.Printf("Register the user %s under the group %s with the limit %d per %d minute(s)", c.User, c.Group, c.AllowedRequest, c.Window)
}
