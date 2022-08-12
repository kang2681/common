package kos

import (
	"errors"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver"
	"github.com/otiai10/copy"
)

var FILE_NOT_FIND = errors.New("File Not Find")

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

// ListFiles 列出 src 根目录下的文件和目录(不包含子目录下的文件)
func ListFiles(src string) ([]os.FileInfo, error) {
	fi, err := os.Stat(src)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return []os.FileInfo{fi}, nil
	}
	return ioutil.ReadDir(src)
}

type FilesTree struct {
	Files []os.FileInfo // 文件
	Dirs  []*FilesTree  // 目录
}

// ListFilesTree 列出目录树
func ListFilesTree(src string) (*FilesTree, error) {
	list, err := ListFiles(src)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	rs := &FilesTree{}
	for _, v := range list {
		if v.IsDir() {
			rs.Files = append(rs.Files, v)
			continue
		}
		dirTree, err := ListFilesTree(filepath.Join(src, v.Name()))
		if err != nil {
			return nil, err
		}
		rs.Dirs = append(rs.Dirs, dirTree)
	}
	return rs, nil
}

// 文件查找
func SearchFile(src string, keyword string) (os.FileInfo, error) {
	list, err := ListFiles(src)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		if v.IsDir() {
			continue
		}
		if strings.Contains(src, keyword) {
			return v, nil
		}
	}
	return nil, FILE_NOT_FIND
}

//func SearchFileTree(src string, keyword string) ([]os.FileInfo, error) {
//	list, err := ListFiles(src)
//	if err != nil {
//		return nil, err
//	}
//	for _, v := range list {
//		if v.IsDir() {
//			continue
//		}
//		if strings.Contains(src, keyword) {
//			return &v, nil
//		}
//	}
//	return nil, FILE_NOT_FIND
//}

func SearchFileCallback(dir string, fn func(d fs.DirEntry) error) []fs.DirEntry {
	if !IsDir(dir) {
		return nil
	}
	arr := make([]fs.DirEntry, 0, 100)
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return fs.SkipDir
		}
		if err := fn(d); err != nil {
			return err
		}
		arr = append(arr, d)
		return nil
	})
	return arr
}
