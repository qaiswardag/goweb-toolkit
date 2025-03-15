package toolkit

import (
	"errors"
	"os"
)

func (t *Tools) EnsureDirCreated(dirName string) error {
	if _, err := os.Stat(dirName); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(dirName, 0777)
		}
		return err
	}
	return errors.New("directory already exists")
}
