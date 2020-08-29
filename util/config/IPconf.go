package config

var (
	selfIP string
)

func SetSelfIP(ip string) {
	selfIP = ip
}

func GetSelfIP() string {
	return selfIP
}
