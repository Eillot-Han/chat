package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	username := "root"  //账号
	password := "Wyy20020508..." //密码
	host := "sh-cynosdbmysql-grp-93vqbnis.sql.tencentcdb.com" //数据库地址，可以是Ip或者域名	
	port := "29672" //数据库端口
	Dbname := "han_im" //数据库名
	config := "charset=utf8&parseTime=true&loc=Local" //数据库连接配置
	dsn := username + ":" + password + "@tcp(" + host + ":"+ port + ")/" + Dbname + "?" + config

	mysqlConfig := mysql.Config{
		DSN:                       dsn, 
		DefaultStringSize:         191, 
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true, 
		SkipInitializeWithVersion: false, 
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}