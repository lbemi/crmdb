package bootstrap

import (
	"fmt"
	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/model/config"
	"github.com/lbemi/lbemi/pkg/model/rules"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"

	"io"
	"os"
	"strconv"
	"time"
)

func InitializeDB(c *config.Config) *gorm.DB {
	switch c.Driver {
	case "mysql":
		return initMysqlGorm(c)
	default:
		return initMysqlGorm(c)

	}

}

func initMysqlGorm(c *config.Config) *gorm.DB {
	dataConfig := c.Database
	if dataConfig.Database == "" {
		return nil
	}

	dsn := dataConfig.User + ":" + dataConfig.Password + "@tcp(" + dataConfig.Host + ":" + strconv.Itoa(dataConfig.Port) + ")/" +
		dataConfig.Database + "?charset=" + dataConfig.Charset + "&parseTime=True&loc=Local"

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
		Logger:                                   getGormLogger(c),
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "tb_",
		//	SingularTable: true,
		//},
	}); err != nil {
		//logs.Logger.Error("mysql connect failed. err:", err)
		fmt.Println("mysql connect failed. err:", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dataConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dataConfig.MaxOpenConns)
		if dataConfig.IsInitialize {
			migration(db)
		}
		return db
	}
}

func migration(db *gorm.DB) {
	fmt.Println("初始化数据库...")
	err := db.AutoMigrate(
		&sys.Menu{},
		&sys.User{},
		&sys.Role{},
		&sys.RoleMenu{},
		&sys.UserRole{},
		&rules.Rule{},
		&asset.Host{},
		&cloud.Config{},
	)
	if err != nil {
		fmt.Println("初始化数据库失败。。。。。", err)
		return
	}

}

func getGormLogWriter(c *config.Config) logger.Writer {
	logConfig := c.Log
	var writer io.Writer
	if c.Database.EnableFileLogWrite {
		writer = &lumberjack.Logger{
			Filename:   logConfig.RootDir + "/" + c.LogFilename,
			MaxSize:    logConfig.MaxSize,
			MaxAge:     logConfig.MaxAge,
			MaxBackups: logConfig.MaxBackup,
			Compress:   logConfig.Compress,
		}
	} else {
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger(c *config.Config) logger.Interface {
	var logMode logger.LogLevel
	switch c.LogMode {
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
	return logger.New(getGormLogWriter(c), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logMode,
		IgnoreRecordNotFoundError: true,
		Colorful:                  !c.EnableFileLogWrite,
	})
}
