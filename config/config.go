package config

import (
	"os"

	"github.com/spf13/viper"
)

// 声明配置
var (
	ServerPort    string
	ServerHost    string
	RedisPort     string
	RedisHost     string
	RedisPassword string
	MysqlPort     string
	MysqlHost     string
	MysqlPwd      string
	MysqlDbName   string
	MysqlUser     string
	LocalPath     string
	Email         string
	EmailAddr     string
	EmailAuth     string
	EmailFrom     string
	EmailHost     string
	EmailExpire   string
	PgsqlDbName   string
	PgsqlHost     string
	PgsqlPort     string
	PgsqlPwd      string
	PgsqlUser     string
	//MachineID   int
)

// 注册配置
func ConfigInit() {
	InitConfig()
	ServerPort = viper.GetString("Server.port")
	ServerHost = viper.GetString("Server.host")
	RedisHost = viper.GetString("Redis.host")
	RedisPort = viper.GetString("Redis.Port")
	RedisPassword = viper.GetString("Redis.password")
	MysqlHost = viper.GetString("Mysql.host")
	MysqlPort = viper.GetString("Mysql.port")
	MysqlPwd = viper.GetString("Mysql.password")
	MysqlDbName = viper.GetString("Mysql.dbname")
	MysqlUser = viper.GetString("Mysql.user")
	PgsqlHost = viper.GetString("pgsql.host")
	PgsqlPort = viper.GetString("pgsql.port")
	PgsqlPwd = viper.GetString("pgsql.password")
	PgsqlDbName = viper.GetString("pgsql.dbname")
	PgsqlUser = viper.GetString("pgsql.user")
	LocalPath = viper.GetString("localfile.path")
	Email = viper.GetString("email.email")
	EmailAddr = viper.GetString("email.addr")
	EmailExpire = viper.GetString("email.expire")
	EmailAuth = viper.GetString("email.auth")
	EmailFrom = viper.GetString("email.from")
	EmailHost = viper.GetString("email.host")
}

// 获取配置
func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir + "/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
