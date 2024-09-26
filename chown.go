//go:build !linux
// +build !linux

package logWriter

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
