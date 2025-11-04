package types

import "time"

// ContainerInfo 容器信息
type ContainerInfo struct {
	ID     string            `json:"id"`
	Name   string            `json:"name"`
	Image  string            `json:"image"`
	Labels map[string]string `json:"labels"`
	State  string            `json:"state"`
}

// ImageCheckResult 镜像检查结果
type ImageCheckResult struct {
	Name       string    `json:"name"`
	LocalHash  string    `json:"local_hash"`
	RemoteHash string    `json:"remote_hash"`
	IsUpdated  bool      `json:"is_updated"`
	CheckedAt  time.Time `json:"checked_at"`
	Error      string    `json:"error,omitempty"`
}

// BatchCheckResult 批量检查结果
type BatchCheckResult struct {
	Containers []ContainerInfo     `json:"containers"`
	Images     []*ImageCheckResult `json:"images"`
	Summary    struct {
		TotalContainers int           `json:"total_containers"`
		TotalImages     int           `json:"total_images"`
		Updated         int           `json:"updated"`
		Failed          int           `json:"failed"`
		UpToDate        int           `json:"up_to_date"`
		Duration        time.Duration `json:"duration"`
	} `json:"summary"`
}

// CheckCallback 检查回调函数类型
type CheckCallback func(*ImageCheckResult)

// CheckMode 检查模式
type CheckMode int

const (
	CheckByName  CheckMode = iota // 按名称检查
	CheckByLabel                  // 按标签检查
	CheckAll                      // 检查所有容器
)
