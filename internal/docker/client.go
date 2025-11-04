package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

// ClientManager 统一的 Docker 客户端管理器
type ClientManager struct {
	cli *client.Client
}

// NewClientManager 创建新的 Docker 客户端管理器
func NewClientManager() (*ClientManager, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, fmt.Errorf("创建 Docker 客户端失败: %w", err)
	}

	return &ClientManager{cli: cli}, nil
}

// GetClient 获取 Docker 客户端实例
func (cm *ClientManager) GetClient() *client.Client {
	return cm.cli
}

// Close 关闭 Docker 客户端连接
func (cm *ClientManager) Close() error {
	if cm.cli != nil {
		return cm.cli.Close()
	}
	return nil
}

// Ping 检查 Docker 服务是否可用
func (cm *ClientManager) Ping(ctx context.Context) error {
	_, err := cm.cli.Ping(ctx)
	if err != nil {
		return fmt.Errorf("发现 Docker 服务不可用: %w", err)
	}
	return nil
}
