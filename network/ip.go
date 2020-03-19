package network

import (
	"net"
	"strings"

	"github.com/axolotlteam/thunder/logger"
	"github.com/parnurzeal/gorequest"
)

//
const (
	IPIFY = "https://api.ipify.org?format=json"
)

// LocalIP -
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logger.WithError(err).Panic("get ip error")
	}
	for _, v := range addrs {
		if ip, ok := v.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}
	return ""
}

// HostIP -
func HostIP() string {
	con, _ := net.Dial("udp", "8.8.8.8:80")
	defer con.Close()
	ip := con.LocalAddr().String()
	xip := strings.Split(ip, ":")
	return xip[0]
}

// PublishIP -
func PublishIP() string {
	r := gorequest.New()
	x := struct {
		IP string `json:"ip"`
	}{}

	_, _, errs := r.Get(IPIFY).EndStruct(&x)
	for _, err := range errs {
		if err != nil {
			logger.WithError(err).Error("connect ipify failed")
			return ""
		}
	}
	return x.IP
}
