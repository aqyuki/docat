package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func TargetDirectoryParser(args []string) (string, error) {
	if len(args) > 1 {
		return "", fmt.Errorf("too many args")
	}

	targetDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if len(args) == 1 {
		dir, err := filepath.Abs(args[0])
		if err != nil {
			return "", err
		}
		targetDir = dir
	}

	return targetDir, nil
}
