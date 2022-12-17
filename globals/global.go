package globals

import "gorm.io/gorm"

/**************配置文件对应结构体***************/
type Configuration struct {
	Api   Api   `yaml:"api"`
	Mysql Mysql `yaml:"mysql"`
	Jwt   Jwt   `yaml:"jwt"`
}
type Api struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
type Mysql struct {
	Dsn          string `yaml:"dsn"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}
type Jwt struct {
	ExpireDuration int    `yaml:"expire_duration"`
	Secret         string `yaml:"secret"`
	Issuer         string `yaml:"issuer"`
}

/**************配置文件对应结构体***************/

// type Application struct {
// 	Config Configuration
// }

var Confok = new(Configuration)

// 定义数据库连接
var DB *gorm.DB

const (
	ERROR      = -1
	ERRORLOGIN = 1001
	SUCCESS    = 0
)
