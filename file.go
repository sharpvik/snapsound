package main

import (
	"io"
	"os"
	"path/filepath"
)

func readFile(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func saveFile(name string, bytes []byte) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(bytes)
	return err
}

func trimExtension(name string) string {
	return name[:len(name)-len(filepath.Ext(name))]
}
