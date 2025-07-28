package upload

import (
    "mime/multipart"
    "path/filepath"
    "time"

    "bosh-admin/core/exception"
    "bosh-admin/core/log"
    "bosh-admin/global"

    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// NewBucket 创建新存储桶连接
func NewBucket() (*oss.Bucket, error) {
    // 创建oss client实例
    client, err := oss.New(global.Config.AliyunOss.Endpoint, global.Config.AliyunOss.AccessKeyId, global.Config.AliyunOss.AccessKeySecret)
    if err != nil {
        return nil, err
    }
    // 获取存储桶
    bucket, err := client.Bucket(global.Config.AliyunOss.BucketName)
    if err != nil {
        return nil, err
    }
    return bucket, nil
}

func AliyunOssUpload(src multipart.File, filename, where string) (string, string, error) {
    bucket, err := NewBucket()
    if err != nil {
        log.Error("创建存储桶失败:", err.Error())
        return "", "", exception.NewException("创建存储桶失败")
    }
    if where == "" {
        where = "default"
    }
    dirPath := filepath.Join(global.Config.AliyunOss.BasePath, where, time.Now().Format(time.DateOnly))
    storePath := filepath.Join(dirPath, filename)
    fullPath := filepath.Join(global.Config.AliyunOss.BucketUrl, storePath)
    if err = bucket.PutObject(storePath, src); err != nil {
        log.Error("上传文件失败:", err.Error())
        return "", "", exception.NewException("上传文件失败")
    }
    return storePath, fullPath, nil
}
