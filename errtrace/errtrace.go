package errtrace

import (
	"fmt"

	"braces.dev/errtrace"
)

func Wrap(err error) error {
	return errtrace.Wrap(err)
}

func Print(err error) {
	fmt.Printf("%+v", err)
}
