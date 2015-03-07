package cloudElements

import "fmt"

// Credentials stores all token information to authenticate to
// Cloud Elements, as well as each element
type Credentials struct {
	Elements     map[int]string
	Organization string
	User         string // User Secret
}

// Authorization is formatted as:
// User <secret>, Organization <token>, Element <token>
func (c Credentials) Authorization(element int) string {
	return fmt.Sprintf(`User %s, Organization %s, Element %s`, c.User, c.Organization, c.Elements[element])
}
