package common

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"unique"

	"github.com/kamijoucen/notesync/pkg/definition"
)

// ScanLocalFiles 扫描本地仓库文件
// key 文件唯一标识
func ScanLocalFiles(basePath string) (map[string]*definition.FileTreeNode, error) {
	files := make(map[string]*definition.FileTreeNode)
	if err := scanfiles1(basePath, files); err != nil {
		return nil, err
	}
	return files, nil
}

func scanfiles1(dirpath string, files map[string]*definition.FileTreeNode) error {
	fInfo, err := os.Stat(dirpath)
	if err != nil {
		return err
	}
	if !fInfo.IsDir() {
		return errors.New(dirpath + " is not a directory")
	}
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			if err := scanfiles1(dirpath+"/"+entry.Name(), files); err != nil {
				return err
			}
		} else {
			fileHash, err := GetFileHash(path.Join(dirpath, entry.Name()))
			if err != nil {
				return err
			}
			fileFlag := dirpath + "/" + entry.Name()
			files[fileFlag] = &definition.FileTreeNode{
				Flag:    unique.Make(fileFlag),
				Name:    entry.Name(),
				IsDir:   false,
				Hash:    fileHash,
				Path:    dirpath,
				PathArr: []string{dirpath}, // TODO
			}
		}
	}
	return nil
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
