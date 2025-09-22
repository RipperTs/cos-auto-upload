package cosclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"

	"cos-auto-upload/internal/config"
)

// New 创建 COS Client
func New(cfg *config.Config) (*cos.Client, error) {
	var base string
	if cfg.BaseURL != "" {
		base = cfg.BaseURL
	} else {
		base = fmt.Sprintf("https://%s.cos.%s.myqcloud.com", cfg.Bucket, cfg.Region)
	}

	u, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("解析COS基础URL失败: %w", err)
	}

	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.SecretID,
			SecretKey: cfg.SecretKey,
		},
	})

	return client, nil
}
