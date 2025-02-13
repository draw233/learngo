package config

// 定义一个结构体
type Config struct {
	Host string
	Port int
}

// 定义一个函数类型，用于设置配置项
type Option func(*Config)

// 定义一些选项
func WithHost(host string) Option {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port int) Option {
	return func(c *Config) {
		c.Port = port
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: 8080,
	}
}

// 创建一个新配置的函数
func NewConfig(options ...Option) *Config {
	// 创建默认配置
	config := NewDefaultConfig()
	// 应用选项
	for _, option := range options {
		option(config)
	}

	return config
}
