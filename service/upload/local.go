package upload

import (
    "fmt"
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
    "time"

    "bosh-admin/core/exception"
    "bosh-admin/core/log"
    "bosh-admin/global"
)

func LocalUpload(src multipart.File, filename, where string) (string, string, error) {
    if where == "" {
        where = "default"
    }
    date := time.Now().Format(time.DateOnly)
    dirPath := filepath.Join(global.Config.Local.StorePath, where, date)
    storePath := filepath.Join(dirPath, filename)
    fullPath := fmt.Sprintf("%s/%s/%s/%s", global.Config.Local.Path, where, date, filename)
    if global.Config.Server.Env == global.DEV {
        fullPath = fmt.Sprintf("http://%s:%d/%s", global.Config.Server.Url, global.Config.Server.Port, fullPath)
    } else {
        fullPath = fmt.Sprintf("%s/%s", global.Config.Server.Url, fullPath)
    }
    err := os.MkdirAll(dirPath, os.ModePerm)
    if err != nil {
        log.Error("创建目录失败:", err.Error())
        return "", "", exception.NewException("创建目录失败")
    }
    out, err := os.Create(storePath)
    if err != nil {
        log.Error("创建文件失败:", err.Error())
        return "", "", exception.NewException("创建文件失败")
    }
    defer func(out *os.File) {
        _ = out.Close()
    }(out)
    _, err = io.Copy(out, src)
    if err != nil {
        log.Error("写入文件失败:", err.Error())
        return "", "", exception.NewException("写入文件失败")
    }
    return storePath, fullPath, nil
}
