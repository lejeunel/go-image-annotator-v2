package cli

import (
	"fmt"
)

type ErrorPresenter struct{}

func (p ErrorPresenter) Error(err error) {
	fmt.Println(err.Error())
}
