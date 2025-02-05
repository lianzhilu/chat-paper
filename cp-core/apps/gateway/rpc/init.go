package rpc

import (
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
)

var (
	serviceConfig = config.GetRuntimeConfig()
	etcdAddrs     = []string{fmt.Sprintf("%s:%s", serviceConfig.EtcdConfig.EtcdHost, serviceConfig.EtcdConfig.EtcdPort)}
)

func init() {
	InitUserClient()
	InitArticleClient()
}
