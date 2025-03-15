package toolkit

import (
	"errors"
	"os"
)

func (t *Tools) EnsureDirCreated(dirName string, perm os.FileMode) error {
	if _, err := os.Stat(dirName); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(dirName, perm)
		}
		return err
	}
	return errors.New("directory already exists")
}
