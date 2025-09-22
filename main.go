package main

import (
	"fmt"
	"os"
	"path/filepath"

	"cos-auto-upload/internal/config"
	"cos-auto-upload/internal/uploader"
)

// 命令用法打印
func usage() {
	fmt.Println("用法: cos-auto-upload <本地文件或目录路径> <COS目标路径>")
	fmt.Println("示例: cos-auto-upload ./dist assets/build/")
	fmt.Println("说明: COS配置从当前工作目录下的 config.json 读取")
}

func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}

	localPath := os.Args[1]
	dest := os.Args[2]

	// 规范化本地路径
	absLocal, err := filepath.Abs(localPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "解析本地路径失败: %v\n", err)
		os.Exit(1)
	}

	// 读取配置文件
	cfg, err := config.Load("config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取配置失败: %v\n", err)
		os.Exit(1)
	}

	up, err := uploader.New(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "初始化COS客户端失败: %v\n", err)
		os.Exit(1)
	}

	if err := up.Upload(absLocal, dest); err != nil {
		fmt.Fprintf(os.Stderr, "上传失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("上传完成")
}
