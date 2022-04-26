package websocket

import (
	"binance/pkg/helpers"
	"binance/pkg/redis"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

var done chan interface{}
var interrupt chan os.Signal

func Client(wss string, account string, timestamp int64) bool  {
	done = make(chan interface{}) // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT
	fmt.Printf("wss url:%s, timestamp:%d \n", wss, timestamp)

	conn, _, err := websocket.DefaultDialer.Dial(wss, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
			}

			mapMes := helpers.JsonToMap(message)

			nowTime := time.Now().Unix()
			diffTime := nowTime - timestamp

			if diffTime > 3600 {
				break
			}

			if _, ok := mapMes["e"]; !ok {
				// fmt.Printf("error:%s \n", mapMes["error"])
				continue
			}

			redisKey := redis.GetLpushKeyBinance(account)
			lPush := redis.Redis.Rpush(redisKey, string(message))
			if !lPush {
				fmt.Printf("写入队列失败:%s \n", string(message))
			}
			fmt.Printf("wss message:%s \n", string(message))
		}
	}()

	// 2秒钟发一次心跳
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return false
		case t := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return false
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return false
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return false
		}
	}

	return true
}
