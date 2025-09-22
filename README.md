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
> 确保 `config.json` 与二进制文件在同一目录下.

```bash
# 上传单个文件
./cos-auto-upload /path/to/file.txt [remote/prefix/]

# 上传目录（包含子目录） (注意最外层的目录名不会被包含在远程路径中)
./cos-auto-upload /path/to/directory/ [remote/prefix/]
```

## 构建

使用 Makefile（推荐）：

```bash
# 默认：编译所有常见平台产物到 bin/
make

# 指定平台矩阵
make build-all PLATFORMS="linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64"

# 单平台（仅当前环境）
make build                 # 输出 bin/cos-auto-upload
# 或覆盖目标平台：
make build OS=linux ARCH=amd64
```

产物命名：
- 多平台：`bin/cos-auto-upload-<os>-<arch>[.exe]`
- 例如：
  - `bin/cos-auto-upload-linux-amd64`
  - `bin/cos-auto-upload-darwin-arm64`
  - `bin/cos-auto-upload-windows-amd64.exe`
  - 单平台（当前环境）：`bin/cos-auto-upload`
