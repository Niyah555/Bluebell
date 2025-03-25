package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

// AppConfig 应用配置结构体
type AppConfig struct {
	Name      string `mapstructure:"name"`       // 应用名称
	Mode      string `mapstructure:"mode"`       // 运行模式
	Version   string `mapstructure:"version"`    // 版本号
	StartTime string `mapstructure:"start_time"` // 启动时间
	MachineID int64  `mapstructure:"machine_id"` // 机器ID
	Port      int    `mapstructure:"port"`       // 端口号

	*LogConfig     `mapstructure:"log"`   // 日志配置
	*MySQLConfig   `mapstructure:"mysql"` // MySQL配置
	*RedisConfig   `mapstructure:"redis"` // Redis配置
	*GoEmailConfig `mapstructure:"email"` //go-email 配置
}

// MySQLConfig MySQL配置结构体
type MySQLConfig struct {
	Host         string `mapstructure:"host"`           // 主机地址
	User         string `mapstructure:"user"`           // 用户名
	Password     string `mapstructure:"password"`       // 密码
	DB           string `mapstructure:"dbname"`         // 数据库名
	Port         int    `mapstructure:"port"`           // 端口号
	MaxOpenConns int    `mapstructure:"max_open_conns"` // 最大打开连接数
	MaxIdleConns int    `mapstructure:"max_idle_conns"` // 最大空闲连接数
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`           // 主机地址
	Password     string `mapstructure:"password"`       // 密码
	Port         int    `mapstructure:"port"`           // 端口号
	DB           int    `mapstructure:"db"`             // 数据库编号
	PoolSize     int    `mapstructure:"pool_size"`      // 连接池大小
	MinIdleConns int    `mapstructure:"min_idle_conns"` // 最小空闲连接数
}

// LogConfig 日志配置结构体
type LogConfig struct {
	Level      string `mapstructure:"level"`       // 日志级别
	Filename   string `mapstructure:"filename"`    // 日志文件名
	MaxSize    int    `mapstructure:"max_size"`    // 单个日志文件最大尺寸
	MaxAge     int    `mapstructure:"max_age"`     // 日志文件保留天数
	MaxBackups int    `mapstructure:"max_backups"` // 日志文件最大保留数
}

// GoEmailConfig 配置结构体
type GoEmailConfig struct {
	SmtpHost string `mapstructure:"smtp_host"`
	SmtpPort int    `mapstructure:"smtp_port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func Init(filePath string) (err error) {
	// 方式1：直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	//viper.SetConfigFile("./conf/config.yaml")
	// 绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("/Users/liwenzhou/Desktop/bluebell/conf/config.yaml")

	// 方式2：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	// 配置文件名不需要带后缀
	// 配置文件位置可配置多个
	//viper.SetConfigName("config") // 指定配置文件名（不带后缀）
	//viper.AddConfigPath(".") // 指定查找配置文件的路径（这里使用相对路径）
	//viper.AddConfigPath("./conf")      // 指定查找配置文件的路径（这里使用相对路径）

	// 基本上是配合远程配置中心使用的，告诉viper当前的数据使用什么格式去解析
	//viper.SetConfigType("json")

	viper.SetConfigFile(filePath)
	// 读取配置信息
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
