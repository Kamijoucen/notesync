package common

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/kamijoucen/notesync/pkg/definition"
)

func ScanLocalFiles(basePath string) ([]*definition.FileTreeNode, error) {
	fInfo, err := os.Stat(basePath)
	if err != nil {
		return nil, err
	}
	if !fInfo.IsDir() {
		return nil, errors.New(basePath + " is not a directory")
	}
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}
	var nodes []*definition.FileTreeNode
	for _, entry := range entries {
		if entry.IsDir() {
			subNodes, err := ScanLocalFiles(basePath + "/" + entry.Name())
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, &definition.FileTreeNode{
				Name:     entry.Name(),
				IsDir:    true,
				Hash:     "",
				Children: subNodes,
			})
		} else {
			fHash, err := GetFileHash(basePath + "/" + entry.Name())
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, &definition.FileTreeNode{
				Name:  entry.Name(),
				Hash:  fHash,
				IsDir: false,
			})
		}
	}
	return nil, nil
}

func GetFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	hashes := sha256.New()
	if _, err := io.Copy(hashes, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hashes.Sum(nil)), nil
}
