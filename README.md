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
