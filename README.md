# cos-auto-upload

将指定文件/目录上传到腾讯云 COS 的指定路径（前缀）。

## 配置

在可执行文件所在工作目录准备 `config.json`（可参考 `config.example.json`）：

```json
{
  "secret_id": "AKID...",
  "secret_key": "...",
  "bucket": "<bucket-name-appid>",
  "region": "ap-guangzhou",
  "base_url": ""  
}
```

- `base_url` 可选：默认使用 `https://<bucket>.cos.<region>.myqcloud.com`。

## 用法

本地运行：

```
go run . <本地文件或目录> <COS目标路径>
```

构建二进制：

```
go build -o bin/cos-auto-upload .
```

使用 Makefile（推荐，一键编译）：

```
make            # 默认编译所有常见平台

# 指定平台矩阵（覆盖默认）
make build-all PLATFORMS="linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64"

# 单平台（仅当前环境）
make build      # 或覆盖：make build OS=linux ARCH=amd64
```

运行二进制：

```
./bin/cos-auto-upload <本地文件或目录> <COS目标路径>
```

示例：

- 上传单个文件到 `assets/logo.png`
  ```
  ./bin/cos-auto-upload ./logo.png assets/logo.png
  ```
- 上传目录（递归）到前缀 `assets/build/`，会保留相对路径结构
  ```
  ./bin/cos-auto-upload ./dist assets/build/
  ```

## 开发

- 格式化与静态检查：`go fmt ./... && go vet ./...`
- 单元测试（如有）：`go test ./...`
