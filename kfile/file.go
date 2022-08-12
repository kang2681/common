package kfile

import (
	"os"
	"time"
)

// ModTime 获取文件修改时间
func ModTime(src string) (modTime time.Time, size int64, err error) {
	var fi os.FileInfo
	fi, err = os.Stat(src)
	if err != nil {
		return
	}
	modTime = fi.ModTime()
	size = fi.Size()
	return
}
