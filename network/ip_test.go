package network

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HostIP(t *testing.T) {
	ip := HostIP()
	assert.NotEmpty(t, ip)
	fmt.Printf("IP:%s\n", ip)
}

func Test_LocalIP(t *testing.T) {
	ip := LocalIP()
	assert.NotEmpty(t, ip)
	fmt.Printf("IP:%s\n", ip)

}

func Test_PublicIP(t *testing.T) {
	ip := PublishIP()
	assert.NotEmpty(t, ip)
	fmt.Printf("IP:%s\n", ip)

}
