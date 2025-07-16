package fs

import (
	"errors"
	"fmt"
	"os"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		// some fs or permission error
		return false, fmt.Errorf("failed to check for file %s: %w", path, err)
	}
}
