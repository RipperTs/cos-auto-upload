package uploader

import (
	"context"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"

	"cos-auto-upload/internal/config"
	"cos-auto-upload/internal/cosclient"
)

type Uploader struct {
	client *cos.Client
}

func New(cfg *config.Config) (*Uploader, error) {
	c, err := cosclient.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Uploader{client: c}, nil
}

// Upload 将本地文件或目录上传到COS目标路径（前缀或对象键）
// - 当 localPath 是文件：
//   - 若 dest 以 "/" 结尾，则对象键为 dest + basename(localPath)
//   - 否则对象键为 dest
// - 当 localPath 是目录：
//   - 递归上传目录下所有文件，对象键为 dest(作为前缀, 自动补尾部斜杠) + 相对路径
func (u *Uploader) Upload(localPath, dest string) error {
	info, err := os.Stat(localPath)
	if err != nil {
		return fmt.Errorf("读取本地路径信息失败: %w", err)
	}

	// 规范化目标键/前缀：去除开头的 '/'
	if strings.HasPrefix(dest, "/") {
		dest = strings.TrimLeft(dest, "/")
	}

	if info.IsDir() {
		return u.uploadDir(localPath, dest)
	}
	return u.uploadSingleFile(localPath, dest)
}

func (u *Uploader) uploadDir(dir, dest string) error {
	// 目标前缀标准化，确保以 "/" 结尾
	prefix := dest
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	root := filepath.Clean(dir)
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return fmt.Errorf("计算相对路径失败: %w", err)
		}

		// 将本地路径分隔符统一为 COS 的 "/"
		key := prefix + toSlash(rel)
		if err := u.putObject(path, key); err != nil {
			return err
		}
		fmt.Printf("上传成功: %s -> %s\n", path, key)
		return nil
	})
}

func (u *Uploader) uploadSingleFile(filePath, dest string) error {
	key := dest
	if strings.HasSuffix(dest, "/") {
		key = dest + filepath.Base(filePath)
	}
	if err := u.putObject(filePath, key); err != nil {
		return err
	}
	fmt.Printf("上传成功: %s -> %s\n", filePath, key)
	return nil
}

func (u *Uploader) putObject(filePath, key string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	defer f.Close()

	// 依据扩展名推断 Content-Type
	contentType := mime.TypeByExtension(strings.ToLower(filepath.Ext(filePath)))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err = u.client.Object.Put(context.Background(), key, f, &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: contentType,
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{},
	})
	if err != nil {
		return fmt.Errorf("上传对象失败(%s): %w", key, err)
	}
	return nil
}

func toSlash(p string) string {
	return strings.ReplaceAll(p, string(filepath.Separator), "/")
}
