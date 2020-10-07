package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"io/ioutil"
	"keeper/config"
	"keeper/postdata"
	"keeper/traffic"
	"log"
	"strconv"
	"time"
)

type serverConfig struct {
	Server      string `json:"server"`
	Server_port int    `json:"server_port"`
	Local_port  int    `json:"local_port"`
	Password    string `json:"password"`
	Timeout     int    `json:"timeout"`
	Method      string `json:"method"`
}

var pause = make(chan int)

func main() {
	flag.StringVar(&config.Config.Name, "name", "", "name need")
	flag.StringVar(&config.Config.Ip, "ip", "", "ip need")
	flag.IntVar(&config.Config.Port, "port", 0, "port need")
	flag.IntVar(&config.Config.BandWidth, "bandwidth", 0, "bandwidth need")

	flag.Parse()

	if config.Config.Name == "" {
		log.Fatal("服务器名必须")
	}

	if config.Config.Ip == "" {
		log.Fatal("服务器ip必须")
	}

	if config.Config.Port == 0 {git a
		log.Fatal("服务器端口必须")
	}

	pass := randomStr()

	postdata.Content.SetInfo(config.Config.Name, config.Config.Ip, strconv.Itoa(config.Config.Port), pass)

	cfg := serverConfig{Server: "0.0.0.0", Server_port: config.Config.Port, Local_port: 1080, Password: pass, Timeout: 60, Method: "chacha20-ietf-poly1305"}
	cfgString, err := json.Marshal(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(cfgString))
	ioutil.WriteFile("/root/keeper/config.json", cfgString, 0666)

	// var outInfo bytes.Buffer
	// cmd := exec.Command("/etc/init.d/shadowsocks-libev", "restart")

	// // 设置接收
	// cmd.Stdout = &outInfo

	// // 执行
	// cmd.Run()

	// fmt.Println(outInfo.String())

	// /etc/init.d/shadowsocks-libev restart
	go traffic.Init()

	go postdata.Init()

	<-pause
}

func randomStr() string {
	crutime := time.Now().UnixNano()
	ctx := md5.New()
	ctx.Write([]byte(strconv.FormatInt(crutime, 10)))
	return hex.EncodeToString(ctx.Sum(nil))
}
