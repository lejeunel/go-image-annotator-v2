package shared

import (
	"net/url"
)

// URLWithQuery returns a copy of u with the given key/value set
func URLWithQuery(u url.URL, key, value string) url.URL {
	copy := u
	q := copy.Query()          // get a copy of the query
	q.Set(key, value)          // set or replace the param
	copy.RawQuery = q.Encode() // assign back
	return copy
}
