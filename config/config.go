package config

type Config struct {
	System System `mapstructure:"system" json:"system" yaml:"system"` // 系统配置
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`          // 日志配置
}

// System 系统配置
type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`    // 环境值
	Name string `mapstructure:"name" json:"name" yaml:"name"` // 系统名称
	Url  string `mapstructure:"url" json:"url" yaml:"url"`    // 系统地址
	Port int    `mapstructure:"port" json:"port" yaml:"port"` // 系统端口
}

// Log 日志配置
type Log struct {
	RootDir         string `mapstructure:"rootDir" json:"rootDir" yaml:"rootDir"`                         // 日志根目录
	Format          string `mapstructure:"format" json:"format" yaml:"format"`                            // 写入格式
	TimestampFormat string `mapstructure:"timestampFormat" json:"timestampFormat" yaml:"timestampFormat"` // 日志时间格式
	MaxBackups      int    `mapstructure:"maxBackups" json:"maxBackups" yaml:"maxBackups"`                // 旧文件的最大个数
	MaxSize         int    `mapstructure:"maxSize" json:"maxSize" yaml:"maxSize"`                         // 日志文件最大大小(MB)
	MaxAge          int    `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`                            // 旧文件的最大保留天数
	Compress        bool   `mapstructure:"compress" json:"compress" yaml:"compress"`                      // 是否压缩
}
