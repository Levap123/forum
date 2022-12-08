package errors

import "fmt"

func Fail(err error, place string) error {
	return fmt.Errorf("%s: %v", place, err)
}
