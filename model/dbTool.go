package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var _db *gorm.DB

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init()  {
	var err  error
	config := viper.New()					//初始化viper读取配置
	config.AddConfigPath("./config/")
	config.SetConfigName("database")	//设置文件名
	config.SetConfigType("ini")			//设置文件类型
	errConfig:= config.ReadInConfig()				//读取文件
	if errConfig != nil {
		log.Println("config,ReadFail")
	}
	//获取数据库配置
	host := config.GetString("mysql.host")		//读取对应配置
	userName := config.GetString("mysql.username")
	pwd := config.GetString("mysql.pwd")
	dbName := config.GetString("mysql.dbName")
	port := config.GetString("mysql.port")

	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	//类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", userName, pwd, host, port, dbName)
	//链接MYSQL
	_db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Println("连接数据库失败,error="+err.Error())
	}else{
		//log.Println(db)
	}
	sqlDB,_ := _db.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100)   //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

}

func GetDb() *gorm.DB  {
	return _db
}