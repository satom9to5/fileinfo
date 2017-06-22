// +build windows

package fileinfo

import (
	"syscall"
)

func (fi *FileInfo) getIno() (uint64, error) {
	h, err := syscall.CreateFile(syscall.StringToUTF16Ptr(fi.Path()),
		syscall.FILE_LIST_DIRECTORY,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE,
		nil, syscall.OPEN_EXISTING,
		syscall.FILE_FLAG_BACKUP_SEMANTICS|syscall.FILE_FLAG_OVERLAPPED, 0)

	defer syscall.CloseHandle(h)

	if err != nil {
		return 0, err
	}

	var hfi syscall.ByHandleFileInformation
	if err = syscall.GetFileInformationByHandle(h, &hfi); err != nil {
		return 0, err
	}

	return uint64(hfi.FileIndexHigh)<<32 | uint64(hfi.FileIndexLow), nil
}
