package ziphelper

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"archive/zip"
)

func CreateZip(files ...FileEntry) (string, error) {
	// Check entries before start creating zip file
	for i, f := range files {
		if len(f.SrcFile) == 0 {
			return "", fmt.Errorf("no src file at file %d", i)
		}
		if len(f.Filename) == 0 {
			return "", fmt.Errorf("no filename at file %d", i)
		}
	}

	zipFile, err := ioutil.TempFile("", "zipped")
	if err != nil {
		return "", fmt.Errorf("temp file error: %w", err)
	}

	fzip, err := os.Create(zipFile.Name())
	if err != nil {
		return "", fmt.Errorf("zip file create error: %w", err)
	}

	zipw := zip.NewWriter(fzip)
	defer zipw.Close()

	for _, entry := range files {
		srcFileReader, err := os.Open(entry.SrcFile)
		if err != nil {
			return "", fmt.Errorf("src file open error: %w", err)
		}
		defer srcFileReader.Close()

		w, err := zipw.Create(entry.Filename)
		if err != nil {
			return "", fmt.Errorf("zip writer create error: %w", err)
		}
		if _, err = io.Copy(w, srcFileReader); err != nil {
			return "", fmt.Errorf("zip copy error: %w", err)
		}
	}

	return zipFile.Name(), nil
}
