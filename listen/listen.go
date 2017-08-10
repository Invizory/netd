// Package listen provides convenient ways to create and manage listeners.
package listen

import (
	"errors"
)

func newError(text string) error {
	return errors.New("listen: " + text)
}
