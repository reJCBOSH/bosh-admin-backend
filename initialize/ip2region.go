package initialize

import (
    "fmt"

    "bosh-admin/global"

    "github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

// InitIp2Region 初始化ip2region
func InitIp2Region() {
    cBuff, err := xdb.LoadContentFromFile("ip2region.xdb")
    if err != nil {
        panic(fmt.Sprintf("加载ip2region.xdb文件失败: %v", err))
    }
    global.XdbSearcher, err = xdb.NewWithBuffer(cBuff)
    if err != nil {
        panic(fmt.Sprintf("初始化XdbSearcher失败: %v", err))
    }
}
