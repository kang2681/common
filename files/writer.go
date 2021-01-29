package files

import (
	"io"
	"os"
)

// Append 文件追加内容
func Append(src string, data []byte) (int, error) {
	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// AppendString 文件追加内容
func AppendString(src, data string) (int, error) {
	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.WriteString(data)
}

// AppendFrom 从 Reader 中追加文件中
func AppendFrom(src string, r io.Reader) (n int64, err error) {
	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return io.Copy(f, r)
}

// WriteFrom 从 Reader 中覆盖写入文件中
func WriteFrom(src string, r io.Reader) (n int64, err error) {
	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return io.Copy(f, r)
}
