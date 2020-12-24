package srv

import (
	"fmt"
)

// NewTags -
func NewTags(opts *Options) []string {

	tags := []string{
		"traefik.enable=true",
		fmt.Sprintf("traefik.http.routers.%s.rule=Host(`%s.traefik`)", opts.name, opts.name),
		fmt.Sprintf("traefik.http.routers.%s.service=%s-service", opts.name, opts.name),
		fmt.Sprintf("traefik.http.routers.%s.middlewares=latency-check,do-retry", opts.name),
		fmt.Sprintf("traefik.http.services.%s-service.loadbalancer.passhostheader=true", opts.name),
		fmt.Sprintf("traefik.http.services.%s-service.loadbalancer.server.scheme=%s", opts.name, traefikScheme()),
	}
	// 中間件標籤
	tags = append(tags, httpMiddlewareTags()...)
	tags = append(tags, opts.name, opts.id)

	return tags
}

func httpMiddlewareTags() []string {
	return []string{
		// 錯誤閘道
		"traefik.http.middlewares.latency-check.circuitbreaker.expression=NetworkErrorRatio() > 0.50",
		// 重傳 3次
		"traefik.http.middlewares.do-retry.retry.attempts=3",
	}
}

func traefikScheme() string {
	scheme := "http"

	switch ServerType {
	case GRPC:
		scheme = "h2c"
	}
	return scheme
}
