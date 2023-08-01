package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	http.HandleFunc("/clip", clipVideo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func clipVideo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 解析请求参数
	url := r.FormValue("url")
	startTime := r.FormValue("start_time")
	endTime := r.FormValue("end_time")

	// 检查参数是否有效
	if url == "" || startTime == "" || endTime == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "缺少必要参数")
		return
	}

	// 调用FFmpeg进行视频剪辑
	outputFile, err := clipWithFFmpeg(url, startTime, endTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "剪辑视频时发生错误: %s", err)
		return
	}

	// 返回剪辑后视频的URL
	downloadURL := generateDownloadURL(outputFile)
	fmt.Fprint(w, downloadURL)
}

func clipWithFFmpeg(url, startTime, endTime string) (string, error) {
	// 生成唯一的输出文件名
	outputFileName := fmt.Sprintf("%d_clip.mp4", time.Now().UnixNano())

	// 构建FFmpeg命令
	cmd := exec.Command("ffmpeg",
		"-i", url,
		"-ss", startTime,
		"-to", endTime,
		"-c", "copy",
		outputFileName,
	)

	// 执行命令
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// 返回输出文件路径
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	outputFile := filepath.Join(exeDir, outputFileName)
	return outputFile, nil
}

func generateDownloadURL(filePath string) string {
	// 假设服务器上的视频文件都在 /videos 目录下
	baseURL := "https://example.com/videos/"
	fileName := filepath.Base(filePath)
	downloadURL := baseURL + fileName
	return downloadURL
}
