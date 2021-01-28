package osext

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
	"github.com/otiai10/copy"
)

// CopyFile
// src 文件不存在 将报错
// 如果 dst 文件存在，将被覆盖
// 如果 dst 文件不存在 文件会自动创建
// 如果 dst 目录不存在 目录会自动创建 权限为 0777
func CopyFile(src, dst string) (int64, error) {
	f, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	// 创建dst 文件目录
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return 0, err
	}
	rf, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer rf.Close()
	return io.Copy(rf, f)
}

// DirCopy 文件夹复制
// 只复制 src 根目录下的文件，子目录下的文件不会复制
func CopyDir(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, v := range files {
		if v.IsDir() {
			continue
		}
		if _, err := CopyFile(filepath.Join(src, v.Name()), filepath.Join(dst, v.Name())); err != nil {
			return err
		}
	}
	return nil
}

// DirCopyTree 文件夹复制
// 递归复制 src 目录和文件
func CopyDirTree(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, v := range files {
		if v.IsDir() {
			if err := CopyDir(filepath.Join(src, v.Name()), filepath.Join(dst, v.Name())); err != nil {
				return err
			}
		} else {
			if _, err := CopyFile(filepath.Join(src, v.Name()), filepath.Join(dst, v.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}

// Copy 文件或目录复制
// 自动识别
func Copy(src, dst string, opt copy.Options) error {
	return copy.Copy(src, dst, opt)
}

// Archive 压缩
func Archive(sources []string, destination string) error {
	return archiver.Archive(sources, destination)
}

// Unarchive 解压
func Unarchive(sources, destination string) error {
	return archiver.Unarchive(sources, destination)
}

//
//func FileAppendString(src, data string) (int, error) {
//	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		return 0, err
//	}
//	defer f.Close()
//	return f.WriteString(data)
//}
//
//func FileAppend(src string, data []byte) (int, error) {
//	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		return 0, err
//	}
//	defer f.Close()
//	return f.Write(data)
//}
//func FileAppendFrom(src string, r io.Reader) (n int, err error) {
//	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		return 0, err
//	}
//	defer f.Close()
//	buf := make([]byte, 1024)
//	for {
//		m, e := r.Read(buf)
//		if m < 0 {
//			return n, fmt.Errorf("reader returned negative count from Read")
//		}
//		if m < 0 {
//			s, err := f.Write(buf[0:m])
//			if err != nil {
//				return n, err
//			}
//			if s != m {
//				return n, fmt.Errorf("file writer returned write count not equal read count")
//			}
//		}
//		n = n + m
//		if e == io.EOF {
//			return n, nil
//		}
//		if e != nil {
//			return n, e
//		}
//	}
//}
