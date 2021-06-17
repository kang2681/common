package osext

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"
)

type CmdContentResp struct {
	Level string // 消息级别
	Msg   string // 消息内容
	Err   error  // 错误
}

func CmdPipeline(cmd *exec.Cmd, dataChan chan CmdContentResp) {
	defer close(dataChan)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		dataChan <- CmdContentResp{Level: "error", Msg: fmt.Sprintf("stdout pipe error %s", err.Error()), Err: err}
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		dataChan <- CmdContentResp{Level: "error", Msg: fmt.Sprintf("stderr pipe error %s", err.Error()), Err: err}
		return
	}
	if err := cmd.Start(); err != nil {
		dataChan <- CmdContentResp{Level: "error", Msg: fmt.Sprintf("cmd start error %s", err.Error()), Err: err}
		return
	}
	var readContent = func(wg *sync.WaitGroup, reader io.Reader, level string, dataChan chan CmdContentResp) {
		defer wg.Done()
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			dataChan <- CmdContentResp{Level: level, Msg: scanner.Text(), Err: nil}
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go readContent(&wg, stdout, "info", dataChan)
	wg.Add(1)
	go readContent(&wg, stderr, "error", dataChan)

	if err := cmd.Wait(); err != nil {
		dataChan <- CmdContentResp{Level: "error", Msg: fmt.Sprintf("cmd wait error %s", err.Error()), Err: err}
		return
	}
	wg.Wait()
}
