package ziphelper

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alexmullins/zip"
)

func CreatePasswordProtectedZip(password string, files ...FileEntry) (string, error) {
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

		w, err := zipw.Encrypt(entry.Filename, password)
		if err != nil {
			return "", fmt.Errorf("zip encrypt error: %w", err)
		}
		if _, err = io.Copy(w, srcFileReader); err != nil {
			return "", fmt.Errorf("zip copy error: %w", err)
		}
		if err = zipw.Flush(); err != nil {
			return "", fmt.Errorf("zip flush error: %w", err)
		}
	}

	return zipFile.Name(), nil
}

func UnzipPasswordProtectedZip(password string, srcFile string, dstDir string) ([]string, error) {
	r, err := zip.OpenReader(srcFile)
	if err != nil {
		return nil, fmt.Errorf("open file error: %w", err)
	}
	defer r.Close()

	unzippedFiles := []string{}
	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(password)
		}
		r, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("extract error at %s: %w", f.Name, err)
		}
		buf, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, fmt.Errorf("zip entry read error at %s: %w", f.Name, err)
		}
		defer r.Close()

		outPath := filepath.Join(dstDir, f.Name)
		err = ioutil.WriteFile(outPath, buf, 0644)
		if err != nil {
			return nil, fmt.Errorf("unzip write error at %s: %w", f.Name, err)
		}
		unzippedFiles = append(unzippedFiles, outPath)
	}
	return unzippedFiles, nil
}
