package config

type System struct {
	Mode         string `mapstructure:"mode" json:"mode" yaml:"mode"`             // 环境值
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`             // 端口值
	OssType      string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // 文件
	LimitCountIP int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
