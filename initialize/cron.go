package initialize

import (
    "bosh-admin/core/log"

    "github.com/robfig/cron/v3"
)

func InitCron() {
    var err error
    // 初始化定时任务
    c := cron.New()
    // 每小时执行一次
    _, err = c.AddFunc("0 0 0/1 * * ?", func() {
        // 定时任务逻辑
    })
    if err != nil {
        log.Error("添加每小时定时任务失败", err)
    }
    // 每天执行一次
    _, err = c.AddFunc("0 0 0 1/1 * ?", func() {
        // 定时任务逻辑
    })
    if err != nil {
        log.Error("添加每天定时任务失败", err)
    }
    // 每月执行一次
    _, err = c.AddFunc("0 0 0 1 * ?", func() {
        // 定时任务逻辑
    })
    if err != nil {
        log.Error("添加每月定时任务失败", err)
    }
    // 启动定时任务
    c.Start()
}
