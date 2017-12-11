package utils

import (
	"fmt"

	errors "github.com/go-errors/errors"
)

// Panic panic with error stack
func Panic(err error) {
	panic(errors.Wrap(err.Error()+"\n"+string(errors.Wrap(err, 1).Stack()), 1))
}

// Panicf formats according to a format specifier and panic the err
func Panicf(format string, i ...interface{}) {
	err := fmt.Errorf(format, i...)
	panic(errors.Wrap(err.Error()+"\n"+string(errors.Wrap(err, 1).Stack()), 1))
}
