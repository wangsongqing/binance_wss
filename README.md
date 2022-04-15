## 自己写的一套币安量化交易，功能后续会逐步完善

## websock 跟单系统

## 运行方式
 - 运行项目方式: go run main.go -controllerName="xxxxx"
 - 打包命令: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o binance
 - 同步到服务器: scp binance root@127.0.0.1:/www/data
 - 在目录下面新建 .env 文件，把相应的配置写进去，然后就可以愉快的玩耍了