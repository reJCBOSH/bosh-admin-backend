package config

type Config struct {
	System System `mapstructure:"system" json:"system" yaml:"system"` // 系统配置
}

type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`    // 环境值
	Name string `mapstructure:"name" json:"name" yaml:"name"` // 系统名称
	Url  string `mapstructure:"url" json:"url" yaml:"url"`    // 系统地址
	Port int    `mapstructure:"port" json:"port" yaml:"port"` // 系统端口
}
