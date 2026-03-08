package testing

import (
	"github.com/go-test/deep"
	"testing"
)

func AssertDeepEqual(t testing.TB, this any, that any, structName string) {
	diff := deep.Equal(this, that)
	if diff != nil {
		t.Fatalf("expected to retrieve identical %v structs, but got different fields: %v",
			structName, diff)
	}

}
