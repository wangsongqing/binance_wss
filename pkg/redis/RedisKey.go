package redis


func GetLpushKeyBinance(key string) string {
	return "binance_go_push_list:" + key
}
