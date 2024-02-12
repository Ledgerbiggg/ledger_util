package video

import (
	"net/http"
	"os"
	"strconv"
	"testing"
)

func TestVideo(t *testing.T) {
	http.HandleFunc("/api/video", func(w http.ResponseWriter, r *http.Request) {
		// 打开视频文件
		videoFile, err := os.Open("videos/video2.mp4")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer videoFile.Close()

		// 获取文件信息
		fileInfo, err := videoFile.Stat()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 设置响应头，指定文件内容类型
		//w.Header().Set("Content-Type", "video/mp4")
		w.Header().Set("Content-Type", "video/mp4")

		// 设置响应头，指定文件总长度
		w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

		// 处理 Range 请求头
		rangeHeader := r.Header.Get("Range")
		if rangeHeader != "" {
			// 解析 Range 请求头，获取起始和结束位置
			start, _, err := parseRange(rangeHeader, fileInfo.Size())
			if err != nil {
				http.Error(w, "Invalid range request", http.StatusRequestedRangeNotSatisfiable)
				return
			}
			// 设置响应状态码为 206 Partial Content
			//w.WriteHeader(206)

			// 设置读取文件的起始位置
			videoFile.Seek(start, 0)

			// 发送部分内容给客户端
			http.ServeContent(w, r, "", fileInfo.ModTime(), videoFile)
		} else {
			// 如果没有 Range 请求头，直接发送整个文件内容
			http.ServeContent(w, r, "", fileInfo.ModTime(), videoFile)
		}
	})
	http.ListenAndServe(":9999", nil)
}

//// 解析 Range 请求头，返回起始和结束位置
//func parseRange(rangeHeader string, fileSize int64) (int64, int64, error) {
//	// Range 请求头的格式为 "bytes=start-end"
//	rangeParts := strings.SplitN(rangeHeader, "=", 2)
//	byteRange := strings.SplitN(rangeParts[1], "-", 2)
//
//	// 解析起始位置
//	start, err := strconv.ParseInt(byteRange[0], 10, 64)
//	if err != nil {
//		return 0, 0, err
//	}
//
//	var end int64
//	// 如果请求头中有指定结束位置，则解析结束位置
//	if byteRange[1] != "" {
//		end, err = strconv.ParseInt(byteRange[1], 10, 64)
//		if err != nil {
//			return 0, 0, err
//		}
//	} else {
//		// 如果请求头中没有指定结束位置，则设置结束位置为文件大小减一
//		end = fileSize - 1
//	}
//
//	return start, end, nil
//}
