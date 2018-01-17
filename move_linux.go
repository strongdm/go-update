// +build linux

package update

import (
	"io"
	"os"
)

func moveFile(from, to string) error {
	fromInfo, err := os.Stat(from)
	if err != nil {
		return err
	}
	fromFd, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFd.Close()

	toFd, err := os.Create(to)
	if err != nil {
		return err
	}
	defer toFd.Close()

	if _, err := io.Copy(toFd, fromFd); err != nil {
		return err
	}

	fromFd.Close()
	toFd.Close()
	os.Chmod(to, fromInfo.Mode())
	return os.Remove(from)
}
