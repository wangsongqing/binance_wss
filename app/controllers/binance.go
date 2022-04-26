package controllers

import (
	BinanceConfig "binance/config"
	"binance/pkg/config"
	"binance/pkg/curl"
	BinanceHmac "binance/pkg/hash"
	"binance/pkg/helpers"
	"binance/pkg/websocket"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
)

type auth struct {
	Timestamp string `json: timestamp`
	Signature string `json: signature`
}

func Run(key string) {
	for {
		reset := initRun(key)
		if reset == true {
			continue
		}

		fmt.Println("listKey 过期了")
	}
}

func initRun(key string) bool {
	listenKeyUrl := BinanceConfig.GetListKeyUrl("/fapi/v1/listenKey")

	apiSecret := helpers.String2Bytes(helpers.FmtStrFromInterface(config.Env(strings.ToUpper(key) + "_SECRET")))

	signature := BinanceHmac.SetSignature(apiSecret, []byte(""))
	timestamp := time.Now().UnixNano() / 1e6
	s := strconv.FormatInt(timestamp, 10)

	var data auth
	data.Timestamp = s
	data.Signature = signature
	bs, _ := json.Marshal(data)

	reader := bytes.NewReader(bs)

	header := map[string]string{
		"X-MBX-APIKEY": helpers.FmtStrFromInterface(config.Env(strings.ToUpper(key) + "_KEY")),
	}

	body := curl.POST(listenKeyUrl, reader, header)
	listenKey := gjson.GetBytes(body, "listenKey")

	socketUrl := fmt.Sprintf("wss://fstream.binance.com/ws/%s", listenKey)

	nowTimes := time.Now().Unix()
	return websocket.Client(socketUrl, key, nowTimes)
}