package fileinfo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	PathSep = string(os.PathSeparator)
)

// os.FileInfo + inode
type FileInfo struct {
	fi  os.FileInfo
	dir string
	ino uint64 // inode
}

func Stat(name string) (*FileInfo, error) {
	info := &FileInfo{}

	fi, err := os.Stat(name)
	if err != nil {
		return nil, err
	}

	info.fi = fi
	info.dir = filepath.Dir(name)

	if ino, err := info.getIno(); err == nil {
		info.ino = ino
	}

	return info, nil
}

func IsDir(name string) bool {
	fi, err := Stat(name)
	if err != nil {
		return false
	}

	return fi.IsDir()
}

func Split(absPath string) (dir, file string) {
	return filepath.Dir(absPath), filepath.Base(absPath)
}

func SplitPath(absPath, rootPath string) []string {
	str := strings.Replace(absPath, rootPath, "", -1)

	if strings.Index(str, PathSep) == 0 {
		return strings.Split(str, PathSep)[1:]
	} else {
		return strings.Split(str, PathSep)
	}
}

func (fi FileInfo) String() string {
	t := "file"
	if fi.IsDir() {
		t = "directory"
	}

	return fmt.Sprintf("Name:%s, Size:%d, Dir:%s, Ino:%d, Type:%s", fi.Name(), fi.Size(), fi.Dir(), fi.Ino(), t)
}

func (fi *FileInfo) Name() string {
	return fi.fi.Name()
}

func (fi *FileInfo) Size() int64 {
	return fi.fi.Size()
}

func (fi *FileInfo) IsDir() bool {
	return fi.fi.IsDir()
}

func (fi *FileInfo) Dir() string {
	return fi.dir
}

func (fi *FileInfo) Path() string {
	return fi.dir + PathSep + fi.Name()
}

func (fi *FileInfo) Ino() uint64 {
	return fi.ino
}
