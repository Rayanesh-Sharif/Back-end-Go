package util

import (
	"github.com/go-faster/errors"
	"os"
	"path"
)

// CreateFolder will create a folder if not exists.
// The arguments are combined to create a path.
func CreateFolder(folderPath ...string) error {
	finalPath := path.Join(folderPath...)
	if _, err := os.Stat(finalPath); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(finalPath, 0666)
			if err != nil {
				return errors.Wrap(err, "cannot create folder")
			}
		} else {
			return errors.Wrap(err, "cannot check if folder exists")
		}
	}
	return nil
}
