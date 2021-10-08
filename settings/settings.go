package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序所有的配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host              string `mapstructure:"host"`
	User              string `mapstructure:"user"`
	Password          string `mapstructure:"password"`
	Dbname            string `mapstructure:"dbname"`
	MysqlMaxOpenConns int    `mapstructure:"mysql_max_open_conns"`
	MysqlMaxIdleConns int    `mapstructure:"mysql_max_idle_conns"`
	Port              int    `mapstructure:"port"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml") //使用这个直接指定不会报空指针异常
	//viper.SetConfigName("config")      // 配置文件名称(无扩展名)同时存在yaml和json格式的配置文件会报空指针异常
	//viper.SetConfigType("yaml")        // 如果配置文件的名称中没有扩展名，则需要配置此项(专用远程获取配置信息时指定配置文件类型)
	//viper.AddConfigPath(".")   // 查找配置文件所在的路径

	//viper.SetConfigFile(filepath) //获取命令行参数使用这个
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("ReadInConfig fail: err:%v \n", err)
		return
	}
	//把读取到的配置信息反序列化到Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("unmarshal conf failed, err:%s \n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已经修改")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("unmarshal conf failed, err:%s \n", err)
		}

	})
	return
}
