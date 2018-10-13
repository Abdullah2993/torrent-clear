//+build !windows

package main

import (
	"errors"
)

func Install(path string) error {
	return errors.New("install only works for windows")
}
