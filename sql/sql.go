package sql

import (
	"fmt"
	"studysystem/config"
	"studysystem/models"

	re "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Rediscc *re.Client
	DB      *gorm.DB
	PgDB    *gorm.DB
)

// 注册数据库引擎
func InitSql() {
	InitRedis()
	InitMysql()
	InitPgsql()
}

// --------------------------------注册服务----------------------------------
// 注册redis
func InitRedis() {
	addr := config.RedisHost + ":" + config.RedisPort
	Rediscc = re.NewClient(&re.Options{
		Addr:     addr,
		Password: config.RedisPassword,
		DB:       0, // use default DB
	})
}

// 注册mysql
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlUser, config.MysqlPwd, config.MysqlHost, config.MysqlPort, config.MysqlDbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
	}
}
func InitPgsql() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.PgsqlHost, config.PgsqlUser, config.PgsqlPwd, config.PgsqlDbName, config.PgsqlPort)
	var err error
	PgDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

//注册pgsql

// --------------------------------获取操作服务的句柄----------------------------------
// 得到redis句柄
func GetRedis() *re.Client {
	return Rediscc
}

// 得到mysql的句柄
func GetMysqlDB() *gorm.DB {
	return DB
}

// 得到pgsql的句柄
func GetPgsql() *gorm.DB {
	return PgDB
}

// --------------------------------初始化表----------------------------------
func RForm() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Video{})
	DB.AutoMigrate(&models.Study_route{})
}
