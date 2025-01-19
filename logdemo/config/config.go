package config

import "gopkg.in/lumberjack.v3"

type LogConfig struct {
	Filename   string
	MaxSize    int64
	BackUp     int
	BufferSize int
	Compress   bool
}

func NewDefaultConfig() *LogConfig {
	return &LogConfig{
		Filename:   "foo.log",
		MaxSize:    1 * lumberjack.MB,
		BackUp:     1,
		BufferSize: 1, //不缓冲立即写入
		Compress:   false,
	}
}
