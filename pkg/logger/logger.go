package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Level 定义日志级别
type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var (
	levelNames = map[Level]string{
		DEBUG: "DEBUG",
		INFO:  "INFO",
		WARN:  "WARN",
		ERROR: "ERROR",
	}
	levelColors = map[Level]string{
		DEBUG: "\033[36m", // 青色
		INFO:  "\033[32m", // 绿色
		WARN:  "\033[33m", // 黄色
		ERROR: "\033[31m", // 红色
	}
	resetColor = "\033[0m"
)

// Logger 日志记录器
type Logger struct {
	level  Level
	output io.Writer
	prefix string
}

// New 创建新的日志记录器
func New() *Logger {
	return &Logger{
		level:  INFO,
		output: os.Stdout,
		prefix: "",
	}
}

// log 内部日志方法
func (l *Logger) log(level Level, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	timestamp := time.Now().Format(time.DateTime)
	levelName := levelNames[level]
	color := levelColors[level]

	// 构建日志消息
	message := fmt.Sprintf(format, args...)

	// 格式化输出
	logLine := fmt.Sprintf("%s%s [%-5s] %s%s\n",
		timestamp, color, levelName, message, resetColor)

	fmt.Fprint(l.output, logLine)
}

// Debug 输出调试日志
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info 输出信息日志
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warn 输出警告日志
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Error 输出错误日志
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Fatal 输出错误日志并退出程序
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
	os.Exit(1)
}

// 全局日志实例
var defaultLogger = New()

// Debug 全局调试日志
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Info 全局信息日志
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Warn 全局警告日志
func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// Error 全局错误日志
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// Fatal 全局致命错误日志
func Fatal(format string, args ...interface{}) {
	defaultLogger.Fatal(format, args...)
}

// SetLevel 设置全局日志级别
func SetLevel(levelStr string) {
	switch levelStr {
	case "DEBUG":
		defaultLogger.level = DEBUG
	case "WARN":
		defaultLogger.level = WARN
	case "ERROR":
		defaultLogger.level = ERROR
	default:
		defaultLogger.level = INFO
	}
}
