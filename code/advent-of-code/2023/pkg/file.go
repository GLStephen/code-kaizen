package pkg

import (
	"os"
)

func OpenFile(loadLocation string) (*os.File, error) {
	file, err := os.Open(loadLocation)
	if err != nil {
		return nil, err
	}

	return file, nil
}
