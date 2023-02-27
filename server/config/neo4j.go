package config

type Neo4j struct {
	Addr      string `mapstructure:"addr" json:"addr" yaml:"addr"`                // 服务器地址:端口  bolt://localhost:7687
	User      string `mapstructure:"user" json:"user" yaml:"user"`                // 用户名
	Password  string `mapstructure:"password" json:"password" yaml:"password"`    // 密码
	Encrypted bool   `mapstructure:"encrypted" json:"encrypted" yaml:"encrypted"` // 加密
}
