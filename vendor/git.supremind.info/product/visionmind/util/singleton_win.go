// +build windows

package util

import (
	"fmt"
	"os"
	"time"
)

// Flock holds the lock data
type Flock struct {
	lockHolder *os.File
	lockFile   string
	lockTime   time.Time
}

// NewFlock create a Flock
func NewFlock(file string) (f *Flock, err error) {
	err = fmt.Errorf("not implemented")
	return
}
