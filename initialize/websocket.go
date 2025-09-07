package initialize

import (
	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/websocket"
)

func InitWebsocket() {
	global.WsHub = websocket.NewHub(global.Logger)
	go global.WsHub.Start()
	log.Info("websocket初始化完成")
}
