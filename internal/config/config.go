package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config 定义COS配置，由根目录下 config.json 提供
type Config struct {
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
	// 可选：自定义域名（含协议），如 https://cdn.example.com
	// 若不提供，将使用 https://<bucket>.cos.<region>.myqcloud.com
	BaseURL string `json:"base_url"`
}

// Load 从给定路径读取并解析配置
func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("打开配置文件失败: %w", err)
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	if cfg.SecretID == "" || cfg.SecretKey == "" || cfg.Bucket == "" || cfg.Region == "" {
		return nil, fmt.Errorf("配置不完整: 需要 secret_id、secret_key、bucket、region")
	}
	return &cfg, nil
}
