package video

import (
	"strconv"
	"strings"
)

// 解析 Range 请求头，返回起始和结束位置
func parseRange(rangeHeader string, fileSize int64) (int64, int64, error) {
	// Range 请求头的格式为 "bytes=start-end"
	rangeParts := strings.SplitN(rangeHeader, "=", 2)
	byteRange := strings.SplitN(rangeParts[1], "-", 2)

	// 解析起始位置
	start, err := strconv.ParseInt(byteRange[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	var end int64
	// 如果请求头中有指定结束位置，则解析结束位置
	if byteRange[1] != "" {
		end, err = strconv.ParseInt(byteRange[1], 10, 64)
		if err != nil {
			return 0, 0, err
		}
	} else {
		// 如果请求头中没有指定结束位置，则设置结束位置为文件大小减一
		end = fileSize - 1
	}

	return start, end, nil
}
