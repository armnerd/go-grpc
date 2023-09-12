package base

import (
	"fmt"
	"log"
	"time"

	"go-grpc/pkg/sqlLogger"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbProvider = wire.NewSet(NewUserDB, NewBlogDB)

type UserDB *gorm.DB
type BlogDB *gorm.DB

func NewUserDB(conf Conf) UserDB {
	dbConf := conf.Data.MysqlUser
	db := connectDB(dbConf.Host, dbConf.Port, dbConf.Database, dbConf.User, dbConf.Pass, dbConf.Charset)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(1)
	return db
}

func NewBlogDB(conf Conf) BlogDB {
	dbConf := conf.Data.MysqlBlog
	db := connectDB(dbConf.Host, dbConf.Port, dbConf.Database, dbConf.User, dbConf.Pass, dbConf.Charset)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(1)
	return db
}

func connectDB(host, port, database, user, pass, charset string) *gorm.DB {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: sqlLogger.NewTraceLogger(logger.Info, time.Second),
	})
	if err != nil {
		log.Fatalf("models.InitDbMySQL err: %v", err)
	}
	return db
}
