package kfile

import (
	"bufio"
	"context"
	"io"
	"os"
)

type Line struct {
	Data   []byte // 行内容
	Err    error  // 错误信息
	Line   int    // 行数
	Offset int    // 偏移量
	IsEOF  bool   // 时候到文件末尾
}

func ReadLine(src string, chanSize int) (<-chan Line, error) {
	return ReadLineWithContext(context.Background(), src, chanSize)
}

func ReadLineWithContext(ctx context.Context, src string, chanSize int) (<-chan Line, error) {
	if chanSize < 0 {
		panic("Channel Size is Negative.")
	}
	f, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	dataChan := make(chan Line, chanSize)
	// 线程读取
	go readLine(ctx, f, dataChan)
	return dataChan, nil
}

func ReadLineFrom(f *os.File, chanSize int) (<-chan Line, error) {
	return ReadLineFromWithContext(context.Background(), f, chanSize)
}

func ReadLineFromWithContext(ctx context.Context, f *os.File, chanSize int) (<-chan Line, error) {
	if chanSize < 0 {
		panic("Channel Size is Negative.")
	}
	dataChan := make(chan Line, chanSize)
	go readLine(ctx, f, dataChan)
	return dataChan, nil
}

// readLine 按行读取内容包含 '\n' 字符
// Offset 、 Line 都是按 os.File 当前的 offset 开始计算的，如果，原先有偏移量，要而且去累加
func readLine(ctx context.Context, f *os.File, dataChan chan Line) {
	defer close(dataChan)
	r := bufio.NewReader(f)
	var curOffset, curLine int
	for {
		select {
		case <-ctx.Done():
			break
		default:
			data, err := r.ReadBytes('\n')
			if err != nil {
				if err != io.EOF {
					dataChan <- Line{Err: err, Offset: curOffset, Line: curLine}
					return
				}
				curLine++
				curOffset += len(data)
				dataChan <- Line{Data: data, Offset: curOffset, Line: curLine, IsEOF: true}
				return
			}
			curLine++
			curOffset += len(data)
			dataChan <- Line{Data: data, Offset: curOffset, Line: curLine}
		}
	}
}
