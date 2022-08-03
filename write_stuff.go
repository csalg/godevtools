package godevtools

import (
	"path/filepath"
	"strings"
)

func AddSuffixToFilename(inputFilename, suffix string) string {
	ext := "." + strings.TrimPrefix(filepath.Ext(inputFilename), ".")
	if ext == "." {
		// There is no extension
		return inputFilename + suffix
	}
	filename := strings.TrimSuffix(inputFilename, ext)

	return filename + suffix + ext
}
