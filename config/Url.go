package config


const BaseUrl = "https://fapi.binance.com"


func GetListKeyUrl(url string) string {
	return BaseUrl + url
}