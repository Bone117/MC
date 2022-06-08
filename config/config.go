package config

type Server struct {
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	//Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	//Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	//AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	//Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	//TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`

	//Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	//Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`

	// 跨域配置
	//Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
