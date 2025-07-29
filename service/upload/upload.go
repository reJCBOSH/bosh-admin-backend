package upload

import (
    "crypto/md5"
    "encoding/hex"
    "errors"
    "io"
    "mime/multipart"

    "bosh-admin/core/exception"
    "bosh-admin/dao"
    "bosh-admin/dao/model"
    "bosh-admin/global"

    "github.com/h2non/filetype"
)

type OssSvc struct{}

func NewOssSvc() *OssSvc {
    return &OssSvc{}
}

func (svc *OssSvc) Upload(file *multipart.FileHeader, where, source, ip string) (*model.Resource, error) {
    src, err := file.Open()
    if err != nil {
        return nil, exception.NewException("打开文件失败", err)
    }
    defer func(src multipart.File) {
        _ = src.Close()
    }(src)
    // 计算MD5校验和
    hash := md5.New()
    if _, err = io.Copy(hash, src); err != nil {
        return nil, exception.NewException("计算MD5校验和失败", err)
    }
    checkSum := hex.EncodeToString(hash.Sum(nil))
    // 重置文件指针到开头
    if _, err = src.Seek(0, 0); err != nil {
        return nil, exception.NewException("重置文件指针失败", err)
    }
    s := dao.NewStatement()
    s.Where("source = ? AND check_sum = ?", source, checkSum)
    resource, err := dao.QueryOne[model.Resource](s)
    if err == nil {
        return &resource, nil
    } else {
        if !errors.Is(err, dao.NotFound) {
            return nil, exception.NewException("查询资源记录失败", err)
        }
    }
    buf, _ := io.ReadAll(src)
    kind, _ := filetype.Match(buf)
    // 重置文件指针到开头
    if _, err = src.Seek(0, 0); err != nil {
        return nil, exception.NewException("重置文件指针失败", err)
    }
    var storePath string
    var fullPath string
    switch global.Config.Server.OssType {
    case "local":
        storePath, fullPath, err = LocalUpload(src, file.Filename, where)
    case "aliyun-oss":
        storePath, fullPath, err = AliyunOssUpload(src, file.Filename, where)
    }
    if err != nil {
        return nil, err
    }
    newResource := model.Resource{
        Source:    source,
        IP:        ip,
        FileName:  file.Filename,
        FileSize:  file.Size,
        FileType:  kind.Extension,
        MimeType:  kind.MIME.Value,
        StorePath: storePath,
        FullPath:  fullPath,
        CheckSum:  checkSum,
    }
    if err = dao.Create(&newResource); err != nil {
        return nil, exception.NewException("创建资源记录失败", err)
    }
    return &newResource, nil
}
