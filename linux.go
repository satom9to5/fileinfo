// +build linux

package fileinfo

import (
	"errors"
	"path"
	"strings"
	"syscall"
)

const (
	PathSep = "/"
)

func DirBase(absPath string) (string, string) {
	return path.Dir(absPath), path.Base(absPath)
}

func SplitPath(absPath, rootPath string) []string {
	str := strings.Replace(absPath, rootPath, "", -1)

	if strings.Index(str, "/") == 0 {
		return strings.Split(str, "/")[1:]
	} else {
		return strings.Split(str, "/")
	}
}

func (fi *FileInfo) getIno() (uint64, error) {
	if st, ok := fi.fi.Sys().(*syscall.Stat_t); ok {
		return st.Ino, nil
	} else {
		return 0, errors.New("get inode failed.")
	}
}
