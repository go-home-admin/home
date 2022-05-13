package election

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/servers"
)

// GetServer 提供统一命名规范的独立服务
func GetServer(leaders ...interface{}) constraint.KernelServer {
	return servers.GetServer(leaders)
}
