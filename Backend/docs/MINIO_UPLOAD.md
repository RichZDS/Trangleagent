# MinIO 图片上传

## 配置

在 `manifest/config/config.yaml` 中配置 MinIO：

```yaml
minio:
  endpoint: "localhost:9000"
  accessKey: "minioadmin"
  secretKey: "minioadmin"
  bucket: "trangleagent"
  useSSL: false
```

## 启动 MinIO

```bash
docker run -d -p 9000:9000 -p 9001:9001 --name minio \
  -e MINIO_ROOT_USER=minioadmin \
  -e MINIO_ROOT_PASSWORD=minioadmin \
  minio/minio server /data --console-address ":9001"
```

控制台：http://localhost:9001

## Bucket 策略

上传的图片需可公开访问。在 MinIO 控制台为 `trangleagent` bucket 设置策略：

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {"AWS": ["*"]},
      "Action": ["s3:GetObject"],
      "Resource": ["arn:aws:s3:::trangleagent/*"]
    }
  ]
}
```

或使用 Access Rules 设置为 `public`。
