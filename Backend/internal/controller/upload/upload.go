package upload

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// UploadImage 上传图片到 MinIO
func UploadImage(r *ghttp.Request) {
	file, header, err := r.Request.FormFile("file")
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": "请选择要上传的图片",
		})
		return
	}
	defer file.Close()

	// 校验文件类型
	ext := filepath.Ext(header.Filename)
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	if !allowed[ext] {
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": "仅支持 jpg/png/gif/webp 格式",
		})
		return
	}
	if header.Size > 5*1024*1024 { // 5MB
		r.Response.WriteJsonExit(g.Map{
			"code":    400,
			"message": "图片大小不能超过 5MB",
		})
		return
	}

	ctx := context.Background()
	endpoint := g.Cfg().MustGetWithEnv(ctx, "minio.endpoint", "localhost:9000").String()
	accessKey := g.Cfg().MustGetWithEnv(ctx, "minio.accessKey", "minioadmin").String()
	secretKey := g.Cfg().MustGetWithEnv(ctx, "minio.secretKey", "minioadmin").String()
	bucket := g.Cfg().MustGetWithEnv(ctx, "minio.bucket", "trangleagent").String()
	useSSL := g.Cfg().MustGetWithEnv(ctx, "minio.useSSL", false).Bool()

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		g.Log().Error(ctx, "MinIO 连接失败:", err)
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "存储服务不可用",
		})
		return
	}

	// 确保 bucket 存在
	exists, _ := client.BucketExists(ctx, bucket)
	if !exists {
		if err = client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			g.Log().Error(ctx, "创建 bucket 失败:", err)
		}
	}

	objectName := fmt.Sprintf("forum/%s/%s%s", time.Now().Format("20060102"), uuid.New().String(), ext)
	_, err = client.PutObject(ctx, bucket, objectName, file, header.Size, minio.PutObjectOptions{ContentType: header.Header.Get("Content-Type")})
	if err != nil {
		g.Log().Error(ctx, "上传失败:", err)
		r.Response.WriteJsonExit(g.Map{
			"code":    500,
			"message": "上传失败",
		})
		return
	}

	// 生成访问 URL（MinIO 默认可公开读或需配置 policy）
	protocol := "http"
	if useSSL {
		protocol = "https"
	}
	url := fmt.Sprintf("%s://%s/%s/%s", protocol, endpoint, bucket, objectName)
	r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"data": g.Map{"url": url},
	})
}
