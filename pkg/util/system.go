package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func ValidateContext(dirPath, componentName string) error {
	if err := validateDir(dirPath); err != nil {
		return err
	}
	amd64BinPath := filepath.Join(dirPath, fmt.Sprintf("bin/linux/amd64/%s", componentName))
	if err := validateBin(amd64BinPath); err != nil {
		return err
	}
	return nil
}

func validateDir(dirPath string) error {
	if !filepath.IsAbs(dirPath) {
		return fmt.Errorf("path %s is not absolute", dirPath)
	}
	if _, err := os.Stat(dirPath); err != nil {
		return err
	}
	return nil
}

func validateBin(componentBinPath string) error {
	if _, err := os.Stat(componentBinPath); err != nil {
		return err
	}
	return nil
}
