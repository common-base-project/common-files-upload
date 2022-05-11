package connection

import (
	"fmt"
	"gorm.io/gorm/schema"
	_ "unicorn-files/pkg/config"
	"unicorn-files/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

// DbEnable 是否开启数据库存储
var DbEnable bool

func Initial() {
	DbEnable = viper.GetBool("db.db_enable")
	if !DbEnable {
		return
	}
	DB.Init()
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	logger.Info("database url: ", config)
	//db, err := gorm.Open("mysql", config)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: config,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		//Logger: logger.GetLogger(),
	})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败，连接地址: %s，error: %s", viper.GetString(`db.addr`), err))
	}

	// 设置字符集
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	// 是否开启详细日志记录
	//db.LogMode(viper.GetBool(`db.gorm.logMode`))

	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(viper.GetInt(`db.gorm.maxOpenConn`))

	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用
	sqlDB.SetMaxIdleConns(viper.GetInt(`db.gorm.maxIdleConn`))

	// 创建表的时候去掉复数
	//db.SingularTable(viper.GetBool(`db.gorm.singularTable`))
}

func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func (db *Database) Init() {
	if !DbEnable {
		return
	}
	DB = &Database{
		Self: InitSelfDB(),
	}
}

func (db *Database) Close() {
	if !DbEnable {
		return
	}
	sqlDB, err := db.Self.DB()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = sqlDB.Close()
	//err := DB.Self.Close()
	if err != nil {
		logger.Error("关闭连接失败，错误信息: %s", err)
	}
}
