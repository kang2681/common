package kfile

import (
	"bytes"
	"net/http"
	"strings"
)

// GetCsvData获取Csv格式数据
func WriteCSV(data [][]string) []byte {
	buf := bytes.NewBufferString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	//添加数据行
	for _, dataArray := range data {
		buf.WriteString(strings.Join(dataArray, ",") + "\n")
	}
	return buf.Bytes()
}

// Export导出文件
func WriteHttpCSV(w http.ResponseWriter, fileName string, data []byte) {
	w.Header().Set("Content-Description", "File Transfer")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "No-cache")
	w.Write(data)
}
