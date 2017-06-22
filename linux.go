// +build linux

package fileinfo

import (
	"errors"
	"syscall"
)

func (fi *FileInfo) getIno() (uint64, error) {
	if st, ok := fi.fi.Sys().(*syscall.Stat_t); ok {
		return st.Ino, nil
	} else {
		return 0, errors.New("get inode failed.")
	}
}
