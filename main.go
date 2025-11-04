package main

import (
	"context"
	"watchducker/cmd"
	"watchducker/pkg/config"
	"watchducker/pkg/logger"
)

func main() {
	if err := config.Load(); err != nil {
		logger.Fatal("初始化失败: %v", err)
	}

	ctx := context.Background()

	// 如果指定了 cron 表达式，则进入定时执行模式
	if config.Get().CronExpression() != "" {
		cmd.RunCronScheduler(ctx)
		return
	}

	cmd.RunOnce(ctx)
}
