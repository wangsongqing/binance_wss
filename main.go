package main

import (
	"binance/app/controllers"
	"binance/bootstrap"
	btsConfig "binance/config"
	"binance/pkg/config"
	"flag"
	"fmt"
)



func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

// go run main.go -controllerName="xxxxx"
func main() {

	controllerName := flag.String("controllerName", "", "")

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	flag.Parse()

	// 初始化 Redis
	bootstrap.SetupRedis()

	switch *controllerName {
		case "xmwme":
			controllers.Run("xmwme")
			break
		case "free":
			controllers.Run("free")
		break
		default:
			fmt.Printf("没有找到控制器: %s, 运行项目方式: go run main.go -controllerName=\"xxxxx\"", *controllerName)
		break
	}



}
