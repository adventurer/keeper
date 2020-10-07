package traffic

import (
	"keeper/config"
	"keeper/postdata"
	"time"

	"github.com/shirou/gopsutil/net"
)

var reciveTmp uint64
var sendTmp uint64
var recive uint64
var send uint64

// Init dfsfs
func Init() {
	i := 0
	for {

		psnet, _ := net.IOCounters(true)
		for _, item := range psnet {

			if item.Name == "eth0" {

				i++
				recive += (item.BytesRecv - reciveTmp)
				send += (item.BytesSent - sendTmp)
				if reciveTmp > 0 {

				}
				if i%5 == 0 {
					// log.Println("set traffic...")
					postdata.Content.SetTraffic(recive/1024*8/uint64(config.Config.BankWidth), send/1024*8/5)
					recive = 0
					send = 0
				}
				reciveTmp = item.BytesRecv
				sendTmp = item.BytesSent

				// log.Printf("device: %s,recive: %d(kb),send %d(kb)", item.Name,
				// 	(item.BytesRecv-recive)/1024*8, (item.BytesSent-send)/1024*8)

			}

		}

		time.Sleep(time.Second * 1)
	}
}
