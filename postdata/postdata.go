package postdata

import (
	"keeper/config"
	"log"
	"sync"
	"time"

	"github.com/kirinlabs/HttpRequest"
)

type ReportData struct {
	Recive uint64
	Send   uint64
	Name   string
	Ip     string
	Port   string
	Pass   string
	Load   float64
	lock   sync.Mutex
}

var Content ReportData

func Init() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			// log.Printf("send: %d,recive: %d,load: %.2f,name: %s,ip: %s,port: %s,pass: %s\n",
			// Content.Send, Content.Recive, Content.Load, Content.Name, Content.Ip, Content.Port, Content.Pass)

			res, err := HttpRequest.NewRequest().Post("https://admin.jiasu.zhifool.com/api/v1/node/register", map[string]interface{}{
				"Name":     Content.Name,
				"Address":  Content.Ip + ":" + Content.Port,
				"Password": Content.Pass,
				"Recive":   Content.Recive,
				"Send":     Content.Send,
				"Load":     Content.Load,
			})
			body, err := res.Body()
			if err != nil {
				log.Println(err.Error())
				continue
			}
			log.Println(string(body))
		}
	}()

}

func (p *ReportData) SetTraffic(recive, send uint64) {
	defer p.lock.Unlock()
	p.lock.Lock()
	p.Recive = recive
	p.Send = send
	p.Load = float64(send) / (float64(config.Config.BandWidth) * 1024)
}

func (p *ReportData) SetInfo(Name, Ip, Port, Pass string) {
	defer p.lock.Unlock()
	p.lock.Lock()
	p.Name = Name
	p.Ip = Ip
	p.Port = Port
	p.Pass = Pass
}
