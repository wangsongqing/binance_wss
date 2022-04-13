package main

import (
	BinanceConfig "binance/config"
	btsConfig "binance/config"
	"binance/pkg/config"
	"binance/pkg/curl"
	BinanceHmac "binance/pkg/hash"
	"binance/pkg/helpers"
	"binance/pkg/websocket"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	_ "strings"
	"time"
)

type auth struct {
	Timestamp string `json: timestamp`
	Signature string `json: signature`
}

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

// go run main.go
func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	listenKeyUrl := BinanceConfig.GetListKeyUrl("/fapi/v1/listenKey")

	apiSecret := helpers.String2Bytes(helpers.FmtStrFromInterface(config.Env("API_SECRET")))

	signature := BinanceHmac.SetSignature(apiSecret, []byte(""))
	timestamp := time.Now().UnixNano() / 1e6
	s := strconv.FormatInt(timestamp, 10)

	var data auth
	data.Timestamp = s
	data.Signature = signature
	bs, _ := json.Marshal(data)

	reader := bytes.NewReader(bs)

	header := map[string]string{
		"X-MBX-APIKEY": helpers.FmtStrFromInterface(config.Env("API_KEY")),
	}

	body := curl.POST(listenKeyUrl, reader, header)
	listenKey := helpers.JsonToMap(body)


	socketUrl := fmt.Sprintf("wss://fstream.binance.com/ws/%s", listenKey["listenKey"])
	websocket.Client(socketUrl)

}
