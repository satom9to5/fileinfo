// +build windows

package fileinfo

import (
	"os"
	"strings"
	"syscall"
)

const (
	PathSep = "\\"
)

func DirBase(absPath string) (string, string) {
	absPath = strings.Replace(absPath, PathSep, "/", -1)

	return path.Dir(absPath), path.Base(absPath)
}

func SplitPath(absPath, rootPath string) []string {
	str := strings.Replace(strings.Replace(absPath, rootPath, "", -1), PathSep, "/", -1)

	if strings.Index(str, "/") == 0 {
		return strings.Split(str, "/")[1:]
	} else {
		return strings.Split(str, "/")
	}
}

func (fi *FileInfo) getIno() (uint64, error) {
	h, err := syscall.CreateFile(syscall.StringToUTF16Ptr(fi.Path()),
		syscall.FILE_LIST_DIRECTORY,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE,
		nil, syscall.OPEN_EXISTING,
		syscall.FILE_FLAG_BACKUP_SEMANTICS|syscall.FILE_FLAG_OVERLAPPED, 0)

	if err != nil {
		return 0, err
	}

	var hfi syscall.ByHandleFileInformation
	if err = syscall.GetFileInformationByHandle(h, &hfi); err != nil {
		syscall.CloseHandle(h)
		return 0, err
	}

	return uint64(hfi.FileIndexHigh)<<32 | uint64(fi.FileIndexLow), nil
}
