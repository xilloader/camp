package datasource

import (
	"context"
	"fmt"
	"github.com/xilloader/glog"
	"gorm.io/gorm/utils"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Mysql struct {
	Name     string
	Host     string
	Port     string
	Prefix   string
	Database string
	Password string
}

// 连接

func (m Mysql) GetConn() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Name, m.Password, m.Host, m.Port, m.Database)
}

var MysqlDB *gorm.DB

//初始化数据库连接
func InitMysql(m Mysql, logr logger.Interface) {
	var sql string
	sql = m.GetConn()
	if logr == nil {
		logr = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags|log.Llongfile), logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	}
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(sql), &gorm.Config{
		Logger:                                   logr,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sdb, e := MysqlDB.DB()
	if e != nil {
		panic(e)
	}
	// 启用gger，显示详细日志
	sdb.SetMaxIdleConns(10)
	sdb.SetMaxOpenConns(1024)
	sdb.SetConnMaxLifetime(time.Second * 600)
}

func LocalMysql(dbName, prefix string) {
	InitMysql(Mysql{
		Name:     "root",
		Host:     "127.0.0.1",
		Port:     "3306",
		Prefix:   prefix,
		Database: dbName,
		Password: "332214",
	}, nil)
}

//获取MySql连接
func GetDB() *gorm.DB {
	return MysqlDB
}

func GormGlogLogger(level glog.Level) logger.Interface {
	return logger.New(glog.V(level),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		})
}

func NewGormLogger(lc logger.Config) logger.Interface {
	return &gormLogger{
		lc,
	}
}

type gormLogger struct {
	logger.Config
}

func (gl *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *gl
	newlogger.LogLevel = level
	return &newlogger
}

func (gl *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	glog.InfolnWithTrace(ctx, msg, append([]interface{}{utils.FileWithLineNum()}, data...))
}

func (gl *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	glog.WarninglnWithTrace(ctx, msg, append([]interface{}{utils.FileWithLineNum()}, data...))
}

func (gl *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	glog.ErrorWithTrace(ctx, msg, append([]interface{}{utils.FileWithLineNum()}, data...))
}

func (gl *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if gl.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && gl.LogLevel >= logger.Error:
			sql, rows := fc()
			glog.ErrorWithTrace(ctx, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		case elapsed > gl.SlowThreshold && gl.SlowThreshold != 0 && gl.LogLevel >= logger.Warn:
			sql, rows := fc()
			glog.WarninglnWithTrace(ctx, "SlowThreshold:"+string(gl.SlowThreshold), utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		case gl.LogLevel >= logger.Info:
			sql, rows := fc()
			glog.InfolnWithTrace(ctx, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
