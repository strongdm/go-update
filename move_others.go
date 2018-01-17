// +build !linux

package update

import "os"

func moveFile(from, to string) error {
	return os.Rename(from, to)
}
