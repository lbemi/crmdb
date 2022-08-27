package bootstrap

import (
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func InitializeDB() *gorm.DB {
	switch global.App.Config.Database.Driver {
	case "mysql":
		return initMysqlGorm()
	default:
		return initMysqlGorm()

	}

}

func initMysqlGorm() *gorm.DB {
	dbConfig := global.App.Config.Database
	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{

		DisableForeignKeyConstraintWhenMigrating: true, //禁用自动创建外键约束
		Logger:                                   getGormLogger(),
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "tb_",
		//	SingularTable: true,
		//},
	}); err != nil {
		global.App.Log.Error("mysql connect failed. err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		if dbConfig.IsInitialize {
			migration(db)
		}
		return db
	}
}

func migration(db *gorm.DB) {
	global.App.Log.Info("初始化数据库...")
	err := db.AutoMigrate(&sys.Menu{}, sys.User{}, sys.Role{}, sys.RoleMenu{}, sys.UserRole{})
	if err != nil {
		global.App.Log.Error("初始化数据库失败。。。。。", zap.Any("err", err))
		return
	}

}

func getGormLogWriter() logger.Writer {
	var writer io.Writer
	if global.App.Config.Database.EnableFileLogWrite {
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.RootDir + "/" + global.App.Config.Database.LogFilename,
			MaxSize:    global.App.Config.MaxSize,
			MaxAge:     global.App.Config.MaxAge,
			MaxBackups: global.App.Config.MaxBackup,
			Compress:   global.App.Config.Compress,
		}
	} else {
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch global.App.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}
	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logMode,
		IgnoreRecordNotFoundError: true,
		Colorful:                  !global.App.Config.Database.EnableFileLogWrite,
	})
}
