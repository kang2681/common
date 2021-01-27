package funcs

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// IsFile
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// PathExists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// IsDir
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// FileCopy
// src 文件不存在 将报错
// 如果 dst 文件存在，将被覆盖
// 如果 dst 文件不存在 文件会自动创建
// 如果 dst 目录不存在 目录会自动创建 权限为 0777
func FileCopy(src, dst string) (int64, error) {
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
func DirCopy(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, v := range files {
		if v.IsDir() {
			continue
		}
		if _, err := FileCopy(filepath.Join(src, v.Name()), filepath.Join(dst, v.Name())); err != nil {
			return err
		}
	}
	return nil
}

// DirCopyTree 文件夹复制
// 递归复制 src 目录和文件
func DirCopyTree(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, v := range files {
		if v.IsDir() {
			if err := DirCopy(filepath.Join(src, v.Name()), filepath.Join(dst, v.Name())); err != nil {
				return err
			}
			continue
		}
		if _, err := FileCopy(filepath.Join(src, v.Name()), filepath.Join(dst, v.Name())); err != nil {
			return err
		}
	}
	return nil
}

func FileAppendString(src, data string) (int, error) {
	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.WriteString(data)
}

func FileAppend(src string, data []byte) (int, error) {
	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}
func FileAppendFrom(src string, r io.Reader) (n int, err error) {
	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		m, e := r.Read(buf)
		if m < 0 {
			return n, fmt.Errorf("reader returned negative count from Read")
		}
		if m < 0 {
			s, err := f.Write(buf[0:m])
			if err != nil {
				return n, err
			}
			if s != m {
				return n, fmt.Errorf("file writer returned write count not equal read count")
			}
		}
		n = n + m
		if e == io.EOF {
			return n, nil
		}
		if e != nil {
			return n, e
		}
	}
}

// ZipDeCompress zip解压
func ZipDeCompress(zipFile, destDir string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if err := doZipDeCompress(file, destDir); err != nil {
			return fmt.Errorf("unzip Error %w", err)
		}
	}
	return nil
}

func doZipDeCompress(file *zip.File, destDir string) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	if file.FileInfo().IsDir() {
		return nil
	}
	filename := filepath.Join(destDir, file.Name)
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return fmt.Errorf("Mkdir ALL %s Error %w,filename:%s", filepath.Dir(filename), err, filename)
	}
	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = io.Copy(w, rc)
	if err != nil {
		return err
	}
	return nil
}
