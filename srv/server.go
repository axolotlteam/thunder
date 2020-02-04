package srv

import (
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

var (
	// DefaultServer -
	DefaultServer = newRPCServer()
	// DefaultHost -
	DefaultHost = newHost()
	// DefaultPort -
	DefaultPort = 0
	// DefaultHook -
	DefaultHook = func() {}
)

func newRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func newListener(opt *Options) net.Listener {
	port := strconv.Itoa(opt.Port)
	listener, err := net.Listen("tcp", opt.Host+":"+port)
	if err != nil {
		log.Panic(err)
	}

	opt.Port = listener.Addr().(*net.TCPAddr).Port

	return listener
}

func newHost() string {
	return localIP()
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
