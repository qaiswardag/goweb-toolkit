package toolkit

import (
	"os"
)

func (t *Tools) EnsureDirCreated(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return os.Mkdir(dirName, os.ModePerm)
	}
	return nil
}
