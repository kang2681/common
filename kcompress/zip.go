package kcompress

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unzip() {

}

// ZipDeCompress zip解压
func ZipDeCompress(zipFile, destDir string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if err := doZipDeCompress(file, destDir); err != nil {
			return fmt.Errorf("unzip Error %w", err)
		}
	}
	return nil
}

func doZipDeCompress(file *zip.File, destDir string) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	if file.FileInfo().IsDir() {
		return nil
	}
	filename := filepath.Join(destDir, file.Name)
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return fmt.Errorf("Mkdir ALL %s Error %w,filename:%s", filepath.Dir(filename), err, filename)
	}
	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = io.Copy(w, rc)
	if err != nil {
		return err
	}
	return nil
}

func Zip() {

}
