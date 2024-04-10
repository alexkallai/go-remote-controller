package qrcodeutils

import (
	"net"
	"os"

	"github.com/mdp/qrterminal"
)

func PrintQrCode(text string) {
	config := qrterminal.Config{
		Level:      qrterminal.L,
		Writer:     os.Stdout,
		HalfBlocks: false,
		BlackChar:  "  ",
		WhiteChar:  "██",
		QuietZone:  1,
	}
	qrterminal.GenerateWithConfig(text, config)
}

func GetLocalIpAdresses() []string {
	var ipAddresses []string
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ipAddresses = append(ipAddresses, ipv4.String())
		}
	}
	return ipAddresses
}
