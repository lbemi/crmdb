package bootstrap

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	opsLog "github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/cloud"
	"github.com/lbemi/lbemi/pkg/model/config"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/model/rules"
	"github.com/lbemi/lbemi/pkg/model/sys"
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dataConfig.User,
		dataConfig.Password,
		dataConfig.Host,
		dataConfig.Port,
		dataConfig.Database,
		dataConfig.Charset,
	)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用自动创建外键约束
		Logger:                                   getGormLogger(c),
	})

	if err != nil {
		opsLog.Logger.Errorf("mysql connect failed. err: %v", err)
		os.Exit(-13)
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dataConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dataConfig.MaxOpenConns)
	if dataConfig.IsInitialize {
		migration(db)
	}
	return db
}

func migration(db *gorm.DB) {
	entities := []interface{}{
		&sys.Menu{},
		&sys.User{},
		&sys.Role{},
		&sys.RoleMenu{},
		&sys.UserRole{},
		&rules.Rule{},
		&asset.Host{},
		&asset.Group{},
		&asset.HostGroup{},
		&asset.HostAccount{},
		&asset.Account{},
		&cloud.Cluster{},
		&sys.UserResource{},
		&logsys.LogLogin{},
		&logsys.LogOperator{},
	}
	opsLog.Logger.Info("Initializing database ...")
	if err := db.AutoMigrate(entities...); err != nil {
		opsLog.Logger.Errorf("Failed to initialize database. err: %v", err)
		return
	}
}

func getGormLogger(c *config.Config) logger.Interface {
	logMode := getLoggerMode(c.LogMode)
	config := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logMode,
		IgnoreRecordNotFoundError: true,
		Colorful:                  !c.EnableFileLogWrite,
	}
	return logger.New(getGormLogWriter(c), config)
}

func getLoggerMode(logMode string) logger.LogLevel {
	switch logMode {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}

func getGormLogWriter(c *config.Config) logger.Writer {
	logConfig := c.Log

	var writer io.Writer
	if c.Database.EnableFileLogWrite {
		writer = &lumberjack.Logger{
			Filename:   filepath.Join(logConfig.RootDir, c.LogFilename),
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
